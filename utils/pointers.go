package utils

import "time"

// StringPointer returns string pointer
func StringPointer(str string) *string {
	return &str
}

// Uint64Pointer returns uint64 pointer
func Uint64Pointer(num uint64) *uint64 {
	return &num
}

// Int64Pointer returns int64 pointer
func Int64Pointer(num int64) *int64 {
	return &num
}

// IntPointer returns int pointer
func IntPointer(num int) *int {
	return &num
}

// Float64Pointer returns float64 pointer
func Float64Pointer(num float64) *float64 {
	return &num
}

// BoolPointer returns bool pointer
func BoolPointer(val bool) *bool {
	return &val
}

// BoolPointer returns time.Duration pointer
func DurationPointer(d time.Duration) *time.Duration {
	return &d
}
