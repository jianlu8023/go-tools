package sm3

import (
	"encoding/hex"
	"hash"

	"github.com/tjfoc/gmsm/sm3"
)

// Sum 计算数据的SM3哈希值
func Sum(data []byte) []byte {
	h := sm3.New()
	h.Write(data)
	return h.Sum(nil)
}

// SumHex 计算数据的SM3哈希值并返回十六进制字符串
func SumHex(data []byte) string {
	return hex.EncodeToString(Sum(data))
}

// SumString 计算字符串的SM3哈希值
func SumString(s string) []byte {
	return Sum([]byte(s))
}

// SumStringHex 计算字符串的SM3哈希值并返回十六进制字符串
func SumStringHex(s string) string {
	return SumHex([]byte(s))
}

// New 创建一个新的SM3哈希对象
func New() hash.Hash {
	return sm3.New()
}

// Digest 是hash.Hash接口的别名，提供更方便的使用
// 可以使用标准的hash.Hash接口方法
// - Write([]byte) (int, error) - 写入数据
// - Sum([]byte) []byte - 计算哈希值并附加到提供的切片
// - Reset() - 重置哈希状态
// - Size() int - 返回哈希大小（32字节）
// - BlockSize() int - 返回块大小（64字节）
type Digest = hash.Hash

// HashSm3 hash string
// content: 待计算hash内容
// []byte: 返回hash值
// error: 错误信息
// Deprecated: 请使用 SumString 替代
func HashSm3(content string) ([]byte, error) {
	return SumString(content), nil
}

// HashSm3Hex 计算字符串的SM3哈希值并返回十六进制字符串
// Deprecated: 请使用 SumStringHex 替代
func HashSm3Hex(content string) (string, error) {
	return SumStringHex(content), nil
}

// HashSm3HexBytes 计算字节切片的SM3哈希值并返回十六进制字符串
// Deprecated: 请使用 SumHex 替代
func HashSm3HexBytes(content []byte) (string, error) {
	return SumHex(content), nil
}
