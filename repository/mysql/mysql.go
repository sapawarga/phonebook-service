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
	var query bytes.Buffer
	var queryParams []interface{}
	var result []*model.PhoneBookResponse
	var err error
	var first = true

	query.WriteString(`
		SELECT
			id, name, description, address, phone_numbers, kabkota_id, kec_id, kel_id, latitude, longitude, cover_image_path,
			status, FROM_UNIXTIME(created_at) as created_at, FROM_UNIXTIME(updated_at) as updated_at, category_id
		FROM sapawarga.phonebooks`)

	if params.Search != nil {
		if first {
			query.WriteString(" WHERE ")
		} else {
			query.WriteString(" AND ")
		}
		query.WriteString(fmt.Sprintf(`(name LIKE LOWER(%s) `, "'%'"+helper.GetStringFromPointer(params.Search)+"'%'"))
		queryParams = append(queryParams, params.Search)
		query.WriteString(fmt.Sprintf(` OR phone_numbers LIKE %s )`, "'%'"+helper.GetStringFromPointer(params.Search)+"'%'"))
		queryParams = append(queryParams, params.Search)
		first = false
	}

	if params.RegencyID != nil {
		if first {
			query.WriteString(" WHERE ")
		} else {
			query.WriteString(" AND ")
		}
		query.WriteString(" kabkota_id = ? ")
		queryParams = append(queryParams, params.RegencyID)
		first = false
	}

	if params.DistrictID != nil {
		if first {
			query.WriteString(" WHERE ")
		} else {
			query.WriteString(" AND ")
		}
		query.WriteString(" kec_id = ? ")
		queryParams = append(queryParams, params.DistrictID)
		first = false
	}

	if params.VillageID != nil {
		if first {
			query.WriteString(" WHERE ")
		} else {
			query.WriteString(" AND ")
		}
		query.WriteString(" kel_id = ?")
		queryParams = append(queryParams, params.VillageID)
		first = false
	}

	if params.Status != nil {
		if first {
			query.WriteString(" WHERE ")
		} else {
			query.WriteString(" AND ")
		}
		query.WriteString(" status = ?")
		queryParams = append(queryParams, params.Status)
	}

	query.WriteString(" LIMIT ?, ?")
	queryParams = append(queryParams, params.Offset, params.Limit)

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

// GetPhonebookDetailByID ...
func (r *PhonebookRepository) GetPhonebookDetailByID(ctx context.Context, id int64) (*model.PhoneBookResponse, error)

// GetCategoryNameByID ...
func (r *PhonebookRepository) GetCategoryNameByID(ctx context.Context, id int64) (string, error)

// GetLocationNameByID ...
func (r *PhonebookRepository) GetLocationNameByID(ctx context.Context, id int64) (string, error)

// Insert ...
func (r *PhonebookRepository) Insert(ctx context.Context, params *model.AddPhonebook) error

// Update ...
func (r *PhonebookRepository) Update(ctx context.Context, params *model.UpdatePhonebook) error

// Delete ...
func (r *PhonebookRepository) Delete(ctx context.Context, id int64) error
