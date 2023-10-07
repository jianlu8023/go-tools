package crypto

import (
	"fmt"

	"github.com/jianlu8023/go-tools/pkg/crypto/rsa"
	"github.com/jianlu8023/go-tools/pkg/encoding/base64"
)

// CryptoAndEncodingDemo
// @Description: CryptoAndEncodingDemo
// @author ght
// @date 2023-10-07 16:50:08
func CryptoAndEncodingDemo() {
	data := []byte("hello world")
	fmt.Println(data)
	encrypt := rsa.RsaEncrypt(data)
	fmt.Println(encrypt)
	encode := base64.ByteToBase64(encrypt)
	fmt.Println(encode)
	decrypt := rsa.RsaDecrypt(encrypt)
	fmt.Println(decrypt)
}
