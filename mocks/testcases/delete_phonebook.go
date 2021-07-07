package testcases

import (
	"database/sql"
	"errors"

	"github.com/sapawarga/phonebook-service/model"
)

// DeletePhonebook ...
type DeletePhonebook struct {
	Description                string
	UsecaseRequest             int64
	DeleteRepositoryRequest    int64
	GetDetailRepositoryRequest int64
	MockDetailRepository       GetDetailResponseRepository
	MockDeleteRepository       error
	MockUsecase                error
}

// DeletePhonebookTestcases ...
var DeletePhonebookTestcases = []DeletePhonebook{
	{
		Description:                "success_delete_phonebook",
		UsecaseRequest:             1,
		DeleteRepositoryRequest:    1,
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
				CreatedAt:      sql.NullInt64{Int64: currentTime, Valid: true},
				UpdatedAt:      sql.NullInt64{Int64: currentTime, Valid: true},
				CategoryID:     sql.NullInt64{Int64: 1, Valid: true},
			},
			Error: nil,
		},
		MockDeleteRepository: nil,
		MockUsecase:          nil,
	}, {
		Description:                "failed_get_detail_phonebook",
		UsecaseRequest:             1,
		DeleteRepositoryRequest:    1,
		GetDetailRepositoryRequest: 1,
		MockDetailRepository: GetDetailResponseRepository{
			Result: nil,
			Error:  errors.New("invalid_id_detail"),
		},
		MockDeleteRepository: nil,
		MockUsecase:          errors.New("invalid_id_detail"),
	}, {
		Description:                "failed_delete_phonebook",
		UsecaseRequest:             1,
		DeleteRepositoryRequest:    1,
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
				CreatedAt:      sql.NullInt64{Int64: currentTime, Valid: true},
				UpdatedAt:      sql.NullInt64{Int64: currentTime, Valid: true},
				CategoryID:     sql.NullInt64{Int64: 1, Valid: true},
			},
			Error: nil,
		},
		MockDeleteRepository: errors.New("failed_on_delete_phonebook"),
		MockUsecase:          errors.New("failed_on_delete_phonebook"),
	},
}

// DeletePhonebookDescription :
func DeletePhonebookDescription() []string {
	var arr = []string{}
	for _, data := range DeletePhonebookTestcases {
		arr = append(arr, data.Description)
	}
	return arr
}
