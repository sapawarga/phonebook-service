package mysql

import (
	"bytes"
	"context"
	"fmt"

	"github.com/sapawarga/phonebook-service/helper"
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
	var query bytes.Buffer
	var queryParams []interface{}
	var result []*model.PhoneBookResponse
	var err error

	query.WriteString(`
		SELECT
			id, name, description, address, phone_numbers, kabkota_id, kec_id, kel_id, latitude, longitude, cover_image_path,
			status, FROM_UNIXTIME(created_at) as created_at, FROM_UNIXTIME(updated_at) as updated_at, category_id
		FROM sapawarga.phonebooks`)

	query.WriteString("WHERE ")
	if params.Name != nil {
		query.WriteString(fmt.Sprintf(` name LIKE LOWER(%s) `, "'%'"+helper.GetStringFromPointer(params.Name)+"'%'"))
		queryParams = append(queryParams, params.Name)
	} else if params.PhoneNumber != nil {
		query.WriteString(fmt.Sprintf(` AND phone_numbers LIKE %s `, "'%'"+helper.GetStringFromPointer(params.PhoneNumber)+"'%'"))
		queryParams = append(queryParams, params.Name)
	}

	if ctx != nil {
		err = r.conn.SelectContext(ctx, &result, query.String(), queryParams...)
	} else {
		err = r.conn.Select(&result, query.String(), queryParams...)
	}

	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetMetaDataPhoneBook ...
func (r *PhonebookRepository) GetMetaDataPhoneBook(ctx context.Context, params *model.GetListRequest) (int64, error) {
	// TODO: create query for get metadata
	return 0, nil
}
