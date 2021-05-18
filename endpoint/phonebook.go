package endpoint

import (
	"context"

	"github.com/sapawarga/phonebook-service/helper"
	"github.com/sapawarga/phonebook-service/model"
	"github.com/sapawarga/phonebook-service/usecase"

	"github.com/go-kit/kit/endpoint"
)

// MakeGetList ...
func MakeGetList(ctx context.Context, usecase usecase.Provider) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*GetListRequest)
		params := &model.ParamsPhoneBook{
			Status:     req.Status,
			Search:     helper.SetPointerString(req.Search),
			RegencyID:  helper.SetPointerInt64(req.RegencyID),
			DistrictID: helper.SetPointerInt64(req.DistrictID),
			VillageID:  helper.SetPointerInt64(req.VillageID),
			Limit:      helper.SetPointerInt64(req.Limit),
			Page:       helper.SetPointerInt64(req.Page),
		}

		resp, err := usecase.GetList(ctx, params)

		if err != nil {
			return nil, err
		}

		meta := &Metadata{
			Page:  resp.Page,
			Total: resp.Total,
		}

		return &PhoneBookWithMeta{
			Data:     resp.PhoneBooks,
			Metadata: meta,
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

		return &PhonebookDetail{
			ID:             resp.ID,
			Name:           resp.Name,
			CategoryID:     resp.CategoryID,
			CategoryName:   resp.CategoryName,
			Address:        resp.Address,
			Description:    resp.Description,
			PhoneNumbers:   resp.PhoneNumbers,
			RegencyID:      resp.RegencyID,
			RegencyName:    resp.RegencyName,
			DistrictID:     resp.DistrictID,
			DistrictName:   resp.DistrictName,
			VillageID:      resp.VillageID,
			VillageName:    resp.VillageName,
			Latitude:       resp.Latitude,
			Longitude:      resp.Longitude,
			CoverImagePath: resp.CoverImagePath,
			Status:         resp.Status,
			CreatedAt:      resp.CreatedAt,
			UpdatedAt:      resp.UpdatedAt,
		}, nil
	}
}

// MakeAddPhonebook ...
func MakeAddPhonebook(ctx context.Context, usecase usecase.Provider) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*AddPhonebookRequest)
		if err := usecase.Insert(ctx, &model.AddPhonebook{
			Name:           req.Name,
			PhoneNumbers:   req.PhoneNumbers,
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
		if err := usecase.Update(ctx, &model.UpdatePhonebook{
			ID:             req.ID,
			Name:           req.Name,
			PhoneNumbers:   req.PhoneNumbers,
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
