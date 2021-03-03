package usecase

import (
	"context"

	"github.com/sapawarga/phonebook-service/model"
)

// Provider interface for PhoneBook
type Provider interface {
	GetList(ctx context.Context, params *model.ParamsPhoneBook) (*model.PhoneBookWithMeta, error)
}
