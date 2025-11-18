package copy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Person struct {
	Name    string
	Age     int
	Hobbies []string
}

func TestDeepCopy(t *testing.T) {
	// 测试基本结构体复制
	src := Person{
		Name:    "Alice",
		Age:     30,
		Hobbies: []string{"reading", "swimming"},
	}

	var dst Person
	err := DeepCopy(src, &dst)

	assert.NoError(t, err, "深拷贝不应该返回错误")
	assert.Equal(t, src.Name, dst.Name, "姓名应该匹配")
	assert.Equal(t, src.Age, dst.Age, "年龄应该匹配")
	assert.Equal(t, len(src.Hobbies), len(dst.Hobbies), "爱好数量应该匹配")

	for i, hobby := range src.Hobbies {
		assert.Equal(t, hobby, dst.Hobbies[i], "爱好应该匹配")
	}

	// 验证深拷贝：修改目标对象不应影响源对象
	dst.Hobbies[0] = "dancing"
	assert.NotEqual(t, src.Hobbies[0], dst.Hobbies[0], "源对象不应受目标对象修改影响")

	// 测试基本类型
	srcInt := 42
	var dstInt int
	err = DeepCopy(srcInt, &dstInt)

	assert.NoError(t, err, "基本类型拷贝不应该返回错误")
	assert.Equal(t, srcInt, dstInt, "整数应该匹配")

	// 测试字符串
	srcStr := "hello world"
	var dstStr string
	err = DeepCopy(srcStr, &dstStr)

	assert.NoError(t, err, "字符串拷贝不应该返回错误")
	assert.Equal(t, srcStr, dstStr, "字符串应该匹配")

	// 测试切片
	srcSlice := []int{1, 2, 3, 4, 5}
	var dstSlice []int
	err = DeepCopy(srcSlice, &dstSlice)

	assert.NoError(t, err, "切片拷贝不应该返回错误")
	assert.Equal(t, len(srcSlice), len(dstSlice), "切片长度应该匹配")

	for i, v := range srcSlice {
		assert.Equal(t, v, dstSlice[i], "切片元素应该匹配")
	}

	// 验证深拷贝：修改目标切片不应影响源切片
	dstSlice[0] = 100
	assert.NotEqual(t, srcSlice[0], dstSlice[0], "源切片不应受目标切片修改影响")
}
