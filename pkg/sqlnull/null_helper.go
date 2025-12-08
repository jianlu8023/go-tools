package sqlnull

import (
	"time"
)

func StringPtr(s string) *string {
	return &s
}

func Int64Ptr(i int64) *int64 {
	return &i
}

func BoolPtr(b bool) *bool {
	return &b
}

func Float64Ptr(f float64) *float64 {
	return &f
}

func Int32Ptr(i int32) *int32 {
	return &i
}

func Int16Ptr(i int16) *int16 {
	return &i
}

func TimePtr(t time.Time) *time.Time {
	return &t
}
