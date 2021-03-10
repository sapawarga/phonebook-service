package model

// ParamsPhoneBook ...
type ParamsPhoneBook struct {
	Search     *string
	RegencyID  *int64
	DistrictID *int64
	VillageID  *int64
	Status     *int64
	Limit      *int64
	Page       *int64
}

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
}
