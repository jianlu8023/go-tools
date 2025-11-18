package check

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReduce(t *testing.T) {
	// 测试整数求和
	numbers := []int{1, 2, 3, 4, 5}
	sum := Reduce(func(a, b int) int { return a + b }, numbers...)
	assert.Equal(t, 15, sum, "整数求和结果应该为15")

	// 测试整数求积
	product := Reduce(func(a, b int) int { return a * b }, numbers...)
	assert.Equal(t, 120, product, "整数求积结果应该为120")

	// 测试字符串拼接
	strings := []string{"hello", " ", "world", "!"}
	concat := Reduce(func(a, b string) string { return a + b }, strings...)
	assert.Equal(t, "hello world!", concat, "字符串拼接结果应该为'hello world!'")

	// 测试单个元素
	single := Reduce(func(a, b int) int { return a + b }, 42)
	assert.Equal(t, 42, single, "单个元素归约结果应该为42")

	// 测试空切片
	var empty []int
	zero := Reduce(func(a, b int) int { return a + b }, empty...)
	assert.Equal(t, 0, zero, "空切片归约结果应该为0")

	// 测试最大值
	max := Reduce(func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}, numbers...)
	assert.Equal(t, 5, max, "最大值结果应该为5")
}
