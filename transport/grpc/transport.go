package grpc

import (
	"context"

	"github.com/sapawarga/phonebook-service/endpoint"
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

	return &grpcServer{
		phonebookGetListHandler,
		phonebookGetDetailHandler,
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
			Id:          v.ID,
			PhoneNumber: v.PhoneNumbers,
			Description: v.Description,
			Name:        v.Name,
			Address:     v.Address,
			Latitude:    v.Latitude,
			Longitude:   v.Longitude,
			Status:      v.Status,
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
		Id:          resp.ID,
		PhoneNumber: resp.PhoneNumbers,
		Description: resp.Description,
		Name:        resp.Name,
		Address:     resp.Address,
		Latitude:    resp.Latitude,
		Longitude:   resp.Longitude,
		Status:      resp.Status,
		CreatedAt:   resp.CreatedAt.String(),
		UpdatedAt:   resp.UpdatedAt.String(),
		Category:    resp.CategoryName,
		District:    resp.DistrictName,
		Village:     resp.VillageName,
	}, nil
}
