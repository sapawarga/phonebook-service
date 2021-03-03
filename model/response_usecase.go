package model

// PhoneBookWithMeta ...
type PhoneBookWithMeta struct {
	PhoneBooks []*PhoneBookResponse
	Page       int64
	Total      int64
}
