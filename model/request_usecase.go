package model

// AddPhonebook ...
type AddPhonebook struct {
	Name           string
	PhoneNumbers   *string
	Address        *string
	Description    *string
	RegencyID      *int64
	DistrictID     *int64
	VillageID      *int64
	Latitude       *string
	Longitude      *string
	CoverImagePath *string
	Status         *int64
	CategoryID     *int64
	Sequence       *int64
}

// UpdatePhonebook ...
type UpdatePhonebook struct {
	ID             int64
	Name           string
	PhoneNumbers   *string
	Address        *string
	Description    *string
	RegencyID      *int64
	DistrictID     *int64
	VillageID      *int64
	Latitude       *string
	Longitude      *string
	CoverImagePath *string
	Status         *int64
	CategoryID     *int64
	Sequence       *int64
}
