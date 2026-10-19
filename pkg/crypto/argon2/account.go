package argon2

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
	"unicode/utf8"

	"golang.org/x/crypto/argon2"
)

const (
	// MinAccountPasswordLength 账户密码最小长度（按字符数）
	MinAccountPasswordLength = 8
	// MaxAccountPasswordLength 账户密码最大长度（按字符数）
	MaxAccountPasswordLength = 128

	accountPasswordMemory    = 19 * 1024 // 19MiB
	accountPasswordTime      = 2
	accountPasswordThreads   = 1
	accountPasswordSaltBytes = 16
	accountPasswordKeyBytes  = 32
)

var (
	// ErrAccountPasswordLength 密码长度不在 8-128 字符范围
	ErrAccountPasswordLength = errors.New("password must contain between 8 and 128 characters")
)

// ValidateAccountPassword 校验新设账户密码（8-128 字符，合法 UTF-8，不做空白裁剪）
func ValidateAccountPassword(password string) error {
	if !utf8.ValidString(password) || utf8.RuneCountInString(password) < MinAccountPasswordLength || utf8.RuneCountInString(password) > MaxAccountPasswordLength {
		return ErrAccountPasswordLength
	}
	return nil
}

// HashAccountPassword 使用 Argon2id（m=19456,t=2,p=1,16B盐,32B密钥）哈希账户密码，
// 返回 PHC 格式字符串 $argon2id$v=19$m=19456,t=2,p=1$<salt>$<key>（RawStd base64）
func HashAccountPassword(password string) (string, error) {
	if err := ValidateAccountPassword(password); err != nil {
		return "", err
	}
	salt := make([]byte, accountPasswordSaltBytes)
	if _, err := rand.Read(salt); err != nil {
		return "", fmt.Errorf("generate account password salt: %w", err)
	}
	key := idKey(password, salt)
	return fmt.Sprintf("$argon2id$v=19$m=19456,t=2,p=1$%s$%s",
		base64.RawStdEncoding.EncodeToString(salt),
		base64.RawStdEncoding.EncodeToString(key)), nil
}

// VerifyAccountPassword 校验密码与 HashAccountPassword 产生的 PHC 格式哈希（常量时间比较）
func VerifyAccountPassword(password, encoded string) bool {
	// 限制明文与参数长度，避免触发内存型 KDF 前被异常输入拖垮
	if len(password) > MaxAccountPasswordLength*utf8.UTFMax || len(encoded) > 256 {
		return false
	}
	parts := strings.Split(encoded, "$")
	if len(parts) != 6 || parts[0] != "" || parts[1] != "argon2id" || parts[2] != "v=19" || parts[3] != "m=19456,t=2,p=1" {
		return false
	}
	salt, err := base64.RawStdEncoding.Strict().DecodeString(parts[4])
	if err != nil || len(salt) != accountPasswordSaltBytes {
		return false
	}
	expected, err := base64.RawStdEncoding.Strict().DecodeString(parts[5])
	if err != nil || len(expected) != accountPasswordKeyBytes {
		return false
	}
	actual := idKey(password, salt)
	return subtle.ConstantTimeCompare(actual, expected) == 1
}

// idKey 以固定账户参数执行 Argon2id
func idKey(password string, salt []byte) []byte {
	return argon2.IDKey([]byte(password), salt, accountPasswordTime, accountPasswordMemory, accountPasswordThreads, accountPasswordKeyBytes)
}
