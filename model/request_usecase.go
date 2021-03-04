package model

// ParamsPhoneBook ...
type ParamsPhoneBook struct {
	Name         string
	PhoneNumber  string
	RegencyCode  string
	DistrictCode string
	VillageCode  string
	Status       int64
	Limit        int64
	Page         int64
}
