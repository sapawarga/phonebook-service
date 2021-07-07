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
	UsecaseParams         model.GetListRequest
	GetListParams         model.GetListRequest
	GetMetaDataParams     model.GetListRequest
	GetCategoryNameParams int64
	MockUsecase           ResponseFromUsecase
	MockGetList           ResponseGetList
	MockGetMetadata       ResponseGetMetadata
	MockCategorydata      CategoryResponse
}

var category = &model.Category{
	ID:   10,
	Name: "category",
}

var meta = &model.Metadata{
	TotalCount:  2,
	PageCount:   1,
	CurrentPage: 1,
	PerPage:     0,
}

// GetPhoneBookData ...
var GetPhoneBookData = []GetPhoneBook{
	{
		Description: "success get phone book",
		UsecaseParams: model.GetListRequest{
			Search: helper.SetPointerString("kantor"),
			Limit:  helper.SetPointerInt64(10),
			Offset: helper.SetPointerInt64(0),
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
						Category:     category,
					},
					{
						ID:           2,
						Name:         "kantor",
						PhoneNumbers: `[{"phone_number": "423443"}]`,
						Description:  "kantor makanan",
						Category:     category,
					},
				},
				Metadata: meta,
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
		Description: "success get phone book by long and lat",
		UsecaseParams: model.GetListRequest{
			Search:    helper.SetPointerString("kantor"),
			Limit:     helper.SetPointerInt64(10),
			Offset:    helper.SetPointerInt64(0),
			Longitude: helper.SetPointerString("-6.00009"),
			Latitude:  helper.SetPointerString("+90.00009"),
		},
		GetListParams: model.GetListRequest{
			Search:    helper.SetPointerString("kantor"),
			Limit:     helper.SetPointerInt64(10),
			Offset:    helper.SetPointerInt64(0),
			Longitude: helper.SetPointerString("-6.00009"),
			Latitude:  helper.SetPointerString("+90.00009"),
		},
		GetMetaDataParams: model.GetListRequest{
			Search:    helper.SetPointerString("kantor"),
			Limit:     helper.SetPointerInt64(10),
			Offset:    helper.SetPointerInt64(0),
			Longitude: helper.SetPointerString("-6.00009"),
			Latitude:  helper.SetPointerString("+90.00009"),
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
						Category:     category,
					},
					{
						ID:           2,
						Name:         "kantor",
						PhoneNumbers: `[{"phone_number": "423443"}]`,
						Description:  "kantor makanan",
						Category:     category,
						Distance:     40,
					},
				},
				Metadata: meta,
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
					Distance:     40,
				},
			},
			Error: nil,
		},
	}, {
		Description: "success when get nil data",
		UsecaseParams: model.GetListRequest{
			Search: helper.SetPointerString("random name"),
			Limit:  helper.SetPointerInt64(10),
			Offset: helper.SetPointerInt64(0),
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
				PhoneBooks: []*model.Phonebook{},
				Metadata: &model.Metadata{
					TotalCount:  0,
					PageCount:   0,
					CurrentPage: 0,
					PerPage:     0,
				},
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
		UsecaseParams: model.GetListRequest{
			Search: helper.SetPointerString("kantor"),
			Limit:  helper.SetPointerInt64(10),
			Offset: helper.SetPointerInt64(0),
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
		UsecaseParams: model.GetListRequest{
			Search: helper.SetPointerString("kantor"),
			Limit:  helper.SetPointerInt64(10),
			Offset: helper.SetPointerInt64(0),
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
