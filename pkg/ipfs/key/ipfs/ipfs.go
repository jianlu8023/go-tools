package ipfs

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

// GenerateKey 生成ipfs的私钥
// @return string 私钥
// @return error 错误信息
func GenerateKey() (string, error) {
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		return "", err
	}
	var str string
	str += fmt.Sprintln("/key/swarm/psk/1.0.0/")
	str += fmt.Sprintln("/base16/")
	str += fmt.Sprintln(hex.EncodeToString(key))
	return str, nil
}
