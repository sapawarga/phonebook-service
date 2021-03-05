package grpc

import (
	"context"

	"github.com/sapawarga/phonebook-service/endpoint"
	transportPhonebook "github.com/sapawarga/phonebook-service/transport/grpc/phonebook"
	"github.com/sapawarga/phonebook-service/usecase"

	kitgrpc "github.com/go-kit/kit/transport/grpc"
)

// MakeHandler ...
func MakeHandler(ctx context.Context, fs usecase.Provider) transportPhonebook.PhoneBookHandlerServer {
	phonebookGetListHandler := kitgrpc.NewServer(
		endpoint.MakeGetList(ctx, fs),
		decodeGetListRequest,
		encodeGetListResponse,
	)

	return &grpcServer{
		phonebookGetListHandler,
	}
}

func decodeGetListRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*transportPhonebook.GetListRequest)
	return &endpoint.GetListRequest{
		Name:        req.GetName(),
		PhoneNumber: req.GetPhoneNumber(),
		RegencyID:   req.GetRegencyId(),
		DistrictID:  req.GetDistrictId(),
		VillageID:   req.GetVillageId(),
		Status:      req.GetStatus(),
		Limit:       req.GetLimit(),
		Page:        req.GetPage(),
	}, nil
}

func encodeGetListResponse(ctx context.Context, r interface{}) (interface{}, error) {
	resp := r.(*endpoint.PhoneBookWithMeta)
	data := resp.Data
	meta := resp.Metadata

	resultData := make([]*transportPhonebook.PhoneBook, 0)
	for _, v := range data {
		result := &transportPhonebook.PhoneBook{
			Id:             v.ID,
			PhoneNumber:    v.PhoneNumber,
			Description:    v.Description,
			Name:           v.Name,
			Address:        v.Address,
			RegencyId:      v.RegencyID,
			DistrictId:     v.DistrictID,
			VillageId:      v.VillageID,
			Latitude:       v.Latitude,
			Longitude:      v.Longitude,
			CoverImagePath: v.CoverImagePath,
			Status:         v.Status,
			CreatedAt:      v.CreatedAt.String(),
			UpdatedAt:      v.UpdatedAt.String(),
			CategoryId:     v.CategoryID,
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
