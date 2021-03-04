package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/sapawarga/phonebook-service/model"
	"github.com/sapawarga/phonebook-service/usecase"
)

// MakeGetList ...
func MakeGetList(ctx context.Context, usecase usecase.Provider) endpoint.Endpoint {
	return func(ctx context.Context, params interface{}) (interface{}, error) {
		req := params.(*GetListRequest)
		resp, err := usecase.GetList(ctx, &model.ParamsPhoneBook{
			Name:         req.Name,
			PhoneNumber:  req.PhoneNumber,
			RegencyCode:  req.RegencyCode,
			DistrictCode: req.DistrictCode,
			VillageCode:  req.VillageCode,
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
