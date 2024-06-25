package str_to_int64

import (
	"strconv"
)

// ToInt64 字符串转int64
// @param str string
// @return int64
// @example ToInt64("123") => 123
func ToInt64(str string) int64 {
	parse, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		panic(err)
	}
	return parse
}
