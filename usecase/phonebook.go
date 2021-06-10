package usecase

import (
	"context"
	"errors"
	"math"

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
	var limit, page, offset int64 = 10, 1, 0
	if params.Limit != nil {
		limit = helper.GetInt64FromPointer(params.Limit)
	}
	if params.Page != nil {
		page = helper.GetInt64FromPointer(params.Page)
	}
	offset = (page - 1) * limit
	req := &model.GetListRequest{
		Search:     params.Search,
		RegencyID:  params.RegencyID,
		DistrictID: params.DistrictID,
		VillageID:  params.VillageID,
		Status:     params.Status,
		Longitude:  params.Longitude,
		Latitude:   params.Latitude,
		Limit:      &limit,
		Offset:     &offset,
	}

	result, err := pb.getPhonebookAndMetadata(ctx, req, logger)
	if err != nil {
		level.Error(logger).Log("error", err)
		return nil, err
	}

	return &model.PhoneBookWithMeta{
		PhoneBooks: result.PhoneBooks,
		Page:       page,
		Total:      result.Total,
		TotalPage:  int64(math.Ceil(float64(result.Total/limit))) + 1}, nil
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
		PhoneNumbers:   resp.PhoneNumbers.String,
		Latitude:       resp.Latitude.String,
		Longitude:      resp.Longitude.String,
		CoverImagePath: resp.CoverImagePath.String,
		Status:         resp.Status.Int64,
		CreatedAt:      resp.CreatedAt.Time,
		UpdatedAt:      resp.UpdatedAt.Time,
	}

	result, err = pb.appendDetailPhonebook(ctx, resp, result)
	if err != nil {
		level.Error(logger).Log("error_add_detail", err)
		return nil, err
	}

	return result, nil
}

// Insert ...
func (pb *PhoneBook) Insert(ctx context.Context, params *model.AddPhonebook) error {
	logger := kitlog.With(pb.logger, "method", "Insert")
	if params.CategoryID != nil {
		_, err := pb.repo.GetCategoryNameByID(ctx, helper.GetInt64FromPointer(params.CategoryID))
		if err != nil {
			level.Error(logger).Log("error_get_category", err)
			return err
		}
	}

	if err := pb.repo.Insert(ctx, params); err != nil {
		level.Error(logger).Log("error", err)
		return err
	}
	return nil
}

// Update ...
func (pb *PhoneBook) Update(ctx context.Context, params *model.UpdatePhonebook) error {
	// TODO: update phonebook
	logger := kitlog.With(pb.logger, "method", "Update")
	if _, err := pb.repo.GetPhonebookDetailByID(ctx, params.ID); err != nil {
		level.Error(logger).Log("error_get_detail", err)
		return err
	}

	if err := pb.repo.Update(ctx, params); err != nil {
		level.Error(logger).Log("error_update", err)
		return err
	}

	return nil
}

// Delete ...
func (pb *PhoneBook) Delete(ctx context.Context, id int64) error {
	// TODO: delete phonebook
	logger := kitlog.With(pb.logger, "method", "Delete")
	if _, err := pb.repo.GetPhonebookDetailByID(ctx, id); err != nil {
		level.Error(logger).Log("error_get_detail", err)
		return err
	}

	if err := pb.repo.Delete(ctx, id); err != nil {
		level.Error(logger).Log("error_delete", err)
		return err
	}

	return nil
}

// CheckHealthReadiness ...
func (pb *PhoneBook) CheckHealthReadiness(ctx context.Context) error {
	logger := kitlog.With(pb.logger, "method", "CheckHealthReadiness")
	if err := pb.repo.CheckHealthReadiness(ctx); err != nil {
		level.Error(logger).Log("error", errors.New("service_not_ready"))
		return errors.New("service_not_ready")
	}
	return nil
}
