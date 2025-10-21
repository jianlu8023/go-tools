package sm2

import (
	"crypto/rand"
	"encoding/pem"
	"os"
	"path/filepath"
	"testing"

	"github.com/tjfoc/gmsm/sm2"
	"github.com/tjfoc/gmsm/x509"
)

const (
	c1c2c3 = false
)

// 使用临时目录进行测试
func tempDir(t *testing.T) string {
	tempDir := filepath.Join(os.TempDir(), "sm2_test")
	if err := os.MkdirAll(tempDir, 0755); err != nil {
		t.Fatal(err)
	}
	return tempDir
}

func TestGenKey(t *testing.T) {
	tempDir := tempDir(t)
	privateKeyPath := filepath.Join(tempDir, "private.pem")
	publicKeyPath := filepath.Join(tempDir, "public.pem")

	privateKey, err := sm2.GenerateKey(rand.Reader)
	if err != nil {
		t.Fatal(err)
	}
	privKeyBytes, err := x509.MarshalSm2PrivateKey(privateKey, nil)
	if err != nil {
		t.Fatal(err)
	}
	privKeyBlock := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: privKeyBytes,
	}

	privKeyFile, err := os.Create(privateKeyPath)
	if err != nil {
		t.Fatal(err)
	}
	defer privKeyFile.Close()
	if err = pem.Encode(privKeyFile, privKeyBlock); err != nil {
		t.Fatal(err)
	}
	publicKey := privateKey.PublicKey

	pubKeyBytes, err := x509.MarshalSm2PublicKey(&publicKey)
	if err != nil {
		t.Fatal(err)
	}

	pubKeyBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubKeyBytes,
	}
	pubKeyFile, err := os.Create(publicKeyPath)
	if err != nil {
		t.Fatal(err)
	}
	defer pubKeyFile.Close()
	if err = pem.Encode(pubKeyFile, pubKeyBlock); err != nil {
		t.Fatal(err)
	}
}

func TestEncryptFile(t *testing.T) {
	tempDir := tempDir(t)
	original := filepath.Join(tempDir, "original.txt")
	encrypted := filepath.Join(tempDir, "encrypted.txt")
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

	if err = EncryptFile(original, encrypted, publicKey, c1c2c3); err != nil {
		t.Fatal(err)
	}
}

// 测试SM2加密解密文本
func TestEncryptDecryptText(t *testing.T) {
	// 生成密钥对
	privateKey, err := sm2.GenerateKey(rand.Reader)
	if err != nil {
		t.Fatal(err)
	}
	publicKey := &privateKey.PublicKey

	text := "Hello, SM2 encryption!"

	// 加密
	ciphertext, err := EncryptText(text, publicKey, c1c2c3)
	if err != nil {
		t.Fatalf("EncryptText error: %v", err)
	}

	// 解密
	decryptedText, err := DecryptText(ciphertext, privateKey, c1c2c3)
	if err != nil {
		t.Fatalf("DecryptText error: %v", err)
	}

	// 验证
	if decryptedText != text {
		t.Errorf("解密结果不匹配. 原文: %s, 解密结果: %s", text, decryptedText)
	}
}

// 测试空输入
func TestEmptyInput(t *testing.T) {
	// 生成密钥对
	privateKey, err := sm2.GenerateKey(rand.Reader)
	if err != nil {
		t.Fatal(err)
	}
	publicKey := &privateKey.PublicKey

	emptyText := ""

	// 测试加密空文本，应该返回错误
	_, err = EncryptText(emptyText, publicKey, c1c2c3)
	if err == nil {
		t.Error("使用空文本加密时应该返回错误")
	}

	// 测试解密空密文，应该返回错误
	_, err = DecryptText([]byte{}, privateKey, c1c2c3)
	if err == nil {
		t.Error("使用空密文解密时应该返回错误")
	}
}

// 测试nil密钥
func TestNilKey(t *testing.T) {
	text := "test text"

	// 测试加密时使用nil公钥
	_, err := EncryptText(text, nil, c1c2c3)
	if err == nil {
		t.Error("使用nil公钥加密时应该返回错误")
	}

	// 测试解密时使用nil私钥
	_, err = DecryptText([]byte(text), nil, c1c2c3)
	if err == nil {
		t.Error("使用nil私钥解密时应该返回错误")
	}
}
