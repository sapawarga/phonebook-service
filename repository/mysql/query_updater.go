package mysql

import (
	"bytes"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/sapawarga/phonebook-service/helper"
	"github.com/sapawarga/phonebook-service/model"
)

func (r *PhonebookRepository) GetListPhonebookByLongLat(ctx context.Context, params *model.GetListRequest) ([]*model.PhoneBookResponse, error) {
	var query bytes.Buffer
	var result []*model.PhoneBookResponse
	var err error

	query.WriteString(` 
	SELECT id, category_id, name, description, address, phone_numbers, kabkota_id , kec_id , kel_id ,status ,
		cover_image_path, latitude, longitude, distance, created_at, updated_at
	`)

	query, queryParams := querySelectLongLat(ctx, query, params, false)

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

func (r *PhonebookRepository) GetListPhonebookByLongLatMeta(ctx context.Context, params *model.GetListRequest) (int64, error) {
	var query bytes.Buffer
	var total int64
	var err error

	query.WriteString(" SELECT COUNT(1) ")

	query, queryParams := querySelectLongLat(ctx, query, params, true)

	if ctx != nil {
		err = r.conn.GetContext(ctx, &total, query.String(), queryParams...)
	} else {
		err = r.conn.Get(&total, query.String(), queryParams...)
	}

	if err != nil || err == sql.ErrNoRows {
		return 0, errors.New("error_getting_total_data")
	}

	return total, nil
}

func (r *PhonebookRepository) CheckHealthReadiness(ctx context.Context) error {
	var err error
	if ctx != nil {
		err = r.conn.PingContext(ctx)
	} else {
		err = r.conn.Ping()
	}

	if err != nil {
		return err
	}

	return nil
}

func querySelectLongLat(ctx context.Context, query bytes.Buffer, params *model.GetListRequest, isCount bool) (newQuery bytes.Buffer, queryParams []interface{}) {
	query.WriteString(`
	FROM (
		SELECT pb.id, pb.category_id, pb.name, pb.description, pb.address,pb.phone_numbers,
			pb.latitude, pb.longitude, pb.kabkota_id ,pb.kec_id ,pb.kel_id, pb.cover_image_path , pb.status , pb.created_at, pb.updated_at ,
			c.radius, c.distance_unit
					* DEGREES(ACOS(COS(RADIANS(c.latpoint))
					* COS(RADIANS(pb.latitude))
					* COS(RADIANS(c.longpoint - pb.longitude))
					+ SIN(RADIANS(c.latpoint))
					* SIN(RADIANS(pb.latitude)))) AS distance
		FROM phonebooks AS pb
		JOIN (   
			SELECT  ? AS latpoint, ? AS longpoint, ? AS radius, 111.045 AS distance_unit
		) AS c ON 1=1
			WHERE pb.latitude
			BETWEEN c.latpoint  - (c.radius / c.distance_unit) AND c.latpoint  + (c.radius / c.distance_unit)
			AND pb.longitude
			BETWEEN c.longpoint - (c.radius / (c.distance_unit * COS(RADIANS(c.latpoint)))) AND c.longpoint + (c.radius / (c.distance_unit * COS(RADIANS(c.latpoint))))
			AND pb.status <> -1
	) AS d	WHERE distance <= radius
	`)
	queryParams = append(queryParams,
		helper.GetStringFromPointer(params.Latitude),
		helper.GetStringFromPointer(params.Longitude),
		helper.RADIUS)
	if !isCount {
		query.WriteString(" ORDER BY distance LIMIT ?, ? ")
		queryParams = append(queryParams, helper.GetInt64FromPointer(params.Offset), helper.GetInt64FromPointer(params.Limit))
	}

	return query, queryParams
}

func querySelectParams(ctx context.Context, query bytes.Buffer, params *model.GetListRequest) (newQuery bytes.Buffer, queryParams []interface{}) {
	if params.Search != nil {
		newQuery.WriteString(fmt.Sprintf(` WHERE (name LIKE LOWER(%s) OR JSON_EXTRACT(phone_numbers, '$[*].phone_number') LIKE %s ) `, "'%"+helper.GetStringFromPointer(params.Search)+"%'", "'%"+helper.GetStringFromPointer(params.Search)+"%'"))
	}

	if params.RegencyID != nil {
		newQuery.WriteString(andWhere(ctx, newQuery, "kabkota_id", "="))
		queryParams = append(queryParams, helper.GetInt64FromPointer(params.RegencyID))
	}

	if params.Name != nil {
		newQuery.WriteString(andWhere(ctx, newQuery, "name", "="))
		queryParams = append(queryParams, helper.GetStringFromPointer(params.Name))
	}

	if params.DistrictID != nil {
		newQuery.WriteString(andWhere(ctx, newQuery, "kec_id", "="))
		queryParams = append(queryParams, helper.GetInt64FromPointer(params.DistrictID))
	}

	if params.VillageID != nil {
		newQuery.WriteString(andWhere(ctx, newQuery, "kel_id", "="))
		queryParams = append(queryParams, helper.GetInt64FromPointer(params.VillageID))
	}

	if params.Status != nil {
		newQuery.WriteString(andWhere(ctx, newQuery, "status", "="))
		queryParams = append(queryParams, helper.GetInt64FromPointer(params.Status))
	}

	return newQuery, queryParams
}

func queryUpdateParams(ctx context.Context, params *model.UpdatePhonebook, queryParams map[string]interface{}) (bytes.Buffer, map[string]interface{}) {
	var query bytes.Buffer
	_, unixTime := helper.GetCurrentTimeUTC()
	if params.Address != nil {
		query.WriteString(" address = :address ")
		queryParams["address"] = helper.GetStringFromPointer(params.Address)
	}
	if params.CategoryID != nil {
		query.WriteString(updateNext(ctx, "category_id"))
		queryParams["category_id"] = helper.GetInt64FromPointer(params.CategoryID)
	}
	if params.CoverImagePath != nil {
		query.WriteString(updateNext(ctx, "cover_image_path"))
		queryParams["cover_image_path"] = helper.GetStringFromPointer(params.CoverImagePath)
	}
	if params.Description != nil {
		query.WriteString(updateNext(ctx, "description"))
		queryParams["description"] = helper.GetStringFromPointer(params.Description)
	}
	if params.DistrictID != nil {
		query.WriteString(updateNext(ctx, "kec_id"))
		queryParams["kec_id"] = helper.GetInt64FromPointer(params.DistrictID)
	}
	if params.Latitude != nil {
		query.WriteString(updateNext(ctx, "latitude"))
		queryParams["latitude"] = helper.GetStringFromPointer(params.Latitude)
	}
	if params.Longitude != nil {
		query.WriteString(updateNext(ctx, "longitude"))
		queryParams["longitude"] = helper.GetStringFromPointer(params.Longitude)
	}
	if params.Name != "" {
		query.WriteString(updateNext(ctx, "name"))
		queryParams["name"] = params.Name
	}
	if params.PhoneNumbers != nil {
		query.WriteString(updateNext(ctx, "phone_numbers"))
		queryParams["phone_numbers"] = helper.GetStringFromPointer(params.PhoneNumbers)
	}
	if params.RegencyID != nil {
		query.WriteString(updateNext(ctx, "kabkota_id"))
		queryParams["kabkota_id"] = helper.GetInt64FromPointer(params.RegencyID)
	}
	if params.Status != nil {
		query.WriteString(updateNext(ctx, "status"))
		queryParams["status"] = helper.GetInt64FromPointer(params.Status)
	}
	if params.VillageID != nil {
		query.WriteString(updateNext(ctx, "kel_id"))
		queryParams["kel_id"] = helper.GetInt64FromPointer(params.VillageID)
	}
	if params.Sequence != nil {
		query.WriteString(updateNext(ctx, "seq"))
		queryParams["seq"] = helper.GetInt64FromPointer(params.Sequence)
	}
	query.WriteString(updateNext(ctx, "updated_at") + " WHERE id = :id ")
	queryParams["updated_at"] = unixTime
	queryParams["id"] = params.ID
	return query, queryParams
}

func andWhere(ctx context.Context, query bytes.Buffer, field string, action string) string {
	qString := query.String()
	query.Reset()
	if strings.Contains(qString, " WHERE ") {
		query.WriteString(fmt.Sprintf(" AND %s %s ? ", field, action))
	} else {
		query.WriteString(fmt.Sprintf(" WHERE %s %s ? ", field, action))
	}
	return query.String()
}

func updateNext(ctx context.Context, field string) string {
	var query bytes.Buffer
	query.WriteString(fmt.Sprintf(" , %s = :%s ", field, field))
	return query.String()
}
