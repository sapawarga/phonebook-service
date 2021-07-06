package usecase

import (
	"context"

	"github.com/sapawarga/phonebook-service/model"
)

// Provider interface for PhoneBook
type Provider interface {
	GetList(ctx context.Context, params *model.ParamsPhoneBook) (*model.PhoneBookWithMeta, error)
	GetDetail(ctx context.Context, id int64) (*model.PhonebookDetail, error)
	IsExistPhoneNumber(ctx context.Context, phone string) (bool, error)
	Insert(ctx context.Context, params *model.AddPhonebook) error
	Update(ctx context.Context, params *model.UpdatePhonebook) error
	Delete(ctx context.Context, id int64) error
	CheckHealthReadiness(ctx context.Context) error
}
