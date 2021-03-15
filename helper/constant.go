package helper

import "errors"

const (
	DELETED          int64 = -1
	INACTIVED        int64 = 0
	ACTIVED          int64 = 10
	DELETED_STRING         = "deleted"
	INACTIVED_STRING       = "inactived"
	ACTIVED_STRING         = "actived"
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
