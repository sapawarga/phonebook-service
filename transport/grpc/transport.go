package grpc

import (
	"context"
	"encoding/json"

	"github.com/sapawarga/phonebook-service/endpoint"
	"github.com/sapawarga/phonebook-service/helper"
	"github.com/sapawarga/phonebook-service/usecase"
	transportPhonebook "github.com/sapawarga/proto-file/phonebook"

	kitgrpc "github.com/go-kit/kit/transport/grpc"
)

// MakeHandler ...
func MakeHandler(ctx context.Context, fs usecase.Provider) transportPhonebook.PhoneBookHandlerServer {
	phonebookGetListHandler := kitgrpc.NewServer(
		endpoint.MakeGetList(ctx, fs),
		decodeGetListRequest,
		encodeGetListResponse,
	)

	phonebookGetDetailHandler := kitgrpc.NewServer(
		endpoint.MakeGetDetail(ctx, fs),
		decodeGetDetailRequest,
		encodeGetDetailResponse,
	)

	phonebookAddHandler := kitgrpc.NewServer(
		endpoint.MakeAddPhonebook(ctx, fs),
		decodeAddPhonebookRequest,
		encodeStatusResponse,
	)

	phonebookUpdateHandler := kitgrpc.NewServer(
		endpoint.MakeUpdatePhonebook(ctx, fs),
		decodeUpdatePhonebookRequest,
		encodeStatusResponse,
	)

	phonebookDeleteHandler := kitgrpc.NewServer(
		endpoint.MakeDeletePhonebook(ctx, fs),
		decodeGetDetailRequest,
		encodeStatusResponse,
	)

	return &grpcServer{
		phonebookGetListHandler,
		phonebookGetDetailHandler,
		phonebookAddHandler,
		phonebookUpdateHandler,
		phonebookDeleteHandler,
	}
}

func decodeGetListRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*transportPhonebook.GetListRequest)

	var status *int64 = nil
	if !req.Status.IsNull {
		status = helper.SetPointerInt64(req.Status.GetValue())
	}

	return &endpoint.GetListRequest{
		Search:     req.GetSearch(),
		RegencyID:  req.GetRegencyId(),
		DistrictID: req.GetDistrictId(),
		VillageID:  req.GetVillageId(),
		Status:     status,
		Limit:      req.GetLimit(),
		Page:       req.GetPage(),
	}, nil
}

func encodeGetListResponse(ctx context.Context, r interface{}) (interface{}, error) {
	resp := r.(*endpoint.PhoneBookWithMeta)
	data := resp.Data.Phonebooks
	meta := resp.Data.Meta

	resultData := make([]*transportPhonebook.PhoneBook, 0)
	for _, v := range data {
		phoneString, _ := json.Marshal(v.PhoneNumbers)

		result := &transportPhonebook.PhoneBook{
			Id:           v.ID,
			PhoneNumbers: string(phoneString),
			Description:  v.Description,
			Name:         v.Name,
			Address:      v.Address,
			Latitude:     v.Latitude,
			Longitude:    v.Longitude,
			Status:       v.Status,
		}
		resultData = append(resultData, result)
	}

	resultMeta := &transportPhonebook.Metadata{
		Page:  int64(meta.PageCount),
		Total: meta.TotalCount,
	}

	return &transportPhonebook.GetListResponse{
		Data:     resultData,
		Metadata: resultMeta,
	}, nil
}

func decodeGetDetailRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*transportPhonebook.GetDetailRequest)
	return &endpoint.GetDetailRequest{
		ID: req.GetId(),
	}, nil
}

