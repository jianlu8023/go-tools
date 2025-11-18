package check

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	// 测试整数平方
	numbers := []int{1, 2, 3, 4, 5}
	squared := Map(func(n int) int { return n * n }, numbers...)

	expected := []int{1, 4, 9, 16, 25}
	assert.Equal(t, expected, squared, "整数平方结果应该匹配期望值")

	// 测试字符串长度
	strings := []string{"apple", "banana", "cherry"}
	lengths := Map(func(s string) int { return len(s) }, strings...)

	expectedLengths := []int{5, 6, 6}
	assert.Equal(t, expectedLengths, lengths, "字符串长度结果应该匹配期望值")

	// 测试类型转换
	ints := []int{1, 2, 3}
	strs := Map(func(n int) string { return string(rune(n + 64)) }, ints...) // 1->'A', 2->'B', 3->'C'

	expectedStrs := []string{"A", "B", "C"}
	assert.Equal(t, expectedStrs, strs, "类型转换结果应该匹配期望值")

	// 测试空切片
	empty := Map(func(n int) int { return n * 2 }, []int{}...)
	assert.Empty(t, empty, "空切片映射结果应该为空")
}
