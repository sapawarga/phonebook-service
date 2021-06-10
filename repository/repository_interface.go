package repository

import (
	"context"

	"github.com/sapawarga/phonebook-service/model"
)

// PhoneBookI ...
type PhoneBookI interface {
	// Read Section
	GetListPhoneBook(ctx context.Context, params *model.GetListRequest) ([]*model.PhoneBookResponse, error)
	GetMetaDataPhoneBook(ctx context.Context, params *model.GetListRequest) (int64, error)
	GetPhonebookDetailByID(ctx context.Context, id int64) (*model.PhoneBookResponse, error)
	GetCategoryNameByID(ctx context.Context, id int64) (string, error)
	GetLocationNameByID(ctx context.Context, id int64) (*string, error)
	GetListPhonebookByLongLat(ctx context.Context, params *model.GetListRequest) ([]*model.PhoneBookResponse, error)
	GetListPhonebookByLongLatMeta(ctx context.Context, params *model.GetListRequest) (int64, error)
	// Create section
	Insert(ctx context.Context, params *model.AddPhonebook) error
	// Update section
	Update(ctx context.Context, params *model.UpdatePhonebook) error
	// Delete section
	Delete(ctx context.Context, id int64) error
	// Check health
	CheckHealthReadiness(ctx context.Context) error
}