func encodeGetDetailResponse(ctx context.Context, r interface{}) (interface{}, error) {
	resp := r.(*endpoint.PhonebookDetail)
	phoneString, _ := json.Marshal(resp.PhoneNumbers)
	return &transportPhonebook.GetDetailResponse{
		Id:           resp.ID,
		PhoneNumbers: string(phoneString),
		Description:  resp.Description,
		Name:         resp.Name,
		Address:      resp.Address,
		Latitude:     resp.Latitude,
		Longitude:    resp.Longitude,
		Status:       resp.Status,
		CreatedAt:    resp.CreatedAt.String(),
		UpdatedAt:    resp.UpdatedAt.String(),
		CategoryId:   resp.Category.ID,
		CategoryName: resp.Category.Name,
		RegencyId:    resp.Regency.ID,
		RegencyName:  resp.Regency.Name,
		DistrictId:   resp.District.ID,
		DistrictName: resp.District.Name,
		VillageId:    resp.Village.ID,
		VillageName:  resp.Village.Name,
	}, nil
}

func decodeAddPhonebookRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*transportPhonebook.AddPhonebookRequest)
	phoneNumbers := []*endpoint.PhoneNumber{}
	err := json.Unmarshal([]byte(req.GetPhoneNumbers()), &phoneNumbers)
	if err != nil {
		return nil, err
	}
	request := &endpoint.AddPhonebookRequest{
		Name:         req.GetName(),
		PhoneNumbers: phoneNumbers,
	}
	if req.Address != "" {
		request.Address = helper.SetPointerString(req.GetAddress())
	}
	if req.CategoryId != 0 {
		request.CategoryID = helper.SetPointerInt64(req.GetCategoryId())
	}
	if req.CoverImagePath != "" {
		request.CoverImagePath = helper.SetPointerString(req.GetCoverImagePath())
	}
	if req.Description != "" {
		request.Description = helper.SetPointerString(req.GetDescription())
	}
	if req.DistrictId != 0 {
		request.DistrictID = helper.SetPointerInt64(req.GetDistrictId())
	}
	if req.Latitude != "" {
		request.Latitude = helper.SetPointerString(req.GetLatitude())
	}
	if req.Longitude != "" {
		request.Longitude = helper.SetPointerString(req.GetLongitude())
	}
	if req.Status != 0 {
		request.Status = helper.SetPointerInt64(req.GetStatus())
	}
	return request, nil
}

func encodeStatusResponse(ctx context.Context, r interface{}) (interface{}, error) {
	resp := r.(*endpoint.StatusResponse)
	return &transportPhonebook.StatusResponse{
		Code:    resp.Code,
		Message: resp.Message,
	}, nil
}

func decodeUpdatePhonebookRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*transportPhonebook.UpdatePhonebookRequest)
	request := &endpoint.UpdatePhonebookRequest{
		ID:   req.GetId(),
		Name: req.GetName(),
	}
	if req.PhoneNumbers != "" {
		phoneNumbers := []*endpoint.PhoneNumber{}
		err := json.Unmarshal([]byte(req.GetPhoneNumbers()), &phoneNumbers)
		if err != nil {
			return nil, err
		}
		request.PhoneNumbers = phoneNumbers
	}
	if req.Address != "" {
		request.Address = helper.SetPointerString(req.GetAddress())
	}
	if req.CategoryId != 0 {
		request.CategoryID = helper.SetPointerInt64(req.GetCategoryId())
	}
	if req.CoverImagePath != "" {
		request.CoverImagePath = helper.SetPointerString(req.GetCoverImagePath())
	}
	if req.Description != "" {
		request.Description = helper.SetPointerString(req.GetDescription())
	}
	if req.RegencyId != 0 {
		request.RegencyID = helper.SetPointerInt64(req.GetRegencyId())
	}
	if req.DistrictId != 0 {
		request.DistrictID = helper.SetPointerInt64(req.GetDistrictId())
	}
	if req.VillageId != 0 {
		request.VillageID = helper.SetPointerInt64(req.GetVillageId())
	}
	if req.Latitude != "" {
		request.Latitude = helper.SetPointerString(req.GetLatitude())
	}
	if req.Longitude != "" {
		request.Longitude = helper.SetPointerString(req.GetLongitude())
	}
	if !req.Status.IsNull {
		request.Status = helper.SetPointerInt64(req.Status.GetValue())
	}
	return request, nil
}
