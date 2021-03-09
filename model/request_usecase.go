package model

// ParamsPhoneBook ...
type ParamsPhoneBook struct {
	Search     string
	RegencyID  int64
	DistrictID int64
	VillageID  int64
	Status     int64
	Limit      int64
	Page       int64
}
