package sha256

import (
	"crypto/sha256"
	"encoding/hex"
	"hash"
)

// Sum 计算数据的SHA256哈希值
func Sum(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[:]
}

// SumHex 计算数据的SHA256哈希值并返回十六进制字符串
func SumHex(data []byte) string {
	hash := sha256.Sum256(data)
	return hex.EncodeToString(hash[:])
}

// SumString 计算字符串的SHA256哈希值
func SumString(s string) []byte {
	hash := sha256.Sum256([]byte(s))
	return hash[:]
}

// SumStringHex 计算字符串的SHA256哈希值并返回十六进制字符串
func SumStringHex(s string) string {
	return SumHex([]byte(s))
}

// New 创建一个新的SHA256哈希对象
func New() hash.Hash {
	return sha256.New()
}

// Digest 是hash.Hash接口的别名，提供更方便的使用
// 可以使用标准的hash.Hash接口方法
// - Write([]byte) (int, error) - 写入数据
// - Sum([]byte) []byte - 计算哈希值并附加到提供的切片
// - Reset() - 重置哈希状态
// - Size() int - 返回哈希大小（32字节）
// - BlockSize() int - 返回块大小（64字节）
type Digest = hash.Hash
