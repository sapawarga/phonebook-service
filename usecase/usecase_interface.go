package usecase

import (
	"context"

	"github.com/sapawarga/phonebook-service/model"
)

// Provider interface for PhoneBook
type Provider interface {
	GetList(ctx context.Context, params *model.ParamsPhoneBook) (*model.PhoneBookWithMeta, error)
	GetDetail(ctx context.Context, id int64) (*model.PhonebookDetail, error)
	Insert(ctx context.Context, params *model.AddPhonebook) error
	Update(ctx context.Context, params *model.UpdatePhonebook) error
}
