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
		qBuffer := checkIsFirst(ctx, query, first)
		query.WriteString(qBuffer.String())
		query.WriteString(fmt.Sprintf(`(name LIKE LOWER(%s) `, "'%"+helper.GetStringFromPointer(params.Search)+"%'"))
		query.WriteString(fmt.Sprintf(` OR phone_numbers LIKE %s )`, "'%"+helper.GetStringFromPointer(params.Search)+"%'"))
		first = false
	}

	if params.RegencyID != nil {
		qBuffer := checkIsFirst(ctx, query, first)
		query.WriteString(qBuffer.String())
		query.WriteString(" kabkota_id = ? ")
		queryParams = append(queryParams, helper.GetInt64FromPointer(params.RegencyID))
		first = false
	}

	if params.DistrictID != nil {
		qBuffer := checkIsFirst(ctx, query, first)
		query.WriteString(qBuffer.String())
		query.WriteString(" kec_id = ? ")
		queryParams = append(queryParams, helper.GetInt64FromPointer(params.DistrictID))
		first = false
	}

	if params.VillageID != nil {
		qBuffer := checkIsFirst(ctx, query, first)
		query.WriteString(qBuffer.String())
		query.WriteString(" kel_id = ?")
		queryParams = append(queryParams, helper.GetInt64FromPointer(params.VillageID))
		first = false
	}

	if params.Status != nil {
		qBuffer := checkIsFirst(ctx, query, first)
		query.WriteString(qBuffer.String())
		query.WriteString(" status = ?")
		queryParams = append(queryParams, helper.GetInt64FromPointer(params.Status))
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
	var queryParams []interface{}
	var total int64
	var err error
	var first = true

	query.WriteString(`
		SELECT
			COUNT(1)
		FROM phonebooks`)

	if params.Search != nil {
		qBuffer := checkIsFirst(ctx, query, first)
		query.WriteString(qBuffer.String())
		query.WriteString(fmt.Sprintf(`(name LIKE LOWER(%s) `, "'%"+helper.GetStringFromPointer(params.Search)+"%'"))
		query.WriteString(fmt.Sprintf(` OR phone_numbers LIKE %s )`, "'%"+helper.GetStringFromPointer(params.Search)+"%'"))
		first = false
	}

	if params.RegencyID != nil {
		qBuffer := checkIsFirst(ctx, query, first)
		query.WriteString(qBuffer.String())
		query.WriteString(" kabkota_id = ? ")
		queryParams = append(queryParams, helper.GetInt64FromPointer(params.RegencyID))
		first = false
	}

	if params.DistrictID != nil {
		qBuffer := checkIsFirst(ctx, query, first)
		query.WriteString(qBuffer.String())
		query.WriteString(" kec_id = ? ")
		queryParams = append(queryParams, helper.GetInt64FromPointer(params.DistrictID))
		first = false
	}

	if params.VillageID != nil {
		qBuffer := checkIsFirst(ctx, query, first)
		query.WriteString(qBuffer.String())
		query.WriteString(" kel_id = ?")
		queryParams = append(queryParams, helper.GetInt64FromPointer(params.VillageID))
		first = false
	}

	if params.Status != nil {
		qBuffer := checkIsFirst(ctx, query, first)
		query.WriteString(qBuffer.String())
		query.WriteString(" status = ?")
		queryParams = append(queryParams, params.Status)
	}

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
	var result = &model.PhoneBookResponse{}
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

	query.WriteString(` SELECT name from categories WHERE id = ? AND type = 'phonebook' AND status = 10 `)
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
		1, :cover_image_path, :status, :created_at, :updated_at, :category_id)`)
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
	var first = true
	var err error
	_, unixTime := helper.GetCurrentTimeUTC()

	query.WriteString(" UPDATE phonebooks SET ")
	if params.Address != nil {
		query.WriteString(" address = :address ")
		queryParams["address"] = params.Address
		first = false
	}
	if params.CategoryID != nil {
		qBuffer := addComma(ctx, query, first)
		query.WriteString(qBuffer.String())
		query.WriteString(" category_id = :category_id ")
		queryParams["category_id"] = params.CategoryID
		first = false
	}
	if params.CoverImagePath != nil {
		qBuffer := addComma(ctx, query, first)
		query.WriteString(qBuffer.String())
		query.WriteString(" cover_image_path = :cover_image_path")
		queryParams["cover_image_path"] = params.CoverImagePath
		first = false
	}
	if params.Description != nil {
		qBuffer := addComma(ctx, query, first)
		query.WriteString(qBuffer.String())
		query.WriteString(" description = :description ")
		queryParams["description"] = params.Description
		first = false
	}
	if params.DistrictID != nil {
		qBuffer := addComma(ctx, query, first)
		query.WriteString(qBuffer.String())
		query.WriteString(" kec_id = :kec_id ")
		queryParams["kec_id"] = params.DistrictID
		first = false
	}
	if params.Latitude != nil {
		qBuffer := addComma(ctx, query, first)
		query.WriteString(qBuffer.String())
		query.WriteString(" latitude = :latitude ")
		queryParams["latitude"] = params.Latitude
		first = false
	}
	if params.Longitude != nil {
		qBuffer := addComma(ctx, query, first)
		query.WriteString(qBuffer.String())
		query.WriteString(" longitude = :longitude ")
		queryParams["longitude"] = params.Longitude
		first = false
	}
	if params.Name != "" {
		qBuffer := addComma(ctx, query, first)
		query.WriteString(qBuffer.String())
		query.WriteString(" name = :name ")
		queryParams["name"] = params.Name
		first = false
	}
	if params.PhoneNumbers != nil {
		qBuffer := addComma(ctx, query, first)
		query.WriteString(qBuffer.String())
		query.WriteString(" phone_numbers = :phone_numbers")
		queryParams["phone_numbers"] = params.PhoneNumbers
		first = false
	}
	if params.RegencyID != nil {
		qBuffer := addComma(ctx, query, first)
		query.WriteString(qBuffer.String())
		query.WriteString(" kabkota_id = :kabkota_id")
		queryParams["kabkota_id"] = params.RegencyID
		first = false
	}
	if params.Status != nil {
		qBuffer := addComma(ctx, query, first)
		query.WriteString(qBuffer.String())
		query.WriteString(" status = :status")
		queryParams["status"] = params.Status
		first = false
	}
	if params.VillageID != nil {
		qBuffer := addComma(ctx, query, first)
		query.WriteString(qBuffer.String())
		query.WriteString(" kel_id = :kel_id")
		queryParams["kel_id"] = params.VillageID
		first = false
	}

	qBuffer := addComma(ctx, query, first)
	query.WriteString(qBuffer.String())
	query.WriteString(" updated_at = :updated_at WHERE id = :id")
	queryParams["updated_at"] = unixTime
	queryParams["id"] = params.ID

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

func checkIsFirst(ctx context.Context, query bytes.Buffer, first bool) bytes.Buffer {
	if first {
		query.WriteString("WHERE")
	} else {
		query.WriteString("AND")
	}
	return query
}

func addComma(ctx context.Context, query bytes.Buffer, first bool) bytes.Buffer {
	if !first {
		query.WriteString(", ")
	}
	return query
}
