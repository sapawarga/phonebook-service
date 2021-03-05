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
	ID             int64
	PhoneNumber    string
	Description    string
	Name           string
	Address        string
	RegencyID      int64
	DistrictID     int64
	VillageID      int64
	Latitude       string
	Longitude      string
	CoverImagePath string
	Status         int64
	CreatedAt      time.Time
	UpdatedAt      time.Time
	CategoryID     int64
}
