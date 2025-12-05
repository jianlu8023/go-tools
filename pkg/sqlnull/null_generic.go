//go:build go1.18
// +build go1.18

package sqlnull

import (
	"database/sql"
)

// NullPtr 将 sql.Null[T] 转换为 *T。
// 若值为 NULL（Valid=false），返回 nil；否则返回指向其值的指针。
func NullPtr[T any](n sql.Null[T]) *T {
	if !n.Valid {
		return nil
	}
	return &n.V
}

// ToNull 将 *T 转换为 sql.Null[T]。
// 若指针为 nil，返回 Valid=false 的 Null[T]；否则返回对应值。
func ToNull[T any](ptr *T) sql.Null[T] {
	if ptr == nil {
		return sql.Null[T]{Valid: false}
	}
	return sql.Null[T]{V: *ptr, Valid: true}
}
