package chacha20poly1305

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGenerateKey(t *testing.T) {
	key, err := GenerateKey()
	if err != nil {
		t.Fatalf("生成密钥失败: %v", err)
	}

	if len(key) != KeySize {
		t.Errorf("密钥长度不正确，期望%d字节，实际%d字节", KeySize, len(key))
	}
}

func TestEncryptDecrypt(t *testing.T) {
	// 生成密钥
	key, err := GenerateKey()
	if err != nil {
		t.Fatalf("生成密钥失败: %v", err)
	}

	plaintext := []byte("Hello, ChaCha20-Poly1305!")
	additionalData := []byte("additional data")

	// 加密
	ciphertext, err := Encrypt(key, plaintext, additionalData)
	if err != nil {
		t.Fatalf("加密失败: %v", err)
	}

	if len(ciphertext) <= NonceSize {
		t.Error("密文长度不正确")
	}

	// 解密
	decrypted, err := Decrypt(key, ciphertext, additionalData)
	if err != nil {
		t.Fatalf("解密失败: %v", err)
	}

	if string(decrypted) != string(plaintext) {
		t.Errorf("解密结果不匹配，原文: %s, 解密结果: %s", string(plaintext), string(decrypted))
	}
}

func TestEncryptDecryptText(t *testing.T) {
	// 生成密钥
	key, err := GenerateKey()
	if err != nil {
		t.Fatalf("生成密钥失败: %v", err)
	}

	text := "Hello, ChaCha20-Poly1305 text encryption!"
	additionalData := []byte("additional data")

	// 加密
	ciphertext, err := EncryptText(key, text, additionalData)
	if err != nil {
		t.Fatalf("文本加密失败: %v", err)
	}

	if len(ciphertext) <= NonceSize {
		t.Error("密文长度不正确")
	}

	// 解密
	decrypted, err := DecryptText(key, ciphertext, additionalData)
	if err != nil {
		t.Fatalf("文本解密失败: %v", err)
	}

	if decrypted != text {
		t.Errorf("解密结果不匹配，原文: %s, 解密结果: %s", text, decrypted)
	}
}

func TestEncryptDecryptFile(t *testing.T) {
	// 创建临时目录
	tempDir := filepath.Join(os.TempDir(), "chacha20poly1305_test")
	if err := os.MkdirAll(tempDir, 0755); err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	originalFile := filepath.Join(tempDir, "original.txt")
	encryptedFile := filepath.Join(tempDir, "encrypted.bin")
	decryptedFile := filepath.Join(tempDir, "decrypted.txt")

	// 生成密钥
	key, err := GenerateKey()
	if err != nil {
		t.Fatalf("生成密钥失败: %v", err)
	}

	content := "Hello, ChaCha20-Poly1305 file encryption!"
	additionalData := []byte("additional data")

	// 创建原始文件
	err = os.WriteFile(originalFile, []byte(content), 0644)
	if err != nil {
		t.Fatalf("创建原始文件失败: %v", err)
	}

	// 加密文件
	err = EncryptFile(key, originalFile, encryptedFile, additionalData)
	if err != nil {
		t.Fatalf("加密文件失败: %v", err)
	}

	// 解密文件
	err = DecryptFile(key, encryptedFile, decryptedFile, additionalData)
	if err != nil {
		t.Fatalf("解密文件失败: %v", err)
	}

	// 验证内容
	decryptedContent, err := os.ReadFile(decryptedFile)
	if err != nil {
		t.Fatalf("读取解密文件失败: %v", err)
	}

	if string(decryptedContent) != content {
		t.Errorf("解密文件内容不匹配，原文: %s, 解密结果: %s", content, string(decryptedContent))
	}
}

func TestKeyLengthValidation(t *testing.T) {
	shortKey := make([]byte, KeySize-1)
	longKey := make([]byte, KeySize+1)
	plaintext := []byte("test data")

	// 测试短密钥
	_, err := Encrypt(shortKey, plaintext, nil)
	if err == nil {
		t.Error("使用短密钥加密应该返回错误")
	}

	_, err = Decrypt(shortKey, make([]byte, NonceSize+10), nil)
	if err == nil {
		t.Error("使用短密钥解密应该返回错误")
	}

	// 测试长密钥
	_, err = Encrypt(longKey, plaintext, nil)
	if err == nil {
		t.Error("使用长密钥加密应该返回错误")
	}

	_, err = Decrypt(longKey, make([]byte, NonceSize+10), nil)
	if err == nil {
		t.Error("使用长密钥解密应该返回错误")
	}
}

func TestEmptyInput(t *testing.T) {
	// 生成密钥
	key, err := GenerateKey()
	if err != nil {
		t.Fatalf("生成密钥失败: %v", err)
	}

	// 测试加密空数据
	_, err = Encrypt(key, []byte{}, nil)
	if err == nil {
		t.Error("加密空数据应该返回错误")
	}

	// 测试解密短数据
	_, err = Decrypt(key, make([]byte, NonceSize-1), nil)
	if err == nil {
		t.Error("解密短数据应该返回错误")
	}
}

func TestIncorrectAdditionalData(t *testing.T) {
	// 生成密钥
	key, err := GenerateKey()
	if err != nil {
		t.Fatalf("生成密钥失败: %v", err)
	}

	plaintext := []byte("test data")
	additionalData := []byte("additional data")
	wrongAdditionalData := []byte("wrong additional data")

	// 加密
	ciphertext, err := Encrypt(key, plaintext, additionalData)
	if err != nil {
		t.Fatalf("加密失败: %v", err)
	}

	// 使用错误的附加数据解密应该失败
	_, err = Decrypt(key, ciphertext, wrongAdditionalData)
	if err == nil {
		t.Error("使用错误的附加数据解密应该返回错误")
	}
}
