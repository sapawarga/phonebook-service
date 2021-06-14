package helper

import (
	"strconv"
	"time"
)

// SetPointerString ...
func SetPointerString(val string) *string {
	if val == "" {
		return nil
	}
	return &val
}

// SetPointerInt64 ...
func SetPointerInt64(val int64) *int64 {
	if val == 0 {
		return nil
	}
	return &val
}

// SetPointerTime ...
func SetPointerTime(val time.Time) *time.Time {
	return &val
}

// GetStringFromPointer ...
func GetStringFromPointer(val *string) string {
	if val == nil {
		return ""
	}
	return *val
}

// GetInt64FromPointer ...
func GetInt64FromPointer(val *int64) int64 {
	if val == nil {
		return 0
	}
	return *val
}

// GetTimeFromPointer ...
func GetTimeFromPointer(val *time.Time) time.Time {
	return *val
}

// GetCurrentTimeUTC ...
func GetCurrentTimeUTC() (standartTime time.Time, unixTime int64) {
	current := time.Now().UTC()
	return current, current.Unix()
}

func ConvertFromStringToInt64(val string) (pointerInt64 *int64, plainInt64 int64) {
	if val == "" {
		return nil, 0
	}
	valInt, _ := strconv.ParseInt(val, 10, 64)
	return &valInt, valInt
}
