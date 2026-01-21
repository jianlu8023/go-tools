package random

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateUID(t *testing.T) {
	// 测试默认字母表和不同大小
	id, err := GenerateUID(defaultAlphabet, 10)
	assert.NoError(t, err)
	assert.Len(t, id, 10)
	for _, char := range id {
		assert.Contains(t, defaultAlphabet, string(char))
	}

	// 测试自定义字母表
	customAlphabet := "abc"
	id2, err := GenerateUID(customAlphabet, 5)
	assert.NoError(t, err)
	assert.Len(t, id2, 5)
	for _, char := range id2 {
		assert.Contains(t, customAlphabet, string(char))
	}

	// 测试错误情况 - 空字母表
	_, err = GenerateUID("", 10)
	assert.Error(t, err)

	// 测试错误情况 - 超长字母表
	longAlphabet := strings.Repeat("a", 256)
	_, err = GenerateUID(longAlphabet, 10)
	assert.Error(t, err)

	// 测试错误情况 - 零大小
	_, err = GenerateUID(defaultAlphabet, 0)
	assert.Error(t, err)

	// 测试错误情况 - 负大小
	_, err = GenerateUID(defaultAlphabet, -1)
	assert.Error(t, err)
}

func TestNewUID(t *testing.T) {
	// 测试默认大小
	id, err := NewUID()
	assert.NoError(t, err)
	assert.Len(t, id, defaultSize)

	// 测试自定义大小
	id2, err := NewUID(32)
	assert.NoError(t, err)
	assert.Len(t, id2, 32)

	// 测试所有字符都在默认字母表中
	for _, char := range id {
		assert.Contains(t, defaultAlphabet, string(char))
	}

	// 验证生成的字符都在默认字母表范围内
	for _, char := range id {
		assert.Contains(t, defaultAlphabet, string(char))
	}
}

func TestUIDCharacteristics(t *testing.T) {
	// 生成多个ID并验证它们都是唯一的
	ids := make(map[string]bool)
	for i := 0; i < 100; i++ {
		id, err := NewUID(16)
		assert.NoError(t, err)
		assert.Len(t, id, 16)
		assert.False(t, ids[id], "ID should be unique")
		ids[id] = true
	}
}
