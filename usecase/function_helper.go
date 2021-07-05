package usecase

import (
	"context"
	"database/sql"
	"fmt"
	"math"

	"github.com/sapawarga/phonebook-service/model"

	kitlog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

func (pb *PhoneBook) getPhonebookAndMetadata(ctx context.Context, params *model.GetListRequest, logger kitlog.Logger) (*model.PhoneBookWithMeta, error) {
	var resp []*model.PhoneBookResponse
	var err error
	var total int64

	if params.Longitude != nil && params.Latitude != nil {
		resp, err = pb.repo.GetListPhonebookByLongLat(ctx, params)
	} else {
		resp, err = pb.repo.GetListPhoneBook(ctx, params)
	}
	if err != nil {
		level.Error(logger).Log("error", err)
		return nil, err
	}
	data, err := pb.appendResultGetList(ctx, resp)
	if err != nil {
		level.Error(logger).Log("error_append_result", err)
		return nil, err
	}
	if params.Longitude != nil && params.Latitude != nil {
		total, err = pb.repo.GetListPhonebookByLongLatMeta(ctx, params)
	} else {
		total, err = pb.repo.GetMetaDataPhoneBook(ctx, params)
	}
	if err != nil {
		level.Error(logger).Log("error", err)
		return nil, err
	}
	return &model.PhoneBookWithMeta{
		PhoneBooks: data,
		Metadata: &model.Metadata{
			TotalCount: total,
			PageCount:  math.Ceil(float64(total) / float64(*params.Limit)),
			PerPage:    *params.Offset,
		},
	}, nil
}

func (pb *PhoneBook) appendResultGetList(ctx context.Context, result []*model.PhoneBookResponse) (listPhonebook []*model.Phonebook, err error) {
	if len(result) == 0 {
		return listPhonebook, nil
	}
	for _, v := range result {
		result := &model.Phonebook{
			ID:            v.ID,
			PhoneNumbers:  v.PhoneNumbers.String,
			Description:   v.Description.String,
			Name:          v.Name.String,
			Address:       v.Address.String,
			Latitude:      v.Latitude.String,
			Longitude:     v.Longitude.String,
			CoverImageURL: fmt.Sprintf("%s/%s", cfg.AppStoragePublicURL, v.CoverImagePath.String),
			Status:        v.Status.Int64,
			RegencyID:     v.RegencyID.Int64,
			DistrictID:    v.DistrictID.Int64,
			VillageID:     v.VillageID.Int64,
			CreatedAt:     v.CreatedAt.Time,
			UpdatedAt:     v.UpdatedAt.Time,
		}
		if v.CategoryID.Valid {
			categoryName, err := pb.repo.GetCategoryNameByID(ctx, v.CategoryID.Int64)
			if err != nil && err != sql.ErrNoRows {
				return nil, err
			}
			result.Category = &model.Category{ID: v.CategoryID.Int64, Name: categoryName}
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
		respDetail.Category = &model.Category{ID: respFromRepo.CategoryID.Int64, Name: categoryName}
	}
	if respFromRepo.RegencyID.Valid {
		regency, err := pb.repo.GetLocationByID(ctx, respFromRepo.RegencyID.Int64)
		if err != nil {
			return nil, err
		}
		respDetail.Regency = regency

	}
	if respFromRepo.DistrictID.Valid {
		district, err := pb.repo.GetLocationByID(ctx, respFromRepo.DistrictID.Int64)
		if err != nil {
			return nil, err
		}
		respDetail.District = district
	}
	if respFromRepo.VillageID.Valid {
		village, err := pb.repo.GetLocationByID(ctx, respFromRepo.VillageID.Int64)
		if err != nil {
			return nil, err
		}
		respDetail.Village = village
	}
	return respDetail, nil
}
