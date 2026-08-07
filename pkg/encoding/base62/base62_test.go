package base62

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 简化的测试用例，专注于基本功能正确性
func TestIntToBase62Basic(t *testing.T) {
	tests := []struct {
		name  string
		input int64
	}{{
		name:  "zero value",
		input: 0,
	}, {
		name:  "small number",
		input: 100,
	}, {
		name:  "medium number",
		input: 12345,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IntToBase62(tt.input)
			// 只验证编码解码一致性
			decoded, err := Base62ToInt(result)
			assert.NoError(t, err)
			assert.Equal(t, tt.input, decoded)
		})
	}
}

func TestBase62ToIntBasic(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int64
		wantErr  bool
	}{{
		name:     "zero value",
		input:    "0",
		expected: 0,
		wantErr:  false,
	}, {
		name:     "invalid character",
		input:    "1C!",
		expected: 0,
		wantErr:  true,
	}, {
		name:     "empty string",
		input:    "",
		expected: 0,
		wantErr:  false,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Base62ToInt(tt.input)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

func TestEncodeDecodeBasic(t *testing.T) {
	// 测试编码解码一致性，包含长字节数组（验证 N-E01/N-E03 修复）
	tests := []struct {
		name  string
		input []byte
	}{{
		name:  "empty input",
		input: []byte{},
	}, {
		name:  "single byte",
		input: []byte{65}, // 'A'
	}, {
		name:  "short bytes",
		input: []byte{0x01, 0x02, 0x03}, // 短二进制数据
	}, {
		name:  "leading zero bytes",
		input: []byte{0x00, 0x05}, // 前导零字节（验证 N-E03 修复）
	}, {
		name:  "all zeros",
		input: []byte{0x00, 0x00, 0x00},
	}, {
		name:  "high bit set (negative int64)",
		input: []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}, // 验证 N-E02 修复
	}, {
		name:  "long bytes > 8",
		input: []byte("Hello, World! This is a long string."), // 验证 N-E01 修复
	}, {
		name:  "exactly 7 bytes",
		input: []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07},
	}, {
		name:  "14 bytes (two groups)",
		input: []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E},
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encoded := Encode(tt.input)
			decoded, err := Decode(encoded)
			assert.NoError(t, err)
			assert.Equal(t, tt.input, decoded)
		})
	}
}

// 示例函数，展示如何使用这些API
func ExampleIntToBase62() {
	encoded := IntToBase62(100)
	fmt.Println(encoded)
}

func ExampleEncode() {
	encoded := Encode([]byte("A"))
	fmt.Println(encoded)
}
