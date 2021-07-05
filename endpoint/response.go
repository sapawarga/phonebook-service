package endpoint

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/sapawarga/phonebook-service/config"
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
	ID            int64           `json:"id"`
	PhoneNumbers  []*PhoneNumber  `json:"phone_numbers"`
	Description   string          `json:"description"`
	Name          string          `json:"name"`
	Address       string          `json:"address"`
	Latitude      string          `json:"latitude"`
	Longitude     string          `json:"longitude"`
	Status        int64           `json:"status,omitempty"`
	StatusLabel   string          `json:"status_label,omitempty"`
	CoverImageURL string          `json:"cover_image_url,omitempty"`
	Category      *model.Category `json:"category"`
	Distance      float64         `json:"distance,omitempty"`
}

// PhonebookDetail ...
type PhonebookDetail struct {
	ID             int64           `json:"id"`
	Name           string          `json:"name"`
	Category       *model.Category `json:"category"`
	Address        string          `json:"address"`
	Description    string          `json:"description"`
	PhoneNumbers   []*PhoneNumber  `json:"phone_numbers"`
	Regency        *model.Location `json:"kabkota,omitempty"`
	District       *model.Location `json:"kecamatan,omitempty"`
	Village        *model.Location `json:"kelurahan,omitempty"`
	Latitude       string          `json:"latitude"`
	Longitude      string          `json:"longitude"`
	CoverImagePath string          `json:"cover_image_url"`
	Status         int64           `json:"status"`
	StatusLabel    string          `json:"status_label"`
	CreatedAt      time.Time       `json:"created_at"`
	UpdatedAt      time.Time       `json:"updated_at"`
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

		encodeData := &Phonebook{
			ID:            v.ID,
			PhoneNumbers:  phoneNumbers,
			Description:   v.Description,
			Name:          v.Name,
			Address:       v.Address,
			CoverImageURL: fmt.Sprintf("%s/%s", cfg.AppStoragePublicURL, v.CoverImageURL),
			Latitude:      v.Latitude,
			Longitude:     v.Longitude,
			Status:        v.Status,
			StatusLabel:   GetStatusLabel[v.Status]["ina"],
			Category:      v.Category,
			Distance:      v.Distance,
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
