package steg

import (
	"github.com/jianlu8023/go-tools/v2/pkg/compress/lz4"
)

func Encode(s []byte, key ...[]byte) (string, error) {
	dst, err := lz4.Compress(s)
	if err != nil {
		return "", err
	}
	if len(key) > 0 {
		dst = encrypt(dst, key[0])
	}
	h := Huffman(dst)
	return h, nil
}

func Decode(s string, key ...[]byte) ([]byte, error) {
	b, err := UnHuffman(s)
	if err != nil {
		return []byte{}, err
	}
	if len(key) != 0 {
		b = decrypt(b, key[0])
	}

	decode, err := lz4.Decompress(b, 0)
	if err != nil {
		return []byte{}, err
	}

	return decode, nil
}
