package pretty

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestMarshal 测试基本的JSON序列化功能
func TestMarshal(t *testing.T) {
	// 准备测试数据
	testData := map[string]interface{}{
		"name":   "test",
		"age":    30,
		"active": true,
		"scores": []int{95, 85, 90},
		"address": map[string]interface{}{
			"city":    "New York",
			"country": "USA",
		},
		"null_value": nil,
	}

	// 测试默认格式化器
	result, err := Marshal(testData)
	assert.NoError(t, err)
	assert.NotEmpty(t, result)

	// 验证格式化后的JSON包含所有字段
	resultStr := string(result)
	assert.Contains(t, resultStr, "name")
	assert.Contains(t, resultStr, "age")
	assert.Contains(t, resultStr, "active")
	assert.Contains(t, resultStr, "scores")
	assert.Contains(t, resultStr, "address")
	assert.Contains(t, resultStr, "null_value")
}

// TestFormat 测试格式化已有的JSON字符串
func TestFormat(t *testing.T) {
	// 准备测试数据
	rawJSON := []byte(`{"name":"test","age":30,"active":true}`)

	result, err := Format(rawJSON)
	assert.NoError(t, err)
	assert.NotEmpty(t, result)

	// 验证格式化后的JSON包含所有字段
	resultStr := string(result)
	assert.Contains(t, resultStr, "name")
	assert.Contains(t, resultStr, "age")
	assert.Contains(t, resultStr, "active")

	// 验证格式化后的JSON有换行和缩进
	assert.Contains(t, resultStr, "\n")
	assert.Contains(t, resultStr, "  ") // 验证默认缩进
}

// TestFormatter_DisabledColor 测试禁用颜色功能
func TestFormatter_DisabledColor(t *testing.T) {
	// 准备测试数据
	testData := map[string]interface{}{
		"name": "test",
	}

	// 创建禁用颜色的格式化器
	formatter := NewFormatter()
	formatter.DisabledColor = true

	result, err := formatter.Marshal(testData)
	assert.NoError(t, err)

	// 验证结果中不包含ANSI颜色代码
	resultStr := string(result)
	assert.NotContains(t, resultStr, "\033")
}

// TestFormatter_StringMaxLength 测试字符串最大长度限制
func TestFormatter_StringMaxLength(t *testing.T) {
	// 准备测试数据
	testData := map[string]interface{}{
		"long_string": "this is a very long string that should be truncated",
	}

	// 创建设置了字符串最大长度的格式化器
	formatter := NewFormatter()
	formatter.StringMaxLength = 10

	result, err := formatter.Marshal(testData)
	assert.NoError(t, err)

	// 验证字符串被截断
	resultStr := string(result)
	assert.Contains(t, resultStr, "this is a ...")
}

// TestAllDataTypes 测试所有数据类型的格式化
func TestAllDataTypes(t *testing.T) {
	// 准备包含所有数据类型的测试数据
	testData := map[string]interface{}{
		"string":  "hello world",
		"number":  42.42,
		"integer": 42,
		"boolean": true,
		"null":    nil,
		"array":   []interface{}{"item1", 123, false, nil},
		"object":  map[string]interface{}{"key": "value"},
	}

	// 使用默认格式化器
	result, err := Marshal(testData)
	assert.NoError(t, err)
	assert.NotEmpty(t, result)

	// 验证所有数据类型都被正确格式化
	resultStr := string(result)
	assert.Contains(t, resultStr, "hello world")
	assert.Contains(t, resultStr, "42.42")
	assert.Contains(t, resultStr, "42")
	assert.Contains(t, resultStr, "true")
	assert.Contains(t, resultStr, "null")
	assert.Contains(t, resultStr, "item1")
	assert.Contains(t, resultStr, "key")
	assert.Contains(t, resultStr, "value")
}

// TestEmptyStructures 测试空的对象和数组
func TestEmptyStructures(t *testing.T) {
	// 准备包含空结构的测试数据
	testData := map[string]interface{}{
		"empty_array":  []interface{}{},
		"empty_object": map[string]interface{}{},
	}

	result, err := Marshal(testData)
	assert.NoError(t, err)

	resultStr := string(result)
	assert.Contains(t, resultStr, "[]")
	assert.Contains(t, resultStr, "{}")
}

// TestFormatInvalidJSON 测试格式化无效的JSON
func TestFormatInvalidJSON(t *testing.T) {
	// 准备无效的JSON
	invalidJSON := []byte(`{invalid json}`)

	result, err := Format(invalidJSON)
	assert.Error(t, err)
	assert.Nil(t, result)
}

// TestFormatter_Marshal 展示基本的JSON格式化功能
func TestFormatter_Marshal(t *testing.T) {
	// 准备测试数据
	testData := map[string]interface{}{
		"name":   "测试用户",
		"age":    30,
		"active": true,
		"scores": []int{95, 85, 90},
		"address": map[string]interface{}{
			"city":    "北京",
			"country": "中国",
		},
		"null_value": nil,
	}

	// 使用默认格式化器
	result, err := Marshal(testData)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(result))
	// 输出将是格式化后的彩色JSON
}

// ExampleFormatter_Configured 展示如何配置Formatter的各种选项
func TestExampleFormatter_Configured(t *testing.T) {
	// 准备测试数据
	testData := map[string]interface{}{
		"name":        "长字符串示例",
		"description": "这是一个非常长的字符串，用于测试字符串截断功能。当设置了StringMaxLength后，超过长度的部分会被截断并显示省略号。",
	}

	// 创建配置了各种选项的格式化器
	formatter := NewFormatter()
	formatter.StringMaxLength = 20 // 限制字符串最大长度为20
	formatter.Indent = 4           // 设置缩进为4个空格
	formatter.Newline = "\r\n"     // 设置换行符为\r\n
	result, err := formatter.Marshal(testData)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(result))
	// 输出将是使用自定义配置的彩色JSON
}

// ExampleFormatter_DisableColor 展示如何禁用彩色输出
func TestExampleFormatter_DisableColor(t *testing.T) {
	// 准备测试数据
	testData := map[string]interface{}{
		"name":  "测试",
		"value": 42,
	}

	// 创建禁用颜色的格式化器
	formatter := NewFormatter()
	formatter.DisabledColor = true

	result, err := formatter.Marshal(testData)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(result))
	// 输出将是没有颜色的格式化JSON
}

// ExampleFormat 展示如何格式化已有的JSON字符串
func TestExampleFormat(t *testing.T) {
	// 已有的JSON字符串
	rawJSON := []byte(`{"name":"test","age":30,"skills":["golang","python","java"]}`)

	result, err := Format(rawJSON)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(result))
	// 输出将是格式化后的彩色JSON
}
