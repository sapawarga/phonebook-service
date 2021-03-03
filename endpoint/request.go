package endpoint

// GetListRequest ...
type GetListRequest struct {
	Name         string `json:"name"`
	PhoneNumber  string `json:"phone_number"`
	RegencyCode  string `json:"regency_code"`
	DistrictCode string `json:"district_code"`
	VillageCode  string `json:"village_code"`
}
