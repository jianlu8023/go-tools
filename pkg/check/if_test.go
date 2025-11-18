package check

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIF(t *testing.T) {
	// 测试返回true的情况
	result1 := IF(true, "yes", "no")
	assert.Equal(t, "yes", result1, "当条件为true时应该返回第一个值")

	// 测试返回false的情况
	result2 := IF(false, "yes", "no")
	assert.Equal(t, "no", result2, "当条件为false时应该返回第二个值")

	// 测试数字类型
	result3 := IF(5 > 3, 100, -100)
	assert.Equal(t, 100, result3, "应该返回正确的数字值")

	// 测试浮点数类型
	result4 := IF(3.14 < 2.0, 1.5, 2.5)
	assert.Equal(t, 2.5, result4, "应该返回正确的浮点数值")

	// 测试布尔类型
	result5 := IF(true, false, true)
	assert.False(t, result5, "应该返回false")
}
