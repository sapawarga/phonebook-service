package endpoint

// GetListRequest ...
type GetListRequest struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	RegencyID   int64  `json:"regency_id"`
	DistrictID  int64  `json:"district_id"`
	VillageID   int64  `json:"village_id"`
	Status      int64  `json:"status"`
	Limit       int64  `json:"limit"`
	Page        int64  `json:"page"`
}
