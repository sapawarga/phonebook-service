package model

import "time"

// PhoneBookWithMeta ...
type PhoneBookWithMeta struct {
	PhoneBooks []*Phonebook
	Page       int64
	Total      int64
}

// Phonebook ...
type Phonebook struct {
	ID           int64
	PhoneNumbers string
	Description  string
	Name         string
	Address      string
	Latitude     string
	Longitude    string
	Status       int64
}

// PhonebookDetail ...
type PhonebookDetail struct {
	ID             int64
	Name           string
	CategoryID     int64
	CategoryName   string
	Address        string
	Description    string
	PhoneNumbers   string
	RegencyID      int64
	RegencyName    string
	DistrictID     int64
	DistrictName   string
	VillageID      int64
	VillageName    string
	Latitude       string
	Longitude      string
	CoverImagePath string
	Status         int64
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
