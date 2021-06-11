package endpoint

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/sapawarga/phonebook-service/model"
)

// PhoneBookWithMeta ...
type PhoneBookWithMeta struct {
	Data     []*Phonebook `json:"data"`
	Metadata *Metadata    `json:"metadata"`
}

// Phonebook ...
type Phonebook struct {
	ID           int64          `json:"id"`
	PhoneNumbers []*PhoneNumber `json:"phone_numbers"`
	Description  string         `json:"description"`
	Name         string         `json:"name"`
	Address      string         `json:"address"`
	Latitude     string         `json:"latitude"`
	Longitude    string         `json:"longitude"`
	Status       int64          `json:"status,omitempty"`
	Category     string         `json:"category_name"`
	Distance     float64        `json:"distance,omitempty"`
}

// Metadata ...
type Metadata struct {
	Page  int64 `json:"page"`
	Total int64 `json:"total"`
}

// PhonebookDetail ...
type PhonebookDetail struct {
	ID             int64          `json:"id"`
	Name           string         `json:"name"`
	CategoryID     int64          `json:"category_id"`
	CategoryName   string         `json:"category_name"`
	Address        string         `json:"address"`
	Description    string         `json:"description"`
	PhoneNumbers   []*PhoneNumber `json:"phone_numbers"`
	RegencyID      *int64         `json:"regency_id,omitempty"`
	RegencyName    *string        `json:"regency_name,omitempty"`
	DistrictID     *int64         `json:"district_id,omitempty"`
	DistrictName   *string        `json:"district_name,omitempty"`
	VillageID      *int64         `json:"village_id,omitempty"`
	VillageName    *string        `json:"village_name,omitempty"`
	Latitude       string         `json:"latitude"`
	Longitude      string         `json:"longitude"`
	CoverImagePath string         `json:"cover_image_path"`
	Status         int64          `json:"status"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
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
		if err := json.Unmarshal([]byte(v.PhoneNumbers), &phoneNumbers); err != nil {
			fmt.Println("error unmarshal", err)
			return nil
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
