package endpoint

import "github.com/sapawarga/phonebook-service/model"

// PhoneBookWithMeta ...
type PhoneBookWithMeta struct {
	Data     []*model.Phonebook `json:"data"`
	Metadata *Metadata          `json:"metadata"`
}

// Metadata ...
type Metadata struct {
	Page  int64 `json:"page"`
	Total int64 `json:"total"`
}
