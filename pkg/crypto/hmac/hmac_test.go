package hmac

import (
	"crypto/sha256"
	"crypto/sha512"
	"testing"
)

func TestSignAndVerify(t *testing.T) {
	key := []byte("secret key")
	message := []byte("Hello, HMAC!")

	// 测试SHA256
	signature, err := Sign(key, message, SHA256)
	if err != nil {
		t.Fatalf("SHA256签名失败: %v", err)
	}

	if len(signature) != sha256.Size {
		t.Errorf("SHA256签名长度不正确，期望%d字节，实际%d字节", sha256.Size, len(signature))
	}

	valid, err := Verify(key, message, signature, SHA256)
	if err != nil {
		t.Fatalf("SHA256验证失败: %v", err)
	}

	if !valid {
		t.Error("SHA256签名验证应该通过")
	}

	// 测试SHA384
	signature, err = Sign(key, message, SHA384)
	if err != nil {
		t.Fatalf("SHA384签名失败: %v", err)
	}

	if len(signature) != sha512.Size384 {
		t.Errorf("SHA384签名长度不正确，期望%d字节，实际%d字节", sha512.Size384, len(signature))
	}

	valid, err = Verify(key, message, signature, SHA384)
	if err != nil {
		t.Fatalf("SHA384验证失败: %v", err)
	}

	if !valid {
		t.Error("SHA384签名验证应该通过")
	}

	// 测试SHA512
	signature, err = Sign(key, message, SHA512)
	if err != nil {
		t.Fatalf("SHA512签名失败: %v", err)
	}

	if len(signature) != sha512.Size {
		t.Errorf("SHA512签名长度不正确，期望%d字节，实际%d字节", sha512.Size, len(signature))
	}

	valid, err = Verify(key, message, signature, SHA512)
	if err != nil {
		t.Fatalf("SHA512验证失败: %v", err)
	}

	if !valid {
		t.Error("SHA512签名验证应该通过")
	}
}

func TestSignAndVerifyText(t *testing.T) {
	key := []byte("secret key")
	text := "Hello, HMAC text signing!"

	// 测试SHA256文本签名
	signature, err := SignText(key, text, SHA256)
	if err != nil {
		t.Fatalf("SHA256文本签名失败: %v", err)
	}

	if len(signature) != sha256.Size {
		t.Errorf("SHA256签名长度不正确，期望%d字节，实际%d字节", sha256.Size, len(signature))
	}

	valid, err := VerifyText(key, text, signature, SHA256)
	if err != nil {
		t.Fatalf("SHA256文本验证失败: %v", err)
	}

	if !valid {
		t.Error("SHA256文本签名验证应该通过")
	}
}

func TestSignAndVerifySHA256Convenience(t *testing.T) {
	key := []byte("secret key")
	message := []byte("Hello, HMAC SHA256!")

	// 测试便利函数
	signature, err := SignSHA256(key, message)
	if err != nil {
		t.Fatalf("SHA256便利函数签名失败: %v", err)
	}

	valid, err := VerifySHA256(key, message, signature)
	if err != nil {
		t.Fatalf("SHA256便利函数验证失败: %v", err)
	}

	if !valid {
		t.Error("SHA256便利函数签名验证应该通过")
	}

	// 测试文本便利函数
	text := "Hello, HMAC SHA256 text!"
	signature, err = SignTextSHA256(key, text)
	if err != nil {
		t.Fatalf("SHA256文本便利函数签名失败: %v", err)
	}

	valid, err = VerifyTextSHA256(key, text, signature)
	if err != nil {
		t.Fatalf("SHA256文本便利函数验证失败: %v", err)
	}

	if !valid {
		t.Error("SHA256文本便利函数签名验证应该通过")
	}
}

func TestVerifyWithWrongSignature(t *testing.T) {
	key := []byte("secret key")
	message := []byte("Hello, HMAC!")
	wrongSignature := make([]byte, sha256.Size)

	valid, err := Verify(key, message, wrongSignature, SHA256)
	if err != nil {
		t.Fatalf("验证失败: %v", err)
	}

	if valid {
		t.Error("使用错误签名验证应该失败")
	}
}

func TestNilKeyErrors(t *testing.T) {
	message := []byte("test message")

	// 测试使用nil密钥签名
	_, err := Sign(nil, message, SHA256)
	if err == nil {
		t.Error("使用nil密钥签名应该返回错误")
	}

	// 测试使用nil密钥验证
	_, err = Verify(nil, message, make([]byte, sha256.Size), SHA256)
	if err == nil {
		t.Error("使用nil密钥验证应该返回错误")
	}
}

func TestEmptyMessageErrors(t *testing.T) {
	key := []byte("secret key")

	// 测试签名空消息
	_, err := Sign(key, []byte{}, SHA256)
	if err == nil {
		t.Error("签名空消息应该返回错误")
	}

	// 测试验证空消息
	_, err = Verify(key, []byte{}, make([]byte, sha256.Size), SHA256)
	if err == nil {
		t.Error("验证空消息应该返回错误")
	}
}

func TestNilHashFuncErrors(t *testing.T) {
	key := []byte("secret key")
	message := []byte("test message")

	// 测试使用nil哈希函数签名
	_, err := Sign(key, message, nil)
	if err == nil {
		t.Error("使用nil哈希函数签名应该返回错误")
	}

	// 测试使用nil哈希函数验证
	_, err = Verify(key, message, make([]byte, sha256.Size), nil)
	if err == nil {
		t.Error("使用nil哈希函数验证应该返回错误")
	}
}
