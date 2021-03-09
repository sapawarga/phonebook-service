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
		Search:     helper.SetPointerString(params.Search),
		RegencyID:  helper.SetPointerInt64(params.RegencyID),
		DistrictID: helper.SetPointerInt64(params.DistrictID),
		VillageID:  helper.SetPointerInt64(params.VillageID),
		Limit:      helper.SetPointerInt64(params.Limit),
		Offset:     helper.SetPointerInt64((params.Page - 1) * 10),
	}

	resp, err := pb.repo.GetListPhoneBook(ctx, req)
	if err != nil {
		level.Error(logger).Log("error", err)
		return nil, err
	}

	data := make([]*model.Phonebook, 0)

	for _, v := range resp {
		result := &model.Phonebook{
			ID:           v.ID,
			PhoneNumbers: v.PhoneNumbers,
			Description:  v.Description.String,
			Name:         v.Name.String,
			Address:      v.Address.String,
			Latitude:     v.Latitude.String,
			Longitude:    v.Longitude.String,
			Status:       v.Status.Int64,
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

// GetDetail ...
func (pb *PhoneBook) GetDetail(ctx context.Context, id int64) (*model.PhonebookDetail, error) {
	logger := kitlog.With(pb.logger, "method", "GetDetail")

	resp, err := pb.repo.GetPhonebookDetailByID(ctx, id)
	if err != nil {
		level.Error(logger).Log("error", err)
		return nil, err
	}

	result := &model.PhonebookDetail{
		ID:             resp.ID,
		Name:           resp.Name.String,
		Address:        resp.Address.String,
		Description:    resp.Description.String,
		PhoneNumbers:   resp.PhoneNumbers,
		Latitude:       resp.Latitude.String,
		Longitude:      resp.Longitude.String,
		CoverImagePath: resp.CoverImagePath.String,
		Status:         resp.Status.Int64,
		CreatedAt:      resp.CreatedAt.Time,
		UpdatedAt:      resp.UpdatedAt.Time,
	}

	if resp.CategoryID.Valid {
		categoryName, err := pb.repo.GetCategoryNameByID(ctx, resp.CategoryID.Int64)
		if err != nil {
			level.Error(logger).Log("error_get_category", err)
			return nil, err
		}
		result.CategoryID = resp.CategoryID.Int64
		result.CategoryName = categoryName
	}

	if resp.RegencyID.Valid {
		regencyName, err := pb.repo.GetLocationNameByID(ctx, resp.RegencyID.Int64)
		if err != nil {
			level.Error(logger).Log("error_get_regency", err)
			return nil, err
		}
		result.RegencyID = resp.RegencyID.Int64
		result.RegencyName = regencyName
	}

	if resp.DistrictID.Valid {
		districtName, err := pb.repo.GetLocationNameByID(ctx, resp.DistrictID.Int64)
		if err != nil {
			level.Error(logger).Log("error_get_district", err)
			return nil, err
		}
		result.DistrictID = resp.DistrictID.Int64
		result.DistrictName = districtName
	}

	if resp.VillageID.Valid {
		villageName, err := pb.repo.GetLocationNameByID(ctx, resp.VillageID.Int64)
		if err != nil {
			level.Error(logger).Log("error_get_village", err)
			return nil, err
		}
		result.VillageID = resp.VillageID.Int64
		result.VillageName = villageName
	}
	return result, nil
}

// Insert ...
func (pb *PhoneBook) Insert(ctx context.Context, params *model.AddPhonebook) error {
	// TODO: insert new phone number
	logger := kitlog.With(pb.logger, "method", "Insert")
	err := pb.repo.Insert(ctx, params)
	if err != nil {
		level.Error(logger).Log("error", err)
		return err
	}
	return nil
}
