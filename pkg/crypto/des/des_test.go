package des

import (
	"os"
	"path/filepath"
	"testing"
)

const (
	key = "12345678" // 8字节密钥
)

// 使用临时目录进行测试
func tempDir(t *testing.T) string {
	tempDir := filepath.Join(os.TempDir(), "des_test")
	if err := os.MkdirAll(tempDir, 0755); err != nil {
		t.Fatal(err)
	}
	return tempDir
}

func TestCBCEncryptFile(t *testing.T) {
	tempDir := tempDir(t)
	original := filepath.Join(tempDir, "original.txt")
	encrypted := filepath.Join(tempDir, "encrypted-cbc.txt")
	content := "this is a des crypto test file "

	if err := os.WriteFile(original, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	if err := CBCEncryptFile(key, original, encrypted); err != nil {
		t.Fatal(err)
	}
}

func TestCBCDecryptFile(t *testing.T) {
	tempDir := tempDir(t)
	original := filepath.Join(tempDir, "original.txt")
	encrypted := filepath.Join(tempDir, "encrypted-cbc.txt")
	decrypted := filepath.Join(tempDir, "decrypted-cbc.txt")
	content := "this is a des crypto test file "

	if err := os.WriteFile(original, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	// 先加密
	if err := CBCEncryptFile(key, original, encrypted); err != nil {
		t.Fatal(err)
	}

	// 再解密
	if err := CBCDecryptFile(key, encrypted, decrypted); err != nil {
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

func TestCTREncryptFile(t *testing.T) {
	tempDir := tempDir(t)
	original := filepath.Join(tempDir, "original.txt")
	encrypted := filepath.Join(tempDir, "encrypted-ctr.txt")
	content := "this is a des crypto test file "

	if err := os.WriteFile(original, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}
	if err := CTREncryptFile(key, original, encrypted); err != nil {
		t.Fatal(err)
	}
}

func TestCTRDecryptFile(t *testing.T) {
	tempDir := tempDir(t)
	original := filepath.Join(tempDir, "original.txt")
	encrypted := filepath.Join(tempDir, "encrypted-ctr.txt")
	decrypted := filepath.Join(tempDir, "decrypted-ctr.txt")
	content := "this is a des crypto test file "

	if err := os.WriteFile(original, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	// 先加密
	if err := CTREncryptFile(key, original, encrypted); err != nil {
		t.Fatal(err)
	}

	// 再解密
	if err := CTRDecryptFile(key, encrypted, decrypted); err != nil {
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

// 测试CBC模式加密解密文本
func TestCBCEncryptDecryptText(t *testing.T) {
	text := "Hello, DES CBC encryption!"
	keyBytes := []byte(key)

	// 加密
	ciphertext, err := CBCEncryptText(keyBytes, text)
	if err != nil {
		t.Fatalf("CBCEncryptText error: %v", err)
	}

	// 解密
	decryptedText, err := CBCDecryptText(keyBytes, ciphertext)
	if err != nil {
		t.Fatalf("CBCDecryptText error: %v", err)
	}

	// 验证
	if decryptedText != text {
		t.Errorf("解密结果不匹配. 原文: %s, 解密结果: %s", text, decryptedText)
	}
}

// 测试CTR模式加密解密文本
func TestCTREncryptDecryptText(t *testing.T) {
	text := "Hello, DES CTR encryption!"
	keyBytes := []byte(key)

	// 加密
	ciphertext, err := CTREncryptText(keyBytes, text)
	if err != nil {
		t.Fatalf("CTREncryptText error: %v", err)
	}

	// 解密
	decryptedText, err := CTRDecryptText(keyBytes, ciphertext)
	if err != nil {
		t.Fatalf("CTRDecryptText error: %v", err)
	}

	// 验证
	if decryptedText != text {
		t.Errorf("解密结果不匹配. 原文: %s, 解密结果: %s", text, decryptedText)
	}
}

// 测试CBC模式加密解密字节切片
func TestCBCEncryptDecryptBytes(t *testing.T) {
	plaintext := []byte("Hello, DES CBC bytes encryption!")
	keyBytes := []byte(key)

	// 加密
	ciphertext, err := CBCEncrypt(keyBytes, plaintext)
	if err != nil {
		t.Fatalf("CBCEncrypt error: %v", err)
	}

	// 解密
	decryptedBytes, err := CBCDecrypt(keyBytes, ciphertext)
	if err != nil {
		t.Fatalf("CBCDecrypt error: %v", err)
	}

	// 验证
	if string(decryptedBytes) != string(plaintext) {
		t.Errorf("解密结果不匹配. 原文: %s, 解密结果: %s", string(plaintext), string(decryptedBytes))
	}
}

// 测试CTR模式加密解密字节切片
func TestCTREncryptDecryptBytes(t *testing.T) {
	plaintext := []byte("Hello, DES CTR bytes encryption!")
	keyBytes := []byte(key)

	// 加密
	ciphertext, err := CTREncrypt(keyBytes, plaintext)
	if err != nil {
		t.Fatalf("CTREncrypt error: %v", err)
	}

	// 解密
	decryptedBytes, err := CTRDecrypt(keyBytes, ciphertext)
	if err != nil {
		t.Fatalf("CTRDecrypt error: %v", err)
	}

	// 验证
	if string(decryptedBytes) != string(plaintext) {
		t.Errorf("解密结果不匹配. 原文: %s, 解密结果: %s", string(plaintext), string(decryptedBytes))
	}
}

// 测试密钥长度验证
func TestKeyLengthValidation(t *testing.T) {
	shortKey := []byte("short")             // 短密钥
	longKey := []byte("this_is_a_long_key") // 长密钥
	text := "test text"

	// 测试CBC模式短密钥
	_, err := CBCEncrypt(shortKey, []byte(text))
	if err == nil {
		t.Error("使用短密钥时应该返回错误")
	}

	// 测试CBC模式长密钥
	_, err = CBCEncrypt(longKey, []byte(text))
	if err == nil {
		t.Error("使用长密钥时应该返回错误")
	}

	// 测试CTR模式短密钥
	_, err = CTREncrypt(shortKey, []byte(text))
	if err == nil {
		t.Error("使用短密钥时应该返回错误")
	}

	// 测试CTR模式长密钥
	_, err = CTREncrypt(longKey, []byte(text))
	if err == nil {
		t.Error("使用长密钥时应该返回错误")
	}
}

// 测试空输入
func TestEmptyInput(t *testing.T) {
	emptyText := ""
	keyBytes := []byte(key)

	// CBC模式
	ciphertext, err := CBCEncryptText(keyBytes, emptyText)
	if err != nil {
		t.Fatalf("CBCEncryptText with empty text error: %v", err)
	}

	decryptedText, err := CBCDecryptText(keyBytes, ciphertext)
	if err != nil {
		t.Fatalf("CBCDecryptText with empty text error: %v", err)
	}

	if decryptedText != emptyText {
		t.Errorf("空文本解密结果不匹配. 原文: '%s', 解密结果: '%s'", emptyText, decryptedText)
	}

	// CTR模式
	ciphertext, err = CTREncryptText(keyBytes, emptyText)
	if err != nil {
		t.Fatalf("CTREncryptText with empty text error: %v", err)
	}

	decryptedText, err = CTRDecryptText(keyBytes, ciphertext)
	if err != nil {
		t.Fatalf("CTRDecryptText with empty text error: %v", err)
	}

	if decryptedText != emptyText {
		t.Errorf("空文本解密结果不匹配. 原文: '%s', 解密结果: '%s'", emptyText, decryptedText)
	}
}
