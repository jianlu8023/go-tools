package sqlnull

import (
	"database/sql"
	"time"
)

// NullStringPtr 将 sql.NullString 转换为 *string。
func NullStringPtr(ns sql.NullString) *string {
	if !ns.Valid {
		return nil
	}
	return &ns.String
}

// StringToNull 将 *string 转换为 sql.NullString。
func StringToNull(s *string) sql.NullString {
	if s == nil {
		return sql.NullString{Valid: false}
	}
	return sql.NullString{String: *s, Valid: true}
}

// StringDefaultToNull 返回一个空字符串的sql.NullString
func StringDefaultToNull() sql.NullString {
	return sql.NullString{String: "", Valid: true}
}

// NullInt64Ptr 将 sql.NullInt64 转换为 *int64。
func NullInt64Ptr(ni sql.NullInt64) *int64 {
	if !ni.Valid {
		return nil
	}
	return &ni.Int64
}

// Int64ToNull 将 *int64 转换为 sql.NullInt64。
func Int64ToNull(i *int64) sql.NullInt64 {
	if i == nil {
		return sql.NullInt64{Valid: false}
	}
	return sql.NullInt64{Int64: *i, Valid: true}
}

// Int64DefaultToNull 返回一个值为0的sql.NullInt64
func Int64DefaultToNull() sql.NullInt64 {
	return sql.NullInt64{Int64: 0, Valid: true}
}

// NullInt32Ptr 将 sql.NullInt32 转换为 *int32。
func NullInt32Ptr(ni sql.NullInt32) *int32 {
	if !ni.Valid {
		return nil
	}
	return &ni.Int32
}

// Int32ToNull 将 *int32 转换为 sql.NullInt32。
func Int32ToNull(i *int32) sql.NullInt32 {
	if i == nil {
		return sql.NullInt32{Valid: false}
	}
	return sql.NullInt32{Int32: *i, Valid: true}
}

// Int32DefaultToNull 返回一个值为0的sql.NullInt32
func Int32DefaultToNull() sql.NullInt32 {
	return sql.NullInt32{Int32: 0, Valid: true}
}

// NullInt16Ptr 将 sql.NullInt16 转换为 *int16。
func NullInt16Ptr(ni sql.NullInt16) *int16 {
	if !ni.Valid {
		return nil
	}
	return &ni.Int16
}

// Int16ToNull 将 *int16 转换为 sql.NullInt16。
func Int16ToNull(i *int16) sql.NullInt16 {
	if i == nil {
		return sql.NullInt16{Valid: false}
	}
	return sql.NullInt16{Int16: *i, Valid: true}
}

// Int16DefaultToNull 返回一个值为0的sql.NullInt16
func Int16DefaultToNull() sql.NullInt16 {
	return sql.NullInt16{Int16: 0, Valid: true}
}

// NullBytePtr 将 sql.NullByte 转换为 *byte。
func NullBytePtr(ni sql.NullByte) *byte {
	if !ni.Valid {
		return nil
	}
	return &ni.Byte
}

// ByteToNull 将 *byte 转换为 sql.NullByte。
func ByteToNull(i *byte) sql.NullByte {
	if i == nil {
		return sql.NullByte{Valid: false}
	}
	return sql.NullByte{Byte: *i, Valid: true}
}

// ByteDefaultToNull 返回一个值为0的sql.NullByte
func ByteDefaultToNull() sql.NullByte {
	return sql.NullByte{Byte: 0, Valid: true}
}

// NullFloat64Ptr 将 sql.NullFloat64 转换为 *float64。
func NullFloat64Ptr(nf sql.NullFloat64) *float64 {
	if !nf.Valid {
		return nil
	}
	return &nf.Float64
}

// Float64ToNull 将 *float64 转换为 sql.NullFloat64。
func Float64ToNull(f *float64) sql.NullFloat64 {
	if f == nil {
		return sql.NullFloat64{Valid: false}
	}
	return sql.NullFloat64{Float64: *f, Valid: true}
}

// Float64DefaultToNull 返回一个值为0.0的sql.NullFloat64
func Float64DefaultToNull() sql.NullFloat64 {
	return sql.NullFloat64{Float64: 0.0, Valid: true}
}

// NullBoolPtr 将 sql.NullBool 转换为 *bool。
func NullBoolPtr(nb sql.NullBool) *bool {
	if !nb.Valid {
		return nil
	}
	return &nb.Bool
}

// BoolToNull 将 *bool 转换为 sql.NullBool。
func BoolToNull(b *bool) sql.NullBool {
	if b == nil {
		return sql.NullBool{Valid: false}
	}
	return sql.NullBool{Bool: *b, Valid: true}
}

// TrueToNull 返回一个值为true的sql.NullBool
func TrueToNull() sql.NullBool {
	return sql.NullBool{Bool: true, Valid: true}
}

// FalseToNull 返回一个值为false的sql.NullBool
func FalseToNull() sql.NullBool {
	return sql.NullBool{Bool: false, Valid: true}
}

// BoolDefaultToNull 返回一个值为false的sql.NullBool
func BoolDefaultToNull() sql.NullBool {
	return sql.NullBool{Bool: false, Valid: true}
}

// NullTimePtr 将 sql.NullTime 转换为 *time.Time。
func NullTimePtr(nt sql.NullTime) *time.Time {
	if !nt.Valid {
		return nil
	}
	return &nt.Time
}

// TimeToNull 将 *time.Time 转换为 sql.NullTime。
func TimeToNull(t *time.Time) sql.NullTime {
	if t == nil {
		return sql.NullTime{Valid: false}
	}
	return sql.NullTime{Time: *t, Valid: true}
}

// TimeDefaultToNull 返回一个零值时间的sql.NullTime
func TimeDefaultToNull() sql.NullTime {
	return sql.NullTime{Time: time.Time{}, Valid: true}
}
