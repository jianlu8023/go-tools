package cluster

import (
	"crypto/rand"
	"encoding/hex"
)

// GenerateKey 生成集群的密钥
// @return string 密钥
// @return error 错误
func GenerateKey() (string, error) {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	secret := hex.EncodeToString(bytes)
	return secret, nil
}
