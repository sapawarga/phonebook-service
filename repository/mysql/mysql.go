package mysql

import (
	"bytes"
	"context"
	"database/sql"
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
		FROM phonebooks`)

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
	var query bytes.Buffer
	var queryParams []interface{}
	var total int64
	var err error
	var first = true

	query.WriteString(`
		SELECT
			COUNT(1)
		FROM phonebooks`)

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
		err = r.conn.GetContext(ctx, &total, query.String(), queryParams...)
	} else {
		err = r.conn.Get(&total, query.String(), queryParams...)
	}

	if err != nil {
		return 0, err
	}

	return total, nil
}

// GetPhonebookDetailByID ...
func (r *PhonebookRepository) GetPhonebookDetailByID(ctx context.Context, id int64) (*model.PhoneBookResponse, error) {
	var query bytes.Buffer
	var result *model.PhoneBookResponse
	var err error

	query.WriteString(`
	SELECT
		id, phone_numbers, description, name, address, kabkota_id, kec_id, kel_id, latitude, longitude, cover_image_path,
		status, FROM_UNIXTIME(created_at) as created_at, FROM_UNIXTIME(updated_at) as updated_at, category_id
	FROM phonebooks`)
	query.WriteString(" WHERE id = ? ")

	if ctx != nil {
		err = r.conn.GetContext(ctx, result, query.String(), id)
	} else {
		err = r.conn.Get(result, query.String(), id)
	}

	if err == sql.ErrNoRows {
		return nil, sql.ErrNoRows
	}

	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetCategoryNameByID ...
func (r *PhonebookRepository) GetCategoryNameByID(ctx context.Context, id int64) (string, error) {
	var query bytes.Buffer
	var result string
	var err error

	query.WriteString(` SELECT name from categories WHERE id = ?`)
	if ctx != nil {
		err = r.conn.GetContext(ctx, &result, query.String(), id)
	} else {
		err = r.conn.Get(&result, query.String(), id)
	}

	if err == sql.ErrNoRows {
		return "", sql.ErrNoRows
	}

	if err != nil {
		return "", err
	}

	return result, nil
}

// GetLocationNameByID ...
func (r *PhonebookRepository) GetLocationNameByID(ctx context.Context, id int64) (string, error) {
	var query bytes.Buffer
	var result string
	var err error

	query.WriteString(` SELECT name from areas WHERE id = ?`)
	if ctx != nil {
		err = r.conn.GetContext(ctx, &result, query.String(), id)
	} else {
		err = r.conn.Get(&result, query.String(), id)
	}

	if err == sql.ErrNoRows {
		return "", sql.ErrNoRows
	}

	if err != nil {
		return "", err
	}

	return result, nil
}

// Insert ...
func (r *PhonebookRepository) Insert(ctx context.Context, params *model.AddPhonebook) error

// Update ...
func (r *PhonebookRepository) Update(ctx context.Context, params *model.UpdatePhonebook) error

// Delete ...
func (r *PhonebookRepository) Delete(ctx context.Context, id int64) error
