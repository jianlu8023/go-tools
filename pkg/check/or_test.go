package check

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrUnwish(t *testing.T) {
	// 测试返回第一个非unwish且非零值的元素
	result1 := OrUnwish(0, 0, 0, 5, 10)
	assert.Equal(t, 5, result1, "应该返回第一个非unwish且非零值的元素")

	// 测试所有元素都是unwish值
	result2 := OrUnwish(5, 5, 5, 5)
	assert.Equal(t, 0, result2, "当所有元素都是unwish值时应该返回零值")

	// 测试字符串类型
	result3 := OrUnwish("", "", "hello", "world")
	assert.Equal(t, "hello", result3, "应该返回第一个非空字符串")

	// 测试所有元素都是空字符串
	result4 := OrUnwish("", "", "")
	assert.Equal(t, "", result4, "当所有元素都是空字符串时应该返回空字符串")
}

func TestOr(t *testing.T) {
	// 测试返回第一个非零值
	result1 := Or(0, 0, 5, 10)
	assert.Equal(t, 5, result1, "应该返回第一个非零值")

	// 测试所有元素都是零值
	result2 := Or(0, 0, 0)
	assert.Equal(t, 0, result2, "当所有元素都是零值时应该返回零值")

	// 测试字符串类型
	result3 := Or("", "", "hello", "world")
	assert.Equal(t, "hello", result3, "应该返回第一个非空字符串")

	// 测试所有元素都是空字符串
	result4 := Or("", "", "")
	assert.Equal(t, "", result4, "当所有元素都是空字符串时应该返回空字符串")
}
