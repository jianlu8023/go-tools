package rsa

import (
	"fmt"
	"testing"

	"github.com/jianlu8023/go-tools/pkg/encoding/base64"
)

// TestRsaEncryptDecrypt
// @Description: TestRsaEncryptDecrypt
// @author ght
// @date 2023-10-07 17:02:45
// @param t:
func TestRsaEncryptDecrypt(t *testing.T) {
	data := []byte("hello world")
	fmt.Println(data)
	encrypt := Encrypt(data)
	fmt.Println(encrypt)
	encode := base64.ByteToBase64(encrypt)
	fmt.Println(encode)
	decrypt := Decrypt(encrypt)
	fmt.Println(decrypt)

}
