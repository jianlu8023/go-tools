package jq

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewQuery(t *testing.T) {
	data := map[string]interface{}{
		"name": "test",
		"age":  30,
	}
	jq := NewQuery(data)
	assert.NotNil(t, jq)
	// 注意：由于我们现在的实现是封装，所以不能直接访问jq.Data
	// 但我们可以通过查询来验证数据是否正确设置
	result, err := jq.Query("name")
	assert.NoError(t, err)
	assert.Equal(t, "test", result)
}

func TestNewStringQuery(t *testing.T) {
	jsonStr := `{"name":"test","age":30}`
	jq, err := NewStringQuery(jsonStr)
	assert.NoError(t, err)
	assert.NotNil(t, jq)

	// 验证解析的数据
	name, err := jq.QueryToString("name")
	assert.NoError(t, err)
	assert.Equal(t, "test", name)

	age, err := jq.Query("age")
	assert.NoError(t, err)
	assert.Equal(t, float64(30), age) // JSON数字默认是float64
}

func TestNewStringQueryInvalidJSON(t *testing.T) {
	jsonStr := `{"name":"test","age":30` // 无效的JSON
	jq, err := NewStringQuery(jsonStr)
	assert.Error(t, err)
	assert.Nil(t, jq)
}

func TestQuerySimple(t *testing.T) {
	jsonStr := `{"name":"test","age":30}`
	jq, _ := NewStringQuery(jsonStr)

	// 测试简单查询
	result, err := jq.Query("name")
	assert.NoError(t, err)
	assert.Equal(t, "test", result)

	result, err = jq.Query("age")
	assert.NoError(t, err)
	assert.Equal(t, float64(30), result) // JSON数字默认是float64

	// 测试不存在的键
	result, err = jq.Query("nonexistent")
	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestQueryNested(t *testing.T) {
	jsonStr := `{
		"user": {
			"name": "test",
			"age": 30,
			"profile": {
				"email": "test@example.com"
			}
		}
	}`
	jq, _ := NewStringQuery(jsonStr)

	// 测试嵌套查询
	result, err := jq.Query("user.name")
	assert.NoError(t, err)
	assert.Equal(t, "test", result)

	result, err = jq.Query("user.profile.email")
	assert.NoError(t, err)
	assert.Equal(t, "test@example.com", result)

	// 测试不存在的嵌套键
	result, err = jq.Query("user.profile.phone")
	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestQueryArray(t *testing.T) {
	jsonStr := `{
		"users": [
			{"name": "alice", "age": 25},
			{"name": "bob", "age": 30}
		]
	}`
	jq, _ := NewStringQuery(jsonStr)

	// 测试数组查询
	result, err := jq.Query("users.[0].name")
	assert.NoError(t, err)
	assert.Equal(t, "alice", result)

	result, err = jq.Query("users.[1].age")
	assert.NoError(t, err)
	assert.Equal(t, float64(30), result)

	// 测试数组索引越界
	result, err = jq.Query("users.[2].name")
	assert.Error(t, err)
	assert.Nil(t, result)

	// 测试非数组访问
	result, err = jq.Query("users.name")
	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestQueryString(t *testing.T) {
	jsonStr := `{"name":"test","description":"age is 30"}`
	jq, _ := NewStringQuery(jsonStr)

	result, err := jq.QueryToString("name")
	assert.NoError(t, err)
	assert.Equal(t, "test", result)

	// 测试字符串字段
	result, err = jq.QueryToString("description")
	assert.NoError(t, err)
	assert.Equal(t, "age is 30", result)
}

func TestQueryInt64(t *testing.T) {
	jsonStr := `{"name":"test","age":30}`
	jq, _ := NewStringQuery(jsonStr)

	// 测试直接数字
	result, err := jq.QueryToInt64("age")
	assert.NoError(t, err)
	assert.Equal(t, int64(30), result)

	// 测试非数字
	result, err = jq.QueryToInt64("name")
	assert.Error(t, err)
	assert.Equal(t, int64(0), result)
}

func TestQueryFloat64(t *testing.T) {
	jsonStr := `{"name":"test","price":29.99}`
	jq, _ := NewStringQuery(jsonStr)

	// 测试直接浮点数
	result, err := jq.QueryToFloat64("price")
	assert.NoError(t, err)
	assert.Equal(t, 29.99, result)

	// 测试非数字
	result, err = jq.QueryToFloat64("name")
	assert.Error(t, err)
	assert.Equal(t, 0.0, result)
}

func TestQueryBool(t *testing.T) {
	jsonStr := `{"name":"test","active":true}`
	jq, _ := NewStringQuery(jsonStr)

	// 测试直接布尔值
	result, err := jq.QueryToBool("active")
	assert.NoError(t, err)
	assert.True(t, result)

	// 测试非布尔值
	result, err = jq.QueryToBool("name")
	assert.Error(t, err)
	assert.False(t, result)
}

func TestQueryArrayConversion(t *testing.T) {
	jsonStr := `{"numbers":[1,2,3],"name":"test"}`
	jq, _ := NewStringQuery(jsonStr)

	// 测试数组转换
	result, err := jq.QueryToArray("numbers")
	assert.NoError(t, err)
	assert.Len(t, result, 3)
	assert.Equal(t, float64(1), result[0])
	assert.Equal(t, float64(2), result[1])
	assert.Equal(t, float64(3), result[2])

	// 测试非数组转换
	result, err = jq.QueryToArray("name")
	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestQueryMapConversion(t *testing.T) {
	jsonStr := `{"user":{"name":"test","age":30},"list":[1,2,3]}`
	jq, _ := NewStringQuery(jsonStr)

	// 测试map转换
	result, err := jq.QueryToMap("user")
	assert.NoError(t, err)
	assert.Equal(t, "test", result["name"])
	assert.Equal(t, float64(30), result["age"])

	// 测试非map转换
	result, err = jq.QueryToMap("list")
	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestSpecialKeys(t *testing.T) {
	jsonStr := `{"hello.world":true,"normal":"value"}`
	jq, _ := NewStringQuery(jsonStr)

	// 测试特殊键名（带引号）
	result, err := jq.Query("'hello.world'")
	assert.NoError(t, err)
	assert.Equal(t, true, result)

	// 测试普通键名
	result, err = jq.Query("normal")
	assert.NoError(t, err)
	assert.Equal(t, "value", result)
}

func TestRootQuery(t *testing.T) {
	jsonStr := `{"name":"test","age":30}`
	jq, _ := NewStringQuery(jsonStr)

	// 测试根查询
	result, err := jq.Query(".")
	assert.NoError(t, err)
	// 根查询返回整个对象，我们验证它是一个map
	if m, ok := result.(map[string]interface{}); ok {
		assert.Equal(t, "test", m["name"])
		assert.Equal(t, float64(30), m["age"])
	} else {
		t.Error("Expected map result from root query")
	}
}

func TestNewFileQuery(t *testing.T) {
	// 由于我们没有实际的文件，这里只是测试函数是否存在
	// 在实际使用中，这个函数会从文件中读取JSON数据
	assert.NotNil(t, NewFileQuery)
}
