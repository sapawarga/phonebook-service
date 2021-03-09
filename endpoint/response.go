package endpoint

import (
	"time"

	"github.com/sapawarga/phonebook-service/model"
)

// PhoneBookWithMeta ...
type PhoneBookWithMeta struct {
	Data     []*model.Phonebook `json:"data"`
	Metadata *Metadata          `json:"metadata"`
}

// Metadata ...
type Metadata struct {
	Page  int64 `json:"page"`
	Total int64 `json:"total"`
}

// PhonebookDetail ...
type PhonebookDetail struct {
	ID             int64     `json:"id"`
	Name           string    `json:"name"`
	CategoryID     int64     `json:"category_id"`
	CategoryName   string    `json:"category_name"`
	Address        string    `json:"address"`
	Description    string    `json:"description"`
	PhoneNumbers   string    `json:"phone_numbers"`
	RegencyID      int64     `json:"regency_id"`
	RegencyName    string    `json:"regency_name"`
	DistrictID     int64     `json:"district_id"`
	DistrictName   string    `json:"district_name"`
	VillageID      int64     `json:"village_id"`
	VillageName    string    `json:"village_name"`
	Latitude       string    `json:"latitude"`
	Longitude      string    `json:"longitude"`
	CoverImagePath string    `json:"cover_image_path"`
	Status         int64     `json:"status"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// StatusResponse ...
type StatusResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
