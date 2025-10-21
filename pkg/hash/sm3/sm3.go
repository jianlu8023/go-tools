package sm3

import (
	"encoding/hex"

	"github.com/tjfoc/gmsm/sm3"
)

// HashSm3 hash string
// content: 待计算hash内容
// []byte: 返回hash值
// error: 错误信息
func HashSm3(content string) ([]byte, error) {
	hash := sm3.New()
	_, err := hash.Write([]byte(content))
	if err != nil {
		return nil, err
	}
	return hash.Sum(nil), nil
}

func HashSm3Hex(content string) (string, error) {
	h := sm3.New()
	_, err := h.Write([]byte(content))
	if err != nil {
		return "", err
	}
	sum := h.Sum(nil)
	hash := hex.EncodeToString(sum[:])
	return hash, nil
}

func HashSm3HexBytes(content []byte) (string, error) {
	h := sm3.New()
	_, err := h.Write(content)
	if err != nil {
		return "", err
	}
	sum := h.Sum(nil)
	hash := hex.EncodeToString(sum[:])
	return hash, nil
}
