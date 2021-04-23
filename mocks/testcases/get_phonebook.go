package testcases

import (
	"database/sql"
	"errors"

	"github.com/sapawarga/phonebook-service/helper"
	"github.com/sapawarga/phonebook-service/model"
)

// ResponseFromUsecase ...
type ResponseFromUsecase struct {
	Result *model.PhoneBookWithMeta
	Error  error
}

// ResponseGetList ...
type ResponseGetList struct {
	Result []*model.PhoneBookResponse
	Error  error
}

// ResponseGetMetadata ...
type ResponseGetMetadata struct {
	Result int64
	Error  error
}

// GetPhoneBook ...
type GetPhoneBook struct {
	Description           string
	UsecaseParams         model.ParamsPhoneBook
	GetListParams         model.GetListRequest
	GetMetaDataParams     model.GetListRequest
	GetCategoryNameParams int64
	MockUsecase           ResponseFromUsecase
	MockGetList           ResponseGetList
	MockGetMetadata       ResponseGetMetadata
	MockCategorydata      CategoryResponse
}

// GetPhoneBookData ...
var GetPhoneBookData = []GetPhoneBook{
	{
		Description: "success get phone book",
		UsecaseParams: model.ParamsPhoneBook{
			Search: helper.SetPointerString("kantor"),
			Limit:  helper.SetPointerInt64(10),
			Page:   helper.SetPointerInt64(1),
		},
		GetListParams: model.GetListRequest{
			Search: helper.SetPointerString("kantor"),
			Limit:  helper.SetPointerInt64(10),
			Offset: helper.SetPointerInt64(0),
		},
		GetMetaDataParams: model.GetListRequest{
			Search: helper.SetPointerString("kantor"),
			Limit:  helper.SetPointerInt64(10),
			Offset: helper.SetPointerInt64(0),
		},
		GetCategoryNameParams: 1,
		MockCategorydata: CategoryResponse{
			Result: "category",
			Error:  nil,
		},
		MockUsecase: ResponseFromUsecase{
			Result: &model.PhoneBookWithMeta{
				PhoneBooks: []*model.Phonebook{
					{
						ID:           1,
						Name:         "kantor",
						PhoneNumbers: `[{"phone_number": "022123"}]`,
						Description:  "kantor cabang MCD",
						Category:     "category",
					},
					{
						ID:           2,
						Name:         "kantor",
						PhoneNumbers: `[{"phone_number": "423443"}]`,
						Description:  "kantor makanan",
						Category:     "category",
					},
				},
				Page:  1,
				Total: 2,
			},
			Error: nil,
		},
		MockGetMetadata: ResponseGetMetadata{
			Result: 2,
			Error:  nil,
		},
		MockGetList: ResponseGetList{
			Result: []*model.PhoneBookResponse{
				{
					ID:           1,
					Name:         sql.NullString{String: "kantor", Valid: true},
					PhoneNumbers: sql.NullString{String: `[{"type":"phone", "phone_number":"+62812312131"]`, Valid: true},
					Description:  sql.NullString{String: "kantor cabang MCD", Valid: true},
					CategoryID:   sql.NullInt64{Int64: 1, Valid: true},
				},
				{
					ID:           2,
					Name:         sql.NullString{String: "kantor", Valid: true},
					PhoneNumbers: sql.NullString{String: `[{"type":"phone", "phone_number":"+62812312131"]`, Valid: true},
					Description:  sql.NullString{String: "kantor makanan", Valid: true},
					CategoryID:   sql.NullInt64{Int64: 1, Valid: true},
				},
			},
			Error: nil,
		},
	}, {
		Description: "success when get nil data",
		UsecaseParams: model.ParamsPhoneBook{
			Search: helper.SetPointerString("random name"),
			Limit:  helper.SetPointerInt64(10),
			Page:   helper.SetPointerInt64(1),
		},
		GetListParams: model.GetListRequest{
			Search: helper.SetPointerString("random name"),
			Limit:  helper.SetPointerInt64(10),
			Offset: helper.SetPointerInt64(0),
		},
		GetMetaDataParams: model.GetListRequest{
			Search: helper.SetPointerString("random name"),
			Limit:  helper.SetPointerInt64(10),
			Offset: helper.SetPointerInt64(0),
		},
		GetCategoryNameParams: 1,
		MockCategorydata: CategoryResponse{
			Result: "category",
			Error:  nil,
		},
		MockUsecase: ResponseFromUsecase{
			Result: &model.PhoneBookWithMeta{
				PhoneBooks: nil,
				Page:       1,
				Total:      0,
			},
			Error: nil,
		},
		MockGetList: ResponseGetList{
			Result: nil,
			Error:  nil,
		},
		MockGetMetadata: ResponseGetMetadata{
			Result: 0,
			Error:  nil,
		},
	}, {
		Description: "failed get phone book",
		UsecaseParams: model.ParamsPhoneBook{
			Search: helper.SetPointerString("kantor"),
			Limit:  helper.SetPointerInt64(10),
			Page:   helper.SetPointerInt64(1),
		},
		GetListParams: model.GetListRequest{
			Search: helper.SetPointerString("kantor"),
			Limit:  helper.SetPointerInt64(10),
			Offset: helper.SetPointerInt64(0),
		},
		GetMetaDataParams: model.GetListRequest{
			Search: helper.SetPointerString("kantor"),
			Limit:  helper.SetPointerInt64(10),
			Offset: helper.SetPointerInt64(0),
		},
		GetCategoryNameParams: 1,
		MockCategorydata: CategoryResponse{
			Result: "category",
			Error:  nil,
		},
		MockUsecase: ResponseFromUsecase{
			Result: nil,
			Error:  errors.New("failed_get_list_phonebook"),
		},
		MockGetMetadata: ResponseGetMetadata{
			Result: 2,
			Error:  nil,
		},
		MockGetList: ResponseGetList{
			Result: nil,
			Error:  errors.New("failed_get_list_phonebook"),
		},
	}, {
		Description: "failed get phone book metadata",
		UsecaseParams: model.ParamsPhoneBook{
			Search: helper.SetPointerString("kantor"),
			Limit:  helper.SetPointerInt64(10),
			Page:   helper.SetPointerInt64(1),
		},
		GetListParams: model.GetListRequest{
			Search: helper.SetPointerString("kantor"),
			Limit:  helper.SetPointerInt64(10),
			Offset: helper.SetPointerInt64(0),
		},
		GetMetaDataParams: model.GetListRequest{
			Search: helper.SetPointerString("kantor"),
			Limit:  helper.SetPointerInt64(10),
			Offset: helper.SetPointerInt64(0),
		},
		GetCategoryNameParams: 1,
		MockCategorydata: CategoryResponse{
			Result: "category",
			Error:  nil,
		},
		MockUsecase: ResponseFromUsecase{
			Result: nil,
			Error:  errors.New("failed_get_metadata"),
		},
		MockGetMetadata: ResponseGetMetadata{
			Result: 0,
			Error:  errors.New("failed_get_metadata"),
		},
		MockGetList: ResponseGetList{
			Result: []*model.PhoneBookResponse{
				{
					ID:           1,
					Name:         sql.NullString{String: "kantor", Valid: true},
					PhoneNumbers: sql.NullString{String: `[{"type":"phone", "phone_number":"+62812312131"]`, Valid: true},
					Description:  sql.NullString{String: "kantor cabang MCD", Valid: true},
					CategoryID:   sql.NullInt64{Int64: 1, Valid: true},
				},
				{
					ID:           2,
					Name:         sql.NullString{String: "kantor", Valid: true},
					PhoneNumbers: sql.NullString{String: `[{"type":"phone", "phone_number":"+62812312131"]`, Valid: true},
					Description:  sql.NullString{String: "kantor makanan", Valid: true},
					CategoryID:   sql.NullInt64{Int64: 1, Valid: true},
				},
			},
			Error: nil,
		},
	},
}

// ListPhonebookDescription :
func ListPhonebookDescription() []string {
	var arr = []string{}
	for _, data := range GetPhoneBookData {
		arr = append(arr, data.Description)
	}
	return arr
}
