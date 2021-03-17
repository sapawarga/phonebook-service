package testcases

import (
	"errors"

	"github.com/sapawarga/phonebook-service/helper"
	"github.com/sapawarga/phonebook-service/model"
)

// InsertPhonebook ...
type InsertPhonebook struct {
	Description            string
	UsecaseRequest         model.AddPhonebook
	GetCategoryNameRequest int64
	RepositoryRequest      model.AddPhonebook
	MockCategory           CategoryResponse
	RepositoryResponse     error
	UsecaseResponse        error
}

var insertPhonebook = model.AddPhonebook{
	Name:           "kantor",
	PhoneNumbers:   helper.SetPointerString(`[{"type": "phone", "phone_number":"14045"}]`),
	Address:        helper.SetPointerString("jalan jalan"),
	Description:    helper.SetPointerString("test"),
	RegencyID:      nil,
	DistrictID:     nil,
	VillageID:      nil,
	Latitude:       helper.SetPointerString("-3.09893"),
	Longitude:      helper.SetPointerString("0.98878"),
	CoverImagePath: helper.SetPointerString("http://localhot:3000"),
	Status:         helper.SetPointerInt64(10),
	CategoryID:     helper.SetPointerInt64(1),
}

// InsertPhonebookTestcases ...
var InsertPhonebookTestcases = []InsertPhonebook{
	{
		Description:            "success insert new phonebook",
		UsecaseRequest:         insertPhonebook,
		RepositoryRequest:      insertPhonebook,
		GetCategoryNameRequest: 1,
		MockCategory: CategoryResponse{
			Result: "phonebook",
			Error:  nil,
		},
		RepositoryResponse: nil,
		UsecaseResponse:    nil,
	}, {
		Description:            "failed_get_category",
		UsecaseRequest:         insertPhonebook,
		RepositoryRequest:      insertPhonebook,
		GetCategoryNameRequest: 1,
		MockCategory: CategoryResponse{
			Result: "",
			Error:  errors.New("failed_get_category"),
		},
		RepositoryResponse: nil,
		UsecaseResponse:    errors.New("failed_get_category"),
	}, {
		Description:            "failed_insert_phonebook",
		UsecaseRequest:         insertPhonebook,
		RepositoryRequest:      insertPhonebook,
		GetCategoryNameRequest: 1,
		MockCategory: CategoryResponse{
			Result: "phonebook",
			Error:  nil,
		},
		RepositoryResponse: errors.New("failed_insert_phonebook"),
		UsecaseResponse:    errors.New("failed_insert_phonebook"),
	},
}

// InsertPhonebookDescription :
func InsertPhonebookDescription() []string {
	var arr = []string{}
	for _, data := range InsertPhonebookTestcases {
		arr = append(arr, data.Description)
	}
	return arr
}
