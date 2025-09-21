package base64

import (
	"encoding/base64"
)

// ToBase64
// @param data:
// @return string:
func ToBase64(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

// ToByte
// @param data:
// @return []byte:
// @return error:
func ToByte(data string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(data)
}
