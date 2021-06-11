package usecase

import (
	"context"
	"database/sql"

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
		Total:      total,
	}, nil
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
