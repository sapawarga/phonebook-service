package endpoint

import (
	"context"
	"fmt"

	"encoding/json"

	"github.com/sapawarga/phonebook-service/helper"
	"github.com/sapawarga/phonebook-service/model"
	"github.com/sapawarga/phonebook-service/usecase"

	"github.com/go-kit/kit/endpoint"
)

// MakeGetList ...
func MakeGetList(ctx context.Context, usecase usecase.Provider) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*GetListRequest)
		var limit, page, offset int64 = 10, 1, 0
		if req.Limit != 0 {
			limit = req.Limit
		}
		if req.Page != 0 {
			page = req.Page
		}
		offset = (page - 1) * limit
		params := &model.GetListRequest{
			Status:     req.Status,
			Search:     helper.SetPointerString(req.Search),
			RegencyID:  helper.SetPointerInt64(req.RegencyID),
			DistrictID: helper.SetPointerInt64(req.DistrictID),
			VillageID:  helper.SetPointerInt64(req.VillageID),
			Limit:      helper.SetPointerInt64(req.Limit),
			Offset:     helper.SetPointerInt64(offset),
			Longitude:  helper.SetPointerString(req.Longitude),
			Latitude:   helper.SetPointerString(req.Latitude),
			SortBy:     helper.SetPointerString(req.SortBy),
			OrderBy:    helper.SetPointerString(helper.AscOrDesc[req.OrderBy]),
			Name:       helper.SetPointerString(req.Name),
			Phone:      helper.SetPointerString(req.PhoneNumber),
		}

		resp, err := usecase.GetList(ctx, params)
		if err != nil {
			return nil, err
		}

		phonebooks := EncodePhonebook(resp.PhoneBooks)

		var meta *Metadata
		if resp.Metadata != nil {
			meta = &Metadata{
				TotalCount:  resp.Metadata.TotalCount,
				PageCount:   resp.Metadata.PageCount,
				CurrentPage: req.Page,
				PerPage:     resp.Metadata.PerPage}
		} else {
			meta = nil
		}

		return &PhoneBookWithMeta{
			Data: &PhonebookWithMeta{
				Phonebooks: phonebooks,
				Meta:       meta,
			},
		}, nil
	}
}

// MakeGetDetail ...
func MakeGetDetail(ctx context.Context, usecase usecase.Provider) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*GetDetailRequest)
		resp, err := usecase.GetDetail(ctx, req.ID)
		if err != nil {
			return nil, err
		}

		phoneNumbers := []*PhoneNumber{}
		if err := json.Unmarshal([]byte(resp.PhoneNumbers), &phoneNumbers); err != nil {
			return nil, err
		}

		category, ok := resp.Category.(*model.Category)

		data := &PhonebookDetail{
			ID:             resp.ID,
			Name:           resp.Name,
			Category:       resp.Category,
			Address:        resp.Address,
			Description:    resp.Description,
			PhoneNumbers:   phoneNumbers,
			Regency:        resp.Regency,
			District:       resp.District,
			Village:        resp.Village,
			Latitude:       resp.Latitude,
			Longitude:      resp.Longitude,
			CoverImagePath: fmt.Sprintf("%s/%s", cfg.AppStoragePublicURL, resp.CoverImageURL),
			Sequence:       resp.Sequence,
			Status:         resp.Status,
			StatusLabel:    GetStatusLabel[resp.Status]["id"],
			CreatedAt:      resp.CreatedAt,
			UpdatedAt:      resp.UpdatedAt,
		}
		if ok {
			data.CategoryID = category.ID
		}
		if resp.Regency != nil {
			data.RegencyID = helper.SetPointerInt64(resp.Regency.ID)
		}
		if resp.District != nil {
			data.DistrictID = helper.SetPointerInt64(resp.District.ID)
		}
		if resp.Village != nil {
			data.VillageID = helper.SetPointerInt64(resp.Village.ID)
		}
		if len(phoneNumbers) == 0 {
			data.PhoneNumbers = nil
		}
		return map[string]interface{}{
			"data": data,
		}, nil
	}
}

// MakeAddPhonebook ...
func MakeAddPhonebook(ctx context.Context, usecase usecase.Provider) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*AddPhonebookRequest)
		phoneNumbers, _ := json.Marshal(req.PhoneNumbers)
		if err := usecase.Insert(ctx, &model.AddPhonebook{
			Name:           req.Name,
			PhoneNumbers:   helper.SetPointerString(string(phoneNumbers)),
			Address:        req.Address,
			Description:    req.Description,
			RegencyID:      req.RegencyID,
			DistrictID:     req.DistrictID,
			VillageID:      req.VillageID,
			CategoryID:     req.CategoryID,
			Latitude:       req.Latitude,
			Longitude:      req.Longitude,
			CoverImagePath: req.CoverImagePath,
			Status:         req.Status,
			Sequence:       req.Sequence,
		}); err != nil {
			return nil, err
		}

		return &StatusResponse{
			Code:    helper.STATUS_CREATED,
			Message: "phonebook_has_been_created_succesfully",
		}, nil

	}
}

// MakeUpdatePhonebook ...
func MakeUpdatePhonebook(ctx context.Context, usecase usecase.Provider) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*UpdatePhonebookRequest)
		var phoneNumbers []byte
		if len(req.PhoneNumbers) > 0 {
			phoneNumbers, _ = json.Marshal(req.PhoneNumbers)
		}
		if err := usecase.Update(ctx, &model.UpdatePhonebook{
			ID:             req.ID,
			Name:           req.Name,
			PhoneNumbers:   helper.SetPointerString(string(phoneNumbers)),
			Address:        req.Address,
			Description:    req.Description,
			RegencyID:      req.RegencyID,
			DistrictID:     req.DistrictID,
			VillageID:      req.VillageID,
			CategoryID:     req.CategoryID,
			Latitude:       req.Latitude,
			Longitude:      req.Longitude,
			CoverImagePath: req.CoverImagePath,
			Status:         req.Status,
			Sequence:       req.Sequence,
		}); err != nil {
			return nil, err
		}

		return &StatusResponse{
			Code:    helper.STATUS_UPDATED,
			Message: "phonebook_has_been_updated_successfuly",
		}, nil
	}
}

// MakeDeletePhonebook ...
func MakeDeletePhonebook(ctx context.Context, usecase usecase.Provider) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*GetDetailRequest)
		if err := usecase.Delete(ctx, req.ID); err != nil {
			return nil, err
		}
		return &StatusResponse{
			Code:    helper.STATUS_DELETED,
			Message: "phonebook_has_been_deleted_successfuly",
		}, nil
	}
}

func MakeCheckHealthy(ctx context.Context) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return &StatusResponse{
			Code:    helper.STATUS_OK,
			Message: "service_is_ok",
		}, nil
	}
}

func MakeCheckReadiness(ctx context.Context, usecase usecase.Provider) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		if err := usecase.CheckHealthReadiness(ctx); err != nil {
			return nil, err
		}
		return &StatusResponse{
			Code:    helper.STATUS_OK,
			Message: "service_is_ready",
		}, nil
	}
}

// MakeIsExistPhoneNumber ...
func MakeIsExistPhoneNumber(ctx context.Context, usecase usecase.Provider) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*IsExistPhoneNumber)

		isExist, err := usecase.IsExistPhoneNumber(ctx, req.PhoneNumber)
		if err != nil {
			return nil, err
		}

		return map[string]map[string]interface{}{
			"data": {"exist": isExist},
		}, nil
	}
}
