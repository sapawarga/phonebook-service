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
	var result []*model.PhoneBookResponse
	var err error

	query.WriteString(`
		SELECT id, name, description, address, phone_numbers, kabkota_id, kec_id, kel_id, latitude, longitude, cover_image_path,
			status, seq, created_at, updated_at, category_id
		FROM phonebooks`)

	selectQuery, queryParams := querySelectParams(ctx, query, params)
	query.WriteString(selectQuery.String())

	if params.SortBy != nil && params.OrderBy != nil {
		query.WriteString(fmt.Sprintf(" ORDER BY %s %s ", helper.GetStringFromPointer(params.SortBy), helper.GetStringFromPointer(params.OrderBy)))
	} else {
		query.WriteString(" ORDER BY seq DESC")
	}

	query.WriteString(" LIMIT ?, ?")
	queryParams = append(queryParams, helper.GetInt64FromPointer(params.Offset), helper.GetInt64FromPointer(params.Limit))

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
	var total int64
	var err error

	query.WriteString(`SELECT COUNT(1) FROM phonebooks`)
	selectQuery, queryParams := querySelectParams(ctx, query, params)
	query.WriteString(selectQuery.String())

	if ctx != nil {
		err = r.conn.GetContext(ctx, &total, query.String(), queryParams...)
	} else {
		err = r.conn.Get(&total, query.String(), queryParams...)
	}

	if err == sql.ErrNoRows {
		return 0, nil
	}

	if err != nil {
		return 0, err
	}

	return total, nil
}

// GetPhonebookDetailByID ...
func (r *PhonebookRepository) GetPhonebookDetailByID(ctx context.Context, id int64) (*model.PhoneBookResponse, error) {
	var query bytes.Buffer
	var result = &model.PhoneBookResponse{}
	var err error

	query.WriteString(`
	SELECT id, phone_numbers, description, name, address, kabkota_id, kec_id, kel_id, latitude, longitude, 
	cover_image_path, status, seq, created_at, updated_at, category_id FROM phonebooks`)
	query.WriteString(" WHERE id = ? ")

	if ctx != nil {
		err = r.conn.GetContext(ctx, result, query.String(), id)
	} else {
		err = r.conn.Get(result, query.String(), id)
	}

	if err == sql.ErrNoRows {
		return nil, nil
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

	query.WriteString(` SELECT name from categories WHERE id = ? AND type = 'phonebook' AND status = 10 `)
	if ctx != nil {
		err = r.conn.GetContext(ctx, &result, query.String(), id)
	} else {
		err = r.conn.Get(&result, query.String(), id)
	}

	if err == sql.ErrNoRows {
		return "", nil
	}

	if err != nil {
		return "", err
	}

	return result, nil
}

// GetLocationByID ...
func (r *PhonebookRepository) GetLocationByID(ctx context.Context, id int64) (*model.Location, error) {
	var query bytes.Buffer
	var result = &model.Location{}
	var err error

	query.WriteString(` SELECT id, name, code_bps FROM areas WHERE id = ?`)
	if ctx != nil {
		err = r.conn.GetContext(ctx, result, query.String(), id)
	} else {
		err = r.conn.Get(result, query.String(), id)
	}

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return result, nil
}

// Insert ...
func (r *PhonebookRepository) Insert(ctx context.Context, params *model.AddPhonebook) error {
	var query bytes.Buffer
	var err error
	_, current := helper.GetCurrentTimeUTC()

	query.WriteString("INSERT INTO phonebooks")
	query.WriteString(`
		(name, description, address, phone_numbers, kabkota_id, kec_id, kel_id, latitude, longitude, 
		seq, cover_image_path, status, created_at, updated_at, category_id)`)
	query.WriteString(`VALUES(
		:name, :description, :address, :phone_numbers, :kabkota_id, :kec_id, :kel_id, :latitude, :longitude, 
		:seq, :cover_image_path, :status, :created_at, :updated_at, :category_id)`)
	queryParams := map[string]interface{}{
		"name":             params.Name,
		"description":      params.Description,
		"address":          params.Address,
		"phone_numbers":    params.PhoneNumbers,
		"kabkota_id":       params.RegencyID,
		"kec_id":           params.DistrictID,
		"kel_id":           params.VillageID,
		"latitude":         params.Latitude,
		"longitude":        params.Longitude,
		"cover_image_path": params.CoverImagePath,
		"status":           params.Status,
		"created_at":       current,
		"updated_at":       current,
		"category_id":      params.CategoryID,
		"seq":              params.Sequence,
	}

	if ctx != nil {
		_, err = r.conn.NamedExecContext(ctx, query.String(), queryParams)
	} else {
		_, err = r.conn.NamedExec(query.String(), queryParams)
	}

	if err != nil {
		return err
	}

	return nil
}

// Update ...
func (r *PhonebookRepository) Update(ctx context.Context, params *model.UpdatePhonebook) error {
	var query bytes.Buffer
	var queryParams = make(map[string]interface{})
	var err error

	query.WriteString(" UPDATE phonebooks SET ")
	updateQuery, queryParams := queryUpdateParams(ctx, params, queryParams)
	query.WriteString(updateQuery.String())

	if ctx != nil {
		_, err = r.conn.NamedExecContext(ctx, query.String(), queryParams)
	} else {
		_, err = r.conn.NamedExec(query.String(), queryParams)
	}

	if err != nil {
		return err
	}

	return nil
}

// Delete ...
func (r *PhonebookRepository) Delete(ctx context.Context, id int64) error {
	var query bytes.Buffer
	var params = make(map[string]interface{})
	var err error

	query.WriteString(" UPDATE phonebooks SET status = :status WHERE id = :id ")
	params["status"] = helper.DELETED
	params["id"] = id
	if ctx != nil {
		_, err = r.conn.NamedExecContext(ctx, query.String(), params)
	} else {
		_, err = r.conn.NamedExec(query.String(), params)
	}

	if err != nil {
		return err
	}

	return nil
}

// IsExistPhoneNumber ...
func (r *PhonebookRepository) IsExistPhoneNumber(ctx context.Context, phone string) (bool, error) {
	var query bytes.Buffer
	var count int
	var err error

	query.WriteString(` SELECT count(1) FROM phonebooks 
	WHERE JSON_CONTAINS(phone_numbers->'$[*].phone_number', json_array(?))`)

	if ctx != nil {
		err = r.conn.GetContext(ctx, &count, query.String(), phone)
	} else {
		err = r.conn.Get(&count, query.String(), phone)
	}

	if err != nil {
		return false, err
	}

	if count == 0 || err == sql.ErrNoRows {
		return false, nil
	}

	return true, nil
}
