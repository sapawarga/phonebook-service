package model

import "database/sql"

// PhoneBookResponse ...
type PhoneBookResponse struct {
	ID             int64          `db:"id"`
	PhoneNumber    string         `db:"phone_number"`
	Description    sql.NullString `db:"description"`
	Name           sql.NullString `db:"name"`
	Address        sql.NullString `db:"address"`
	RegencyID      sql.NullInt64  `db:"kabkota_id"`
	DistrictID     sql.NullInt64  `db:"kec_id"`
	VillageID      sql.NullInt64  `db:"kel_id"`
	Latitude       sql.NullString `db:"latitude"`
	Longitude      sql.NullString `db:"longitude"`
	CoverImagePath sql.NullString `db:"cover_image_path"`
	Status         sql.NullInt64  `db:"status"`
	CreatedAt      sql.NullTime   `db:"created_at"`
	UpdatedAt      sql.NullTime   `db:"updated_at"`
	CategoryID     sql.NullInt64  `db:"category_id"`
}
