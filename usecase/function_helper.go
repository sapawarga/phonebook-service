package usecase

import (
	"context"
	"fmt"
	"math"
	"strconv"

	"github.com/sapawarga/phonebook-service/helper"
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
		level.Error(logger).Log("error_get_list", err)
		return nil, err
	}
	data, err := pb.appendResultGetList(ctx, params, resp)
	if err != nil {
		level.Error(logger).Log("error_append_result", err)
		return nil, err
	}
	meta := &model.Metadata{}
	if params.Longitude == nil && params.Latitude == nil {
		total, err = pb.repo.GetMetaDataPhoneBook(ctx, params)
	}

	if err != nil {
		level.Error(logger).Log("error", err)
		return nil, err
	}

	if meta != nil {
		meta.TotalCount = total
		meta.PageCount = math.Ceil(float64(total) / float64(*params.Limit))
		meta.PerPage = helper.GetInt64FromPointer(params.Limit)
	} else {
		meta = nil
	}
	return &model.PhoneBookWithMeta{
		PhoneBooks: data,
		Metadata:   meta,
	}, nil
}

func (pb *PhoneBook) appendResultGetList(ctx context.Context, params *model.GetListRequest, result []*model.PhoneBookResponse) (listPhonebook []*model.Phonebook, err error) {
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
			CoverImageURL: helper.SetPointerString(v.CoverImagePath.String),
			Status:        v.Status.Int64,
			Sequence:      v.Sequence.Int64,
			CreatedAt:     v.CreatedAt.Int64,
			UpdatedAt:     v.UpdatedAt.Int64,
			Distance:      v.Distance,
		}

		resAppend, err := pb.appendDetailPhonebook(ctx, v, result)
		if err != nil {
			return nil, err
		}

		if params.Latitude != nil && params.Longitude != nil {
			id := resAppend.ID.(int64)
			distance := resAppend.Distance.(float64)
			category := resAppend.Category.(*model.Category)
			resAppend.Category = category.Name
			resAppend.ID = strconv.FormatInt(id, 10)
			resAppend.Distance = fmt.Sprintf("%f", distance)
		}

		listPhonebook = append(listPhonebook, resAppend)
	}

	return listPhonebook, nil
}

func (pb *PhoneBook) appendDetailPhonebook(ctx context.Context, respFromRepo *model.PhoneBookResponse, respDetail *model.Phonebook) (*model.Phonebook, error) {
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
