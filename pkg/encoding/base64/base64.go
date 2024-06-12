package base64

import (
	"encoding/base64"
)

// ByteToBase64
// @param data:
// @return string:
func ByteToBase64(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

// Base64ToByte
// @param data:
// @return []byte:
// @return error:
func Base64ToByte(data string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(data)
}
