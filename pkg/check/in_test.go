package check

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIn(t *testing.T) {
	// 测试整数存在于数组中
	result1 := In(3, 1, 2, 3, 4, 5)
	assert.True(t, result1, "3应该在数组[1, 2, 3, 4, 5]中")

	// 测试整数不存在于数组中
	result2 := In(6, 1, 2, 3, 4, 5)
	assert.False(t, result2, "6不应该在数组[1, 2, 3, 4, 5]中")

	// 测试字符串存在于数组中
	result3 := In("hello", "world", "hello", "go")
	assert.True(t, result3, "'hello'应该在数组['world', 'hello', 'go']中")

	// 测试字符串不存在于数组中
	result4 := In("java", "world", "hello", "go")
	assert.False(t, result4, "'java'不应该在数组['world', 'hello', 'go']中")

	// 测试空数组
	result5 := In(1)
	assert.False(t, result5, "1不应该在空数组中")

	// 测试单个元素数组
	result6 := In(42, 42)
	assert.True(t, result6, "42应该在数组[42]中")

	// 测试浮点数
	result7 := In(3.14, 1.0, 2.5, 3.14, 4.0)
	assert.True(t, result7, "3.14应该在数组[1.0, 2.5, 3.14, 4.0]中")
}
