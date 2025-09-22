package base64

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToBase64(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected string
	}{{
		name:     "basic string",
		input:    []byte("hello"),
		expected: "aGVsbG8=",
	}, {
		name:     "empty input",
		input:    []byte{},
		expected: "",
	}, {
		name:     "binary data",
		input:    []byte{0x00, 0xFF, 0xAA, 0x55},
		expected: "AP+qVQ==",
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToBase64(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestToByte(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []byte
		wantErr  bool
	}{{
		name:     "basic decoding",
		input:    "aGVsbG8=",
		expected: []byte("hello"),
		wantErr:  false,
	}, {
		name:     "empty input",
		input:    "",
		expected: []byte{},
		wantErr:  false,
	}, {
		name:     "invalid base64",
		input:    "invalid!!!",
		expected: nil,
		wantErr:  true,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ToByte(tt.input)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

func TestToBase64URL(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected string
	}{{
		name:     "basic URL safe",
		input:    []byte("hello/world"),
		expected: "aGVsbG8vd29ybGQ=",
	}, {
		name:     "binary data URL safe",
		input:    []byte{0x00, 0xFF, 0xAA, 0x55}, // 包含URL安全字符
		expected: "AP-qVQ==",
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToBase64URL(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestToByteURL(t *testing.T) {
	// 测试URL安全的base64解码
	tests := []struct {
		name     string
		input    string
		expected []byte
		wantErr  bool
	}{{
		name:     "basic URL decoding",
		input:    "aGVsbG8vd29ybGQ=",
		expected: []byte("hello/world"),
		wantErr:  false,
	}, {
		name:     "URL safe chars",
		input:    "AP-qVQ==", // 使用了 '-' 而不是 '+'
		expected: []byte{0x00, 0xFF, 0xAA, 0x55},
		wantErr:  false,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ToByteURL(tt.input)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

func TestToBase64NoPadding(t *testing.T) {
	// 测试不带填充的base64编码
	tests := []struct {
		name     string
		input    []byte
		expected string
	}{{
		name:     "without padding",
		input:    []byte("hello"),
		expected: "aGVsbG8", // 去掉了末尾的 '='
	}, {
		name:     "binary data no padding",
		input:    []byte{0x00, 0xFF, 0xAA}, // 长度是3的倍数
		expected: "AP+q",                   // 不需要填充
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToBase64NoPadding(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestToBase64URLNoPadding(t *testing.T) {
	// 测试不带填充的URL安全base64编码
	tests := []struct {
		name     string
		input    []byte
		expected string
	}{{
		name:     "URL safe no padding",
		input:    []byte("hello/world"),
		expected: "aGVsbG8vd29ybGQ", // 去掉了末尾的 '='
	}, {
		name:     "URL safe binary no padding",
		input:    []byte{0x00, 0xFF, 0xAA, 0x55, 0x11}, // 长度不是3的倍数
		expected: "AP-qVRE",                            // 修正为实际输出
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToBase64URLNoPadding(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

// 示例函数，展示如何使用这些API
func ExampleToBase64() {
	encoded := ToBase64([]byte("hello world"))
	fmt.Println(encoded)
	// Output: aGVsbG8gd29ybGQ=
}

func ExampleToBase64URLNoPadding() {
	encoded := ToBase64URLNoPadding([]byte("hello/world"))
	fmt.Println(encoded)
	// Output: aGVsbG8vd29ybGQ
}
