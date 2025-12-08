package rand

import (
	crand "crypto/rand"
	"encoding/binary"
	"errors"
	"math/big"
)

// GenerateCryptoRandomString generates a random string for cryptographic usage.
// If runes is empty, returns an error.
func GenerateCryptoRandomString(n int, runes string) (string, error) {
	letters := []rune(runes)
	// 处理空字符集或长度为0的情况
	if len(letters) == 0 {
		return "", errors.New("empty runes not allowed")
	}
	if n == 0 {
		return "", nil
	}
	b := make([]rune, n)
	for i := range b {
		v, err := crand.Int(crand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		b[i] = letters[v.Int64()]
	}
	return string(b), nil
}

// CryptoUint64 returns cryptographic random uint64.
func CryptoUint64() (uint64, error) {
	var v uint64
	if err := binary.Read(crand.Reader, binary.LittleEndian, &v); err != nil {
		return 0, err
	}
	return v, nil
}

// GenerateCryptoRandomBytes generates a slice of random bytes for cryptographic usage.
func GenerateCryptoRandomBytes(n int) ([]byte, error) {
	if n < 0 {
		return nil, errors.New("negative length not allowed")
	}
	b := make([]byte, n)
	_, err := crand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GenerateCryptoRandomInt generates a random integer in the range [min, max) for cryptographic usage.
func GenerateCryptoRandomInt(min, max int) (int, error) {
	if min >= max {
		return 0, errors.New("min must be less than max")
	}

	rangeSize := big.NewInt(int64(max - min))
	n, err := crand.Int(crand.Reader, rangeSize)
	if err != nil {
		return 0, err
	}

	return min + int(n.Int64()), nil
}

// GenerateCryptoRandomFloat64 generates a random float64 in the range [0.0, 1.0) for cryptographic usage.
func GenerateCryptoRandomFloat64() (float64, error) {
	// 使用两个uint32来生成一个float64，确保精度
	var buf [8]byte
	_, err := crand.Read(buf[:])
	if err != nil {
		return 0, err
	}

	// 将前53位用于生成[0.0, 1.0)范围的浮点数
	// 53位是IEEE 754 double-precision浮点数的尾数长度
	u := uint64(buf[0])<<56 | uint64(buf[1])<<48 | uint64(buf[2])<<40 | uint64(buf[3])<<32 |
		uint64(buf[4])<<24 | uint64(buf[5])<<16 | uint64(buf[6])<<8 | uint64(buf[7])

	// 清除最高位（符号位）并右移到适当位置
	// 这样得到的值范围是[0, 2^53-1]/2^53 = [0.0, 1.0)
	return float64(u>>11) / (1 << 53), nil
}
