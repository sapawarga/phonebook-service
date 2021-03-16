package testcases

import (
	"database/sql"
	"errors"

	"github.com/sapawarga/phonebook-service/helper"
	"github.com/sapawarga/phonebook-service/model"
)

// UpdatePhonebook ...
type UpdatePhonebook struct {
	Description                string
	UsecaseRequest             model.UpdatePhonebook
	UpdateRepositoryRequest    model.UpdatePhonebook
	GetDetailRepositoryRequest int64
	MockDetailRepository       GetDetailResponseRepository
	MockUpdateRepository       error
	MockUsecase                error
}

var updatePhonebook = model.UpdatePhonebook{
	ID:             1,
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
	CategoryID:     nil,
}

// UpdatePhonebookTestcases ...
var UpdatePhonebookTestcases = []UpdatePhonebook{
	{
		Description:                "success_update_phonebook",
		UsecaseRequest:             updatePhonebook,
		UpdateRepositoryRequest:    updatePhonebook,
		GetDetailRepositoryRequest: 1,
		MockDetailRepository: GetDetailResponseRepository{
			Result: &model.PhoneBookResponse{
				ID:             1,
				PhoneNumbers:   sql.NullString{String: `[{"type":"phone", "phone_number":"+62812312131"]`, Valid: true},
				Description:    sql.NullString{String: "test case", Valid: true},
				Name:           sql.NullString{String: "test kantor", Valid: true},
				Address:        sql.NullString{String: "jalan panjang", Valid: true},
				RegencyID:      sql.NullInt64{Int64: 1, Valid: true},
				DistrictID:     sql.NullInt64{Int64: 10, Valid: true},
				VillageID:      sql.NullInt64{Int64: 100, Valid: true},
				Latitude:       sql.NullString{String: "-6.231928", Valid: true},
				Longitude:      sql.NullString{String: "0.988789", Valid: true},
				CoverImagePath: sql.NullString{String: "http://localhost:9080", Valid: true},
				Status:         sql.NullInt64{Int64: 10, Valid: true},
				CreatedAt:      sql.NullTime{Time: currentTime, Valid: true},
				UpdatedAt:      sql.NullTime{Time: currentTime, Valid: true},
				CategoryID:     sql.NullInt64{Int64: 1, Valid: true},
			},
			Error: nil,
		},
		MockUpdateRepository: nil,
		MockUsecase:          nil,
	}, {
		Description:                "failed_get_detail_phonebook",
		UsecaseRequest:             updatePhonebook,
		UpdateRepositoryRequest:    updatePhonebook,
		GetDetailRepositoryRequest: 1,
		MockDetailRepository: GetDetailResponseRepository{
			Result: nil,
			Error:  errors.New("invalid_id_detail"),
		},
		MockUpdateRepository: nil,
		MockUsecase:          errors.New("invalid_id_detail"),
	}, {
		Description:                "failed_update_phonebook",
		UsecaseRequest:             updatePhonebook,
		UpdateRepositoryRequest:    updatePhonebook,
		GetDetailRepositoryRequest: 1,
		MockDetailRepository: GetDetailResponseRepository{
			Result: &model.PhoneBookResponse{
				ID:             1,
				PhoneNumbers:   sql.NullString{String: `[{"type":"phone", "phone_number":"+62812312131"]`, Valid: true},
				Description:    sql.NullString{String: "test case", Valid: true},
				Name:           sql.NullString{String: "test kantor", Valid: true},
				Address:        sql.NullString{String: "jalan panjang", Valid: true},
				RegencyID:      sql.NullInt64{Int64: 1, Valid: true},
				DistrictID:     sql.NullInt64{Int64: 10, Valid: true},
				VillageID:      sql.NullInt64{Int64: 100, Valid: true},
				Latitude:       sql.NullString{String: "-6.231928", Valid: true},
				Longitude:      sql.NullString{String: "0.988789", Valid: true},
				CoverImagePath: sql.NullString{String: "http://localhost:9080", Valid: true},
				Status:         sql.NullInt64{Int64: 10, Valid: true},
				CreatedAt:      sql.NullTime{Time: currentTime, Valid: true},
				UpdatedAt:      sql.NullTime{Time: currentTime, Valid: true},
				CategoryID:     sql.NullInt64{Int64: 1, Valid: true},
			},
			Error: nil,
		},
		MockUpdateRepository: errors.New("failed_update_a_phonebook"),
		MockUsecase:          errors.New("failed_update_a_phonebook"),
	},
}

// UpdatePhonebookDescription :
func UpdatePhonebookDescription() []string {
	var arr = []string{}
	for _, data := range UpdatePhonebookTestcases {
		arr = append(arr, data.Description)
	}
	return arr
}
