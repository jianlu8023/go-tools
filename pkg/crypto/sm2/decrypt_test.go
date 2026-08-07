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

// TestDecryptFileMultiBlock 测试多块文件的加解密往返，
// 用于验证长度前缀帧格式在跨块场景下能正确对齐解密。
func TestDecryptFileMultiBlock(t *testing.T) {
	tempDir := tempDir(t)
	original := filepath.Join(tempDir, "original_multi.txt")
	encrypted := filepath.Join(tempDir, "encrypted_multi.txt")
	decrypted := filepath.Join(tempDir, "decrypted_multi.txt")

	// 构造超过 encryptBlockSize(3999) 的内容，强制产生多个密文块
	// 10000 字节 => 3 个明文块: 3999 + 3999 + 2002
	content := make([]byte, 10000)
	for i := range content {
		content[i] = byte(i % 256)
	}

	// 生成密钥
	privateKey, err := sm2.GenerateKey(rand.Reader)
	if err != nil {
		t.Fatal(err)
	}
	publicKey := &privateKey.PublicKey

	// 创建原始文件
	if err := os.WriteFile(original, content, 0644); err != nil {
		t.Fatal(err)
	}

	// 加密
	if err = EncryptFile(original, encrypted, publicKey, c1c2c3); err != nil {
		t.Fatal(err)
	}

	// 解密
	if err = DecryptFile(encrypted, decrypted, privateKey, c1c2c3); err != nil {
		t.Fatal(err)
	}

	// 验证内容
	decryptedContent, err := os.ReadFile(decrypted)
	if err != nil {
		t.Fatal(err)
	}

	if len(decryptedContent) != len(content) {
		t.Fatalf("解密内容长度不匹配. 原文长度: %d, 解密结果长度: %d", len(content), len(decryptedContent))
	}

	for i := range content {
		if decryptedContent[i] != content[i] {
			t.Fatalf("解密内容在第 %d 字节不匹配. 原文: %d, 解密结果: %d", i, content[i], decryptedContent[i])
		}
	}
}
