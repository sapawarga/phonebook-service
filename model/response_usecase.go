package model

import "time"

// PhoneBookWithMeta ...
type PhoneBookWithMeta struct {
	PhoneBooks []*Phonebook `json:"items"`
	Metadata   *Metadata    `json:"_meta"`
}

type Metadata struct {
	TotalCount  int64   `json:"totalCount"`
	PageCount   float64 `json:"pageCount"`
	CurrentPage int64   `json:"currentPage"`
	PerPage     int64   `json:"perPage"`
}

// Phonebook ...
type Phonebook struct {
	ID            int64     `json:"id"`
	PhoneNumbers  string    `json:"phone_numbers"`
	Description   string    `json:"description"`
	Name          string    `json:"name"`
	Address       string    `json:"address"`
	Latitude      string    `json:"latitude"`
	Longitude     string    `json:"longitude"`
	RegencyID     int64     `json:"kabkota_id,omitempty"`
	CoverImageURL string    `json:"cover_image_url,omitempty"`
	DistrictID    int64     `json:"kec_id,omitempty"`
	VillageID     int64     `json:"kel_id,omitempty"`
	Status        int64     `json:"status,omitempty"`
	Category      *Category `json:"category"`
	Distance      float64   `json:"distance,omitempty"`
	CreatedAt     time.Time `json:"created_at,omitempty"`
	UpdatedAt     time.Time `json:"updated_at,omitempty"`
}

// PhonebookDetail ...
type PhonebookDetail struct {
	ID             int64     `json:"id"`
	Name           string    `json:"name"`
	Category       *Category `json:"category"`
	Address        string    `json:"address"`
	Description    string    `json:"description"`
	PhoneNumbers   string    `json:"phone_numbers"`
	Regency        *Location `json:"kabkota,omitempty"`
	District       *Location `json:"kecamatan,omitempty"`
	Village        *Location `json:"kelurahan,omitempty"`
	Latitude       string    `json:"latitude"`
	Longitude      string    `json:"longitude"`
	CoverImagePath string    `json:"cover_image_path"`
	Status         int64     `json:"status"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type Category struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Location struct {
	ID      int64  `db:"id" json:"id" `
	BPSCode string `db:"code_bps" json:"code_bps" `
	Name    string `db:"name" json:"name" `
}
