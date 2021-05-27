package usecase

import (
	"context"
	"database/sql"

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
		Limit:      &limit,
		Offset:     &offset,
	}

	resp, err := pb.repo.GetListPhoneBook(ctx, req)
	if err != nil {
		level.Error(logger).Log("error", err)
		return nil, err
	}

	data, err := pb.appendResultGetList(ctx, resp)
	if err != nil {
		level.Error(logger).Log("error_append_result", err)
		return nil, err
	}
	total, err := pb.repo.GetMetaDataPhoneBook(ctx, req)
	if err != nil {
		level.Error(logger).Log("error", err)
		return nil, err
	}

	return &model.PhoneBookWithMeta{
		PhoneBooks: data,
		Page:       page,
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

func (pb *PhoneBook) appendResultGetList(ctx context.Context, result []*model.PhoneBookResponse) (listPhonebook []*model.Phonebook, err error) {
	for _, v := range result {
		result := &model.Phonebook{
			ID:           v.ID,
			PhoneNumbers: v.PhoneNumbers.String,
			Description:  v.Description.String,
			Name:         v.Name.String,
			Address:      v.Address.String,
			Latitude:     v.Latitude.String,
			Longitude:    v.Longitude.String,
			Status:       v.Status.Int64,
		}
		if v.CategoryID.Valid {
			categoryName, err := pb.repo.GetCategoryNameByID(ctx, v.CategoryID.Int64)
			if err != nil && err != sql.ErrNoRows {
				return nil, err
			}
			result.Category = categoryName
		}
		listPhonebook = append(listPhonebook, result)
	}
	return listPhonebook, nil
}

func (pb *PhoneBook) appendDetailPhonebook(ctx context.Context, respFromRepo *model.PhoneBookResponse, respDetail *model.PhonebookDetail) (*model.PhonebookDetail, error) {
	if respFromRepo.CategoryID.Valid {
		categoryName, err := pb.repo.GetCategoryNameByID(ctx, respFromRepo.CategoryID.Int64)
		if err != nil {
			return nil, err
		}
		respDetail.CategoryID = respFromRepo.CategoryID.Int64
		respDetail.CategoryName = categoryName
	}
	if respFromRepo.RegencyID.Valid {
		regencyName, err := pb.repo.GetLocationNameByID(ctx, respFromRepo.RegencyID.Int64)
		if err != nil {
			return nil, err
		}
		respDetail.RegencyID = respFromRepo.RegencyID.Int64
		respDetail.RegencyName = regencyName
	}
	if respFromRepo.DistrictID.Valid {
		districtName, err := pb.repo.GetLocationNameByID(ctx, respFromRepo.DistrictID.Int64)
		if err != nil {
			return nil, err
		}
		respDetail.DistrictID = respFromRepo.DistrictID.Int64
		respDetail.DistrictName = districtName
	}
	if respFromRepo.VillageID.Valid {
		villageName, err := pb.repo.GetLocationNameByID(ctx, respFromRepo.VillageID.Int64)
		if err != nil {
			return nil, err
		}
		respDetail.VillageID = respFromRepo.VillageID.Int64
		respDetail.VillageName = villageName
	}
	return respDetail, nil
}
