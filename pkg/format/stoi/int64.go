package stoi

import (
	"strconv"
)

// Int64 字符串转int64
// @param str string
// @return int64
// @example ToInt64("123") => 123
func Int64(str string) int64 {
	parse, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		panic(err)
	}
	return parse
}
