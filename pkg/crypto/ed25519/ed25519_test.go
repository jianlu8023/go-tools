package ed25519

import (
	"crypto/ed25519"
	"os"
	"path/filepath"
	"testing"
)

func TestGenerateKey(t *testing.T) {
	publicKey, privateKey, err := GenerateKey()
	if err != nil {
		t.Fatalf("生成密钥对失败: %v", err)
	}

	if publicKey == nil {
		t.Error("公钥不能为空")
	}

	if privateKey == nil {
		t.Error("私钥不能为空")
	}

	if len(publicKey) != ed25519.PublicKeySize {
		t.Errorf("公钥长度不正确，期望%d字节，实际%d字节", ed25519.PublicKeySize, len(publicKey))
	}

	if len(privateKey) != ed25519.PrivateKeySize {
		t.Errorf("私钥长度不正确，期望%d字节，实际%d字节", ed25519.PrivateKeySize, len(privateKey))
	}
}

func TestSignAndVerify(t *testing.T) {
	// 生成密钥对
	publicKey, privateKey, err := GenerateKey()
	if err != nil {
		t.Fatalf("生成密钥对失败: %v", err)
	}

	message := []byte("Hello, Ed25519!")

	// 签名
	signature, err := Sign(privateKey, message)
	if err != nil {
		t.Fatalf("签名失败: %v", err)
	}

	if len(signature) != ed25519.SignatureSize {
		t.Errorf("签名长度不正确，期望%d字节，实际%d字节", ed25519.SignatureSize, len(signature))
	}

	// 验证
	valid, err := Verify(publicKey, message, signature)
	if err != nil {
		t.Fatalf("验证失败: %v", err)
	}

	if !valid {
		t.Error("签名验证应该通过")
	}

	// 验证错误的签名
	wrongSignature := make([]byte, ed25519.SignatureSize)
	valid, err = Verify(publicKey, message, wrongSignature)
	if err != nil {
		t.Fatalf("验证失败: %v", err)
	}

	if valid {
		t.Error("错误的签名验证应该失败")
	}
}

func TestSignAndVerifyText(t *testing.T) {
	// 生成密钥对
	publicKey, privateKey, err := GenerateKey()
	if err != nil {
		t.Fatalf("生成密钥对失败: %v", err)
	}

	text := "Hello, Ed25519 text signing!"

	// 签名
	signature, err := SignText(privateKey, text)
	if err != nil {
		t.Fatalf("文本签名失败: %v", err)
	}

	if len(signature) != ed25519.SignatureSize {
		t.Errorf("签名长度不正确，期望%d字节，实际%d字节", ed25519.SignatureSize, len(signature))
	}

	// 验证
	valid, err := VerifyText(publicKey, text, signature)
	if err != nil {
		t.Fatalf("文本验证失败: %v", err)
	}

	if !valid {
		t.Error("文本签名验证应该通过")
	}
}

func TestSaveAndLoadKeys(t *testing.T) {
	// 创建临时目录
	tempDir := filepath.Join(os.TempDir(), "ed25519_test")
	if err := os.MkdirAll(tempDir, 0755); err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	privateKeyFile := filepath.Join(tempDir, "private.pem")
	publicKeyFile := filepath.Join(tempDir, "public.pem")

	// 生成密钥对
	publicKey, privateKey, err := GenerateKey()
	if err != nil {
		t.Fatalf("生成密钥对失败: %v", err)
	}

	// 保存密钥
	err = SavePrivateKey(privateKey, privateKeyFile)
	if err != nil {
		t.Fatalf("保存私钥失败: %v", err)
	}

	err = SavePublicKey(publicKey, publicKeyFile)
	if err != nil {
		t.Fatalf("保存公钥失败: %v", err)
	}

	// 加载密钥
	loadedPrivateKey, err := LoadPrivateKey(privateKeyFile)
	if err != nil {
		t.Fatalf("加载私钥失败: %v", err)
	}

	loadedPublicKey, err := LoadPublicKey(publicKeyFile)
	if err != nil {
		t.Fatalf("加载公钥失败: %v", err)
	}

	// 验证密钥是否一致
	if len(loadedPrivateKey) != len(privateKey) {
		t.Error("加载的私钥长度不匹配")
	}

	if len(loadedPublicKey) != len(publicKey) {
		t.Error("加载的公钥长度不匹配")
	}

	// 验证密钥可以正常使用
	message := []byte("Test message for key loading")
	signature, err := Sign(loadedPrivateKey, message)
	if err != nil {
		t.Fatalf("使用加载的私钥签名失败: %v", err)
	}

	valid, err := Verify(loadedPublicKey, message, signature)
	if err != nil {
		t.Fatalf("使用加载的公钥验证失败: %v", err)
	}

	if !valid {
		t.Error("使用加载的密钥验证应该通过")
	}
}

func TestNilKeyErrors(t *testing.T) {
	message := []byte("Test message")

	// 测试使用nil私钥签名
	_, err := Sign(nil, message)
	if err == nil {
		t.Error("使用nil私钥签名应该返回错误")
	}

	// 测试使用nil公钥验证
	_, err = Verify(nil, message, make([]byte, ed25519.SignatureSize))
	if err == nil {
		t.Error("使用nil公钥验证应该返回错误")
	}
}

func TestEmptyMessageErrors(t *testing.T) {
	// 生成密钥对
	_, privateKey, err := GenerateKey()
	if err != nil {
		t.Fatalf("生成密钥对失败: %v", err)
	}

	// 测试签名空消息
	_, err = Sign(privateKey, []byte{})
	if err == nil {
		t.Error("签名空消息应该返回错误")
	}
}
