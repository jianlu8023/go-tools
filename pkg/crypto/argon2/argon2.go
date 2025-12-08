package argon2

import (
	"crypto/rand"
	"fmt"
	"io"

	"golang.org/x/crypto/argon2"
)

const (
	// DefaultParams 默认参数
	DefaultTime    = 1
	DefaultMemory  = 64 * 1024
	DefaultThreads = 4
	DefaultKeyLen  = 32
)

// Params Argon2参数
type Params struct {
	Time    uint32 // 时间复杂度
	Memory  uint32 // 内存复杂度(KiB)
	Threads uint8  // 并行度
	KeyLen  uint32 // 密钥长度
}

// DefaultParams 返回默认参数
func DefaultParams() Params {
	return Params{
		Time:    DefaultTime,
		Memory:  DefaultMemory,
		Threads: DefaultThreads,
		KeyLen:  DefaultKeyLen,
	}
}

// GenerateSalt 生成随机盐值
// 返回16字节盐值和错误信息
func GenerateSalt() ([]byte, error) {
	salt := make([]byte, 16)
	if _, err := io.ReadFull(rand.Reader, salt); err != nil {
		return nil, fmt.Errorf("生成盐值失败: %v", err)
	}
	return salt, nil
}

// Hash 使用Argon2id算法哈希密码
// password: 密码
// salt: 盐值
// params: 参数
// 返回哈希值和错误信息
func Hash(password, salt []byte, params Params) ([]byte, error) {
	// 验证参数
	if len(password) == 0 {
		return nil, fmt.Errorf("密码不能为空")
	}

	if len(salt) == 0 {
		return nil, fmt.Errorf("盐值不能为空")
	}

	if params.Time == 0 {
		return nil, fmt.Errorf("时间参数必须大于0")
	}

	if params.Memory == 0 {
		return nil, fmt.Errorf("内存参数必须大于0")
	}

	if params.Threads == 0 {
		return nil, fmt.Errorf("线程参数必须大于0")
	}

	if params.KeyLen == 0 {
		return nil, fmt.Errorf("密钥长度参数必须大于0")
	}

	// 使用Argon2id算法哈希密码
	hash := argon2.IDKey(password, salt, params.Time, params.Memory, params.Threads, params.KeyLen)

	return hash, nil
}

// HashPassword 使用Argon2id算法哈希密码（包含盐值生成）
// password: 密码
// params: 参数
// 返回哈希值、盐值和错误信息
func HashPassword(password []byte, params Params) ([]byte, []byte, error) {
	// 验证参数
	if len(password) == 0 {
		return nil, nil, fmt.Errorf("密码不能为空")
	}

	// 生成随机盐值
	salt, err := GenerateSalt()
	if err != nil {
		return nil, nil, fmt.Errorf("生成盐值失败: %v", err)
	}

	// 哈希密码
	hash, err := Hash(password, salt, params)
	if err != nil {
		return nil, nil, fmt.Errorf("哈希密码失败: %v", err)
	}

	return hash, salt, nil
}

// Verify 验证密码与哈希值是否匹配
// password: 密码
// hash: 哈希值
// salt: 盐值
// params: 参数
// 返回验证结果和错误信息
func Verify(password, hash, salt []byte, params Params) (bool, error) {
	// 验证参数
	if len(password) == 0 {
		return false, fmt.Errorf("密码不能为空")
	}

	if len(hash) == 0 {
		return false, fmt.Errorf("哈希值不能为空")
	}

	if len(salt) == 0 {
		return false, fmt.Errorf("盐值不能为空")
	}

	// 重新哈希密码
	computedHash, err := Hash(password, salt, params)
	if err != nil {
		return false, fmt.Errorf("哈希密码失败: %v", err)
	}

	// 比较哈希值
	if len(computedHash) != len(hash) {
		return false, nil
	}

	for i := range computedHash {
		if computedHash[i] != hash[i] {
			return false, nil
		}
	}

	return true, nil
}

// HashText 使用Argon2id算法哈希文本密码
// password: 密码文本
// salt: 盐值
// params: 参数
// 返回哈希值和错误信息
func HashText(password string, salt []byte, params Params) ([]byte, error) {
	return Hash([]byte(password), salt, params)
}

// HashTextPassword 使用Argon2id算法哈希文本密码（包含盐值生成）
// password: 密码文本
// params: 参数
// 返回哈希值、盐值和错误信息
func HashTextPassword(password string, params Params) ([]byte, []byte, error) {
	return HashPassword([]byte(password), params)
}

// VerifyText 验证文本密码与哈希值是否匹配
// password: 密码文本
// hash: 哈希值
// salt: 盐值
// params: 参数
// 返回验证结果和错误信息
func VerifyText(password string, hash, salt []byte, params Params) (bool, error) {
	return Verify([]byte(password), hash, salt, params)
}
