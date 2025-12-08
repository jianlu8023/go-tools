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
