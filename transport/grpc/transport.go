package grpc

import (
	"context"

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
		encodeAddPhonebookResponse,
	)

	return &grpcServer{
		phonebookGetListHandler,
		phonebookGetDetailHandler,
		phonebookAddHandler,
	}
}

func decodeGetListRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*transportPhonebook.GetListRequest)
	return &endpoint.GetListRequest{
		Search:     req.GetSearch(),
		RegencyID:  req.GetRegencyId(),
		DistrictID: req.GetDistrictId(),
		VillageID:  req.GetVillageId(),
		Status:     req.GetStatus(),
		Limit:      req.GetLimit(),
		Page:       req.GetPage(),
	}, nil
}

func encodeGetListResponse(ctx context.Context, r interface{}) (interface{}, error) {
	resp := r.(*endpoint.PhoneBookWithMeta)
	data := resp.Data
	meta := resp.Metadata

	resultData := make([]*transportPhonebook.PhoneBook, 0)
	for _, v := range data {
		result := &transportPhonebook.PhoneBook{
			Id:           v.ID,
			PhoneNumbers: v.PhoneNumbers,
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
		Page:  meta.Page,
		Total: meta.Total,
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
	return &transportPhonebook.GetDetailResponse{
		Id:           resp.ID,
		PhoneNumbers: resp.PhoneNumbers,
		Description:  resp.Description,
		Name:         resp.Name,
		Address:      resp.Address,
		Latitude:     resp.Latitude,
		Longitude:    resp.Longitude,
		Status:       resp.Status,
		CreatedAt:    resp.CreatedAt.String(),
		UpdatedAt:    resp.UpdatedAt.String(),
		CategoryId:   resp.CategoryID,
		CategoryName: resp.CategoryName,
		RegencyId:    resp.RegencyID,
		RegencyName:  resp.RegencyName,
		DistrictId:   resp.DistrictID,
		DistrictName: resp.DistrictName,
		VillageId:    resp.VillageID,
		VillageName:  resp.VillageName,
	}, nil
}

func decodeAddPhonebookRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*transportPhonebook.AddPhonebookRequest)
	request := &endpoint.AddPhonebookRequest{
		Name:         req.GetName(),
		PhoneNumbers: helper.SetPointerString(req.GetPhoneNumbers()),
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

func encodeAddPhonebookResponse(ctx context.Context, r interface{}) (interface{}, error) {
	resp := r.(*endpoint.StatusResponse)
	return &transportPhonebook.StatusResponse{
		Code:    resp.Code,
		Message: resp.Message,
	}, nil
}
