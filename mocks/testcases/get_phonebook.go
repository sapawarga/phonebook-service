package testcases

import (
	"database/sql"

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
	Description       string
	UsecaseParams     model.ParamsPhoneBook
	GetListParams     model.GetListRequest
	GetMetaDataParams model.GetListRequest
	MockUsecase       ResponseFromUsecase
	MockGetList       ResponseGetList
	MockGetMetadata   ResponseGetMetadata
}

// GetPhoneBookData ...
var GetPhoneBookData = []GetPhoneBook{
	{
		Description: "success get phone book",
		UsecaseParams: model.ParamsPhoneBook{
			Search: "kantor",
			Limit:  10,
			Page:   1,
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
		MockUsecase: ResponseFromUsecase{
			Result: &model.PhoneBookWithMeta{
				PhoneBooks: []*model.Phonebook{
					{
						ID:           1,
						Name:         "kantor",
						PhoneNumbers: `[{"phone_number": "022123"}]`,
						Description:  "kantor cabang MCD",
					},
					{
						ID:           2,
						Name:         "kantor",
						PhoneNumbers: `[{"phone_number": "423443"}]`,
						Description:  "kantor makanan",
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
					PhoneNumbers: `[{"phone_number": "022123"}]`,
					Description:  sql.NullString{String: "kantor cabang MCD", Valid: true},
				},
				{
					ID:           2,
					Name:         sql.NullString{String: "kantor", Valid: true},
					PhoneNumbers: `[{"phone_number": "423443"}]`,
					Description:  sql.NullString{String: "kantor makanan", Valid: true},
				},
			},
			Error: nil,
		},
	}, {
		Description: "success when get nil data",
		UsecaseParams: model.ParamsPhoneBook{
			Search: "random name",
			Limit:  10,
			Page:   1,
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
	},
}

// Description :
func Description() []string {
	var arr = []string{}
	for _, data := range GetPhoneBookData {
		arr = append(arr, data.Description)
	}
	return arr
}
