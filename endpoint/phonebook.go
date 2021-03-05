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
			Name:        req.Name,
			PhoneNumber: req.PhoneNumber,
			RegencyID:   req.RegencyID,
			DistrictID:  req.DistrictID,
			VillageID:   req.VillageID,
			Status:      req.Status,
			Limit:       req.Limit,
			Page:        req.Page,
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
