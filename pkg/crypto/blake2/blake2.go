package blake2

import (
	"fmt"
	"hash"
	"io"
	"os"

	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/blake2s"
)

// HashFunc 哈希函数类型
type HashFunc func() (hash.Hash, error)

// 支持的哈希函数
var (
	Blake2b_256 HashFunc = func() (hash.Hash, error) { return blake2b.New256(nil) }
	Blake2b_384 HashFunc = func() (hash.Hash, error) { return blake2b.New384(nil) }
	Blake2b_512 HashFunc = func() (hash.Hash, error) { return blake2b.New512(nil) }
	Blake2s_256 HashFunc = func() (hash.Hash, error) { return blake2s.New256(nil) }
)

// Hash 计算数据的Blake2哈希值
// data: 数据
// hashFunc: 哈希函数
// 返回哈希值和错误信息
func Hash(data []byte, hashFunc HashFunc) ([]byte, error) {
	// 验证参数
	if hashFunc == nil {
		return nil, fmt.Errorf("哈希函数不能为空")
	}

	// 创建哈希实例
	h, err := hashFunc()
	if err != nil {
		return nil, fmt.Errorf("创建哈希实例失败: %v", err)
	}

	// 写入数据
	_, err = h.Write(data)
	if err != nil {
		return nil, fmt.Errorf("写入数据失败: %v", err)
	}

	// 计算哈希值
	hashValue := h.Sum(nil)

	return hashValue, nil
}

// HashText 计算文本的Blake2哈希值
// text: 文本
// hashFunc: 哈希函数
// 返回哈希值和错误信息
func HashText(text string, hashFunc HashFunc) ([]byte, error) {
	return Hash([]byte(text), hashFunc)
}

// HashFile 计算文件的Blake2哈希值
// filename: 文件名
// hashFunc: 哈希函数
// 返回哈希值和错误信息
func HashFile(filename string, hashFunc HashFunc) ([]byte, error) {
	// 验证参数
	if filename == "" {
		return nil, fmt.Errorf("文件名不能为空")
	}

	if hashFunc == nil {
		return nil, fmt.Errorf("哈希函数不能为空")
	}

	// 打开文件
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("打开文件失败: %v", err)
	}
	defer file.Close()

	// 创建哈希实例
	h, err := hashFunc()
	if err != nil {
		return nil, fmt.Errorf("创建哈希实例失败: %v", err)
	}

	// 读取文件内容并计算哈希值
	_, err = io.Copy(h, file)
	if err != nil {
		return nil, fmt.Errorf("读取文件失败: %v", err)
	}

	// 计算哈希值
	hashValue := h.Sum(nil)

	return hashValue, nil
}

// HashBlake2b256 计算数据的Blake2b-256哈希值
// data: 数据
// 返回哈希值和错误信息
func HashBlake2b256(data []byte) ([]byte, error) {
	return Hash(data, Blake2b_256)
}

// HashTextBlake2b256 计算文本的Blake2b-256哈希值
// text: 文本
// 返回哈希值和错误信息
func HashTextBlake2b256(text string) ([]byte, error) {
	return HashText(text, Blake2b_256)
}

// HashFileBlake2b256 计算文件的Blake2b-256哈希值
// filename: 文件名
// 返回哈希值和错误信息
func HashFileBlake2b256(filename string) ([]byte, error) {
	return HashFile(filename, Blake2b_256)
}

// HashBlake2b512 计算数据的Blake2b-512哈希值
// data: 数据
// 返回哈希值和错误信息
func HashBlake2b512(data []byte) ([]byte, error) {
	return Hash(data, Blake2b_512)
}

// HashTextBlake2b512 计算文本的Blake2b-512哈希值
// text: 文本
// 返回哈希值和错误信息
func HashTextBlake2b512(text string) ([]byte, error) {
	return HashText(text, Blake2b_512)
}

// HashFileBlake2b512 计算文件的Blake2b-512哈希值
// filename: 文件名
// 返回哈希值和错误信息
func HashFileBlake2b512(filename string) ([]byte, error) {
	return HashFile(filename, Blake2b_512)
}
