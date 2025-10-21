package sm2

import (
	"crypto/rand"
	"os"
	"path/filepath"
	"testing"

	"github.com/tjfoc/gmsm/sm2"
)

func TestDecryptFile(t *testing.T) {
	tempDir := tempDir(t)
	original := filepath.Join(tempDir, "original.txt")
	encrypted := filepath.Join(tempDir, "encrypted.txt")
	decrypted := filepath.Join(tempDir, "decrypted.txt")
	content := "this is a sm2 crypto test file"

	// 生成密钥
	privateKey, err := sm2.GenerateKey(rand.Reader)
	if err != nil {
		t.Fatal(err)
	}
	publicKey := &privateKey.PublicKey

	// 创建原始文件
	if err := os.WriteFile(original, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	// 先加密
	if err = EncryptFile(original, encrypted, publicKey, c1c2c3); err != nil {
		t.Fatal(err)
	}

	// 再解密
	if err = DecryptFile(encrypted, decrypted, privateKey, c1c2c3); err != nil {
		t.Fatal(err)
	}

	// 验证内容
	decryptedContent, err := os.ReadFile(decrypted)
	if err != nil {
		t.Fatal(err)
	}

	if string(decryptedContent) != content {
		t.Errorf("解密内容不匹配. 原文: %s, 解密结果: %s", content, string(decryptedContent))
	}
}
