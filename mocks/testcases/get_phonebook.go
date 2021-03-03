package testcases

import "github.com/sapawarga/phonebook-service/model"

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
			Name: "kantor",
		},
		GetListParams: model.GetListRequest{
			Name:  "kantor",
			Limit: 10,
			Page:  1,
		},
		GetMetaDataParams: model.GetListRequest{
			Name:  "kantor",
			Limit: 10,
			Page:  1,
		},
		MockUsecase: ResponseFromUsecase{
			Result: &model.PhoneBookWithMeta{
				PhoneBooks: []*model.PhoneBookResponse{
					{
						ID:          1,
						Name:        "kantor",
						PhoneNumber: "022123",
						Description: "kantor cabang McD",
					},
					{
						ID:          2,
						Name:        "kantor",
						PhoneNumber: "14225",
						Description: "kantor makanan",
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
					ID:          1,
					Name:        "kantor McD",
					PhoneNumber: "022123",
					Description: "kantor cabang McD",
				},
				{
					ID:          2,
					Name:        "kantor makanan",
					PhoneNumber: "14225",
					Description: "kantor makanan",
				},
			},
			Error: nil,
		},
	}, {
		Description: "success when get nil data",
		UsecaseParams: model.ParamsPhoneBook{
			Name: "random name",
		},
		GetListParams: model.GetListRequest{
			Name:  "random name",
			Limit: 10,
			Page:  1,
		},
		GetMetaDataParams: model.GetListRequest{
			Name:  "random name",
			Limit: 10,
			Page:  1,
		},
		MockUsecase: ResponseFromUsecase{
			Result: &model.PhoneBookWithMeta{
				PhoneBooks: []*model.PhoneBookResponse{},
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
