package endpoint

// GetListRequest ...
type GetListRequest struct {
	Search      string `httpquery:"search"`
	Name        string `httpquery:"name"`
	PhoneNumber string `httpquery:"phone"`
	RegencyID   int64  `httpquery:"kabkota_id"`
	DistrictID  int64  `httpquery:"kec_id"`
	VillageID   int64  `httpquery:"kel_id"`
	Status      *int64 `httpquery:"status"`
	Latitude    string `httpquery:"latitude"`
	Longitude   string `httpquery:"longitude"`
	Limit       int64  `httpquery:"limit"`
	Page        int64  `httpquery:"page"`
	SortBy      string `httpquery:"sort_by"`
	OrderBy     string `httpquery:"sort_order"`
}

// AddPhonebookRequest ...
type AddPhonebookRequest struct {
	Name           string         `json:"name"`
	Description    *string        `json:"description"`
	PhoneNumbers   []*PhoneNumber `json:"phone_numbers"`
	RegencyID      *int64         `json:"kabkota_id"`
	DistrictID     *int64         `json:"kec_id"`
	VillageID      *int64         `json:"kel_id"`
	Status         *int64         `json:"status"`
	Latitude       *string        `json:"latitude"`
	Longitude      *string        `json:"longitude"`
	CategoryID     *int64         `json:"category_id"`
	CoverImagePath *string        `json:"cover_image_path"`
	Address        *string        `json:"address"`
	Sequence       *int64         `json:"seq"`
}

// GetDetailRequest ...
type GetDetailRequest struct {
	ID int64 `httpparams:"id"`
}

// UpdatePhonebookRequest ...
type UpdatePhonebookRequest struct {
	ID             int64          `json:"id"`
	Name           string         `json:"name"`
	Description    *string        `json:"description"`
	PhoneNumbers   []*PhoneNumber `json:"phone_numbers"`
	RegencyID      *int64         `json:"kabkota_id"`
	DistrictID     *int64         `json:"kec_id"`
	VillageID      *int64         `json:"kel_id"`
	Status         *int64         `json:"status"`
	Latitude       *string        `json:"latitude"`
	Longitude      *string        `json:"longitude"`
	CategoryID     *int64         `json:"category_id"`
	CoverImagePath *string        `json:"cover_image_path"`
	Address        *string        `json:"address"`
	Sequence       *int64         `json:"seq,omitempty"`
}

// PhoneNumber ...
type PhoneNumber struct {
	PhoneNumber string `json:"phone_number"`
	Type        string `json:"type"`
}

// IsExistPhoneNumber ...
type IsExistPhoneNumber struct {
	PhoneNumber string `httpquery:"phone_number"`
}
