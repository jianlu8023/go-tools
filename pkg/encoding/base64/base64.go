package base64

import (
	"encoding/base64"
)

// ByteToBase64
// @Description: ByteToBase64
// @author ght
// @date 2023-10-07 16:37:51
// @param data:
// @return string:
func ByteToBase64(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

// Base64ToByte
// @Description: Base64ToByte
// @author ght
// @date 2023-10-07 16:37:54
// @param data:
// @return []byte:
// @return error:
func Base64ToByte(data string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(data)
}
