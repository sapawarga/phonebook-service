package usecase

import (
	"context"

	"github.com/sapawarga/phonebook-service/helper"
	"github.com/sapawarga/phonebook-service/model"
	"github.com/sapawarga/phonebook-service/repository"

	kitlog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

// PhoneBook ...
type PhoneBook struct {
	repo   repository.PhoneBookI
	logger kitlog.Logger
}

// NewPhoneBook ...
func NewPhoneBook(repo repository.PhoneBookI, logger kitlog.Logger) *PhoneBook {
	return &PhoneBook{
		repo:   repo,
		logger: logger,
	}
}

// GetList ...
func (pb *PhoneBook) GetList(ctx context.Context, params *model.ParamsPhoneBook) (*model.PhoneBookWithMeta, error) {
	logger := kitlog.With(pb.logger, "method", "GetList")
	req := &model.GetListRequest{
		Name:        helper.SetPointerString(params.Name),
		PhoneNumber: helper.SetPointerString(params.PhoneNumber),
		RegencyID:   helper.SetPointerInt64(params.RegencyID),
		DistrictID:  helper.SetPointerInt64(params.DistrictID),
		VillageID:   helper.SetPointerInt64(params.VillageID),
		Limit:       helper.SetPointerInt64(params.Limit),
		Offset:      helper.SetPointerInt64((params.Page - 1) * 10),
	}

	resp, err := pb.repo.GetListPhoneBook(ctx, req)
	if err != nil {
		level.Error(logger).Log("error", err)
		return nil, err
	}

	data := make([]*model.Phonebook, 0)

	for _, v := range resp {
		result := &model.Phonebook{
			ID:             v.ID,
			PhoneNumber:    v.PhoneNumber,
			Description:    v.Description.String,
			Name:           v.Name.String,
			Address:        v.Address.String,
			RegencyID:      v.DistrictID.Int64,
			DistrictID:     v.DistrictID.Int64,
			VillageID:      v.VillageID.Int64,
			Latitude:       v.Latitude.String,
			Longitude:      v.Longitude.String,
			CoverImagePath: v.CoverImagePath.String,
			Status:         v.Status.Int64,
			CreatedAt:      v.CreatedAt.Time,
			UpdatedAt:      v.UpdatedAt.Time,
			CategoryID:     v.CategoryID.Int64,
		}
		data = append(data, result)
	}

	total, err := pb.repo.GetMetaDataPhoneBook(ctx, req)
	if err != nil {
		level.Error(logger).Log("error", err)
		return nil, err
	}

	return &model.PhoneBookWithMeta{
		PhoneBooks: data,
		Page:       params.Page,
		Total:      total,
	}, nil
}
