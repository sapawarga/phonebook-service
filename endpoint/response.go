package endpoint

import (
	"encoding/json"
	"fmt"

	"github.com/sapawarga/phonebook-service/config"
	"github.com/sapawarga/phonebook-service/helper"
	"github.com/sapawarga/phonebook-service/model"
)

// PhoneBookWithMeta ...
type PhoneBookWithMeta struct {
	Data *PhonebookWithMeta `json:"data"`
}

// PhonebookWithMeta ...
type PhonebookWithMeta struct {
	Phonebooks []*Phonebook `json:"items"`
	Meta       *Metadata    `json:"_meta"`
}

// Metadata ...
type Metadata struct {
	TotalCount  int64   `json:"totalCount"`
	PageCount   float64 `json:"pageCount"`
	CurrentPage int64   `json:"currentPage"`
	PerPage     int64   `json:"perPage"`
}

// Phonebook ...
type Phonebook struct {
	ID            interface{}     `json:"id"`
	PhoneNumbers  []*PhoneNumber  `json:"phone_numbers"`
	Description   *string         `json:"description"`
	Name          *string         `json:"name"`
	Address       *string         `json:"address"`
	Latitude      *string         `json:"latitude"`
	Longitude     *string         `json:"longitude"`
	Status        int64           `json:"status"`
	StatusLabel   string          `json:"status_label"`
	CoverImageURL *string         `json:"cover_image_url"`
	CategoryID    *int64          `json:"category_id"`
	Category      interface{}     `json:"category"`
	Sequence      int64           `json:"seq"`
	Distance      interface{}     `json:"distance,omitempty"`
	RegencyID     *int64          `json:"kabkota_id"`
	Regency       *model.Location `json:"kabkota"`
	DistrictID    *int64          `json:"kec_id"`
	District      *model.Location `json:"kecamatan"`
	VillageID     *int64          `json:"kel_id"`
	Village       *model.Location `json:"kelurahan"`
	CreatedAt     int64           `json:"created_at"`
	UpdatedAt     int64           `json:"updated_at"`
}

// PhonebookDetail ...
type PhonebookDetail struct {
	ID             interface{}     `json:"id"`
	Name           string          `json:"name"`
	CategoryID     int64           `json:"category_id"`
	Category       interface{}     `json:"category"`
	Address        string          `json:"address"`
	Description    string          `json:"description"`
	PhoneNumbers   []*PhoneNumber  `json:"phone_numbers"`
	RegencyID      *int64          `json:"kabkota_id"`
	Regency        *model.Location `json:"kabkota"`
	DistrictID     *int64          `json:"kec_id"`
	District       *model.Location `json:"kecamatan"`
	VillageID      *int64          `json:"kel_id"`
	Village        *model.Location `json:"kelurahan"`
	Latitude       string          `json:"latitude"`
	Longitude      string          `json:"longitude"`
	Sequence       int64           `json:"seq"`
	CoverImagePath string          `json:"cover_image_url"`
	Status         int64           `json:"status"`
	StatusLabel    string          `json:"status_label"`
	CreatedAt      int64           `json:"created_at"`
	UpdatedAt      int64           `json:"updated_at"`
}

// StatusResponse ...
type StatusResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

var cfg, _ = config.NewConfig()

// EncodePhonebook ...
func EncodePhonebook(data []*model.Phonebook) []*Phonebook {
	result := make([]*Phonebook, 0)
	for _, v := range data {
		phoneNumbers := []*PhoneNumber{}

		if v.PhoneNumbers != "" {
			_ = json.Unmarshal([]byte(v.PhoneNumbers), &phoneNumbers)
		}

		category, ok := v.Category.(*model.Category)

		encodeData := &Phonebook{
			ID:            v.ID,
			PhoneNumbers:  phoneNumbers,
			Description:   helper.SetPointerString(v.Description),
			Name:          helper.SetPointerString(v.Name),
			Address:       helper.SetPointerString(v.Address),
			CoverImageURL: helper.SetPointerString(fmt.Sprintf("%s/%s", cfg.AppStoragePublicURL, v.CoverImageURL)),
			Latitude:      helper.SetPointerString(v.Latitude),
			Longitude:     helper.SetPointerString(v.Longitude),
			Status:        v.Status,
			StatusLabel:   GetStatusLabel[v.Status]["id"],
			Category:      v.Category,
			Distance:      v.Distance,
			Sequence:      v.Sequence,
			CreatedAt:     v.CreatedAt,
			UpdatedAt:     v.UpdatedAt,
			Regency:       v.Regency,
			District:      v.District,
			Village:       v.Village,
		}
		if ok {
			encodeData.CategoryID = helper.SetPointerInt64(category.ID)
		}
		if v.Regency != nil {
			encodeData.RegencyID = helper.SetPointerInt64(v.Regency.ID)
		}
		if v.District != nil {
			encodeData.DistrictID = helper.SetPointerInt64(v.District.ID)
		}
		if v.Village != nil {
			encodeData.VillageID = helper.SetPointerInt64(v.Village.ID)
		}
		if len(phoneNumbers) == 0 {
			encodeData.PhoneNumbers = nil
		}

		result = append(result, encodeData)
	}

	return result
}

// GetStatusLabel ...
var GetStatusLabel = map[int64]map[string]string{
	-1: {"en": "status deleted", "id": "Dihapus"},
	0:  {"en": "Not Active", "id": "Tidak Aktif"},
	10: {"en": "Active", "id": "Aktif"},
}
