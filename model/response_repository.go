package model

// PhoneBookResponse ...
type PhoneBookResponse struct {
	ID          int64  `db:"id"`
	PhoneNumber string `db:"phone_number"`
	Description string `db:"description"`
	Name        string `db:"name"`
}
