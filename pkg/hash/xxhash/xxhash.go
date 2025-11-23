package xxhash

import (
	"github.com/cespare/xxhash/v2"
)

// Sum64 计算数据的XXHash64哈希值
func Sum64(data []byte) uint64 {
	return xxhash.Sum64(data)
}

// Sum64String 计算字符串的XXHash64哈希值
func Sum64String(s string) uint64 {
	return xxhash.Sum64String(s)
}

// New 创建一个新的XXHash64哈希对象
func New() *xxhash.Digest {
	return xxhash.New()
}

// Digest 是xxhash.Digest的别名，提供更方便的使用
// 可以使用标准的hash.Hash接口方法
// - Write([]byte) (int, error) - 写入数据
// - Sum([]byte) []byte - 计算哈希值并附加到提供的切片
// - Reset() - 重置哈希状态
// - Size() int - 返回哈希大小（8字节）
// - BlockSize() int - 返回块大小（16字节）
type Digest = xxhash.Digest
