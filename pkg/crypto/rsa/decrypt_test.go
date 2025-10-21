package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
	"path/filepath"
	"testing"
)

const (
	keySize = 1024 // 使用较小的密钥以提高测试速度
)

// 使用临时目录进行测试
func tempDir(t *testing.T) string {
	tempDir := filepath.Join(os.TempDir(), "rsa_test")
	if err := os.MkdirAll(tempDir, 0755); err != nil {
		t.Fatal(err)
	}
	return tempDir
}

func TestGenKey(t *testing.T) {
	tempDir := tempDir(t)
	privateKeyPath := filepath.Join(tempDir, "private.pem")
	publicKeyPath := filepath.Join(tempDir, "public.pem")

	privateKey, err := rsa.GenerateKey(rand.Reader, keySize)
	if err != nil {
		t.Fatal(err)
	}
	privKeyBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}
	if privKeyFile, err := os.Create(privateKeyPath); err != nil {
		defer privKeyFile.Close()
		t.Fatal(err)
	} else {
		defer privKeyFile.Close()
		if err := pem.Encode(privKeyFile, privKeyBlock); err != nil {
			t.Fatal(err)
		}
	}

	publicKey := &privateKey.PublicKey
	pubKeyBlock := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(publicKey),
	}
	if pubKeyFile, err := os.Create(publicKeyPath); err != nil {
		defer pubKeyFile.Close()
		t.Fatal(err)
	} else {
		defer pubKeyFile.Close()
		if err := pem.Encode(pubKeyFile, pubKeyBlock); err != nil {
			t.Fatal(err)
		}
	}
}

func TestEncryptFile(t *testing.T) {
	tempDir := tempDir(t)
	original := filepath.Join(tempDir, "original.txt")
	encrypted := filepath.Join(tempDir, "encrypted.txt")
	content := "this is a rsa crypto test file"

	// 生成密钥
	privateKey, err := rsa.GenerateKey(rand.Reader, keySize)
	if err != nil {
		t.Fatal(err)
	}

	publicKey := &privateKey.PublicKey

	// 创建原始文件
	if err := os.WriteFile(original, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	if err := EncryptFile(original, encrypted, publicKey); err != nil {
		t.Fatal(err)
	}
}

func TestDecryptFile(t *testing.T) {
	tempDir := tempDir(t)
	original := filepath.Join(tempDir, "original.txt")
	encrypted := filepath.Join(tempDir, "encrypted.txt")
	decrypted := filepath.Join(tempDir, "decrypted.txt")
	content := "this is a rsa crypto test file"

	// 生成密钥
	privateKey, err := rsa.GenerateKey(rand.Reader, keySize)
	if err != nil {
		t.Fatal(err)
	}

	publicKey := &privateKey.PublicKey

	// 创建原始文件
	if err := os.WriteFile(original, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	// 先加密
	if err := EncryptFile(original, encrypted, publicKey); err != nil {
		t.Fatal(err)
	}

	// 再解密
	if err := DecryptFile(encrypted, decrypted, privateKey); err != nil {
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

// 测试RSA加密解密文本
func TestEncryptDecryptText(t *testing.T) {
	// 生成密钥对
	privateKey, err := rsa.GenerateKey(rand.Reader, keySize)
	if err != nil {
		t.Fatal(err)
	}
	publicKey := &privateKey.PublicKey

	text := "Hello, RSA encryption!"

	// 加密
	ciphertext, err := EncryptText(text, publicKey)
	if err != nil {
		t.Fatalf("EncryptText error: %v", err)
	}

	// 解密
	decryptedText, err := DecryptText(ciphertext, privateKey)
	if err != nil {
		t.Fatalf("DecryptText error: %v", err)
	}

	// 验证
	if decryptedText != text {
		t.Errorf("解密结果不匹配. 原文: %s, 解密结果: %s", text, decryptedText)
	}
}

// 测试RSA加密解密字节切片
func TestEncryptDecryptBytes(t *testing.T) {
	// 生成密钥对
	privateKey, err := rsa.GenerateKey(rand.Reader, keySize)
	if err != nil {
		t.Fatal(err)
	}
	publicKey := &privateKey.PublicKey

	plaintext := []byte("Hello, RSA bytes encryption!")

	// 加密
	ciphertext, err := Encrypt(plaintext, publicKey)
	if err != nil {
		t.Fatalf("Encrypt error: %v", err)
	}

	// 解密
	decryptedBytes, err := Decrypt(ciphertext, privateKey)
	if err != nil {
		t.Fatalf("Decrypt error: %v", err)
	}

	// 验证
	if string(decryptedBytes) != string(plaintext) {
		t.Errorf("解密结果不匹配. 原文: %s, 解密结果: %s", string(plaintext), string(decryptedBytes))
	}
}

// 测试空输入
func TestEmptyInput(t *testing.T) {
	// 生成密钥对
	privateKey, err := rsa.GenerateKey(rand.Reader, keySize)
	if err != nil {
		t.Fatal(err)
	}
	publicKey := &privateKey.PublicKey

	emptyText := ""

	// 加密
	ciphertext, err := EncryptText(emptyText, publicKey)
	if err != nil {
		t.Fatalf("EncryptText with empty text error: %v", err)
	}

	// 解密
	decryptedText, err := DecryptText(ciphertext, privateKey)
	if err != nil {
		t.Fatalf("DecryptText with empty text error: %v", err)
	}

	if decryptedText != emptyText {
		t.Errorf("空文本解密结果不匹配. 原文: '%s', 解密结果: '%s'", emptyText, decryptedText)
	}
}

// 测试超长文本输入
func TestLongTextInput(t *testing.T) {
	// 生成密钥对
	privateKey, err := rsa.GenerateKey(rand.Reader, keySize)
	if err != nil {
		t.Fatal(err)
	}
	publicKey := &privateKey.PublicKey

	// 创建一个超过RSA限制的长文本
	longText := make([]byte, publicKey.Size()-10) // 超过最大限制
	for i := range longText {
		longText[i] = 'a'
	}

	// 尝试加密，应该返回错误
	_, err = Encrypt(longText, publicKey)
	if err == nil {
		t.Error("使用超长文本加密时应该返回错误")
	}
}

// 测试nil密钥
func TestNilKey(t *testing.T) {
	text := "test text"

	// 测试加密时使用nil公钥
	_, err := EncryptText(text, nil)
	if err == nil {
		t.Error("使用nil公钥加密时应该返回错误")
	}

	// 测试解密时使用nil私钥
	_, err = DecryptText([]byte(text), nil)
	if err == nil {
		t.Error("使用nil私钥解密时应该返回错误")
	}
}
