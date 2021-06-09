package model

import "time"

// PhoneBookWithMeta ...
type PhoneBookWithMeta struct {
	PhoneBooks []*Phonebook `json:"data"`
	Page       int64        `json:"page"`
	Total      int64        `json:"total"`
	TotalPage  int64        `json:"total_page"`
}

// Phonebook ...
type Phonebook struct {
	ID           int64   `json:"id"`
	PhoneNumbers string  `json:"phone_numbers"`
	Description  string  `json:"description"`
	Name         string  `json:"name"`
	Address      string  `json:"address"`
	Latitude     string  `json:"latitude"`
	Longitude    string  `json:"longitude"`
	Status       int64   `json:"status,omitempty"`
	Category     string  `json:"category_name"`
	Distance     float64 `json:"distance,omitempty"`
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
	RegencyName    *string   `json:"regency_name,omitempty"`
	DistrictID     int64     `json:"district_id"`
	DistrictName   *string   `json:"district_name,omitempty"`
	VillageID      int64     `json:"village_id"`
	VillageName    *string   `json:"village_name,omitempty"`
	Latitude       string    `json:"latitude"`
	Longitude      string    `json:"longitude"`
	CoverImagePath string    `json:"cover_image_path"`
	Status         int64     `json:"status"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
