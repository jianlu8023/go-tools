package pbkdf2

import (
	"testing"
)

func TestGenerateSalt(t *testing.T) {
	salt, err := GenerateSalt()
	if err != nil {
		t.Fatalf("生成盐值失败: %v", err)
	}

	if len(salt) != 16 {
		t.Errorf("盐值长度不正确，期望16字节，实际%d字节", len(salt))
	}
}

func TestDeriveKey(t *testing.T) {
	password := []byte("mysecretpassword")
	salt := []byte("somesalt12345678")
	params := DefaultParams()

	key, err := DeriveKey(password, salt, params)
	if err != nil {
		t.Fatalf("派生密钥失败: %v", err)
	}

	if len(key) != params.KeyLen {
		t.Errorf("密钥长度不正确，期望%d字节，实际%d字节", params.KeyLen, len(key))
	}
}

func TestDeriveKeyWithSalt(t *testing.T) {
	password := []byte("mysecretpassword")
	params := DefaultParams()

	key, salt, err := DeriveKeyWithSalt(password, params)
	if err != nil {
		t.Fatalf("派生密钥失败: %v", err)
	}

	if len(key) != params.KeyLen {
		t.Errorf("密钥长度不正确，期望%d字节，实际%d字节", params.KeyLen, len(key))
	}

	if len(salt) != 16 {
		t.Errorf("盐值长度不正确，期望16字节，实际%d字节", len(salt))
	}
}

func TestDeriveKeySHA512(t *testing.T) {
	password := []byte("mysecretpassword")
	salt := []byte("somesalt12345678")
	params := Params{
		Iterations: DefaultIterations,
		KeyLen:     64,
		HashFunc:   SHA512,
	}

	key, err := DeriveKey(password, salt, params)
	if err != nil {
		t.Fatalf("派生密钥失败: %v", err)
	}

	if len(key) != params.KeyLen {
		t.Errorf("密钥长度不正确，期望%d字节，实际%d字节", params.KeyLen, len(key))
	}
}

func TestDeriveKeyText(t *testing.T) {
	password := "mysecretpassword"
	salt := []byte("somesalt12345678")
	params := DefaultParams()

	key, err := DeriveKeyText(password, salt, params)
	if err != nil {
		t.Fatalf("文本派生密钥失败: %v", err)
	}

	if len(key) != params.KeyLen {
		t.Errorf("密钥长度不正确，期望%d字节，实际%d字节", params.KeyLen, len(key))
	}
}

func TestDeriveKeyTextWithSalt(t *testing.T) {
	password := "mysecretpassword"
	params := DefaultParams()

	key, salt, err := DeriveKeyTextWithSalt(password, params)
	if err != nil {
		t.Fatalf("文本派生密钥失败: %v", err)
	}

	if len(key) != params.KeyLen {
		t.Errorf("密钥长度不正确，期望%d字节，实际%d字节", params.KeyLen, len(key))
	}

	if len(salt) != 16 {
		t.Errorf("盐值长度不正确，期望16字节，实际%d字节", len(salt))
	}
}

func TestInvalidParameters(t *testing.T) {
	password := []byte("mysecretpassword")
	salt := []byte("somesalt12345678")

	_, err := DeriveKey([]byte{}, salt, DefaultParams())
	if err == nil {
		t.Error("使用空密码派生密钥应该返回错误")
	}

	_, err = DeriveKey(password, []byte{}, DefaultParams())
	if err == nil {
		t.Error("使用空盐值派生密钥应该返回错误")
	}

	invalidParams := Params{
		Iterations: 0,
		KeyLen:     DefaultKeyLen,
		HashFunc:   SHA256,
	}
	_, err = DeriveKey(password, salt, invalidParams)
	if err == nil {
		t.Error("使用零迭代次数派生密钥应该返回错误")
	}

	invalidParams = Params{
		Iterations: DefaultIterations,
		KeyLen:     0,
		HashFunc:   SHA256,
	}
	_, err = DeriveKey(password, salt, invalidParams)
	if err == nil {
		t.Error("使用零密钥长度派生密钥应该返回错误")
	}

	invalidParams = Params{
		Iterations: DefaultIterations,
		KeyLen:     DefaultKeyLen,
		HashFunc:   nil,
	}
	_, err = DeriveKey(password, salt, invalidParams)
	if err == nil {
		t.Error("使用nil哈希函数派生密钥应该返回错误")
	}
}
