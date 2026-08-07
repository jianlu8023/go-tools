package pbkdf2

import (
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"crypto/subtle"
	"fmt"
	"hash"
	"io"

	"golang.org/x/crypto/pbkdf2"
)

const (
	DefaultIterations = 210000
	DefaultKeyLen     = 32
)

type HashFunc func() hash.Hash

var (
	SHA256 HashFunc = sha256.New
	SHA512 HashFunc = sha512.New
)

type Params struct {
	Iterations int
	KeyLen     int
	HashFunc   HashFunc
}

func DefaultParams() Params {
	return Params{
		Iterations: DefaultIterations,
		KeyLen:     DefaultKeyLen,
		HashFunc:   SHA256,
	}
}

func GenerateSalt() ([]byte, error) {
	salt := make([]byte, 16)
	if _, err := io.ReadFull(rand.Reader, salt); err != nil {
		return nil, fmt.Errorf("生成盐值失败: %v", err)
	}
	return salt, nil
}

func DeriveKey(password, salt []byte, params Params) ([]byte, error) {
	if len(password) == 0 {
		return nil, fmt.Errorf("密码不能为空")
	}

	if len(salt) == 0 {
		return nil, fmt.Errorf("盐值不能为空")
	}

	if params.Iterations <= 0 {
		return nil, fmt.Errorf("迭代次数必须大于0")
	}

	if params.KeyLen <= 0 {
		return nil, fmt.Errorf("密钥长度必须大于0")
	}

	if params.HashFunc == nil {
		return nil, fmt.Errorf("哈希函数不能为空")
	}

	key := pbkdf2.Key(password, salt, params.Iterations, params.KeyLen, params.HashFunc)
	return key, nil
}

// Verify 验证密码是否与已派生的密钥匹配
// password: 密码
// salt: 盐值
// params: 参数
// expectedKey: 预期密钥
// 返回验证结果和错误信息
func Verify(password, salt, expectedKey []byte, params Params) (bool, error) {
	if len(expectedKey) == 0 {
		return false, fmt.Errorf("预期密钥不能为空")
	}

	derivedKey, err := DeriveKey(password, salt, params)
	if err != nil {
		return false, err
	}

	return subtle.ConstantTimeCompare(derivedKey, expectedKey) == 1, nil
}

func DeriveKeyWithSalt(password []byte, params Params) ([]byte, []byte, error) {
	if len(password) == 0 {
		return nil, nil, fmt.Errorf("密码不能为空")
	}

	salt, err := GenerateSalt()
	if err != nil {
		return nil, nil, fmt.Errorf("生成盐值失败: %v", err)
	}

	key, err := DeriveKey(password, salt, params)
	if err != nil {
		return nil, nil, fmt.Errorf("派生密钥失败: %v", err)
	}

	return key, salt, nil
}

func DeriveKeyText(password string, salt []byte, params Params) ([]byte, error) {
	return DeriveKey([]byte(password), salt, params)
}

func DeriveKeyTextWithSalt(password string, params Params) ([]byte, []byte, error) {
	return DeriveKeyWithSalt([]byte(password), params)
}
