package model

// GetListRequest penggunaan pointer ini agar dapat memberikan value nil jika tidak digunakan
type GetListRequest struct {
	Search     *string
	Name       *string
	Phone      *string
	RegencyID  *int64
	DistrictID *int64
	VillageID  *int64
	Status     *int64
	Limit      *int64
	Offset     *int64
	Longitude  *string
	Latitude   *string
	OrderBy    *string
	SortBy     *string
}
