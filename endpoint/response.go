package endpoint

import (
	"encoding/json"
	"time"

	"github.com/sapawarga/phonebook-service/model"
)

// PhoneBookWithMeta ...
type PhoneBookWithMeta struct {
	Data *model.PhoneBookWithMeta `json:"data"`
}

// Phonebook ...
type Phonebook struct {
	ID           int64           `json:"id"`
	PhoneNumbers []*PhoneNumber  `json:"phone_numbers"`
	Description  string          `json:"description"`
	Name         string          `json:"name"`
	Address      string          `json:"address"`
	Latitude     string          `json:"latitude"`
	Longitude    string          `json:"longitude"`
	Status       int64           `json:"status,omitempty"`
	Category     *model.Category `json:"category"`
	Distance     float64         `json:"distance,omitempty"`
}

// Metadata ...
type Metadata struct {
	Page  int64 `json:"page"`
	Total int64 `json:"total"`
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
	CoverImagePath string          `json:"cover_image_path"`
	Status         int64           `json:"status"`
	CreatedAt      time.Time       `json:"created_at"`
	UpdatedAt      time.Time       `json:"updated_at"`
}

// StatusResponse ...
type StatusResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func EncodePhonebook(data []*model.Phonebook) []*Phonebook {
	result := make([]*Phonebook, 0)
	for _, v := range data {
		phoneNumbers := []*PhoneNumber{}

		if v.PhoneNumbers != "" {
			_ = json.Unmarshal([]byte(v.PhoneNumbers), &phoneNumbers)
		}

		encodeData := &Phonebook{
			ID:           v.ID,
			PhoneNumbers: phoneNumbers,
			Description:  v.Description,
			Name:         v.Name,
			Address:      v.Address,
			Latitude:     v.Latitude,
			Longitude:    v.Longitude,
			Status:       v.Status,
			Category:     v.Category,
			Distance:     v.Distance,
		}

		result = append(result, encodeData)
	}

	return result
}
