package steg

import (
	"crypto/rand"
	"io"

	"github.com/jianlu8023/go-tools/v2/pkg/compress/lz4"
	"github.com/jianlu8023/go-tools/v2/pkg/crypto/pbkdf2"
)

const saltLen = 16

func Encode(s []byte, key ...[]byte) (string, error) {
	dst, err := lz4.Compress(s)
	if err != nil {
		return "", err
	}
	if len(key) > 0 && len(key[0]) > 0 {
		salt := make([]byte, saltLen)
		if _, err := io.ReadFull(rand.Reader, salt); err != nil {
			return "", err
		}

		params := pbkdf2.Params{
			Iterations: pbkdf2.DefaultIterations,
			KeyLen:     48,
			HashFunc:   pbkdf2.SHA256,
		}
		derived, err := pbkdf2.DeriveKey(key[0], salt, params)
		if err != nil {
			return "", err
		}

		aesKey := derived[:32]
		iv := derived[32:48]

		dst, err = encrypt(dst, aesKey, iv)
		if err != nil {
			return "", err
		}

		dst = append(salt, dst...)
	}
	h := Huffman(dst)
	return h, nil
}

func Decode(s string, key ...[]byte) ([]byte, error) {
	b, err := UnHuffman(s)
	if err != nil {
		return []byte{}, err
	}
	if len(key) > 0 && len(key[0]) > 0 {
		if len(b) < saltLen {
			return []byte{}, nil
		}

		salt := b[:saltLen]
		cipherText := b[saltLen:]

		params := pbkdf2.Params{
			Iterations: pbkdf2.DefaultIterations,
			KeyLen:     48,
			HashFunc:   pbkdf2.SHA256,
		}
		derived, err := pbkdf2.DeriveKey(key[0], salt, params)
		if err != nil {
			return []byte{}, err
		}

		aesKey := derived[:32]
		iv := derived[32:48]

		b, err = decrypt(cipherText, aesKey, iv)
		if err != nil {
			return []byte{}, err
		}
	}

	decode, err := lz4.Decompress(b, 0)
	if err != nil {
		return []byte{}, err
	}

	return decode, nil
}
