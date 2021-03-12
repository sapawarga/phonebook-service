package endpoint

// GetListRequest ...
type GetListRequest struct {
	Search     string `httpquery:"search"`
	RegencyID  int64  `httpquery:"regency_id"`
	DistrictID int64  `httpquery:"district_id"`
	VillageID  int64  `httpquery:"village_id"`
	Status     int64  `httpquery:"status"`
	Latitude   string `httpquery:"latitude"`
	Longitude  string `httpquery:"longitude"`
	Limit      int64  `httpquery:"limit"`
	Page       int64  `httpquery:"page"`
}

// AddPhonebookRequest ...
type AddPhonebookRequest struct {
	Name           string  `json:"name"`
	Description    *string `json:"description"`
	PhoneNumbers   *string `json:"phone_numbers"`
	RegencyID      *int64  `json:"regency_id"`
	DistrictID     *int64  `json:"district_id"`
	VillageID      *int64  `json:"village_id"`
	Status         *int64  `json:"status"`
	Latitude       *string `json:"latitude"`
	Longitude      *string `json:"longitude"`
	CategoryID     *int64  `json:"category_id"`
	CoverImagePath *string `json:"cover_image_path"`
	Address        *string `json:"address"`
}

// GetDetailRequest ...
type GetDetailRequest struct {
	ID int64 `httpparams:"id"`
}

// UpdatePhonebookRequest ...
type UpdatePhonebookRequest struct {
	ID             int64   `json:"id"`
	Name           string  `json:"name"`
	Description    *string `json:"description"`
	PhoneNumbers   *string `json:"phone_numbers"`
	RegencyID      *int64  `json:"regency_id"`
	DistrictID     *int64  `json:"district_id"`
	VillageID      *int64  `json:"village_id"`
	Status         *int64  `json:"status"`
	Latitude       *string `json:"latitude"`
	Longitude      *string `json:"longitude"`
	CategoryID     *int64  `json:"category_id"`
	CoverImagePath *string `json:"cover_image_path"`
	Address        *string `json:"address"`
}
