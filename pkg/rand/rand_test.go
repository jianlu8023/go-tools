package rand

import (
	"testing"
)

func TestGenerateCryptoRandomString(t *testing.T) {
	// 测试基本功能
	s, err := GenerateCryptoRandomString(10, "abcdef")
	if err != nil {
		t.Errorf("GenerateCryptoRandomString returned error: %v", err)
	}
	if len(s) != 10 {
		t.Errorf("expected length 10, got %d", len(s))
	}

	// 测试空字符集
	_, err = GenerateCryptoRandomString(5, "")
	if err == nil {
		t.Error("expected error for empty runes, got nil")
	}

	// 测试长度为0
	s, err = GenerateCryptoRandomString(0, "abc")
	if err != nil {
		t.Errorf("GenerateCryptoRandomString(0) returned error: %v", err)
	}
	if len(s) != 0 {
		t.Errorf("expected empty string for length 0, got %q", s)
	}
}

func TestCryptoUint64(t *testing.T) {
	// 测试多次调用返回不同的值
	v1, err := CryptoUint64()
	if err != nil {
		t.Errorf("CryptoUint64 returned error: %v", err)
	}

	v2, err := CryptoUint64()
	if err != nil {
		t.Errorf("CryptoUint64 returned error: %v", err)
	}

	// 虽然理论上可能相同，但概率极低，可以认为是测试通过
	if v1 == v2 {
		t.Log("Warning: two CryptoUint64 calls returned the same value (low probability event)")
	}
}

func TestGenerateCryptoRandomBytes(t *testing.T) {
	// 测试基本功能
	b, err := GenerateCryptoRandomBytes(10)
	if err != nil {
		t.Errorf("GenerateCryptoRandomBytes returned error: %v", err)
	}
	if len(b) != 10 {
		t.Errorf("expected length 10, got %d", len(b))
	}

	// 测试长度为0
	b, err = GenerateCryptoRandomBytes(0)
	if err != nil {
		t.Errorf("GenerateCryptoRandomBytes(0) returned error: %v", err)
	}
	if len(b) != 0 {
		t.Errorf("expected empty slice for length 0, got length %d", len(b))
	}

	// 测试多次调用返回不同的值
	b1, _ := GenerateCryptoRandomBytes(16)
	b2, _ := GenerateCryptoRandomBytes(16)

	different := false
	for i := 0; i < 16; i++ {
		if b1[i] != b2[i] {
			different = true
			break
		}
	}

	if !different {
		t.Log("Warning: two GenerateCryptoRandomBytes calls returned the same value (low probability event)")
	}
}

func TestGenerateCryptoRandomInt(t *testing.T) {
	// 测试基本功能
	min, max := 5, 15
	v, err := GenerateCryptoRandomInt(min, max)
	if err != nil {
		t.Errorf("GenerateCryptoRandomInt returned error: %v", err)
	}

	if v < min || v >= max {
		t.Errorf("value %d out of range [%d, %d)", v, min, max)
	}

	// 测试无效范围
	_, err = GenerateCryptoRandomInt(10, 5)
	if err == nil {
		t.Error("expected error for invalid range, got nil")
	}

	// 测试相同边界
	_, err = GenerateCryptoRandomInt(5, 5)
	if err == nil {
		t.Error("expected error for equal min and max, got nil")
	}
}

func TestGenerateCryptoRandomFloat64(t *testing.T) {
	// 测试基本功能
	v, err := GenerateCryptoRandomFloat64()
	if err != nil {
		t.Errorf("GenerateCryptoRandomFloat64 returned error: %v", err)
	}

	if v < 0.0 || v >= 1.0 {
		t.Errorf("value %f out of range [0.0, 1.0)", v)
	}

	// 测试多次调用返回不同的值
	v1, _ := GenerateCryptoRandomFloat64()
	v2, _ := GenerateCryptoRandomFloat64()

	if v1 == v2 {
		t.Log("Warning: two GenerateCryptoRandomFloat64 calls returned the same value (low probability event)")
	}
}
