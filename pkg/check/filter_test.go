package check

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	// 测试过滤偶数
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	evenNumbers := Filter(func(n int) bool { return n%2 == 0 }, numbers...)

	expected := []int{2, 4, 6, 8, 10}
	assert.Equal(t, expected, evenNumbers, "过滤偶数结果应该匹配期望值")

	// 测试过滤字符串
	strings := []string{"apple", "banana", "cherry", "date", "elderberry"}
	longStrings := Filter(func(s string) bool { return len(s) > 5 }, strings...)

	expectedStrings := []string{"banana", "cherry", "elderberry"}
	assert.Equal(t, expectedStrings, longStrings, "过滤长字符串结果应该匹配期望值")

	// 测试空切片
	empty := Filter(func(n int) bool { return n > 0 }, []int{}...)
	assert.Empty(t, empty, "空切片过滤结果应该为空")
}
