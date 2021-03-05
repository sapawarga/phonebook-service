package mysql

import (
	"context"

	"github.com/sapawarga/phonebook-service/model"

	"github.com/jmoiron/sqlx"
)

// PhonebookRepository ...
type PhonebookRepository struct {
	conn *sqlx.DB
}

// NewPhonebookRepository ...
func NewPhonebookRepository(conn *sqlx.DB) *PhonebookRepository {
	return &PhonebookRepository{
		conn: conn,
	}
}

// GetListPhoneBook ...
func (r *PhonebookRepository) GetListPhoneBook(ctx context.Context, params *model.GetListRequest) ([]*model.PhoneBookResponse, error) {
	// TODO: create query for get list phone book
	// var query bytes.Buffer
	// var queryParams []interface{}
	// var result []*model.PhoneBookResponse
	// var err error

	// query.WriteString(`
	// 	SELECT
	// 		id,
	// 		name,
	// 		description,
	// 		address,
	// 		phone_numbers,
	// 		kabkota_id,
	// 		kec_id,
	// 		kel_id,
	// 		latitude,
	// 		longitude,
	// 		cover_image_path,
	// 		status,
	// 		FROM_UNIXTIME(created_at) as created_at,
	// 		FROM_UNIXTIME(updated_at) as updated_at,
	// 		category_id
	// 	FROM sapawarga.phonebooks`)
	// query.WriteString(` WHERE name ilike `)
	return nil, nil
}

// GetMetaDataPhoneBook ...
func (r *PhonebookRepository) GetMetaDataPhoneBook(ctx context.Context, params *model.GetListRequest) (int64, error) {
	// TODO: create query for get metadata
	return 0, nil
}
