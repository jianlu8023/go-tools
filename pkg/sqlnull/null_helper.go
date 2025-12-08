package sqlnull

import (
	"time"
)

// StringPtr 返回指向字符串的指针
func StringPtr(s string) *string {
	return &s
}

// StringDefaultPtr 返回指向空字符串的指针
func StringDefaultPtr() *string {
	return StringPtr("")
}

// Int64Ptr 返回指向int64的指针
func Int64Ptr(i int64) *int64 {
	return &i
}

// Int64DefaultPtr 返回指向0的int64指针
func Int64DefaultPtr() *int64 {
	return Int64Ptr(0)
}

// TruePtr 返回指向true的bool指针
func TruePtr() *bool {
	return BoolPtr(true)
}

// FalsePtr 返回指向false的bool指针
func FalsePtr() *bool {
	return BoolPtr(false)
}

// BoolPtr 返回指向bool的指针
func BoolPtr(b bool) *bool {
	return &b
}

// BoolDefaultPtr 返回指向false的bool指针
func BoolDefaultPtr() *bool {
	return BoolPtr(false)
}

// Float64Ptr 返回指向float64的指针
func Float64Ptr(f float64) *float64 {
	return &f
}

// Float64DefaultPtr 返回指向0.0的float64指针
func Float64DefaultPtr() *float64 {
	return Float64Ptr(0.0)
}

// Int32Ptr 返回指向int32的指针
func Int32Ptr(i int32) *int32 {
	return &i
}

// Int32DefaultPtr 返回指向0的int32指针
func Int32DefaultPtr() *int32 {
	return Int32Ptr(0)
}

// Int16Ptr 返回指向int16的指针
func Int16Ptr(i int16) *int16 {
	return &i
}

// Int16DefaultPtr 返回指向0的int16指针
func Int16DefaultPtr() *int16 {
	return Int16Ptr(0)
}

// TimePtr 返回指向time.Time的指针
func TimePtr(t time.Time) *time.Time {
	return &t
}

// TimeDefaultPtr 返回指向零值时间的指针
func TimeDefaultPtr() *time.Time {
	return TimePtr(time.Time{})
}

// IntPtr 返回指向int的指针
func IntPtr(i int) *int {
	return &i
}

// IntDefaultPtr 返回指向0的int指针
func IntDefaultPtr() *int {
	return IntPtr(0)
}

// Int8Ptr 返回指向int8的指针
func Int8Ptr(i int8) *int8 {
	return &i
}

// Int8DefaultPtr 返回指向0的int8指针
func Int8DefaultPtr() *int8 {
	return Int8Ptr(0)
}

// UintPtr 返回指向uint的指针
func UintPtr(i uint) *uint {
	return &i
}

// UintDefaultPtr 返回指向0的uint指针
func UintDefaultPtr() *uint {
	return UintPtr(0)
}

// Uint8Ptr 返回指向uint8的指针
func Uint8Ptr(i uint8) *uint8 {
	return &i
}

// Uint8DefaultPtr 返回指向0的uint8指针
func Uint8DefaultPtr() *uint8 {
	return Uint8Ptr(0)
}

// Uint16Ptr 返回指向uint16的指针
func Uint16Ptr(i uint16) *uint16 {
	return &i
}

// Uint16DefaultPtr 返回指向0的uint16指针
func Uint16DefaultPtr() *uint16 {
	return Uint16Ptr(0)
}

// Uint32Ptr 返回指向uint32的指针
func Uint32Ptr(i uint32) *uint32 {
	return &i
}

// Uint32DefaultPtr 返回指向0的uint32指针
func Uint32DefaultPtr() *uint32 {
	return Uint32Ptr(0)
}

// Uint64Ptr 返回指向uint64的指针
func Uint64Ptr(i uint64) *uint64 {
	return &i
}

// Uint64DefaultPtr 返回指向0的uint64指针
func Uint64DefaultPtr() *uint64 {
	return Uint64Ptr(0)
}

// Float32Ptr 返回指向float32的指针
func Float32Ptr(f float32) *float32 {
	return &f
}

// Float32DefaultPtr 返回指向0.0的float32指针
func Float32DefaultPtr() *float32 {
	return Float32Ptr(0.0)
}
