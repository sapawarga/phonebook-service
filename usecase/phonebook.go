package usecase

import (
	"context"
	"errors"

	"github.com/sapawarga/phonebook-service/config"
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

var cfg, _ = config.NewConfig()

// NewPhoneBook ...
func NewPhoneBook(repo repository.PhoneBookI, logger kitlog.Logger) *PhoneBook {
	return &PhoneBook{
		repo:   repo,
		logger: logger,
	}
}

// GetList ...
func (pb *PhoneBook) GetList(ctx context.Context, params *model.GetListRequest) (*model.PhoneBookWithMeta, error) {
	logger := kitlog.With(pb.logger, "method", "GetList")
	result, err := pb.getPhonebookAndMetadata(ctx, params, logger)
	if err != nil {
		level.Error(logger).Log("error", err)
		return nil, err
	}

	return &model.PhoneBookWithMeta{
		PhoneBooks: result.PhoneBooks,
		Metadata:   result.Metadata}, nil
}

// GetDetail ...
func (pb *PhoneBook) GetDetail(ctx context.Context, id int64) (*model.Phonebook, error) {
	logger := kitlog.With(pb.logger, "method", "GetDetail")
	resp, err := pb.repo.GetPhonebookDetailByID(ctx, id)
	if err != nil {
		level.Error(logger).Log("error", err)
		return nil, err
	}

	result := &model.Phonebook{
		ID:            resp.ID,
		Name:          resp.Name.String,
		Address:       resp.Address.String,
		Description:   resp.Description.String,
		PhoneNumbers:  resp.PhoneNumbers.String,
		Latitude:      resp.Latitude.String,
		Longitude:     resp.Longitude.String,
		CoverImageURL: resp.CoverImagePath.String,
		Status:        resp.Status.Int64,
		CreatedAt:     resp.CreatedAt.Int64,
		UpdatedAt:     resp.UpdatedAt.Int64,
		Sequence:      resp.Sequence.Int64,
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

// IsExistPhoneNumber ...
func (pb *PhoneBook) IsExistPhoneNumber(ctx context.Context, phone string) (bool, error) {
	logger := kitlog.With(pb.logger, "method", "IsExistPhoneNumber")
	isExist, err := pb.repo.IsExistPhoneNumber(ctx, phone)
	if err != nil {
		level.Error(logger).Log("error_is_exist", err)
		return false, err
	}

	return isExist, nil
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
