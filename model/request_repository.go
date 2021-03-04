package model

// GetListRequest ...
type GetListRequest struct {
	Name         string
	PhoneNumber  string
	RegencyCode  string
	DistrictCode string
	VillageCode  string
	Status       int64
	Limit        int64
	Offset       int64
}
