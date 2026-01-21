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

func TestArrayIn(t *testing.T) {
	// 测试整数存在于数组中
	exists1, index1 := ArrayIn(3, []int{1, 2, 3, 4, 5})
	assert.True(t, exists1, "3应该在数组[1, 2, 3, 4, 5]中")
	assert.Equal(t, 2, index1, "3在数组[1, 2, 3, 4, 5]中的索引应该是2")

	// 测试整数不存在于数组中
	exists2, index2 := ArrayIn(6, []int{1, 2, 3, 4, 5})
	assert.False(t, exists2, "6不应该在数组[1, 2, 3, 4, 5]中")
	assert.Equal(t, -1, index2, "6不在数组[1, 2, 3, 4, 5]中，索引应该是-1")

	// 测试字符串存在于数组中
	exists3, index3 := ArrayIn("hello", []string{"world", "hello", "go"})
	assert.True(t, exists3, "'hello'应该在数组['world', 'hello', 'go']中")
	assert.Equal(t, 1, index3, "'hello'在数组['world', 'hello', 'go']中的索引应该是1")

	// 测试字符串不存在于数组中
	exists4, index4 := ArrayIn("java", []string{"world", "hello", "go"})
	assert.False(t, exists4, "'java'不应该在数组['world', 'hello', 'go']中")
	assert.Equal(t, -1, index4, "'java'不在数组['world', 'hello', 'go']中，索引应该是-1")

	// 测试空数组
	exists5, index5 := ArrayIn(1, []int{})
	assert.False(t, exists5, "1不应该在空数组中")
	assert.Equal(t, -1, index5, "空数组中查找任何元素，索引都应该是-1")

	// 测试单个元素数组（存在）
	exists6, index6 := ArrayIn(42, []int{42})
	assert.True(t, exists6, "42应该在数组[42]中")
	assert.Equal(t, 0, index6, "42在数组[42]中的索引应该是0")

	// 测试单个元素数组（不存在）
	exists7, index7 := ArrayIn(43, []int{42})
	assert.False(t, exists7, "43不应该在数组[42]中")
	assert.Equal(t, -1, index7, "43不在数组[42]中，索引应该是-1")

	// 测试浮点数
	exists8, index8 := ArrayIn(3.14, []float64{1.0, 2.5, 3.14, 4.0})
	assert.True(t, exists8, "3.14应该在数组[1.0, 2.5, 3.14, 4.0]中")
	assert.Equal(t, 2, index8, "3.14在数组[1.0, 2.5, 3.14, 4.0]中的索引应该是2")

	// 测试第一个元素
	exists9, index9 := ArrayIn("first", []string{"first", "second", "third"})
	assert.True(t, exists9, "'first'应该在数组['first', 'second', 'third']中")
	assert.Equal(t, 0, index9, "'first'在数组['first', 'second', 'third']中的索引应该是0")

	// 测试最后一个元素
	exists10, index10 := ArrayIn("last", []string{"first", "second", "last"})
	assert.True(t, exists10, "'last'应该在数组['first', 'second', 'last']中")
	assert.Equal(t, 2, index10, "'last'在数组['first', 'second', 'last']中的索引应该是2")
}
