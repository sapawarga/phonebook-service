package model

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
	ID            interface{} `json:"id"`
	PhoneNumbers  string      `json:"phone_numbers"`
	Description   string      `json:"description"`
	Name          string      `json:"name"`
	Address       string      `json:"address"`
	Latitude      string      `json:"latitude"`
	Longitude     string      `json:"longitude"`
	CoverImageURL string      `json:"cover_image_url"`
	Regency       *Location   `json:"kabkota"`
	District      *Location   `json:"kecamatan"`
	Village       *Location   `json:"kelurahan"`
	Status        int64       `json:"status"`
	Sequence      int64       `json:"seq"`
	Category      interface{} `json:"category"`
	Distance      interface{} `json:"distance,omitempty"`
	CreatedAt     int64       `json:"created_at"`
	UpdatedAt     int64       `json:"updated_at"`
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
