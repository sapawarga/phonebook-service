package helper

import "time"

// SetPointerString ...
func SetPointerString(val string) *string {
	return &val
}

// SetPointerInt64 ...
func SetPointerInt64(val int64) *int64 {
	return &val
}

// SetPointerTime ...
func SetPointerTime(val time.Time) *time.Time {
	return &val
}

// GetStringFromPointer ...
func GetStringFromPointer(val *string) string {
	return *val
}

// GetInt64FromPointer ...
func GetInt64FromPointer(val *int64) int64 {
	return *val
}

// GetTimeFromPointer ...
func GetTimeFromPointer(val *time.Time) time.Time {
	return *val
}
