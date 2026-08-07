package base64

import (
	"encoding/base64"
)

// ToBase64 将字节数组编码为标准 base64 字符串
// @param data 输入字节数组
// @return string base64 编码后的字符串
func ToBase64(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

// ToByte 将 base64 字符串解码为字节数组
// @param data base64 编码的字符串
// @return []byte 解码后的字节数组
// @return error 解码过程中的错误
func ToByte(data string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(data)
}

// ToBase64URL 编码为 URL 安全的 base64 字符串
// @param data 输入字节数组
// @return string URL 安全的 base64 编码字符串
func ToBase64URL(data []byte) string {
	return base64.URLEncoding.EncodeToString(data)
}

// ToByteURL 解码 URL 安全的 base64 字符串
// @param data URL 安全的 base64 编码字符串
// @return []byte 解码后的字节数组
// @return error 解码过程中的错误
func ToByteURL(data string) ([]byte, error) {
	return base64.URLEncoding.DecodeString(data)
}

// ToBase64NoPadding 编码为不带填充的标准 base64 字符串
// @param data 输入字节数组
// @return string 不带填充的 base64 编码字符串
func ToBase64NoPadding(data []byte) string {
	return base64.RawStdEncoding.EncodeToString(data)
}

// ToByteNoPadding 解码不带填充的标准 base64 字符串
// @param data 不带填充的 base64 编码字符串
// @return []byte 解码后的字节数组
// @return error 解码过程中的错误
func ToByteNoPadding(data string) ([]byte, error) {
	return base64.RawStdEncoding.DecodeString(data)
}

// ToBase64URLNoPadding 编码为不带填充的 URL 安全 base64 字符串
// @param data 输入字节数组
// @return string 不带填充的 URL 安全 base64 编码字符串
func ToBase64URLNoPadding(data []byte) string {
	return base64.RawURLEncoding.EncodeToString(data)
}

// ToByteURLNoPadding 解码不带填充的 URL 安全 base64 字符串
// @param data 不带填充的 URL 安全 base64 编码字符串
// @return []byte 解码后的字节数组
// @return error 解码过程中的错误
func ToByteURLNoPadding(data string) ([]byte, error) {
	return base64.RawURLEncoding.DecodeString(data)
}
