package argon2

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

func TestHash(t *testing.T) {
	password := []byte("mysecretpassword")
	salt := []byte("somesalt12345678")
	params := DefaultParams()

	hash, err := Hash(password, salt, params)
	if err != nil {
		t.Fatalf("哈希密码失败: %v", err)
	}

	if len(hash) != int(params.KeyLen) {
		t.Errorf("哈希值长度不正确，期望%d字节，实际%d字节", params.KeyLen, len(hash))
	}
}

func TestHashPassword(t *testing.T) {
	password := []byte("mysecretpassword")
	params := DefaultParams()

	hash, salt, err := HashPassword(password, params)
	if err != nil {
		t.Fatalf("哈希密码失败: %v", err)
	}

	if len(hash) != int(params.KeyLen) {
		t.Errorf("哈希值长度不正确，期望%d字节，实际%d字节", params.KeyLen, len(hash))
	}

	if len(salt) != 16 {
		t.Errorf("盐值长度不正确，期望16字节，实际%d字节", len(salt))
	}
}

func TestVerify(t *testing.T) {
	password := []byte("mysecretpassword")
	params := DefaultParams()

	// 哈希密码
	hash, salt, err := HashPassword(password, params)
	if err != nil {
		t.Fatalf("哈希密码失败: %v", err)
	}

	// 验证正确密码
	valid, err := Verify(password, hash, salt, params)
	if err != nil {
		t.Fatalf("验证密码失败: %v", err)
	}

	if !valid {
		t.Error("验证正确密码应该通过")
	}

	// 验证错误密码
	wrongPassword := []byte("wrongpassword")
	valid, err = Verify(wrongPassword, hash, salt, params)
	if err != nil {
		t.Fatalf("验证错误密码失败: %v", err)
	}

	if valid {
		t.Error("验证错误密码应该失败")
	}
}

func TestTextFunctions(t *testing.T) {
	password := "mysecretpassword"
	salt := []byte("somesalt12345678")
	params := DefaultParams()

	// 测试文本哈希
	hash, err := HashText(password, salt, params)
	if err != nil {
		t.Fatalf("文本哈希密码失败: %v", err)
	}

	if len(hash) != int(params.KeyLen) {
		t.Errorf("哈希值长度不正确，期望%d字节，实际%d字节", params.KeyLen, len(hash))
	}

	// 测试文本密码哈希
	hash, salt, err = HashTextPassword(password, params)
	if err != nil {
		t.Fatalf("文本密码哈希失败: %v", err)
	}

	if len(hash) != int(params.KeyLen) {
		t.Errorf("哈希值长度不正确，期望%d字节，实际%d字节", params.KeyLen, len(hash))
	}

	if len(salt) != 16 {
		t.Errorf("盐值长度不正确，期望16字节，实际%d字节", len(salt))
	}

	// 测试文本验证
	valid, err := VerifyText(password, hash, salt, params)
	if err != nil {
		t.Fatalf("文本验证密码失败: %v", err)
	}

	if !valid {
		t.Error("文本验证正确密码应该通过")
	}
}

func TestInvalidParameters(t *testing.T) {
	password := []byte("mysecretpassword")
	salt := []byte("somesalt12345678")

	// 测试空密码
	_, err := Hash([]byte{}, salt, DefaultParams())
	if err == nil {
		t.Error("使用空密码哈希应该返回错误")
	}

	// 测试空盐值
	_, err = Hash(password, []byte{}, DefaultParams())
	if err == nil {
		t.Error("使用空盐值哈希应该返回错误")
	}

	// 测试零时间参数
	invalidParams := Params{
		Time:    0,
		Memory:  DefaultMemory,
		Threads: DefaultThreads,
		KeyLen:  DefaultKeyLen,
	}
	_, err = Hash(password, salt, invalidParams)
	if err == nil {
		t.Error("使用零时间参数哈希应该返回错误")
	}

	// 测试零内存参数
	invalidParams = Params{
		Time:    DefaultTime,
		Memory:  0,
		Threads: DefaultThreads,
		KeyLen:  DefaultKeyLen,
	}
	_, err = Hash(password, salt, invalidParams)
	if err == nil {
		t.Error("使用零内存参数哈希应该返回错误")
	}

	// 测试零线程参数
	invalidParams = Params{
		Time:    DefaultTime,
		Memory:  DefaultMemory,
		Threads: 0,
		KeyLen:  DefaultKeyLen,
	}
	_, err = Hash(password, salt, invalidParams)
	if err == nil {
		t.Error("使用零线程参数哈希应该返回错误")
	}

	// 测试零密钥长度参数
	invalidParams = Params{
		Time:    DefaultTime,
		Memory:  DefaultMemory,
		Threads: DefaultThreads,
		KeyLen:  0,
	}
	_, err = Hash(password, salt, invalidParams)
	if err == nil {
		t.Error("使用零密钥长度参数哈希应该返回错误")
	}
}

func TestVerifyInvalidParameters(t *testing.T) {
	password := []byte("mysecretpassword")
	hash := make([]byte, 32)
	salt := make([]byte, 16)
	params := DefaultParams()

	// 测试空密码验证
	_, err := Verify([]byte{}, hash, salt, params)
	if err == nil {
		t.Error("使用空密码验证应该返回错误")
	}

	// 测试空哈希值验证
	_, err = Verify(password, []byte{}, salt, params)
	if err == nil {
		t.Error("使用空哈希值验证应该返回错误")
	}

	// 测试空盐值验证
	_, err = Verify(password, hash, []byte{}, params)
	if err == nil {
		t.Error("使用空盐值验证应该返回错误")
	}
}
