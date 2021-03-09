package endpoint

import (
	"context"

	"github.com/sapawarga/phonebook-service/model"
	"github.com/sapawarga/phonebook-service/usecase"

	"github.com/go-kit/kit/endpoint"
)

// MakeGetList ...
func MakeGetList(ctx context.Context, usecase usecase.Provider) endpoint.Endpoint {
	return func(ctx context.Context, params interface{}) (interface{}, error) {
		req := params.(*GetListRequest)
		resp, err := usecase.GetList(ctx, &model.ParamsPhoneBook{
			Search:     req.Search,
			RegencyID:  req.RegencyID,
			DistrictID: req.DistrictID,
			VillageID:  req.VillageID,
			Status:     req.Status,
			Limit:      req.Limit,
			Page:       req.Page,
		})

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
	return func(ctx context.Context, params interface{}) (interface{}, error) {
		req := params.(*GetDetailRequest)
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
