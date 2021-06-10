package helper

import "errors"

const (
	DELETED          int64   = -1
	INACTIVED        int64   = 0
	ACTIVED          int64   = 10
	DELETED_STRING           = "deleted"
	INACTIVED_STRING         = "inactived"
	ACTIVED_STRING           = "actived"
	SELECT_QUERY             = "select"
	UPDATE_QUERY             = "update"
	HTTP_GET                 = "GET"
	HTTP_PUT                 = "PUT"
	HTTP_POST                = "POST"
	HTTP_DELETE              = "DELETE"
	STATUS_OK                = "status_ok"
	STATUS_CREATED           = "status_created"
	STATUS_UPDATED           = "status_updated"
	STATUS_DELETED           = "status_deleted"
	RADIUS           float64 = 3.0
)

// StatusEnum ...
var StatusEnum = map[int64]string{
	DELETED:   DELETED_STRING,
	INACTIVED: INACTIVED_STRING,
	ACTIVED:   ACTIVED_STRING,
}

// GetStatusEnum ...
func GetStatusEnum(status string) (*int64, error) {
	for i, v := range StatusEnum {
		if v == status {
			return &i, nil
		}
	}

	return nil, errors.New("status_not_registered")
}
