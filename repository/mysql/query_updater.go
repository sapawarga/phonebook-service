package mysql

import (
	"bytes"
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/sapawarga/phonebook-service/helper"
	"github.com/sapawarga/phonebook-service/model"
)

func (r *PhonebookRepository) GetListPhonebookByLongLat(ctx context.Context, params *model.GetListRequest) ([]*model.PhoneBookResponse, error) {
	var query bytes.Buffer
	var result []*model.PhoneBookResponse
	var err error

	query.WriteString(` 
	SELECT id, category_id, name, description, address, phone_numbers, kabkota_id , kec_id , kel_id ,status , 
		latitude, longitude, distance, FROM_UNIXTIME(created_at) as created_at, FROM_UNIXTIME(updated_at) as updated_at
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
	var first = true
	if params.Search != nil {
		query.WriteString(isFirstQuery(ctx, first, helper.SELECT_QUERY))
		query.WriteString(fmt.Sprintf(`(name LIKE LOWER(%s) OR phone_numbers LIKE %s ) `, "'%"+helper.GetStringFromPointer(params.Search)+"%'", "'%"+helper.GetStringFromPointer(params.Search)+"%'"))
		first = false
	}

	if params.RegencyID != nil {
		query.WriteString(isFirstQuery(ctx, first, helper.SELECT_QUERY) + " kabkota_id = ? ")
		queryParams = append(queryParams, helper.GetInt64FromPointer(params.RegencyID))
		first = false
	}

	if params.DistrictID != nil {
		query.WriteString(isFirstQuery(ctx, first, helper.SELECT_QUERY) + " kec_id = ? ")
		queryParams = append(queryParams, helper.GetInt64FromPointer(params.DistrictID))
		first = false
	}

	if params.VillageID != nil {
		query.WriteString(isFirstQuery(ctx, first, helper.SELECT_QUERY) + " kel_id = ?")
		queryParams = append(queryParams, helper.GetInt64FromPointer(params.VillageID))
		first = false
	}

	if params.Status != nil {
		query.WriteString(isFirstQuery(ctx, first, helper.SELECT_QUERY) + " status = ?")
		queryParams = append(queryParams, helper.GetInt64FromPointer(params.Status))
	}

	return query, queryParams
}

func queryUpdateParams(ctx context.Context, params *model.UpdatePhonebook, queryParams map[string]interface{}) (bytes.Buffer, map[string]interface{}) {
	var query bytes.Buffer
	var first = true
	_, unixTime := helper.GetCurrentTimeUTC()
	if params.Address != nil {
		query.WriteString(" address = :address ")
		queryParams["address"] = helper.GetStringFromPointer(params.Address)
		first = false
	}
	if params.CategoryID != nil {
		query.WriteString(isFirstQuery(ctx, first, helper.UPDATE_QUERY) + " category_id = :category_id ")
		queryParams["category_id"] = helper.GetInt64FromPointer(params.CategoryID)
		first = false
	}
	if params.CoverImagePath != nil {
		query.WriteString(isFirstQuery(ctx, first, helper.UPDATE_QUERY) + " cover_image_path = :cover_image_path")
		queryParams["cover_image_path"] = helper.GetStringFromPointer(params.CoverImagePath)
		first = false
	}
	if params.Description != nil {
		query.WriteString(isFirstQuery(ctx, first, helper.UPDATE_QUERY) + " description = :description ")
		queryParams["description"] = helper.GetStringFromPointer(params.Description)
		first = false
	}
	if params.DistrictID != nil {
		query.WriteString(isFirstQuery(ctx, first, helper.UPDATE_QUERY) + " kec_id = :kec_id ")
		queryParams["kec_id"] = helper.GetInt64FromPointer(params.DistrictID)
		first = false
	}
	if params.Latitude != nil {
		query.WriteString(isFirstQuery(ctx, first, helper.UPDATE_QUERY) + " latitude = :latitude ")
		queryParams["latitude"] = helper.GetStringFromPointer(params.Latitude)
		first = false
	}
	if params.Longitude != nil {
		query.WriteString(isFirstQuery(ctx, first, helper.UPDATE_QUERY) + " longitude = :longitude ")
		queryParams["longitude"] = helper.GetStringFromPointer(params.Longitude)
		first = false
	}
	if params.Name != "" {
		query.WriteString(isFirstQuery(ctx, first, helper.UPDATE_QUERY) + " name = :name ")
		queryParams["name"] = params.Name
		first = false
	}
	if params.PhoneNumbers != nil {
		query.WriteString(isFirstQuery(ctx, first, helper.UPDATE_QUERY) + " phone_numbers = :phone_numbers")
		queryParams["phone_numbers"] = helper.GetStringFromPointer(params.PhoneNumbers)
		first = false
	}
	if params.RegencyID != nil {
		query.WriteString(isFirstQuery(ctx, first, helper.UPDATE_QUERY) + " kabkota_id = :kabkota_id")
		queryParams["kabkota_id"] = helper.GetInt64FromPointer(params.RegencyID)
		first = false
	}
	if params.Status != nil {
		query.WriteString(isFirstQuery(ctx, first, helper.UPDATE_QUERY) + " status = :status")
		queryParams["status"] = helper.GetInt64FromPointer(params.Status)
		first = false
	}
	if params.VillageID != nil {
		query.WriteString(isFirstQuery(ctx, first, helper.UPDATE_QUERY) + " kel_id = :kel_id")
		queryParams["kel_id"] = helper.GetInt64FromPointer(params.VillageID)
		first = false
	}
	query.WriteString(isFirstQuery(ctx, first, helper.UPDATE_QUERY) + " updated_at = :updated_at WHERE id = :id")
	queryParams["updated_at"] = unixTime
	queryParams["id"] = params.ID
	return query, queryParams
}

func isFirstQuery(ctx context.Context, isFirst bool, queryType string) string {
	var query bytes.Buffer
	if queryType == helper.SELECT_QUERY {
		if isFirst {
			query.WriteString(" WHERE ")
		} else {
			query.WriteString(" AND ")
		}
	} else if queryType == helper.UPDATE_QUERY {
		if !isFirst {
			query.WriteString(" , ")
		}
	}

	return query.String()
}
