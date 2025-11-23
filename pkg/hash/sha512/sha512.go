package sha512

import (
	"crypto/sha512"
	"encoding/hex"
	"hash"
)

// Sum 计算数据的SHA512哈希值
func Sum(data []byte) []byte {
	hash := sha512.Sum512(data)
	return hash[:]
}

// SumHex 计算数据的SHA512哈希值并返回十六进制字符串
func SumHex(data []byte) string {
	hash := sha512.Sum512(data)
	return hex.EncodeToString(hash[:])
}

// SumString 计算字符串的SHA512哈希值
func SumString(s string) []byte {
	hash := sha512.Sum512([]byte(s))
	return hash[:]
}

// SumStringHex 计算字符串的SHA512哈希值并返回十六进制字符串
func SumStringHex(s string) string {
	return SumHex([]byte(s))
}

// New 创建一个新的SHA512哈希对象
func New() hash.Hash {
	return sha512.New()
}

// Digest 是hash.Hash接口的别名，提供更方便的使用
// 可以使用标准的hash.Hash接口方法
// - Write([]byte) (int, error) - 写入数据
// - Sum([]byte) []byte - 计算哈希值并附加到提供的切片
// - Reset() - 重置哈希状态
// - Size() int - 返回哈希大小（64字节）
// - BlockSize() int - 返回块大小（128字节）
type Digest = hash.Hash

// Sum512_224 计算数据的SHA-512/224哈希值
func Sum512_224(data []byte) []byte {
	hash := sha512.Sum512_224(data)
	return hash[:]
}

// Sum512_224Hex 计算数据的SHA-512/224哈希值并返回十六进制字符串
func Sum512_224Hex(data []byte) string {
	hash := sha512.Sum512_224(data)
	return hex.EncodeToString(hash[:])
}

// New512_224 创建一个新的SHA-512/224哈希对象
func New512_224() hash.Hash {
	return sha512.New512_224()
}

// Sum512_256 计算数据的SHA-512/256哈希值
func Sum512_256(data []byte) []byte {
	hash := sha512.Sum512_256(data)
	return hash[:]
}

// Sum512_256Hex 计算数据的SHA-512/256哈希值并返回十六进制字符串
func Sum512_256Hex(data []byte) string {
	hash := sha512.Sum512_256(data)
	return hex.EncodeToString(hash[:])
}

// New512_256 创建一个新的SHA-512/256哈希对象
func New512_256() hash.Hash {
	return sha512.New512_256()
}
