package md5

import (
	"crypto/md5"
	"encoding/hex"
)

// MD5 返回str的md5值
// @param str string
// @return string
func MD5(str string) string {
	hash := md5.Sum([]byte(str))
	return hex.EncodeToString(hash[:])
}
