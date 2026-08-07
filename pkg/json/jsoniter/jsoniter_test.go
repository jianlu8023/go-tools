package jsoniter

import (
	"errors"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 测试基本的JSON编解码功能
func TestBasicCodec(t *testing.T) {
	tests := []struct {
		name      string
		data      interface{}
		expectErr bool
	}{{
		name: "basic struct",
		data: struct {
			Name  string
			Age   int
			Email string
		}{"John Doe", 30, "john@example.com"},
		expectErr: false,
	}, {
		name: "nested struct",
		data: struct {
			User struct {
				Name string
			}
			Settings map[string]interface{}
		}{User: struct{ Name string }{"Jane"}, Settings: map[string]interface{}{"theme": "dark"}},
		expectErr: false,
	}, {
		name: "complex types",
		data: struct {
			Data []byte
			Num  *big.Int
		}{[]byte{0x1, 0x2, 0x3, 0xf}, big.NewInt(100)},
		expectErr: false,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 测试编码
			b, err := Marshal(tt.data)
			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, b)
			}

			// 测试解码 (使用相同类型)
			if !tt.expectErr {
				var decoded interface{}
				assert.NoError(t, Unmarshal(b, &decoded))
			}
		})
	}
}

// 测试MarshalPretty函数
func TestMarshalPretty(t *testing.T) {
	data := map[string]interface{}{
		"name": "John",
		"age":  30,
		"address": map[string]string{
			"street": "123 Main St",
			"city":   "Anytown",
		},
	}

	// 测试美化输出
	prettyJSON, err := MarshalPretty(data)
	assert.NoError(t, err)
	assert.Contains(t, string(prettyJSON), "\n")
	assert.Contains(t, string(prettyJSON), "  ") // 检查缩进

	// 验证美化后的JSON可以正常解析
	var decoded map[string]interface{}
	assert.NoError(t, Unmarshal(prettyJSON, &decoded))
	assert.Equal(t, data["name"], decoded["name"])
}

// 测试MarshalString函数
func TestMarshalString(t *testing.T) {
	data := map[string]interface{}{
		"name": "John",
		"age":  30,
	}

	jsonStr, err := MarshalString(data)
	assert.NoError(t, err)
	assert.NotEmpty(t, jsonStr)
	assert.Contains(t, jsonStr, "name")
	assert.Contains(t, jsonStr, "John")
	assert.Contains(t, jsonStr, "age")
	assert.Contains(t, jsonStr, "30")
}

// 测试UnmarshalString函数
func TestUnmarshalString(t *testing.T) {
	jsonStr := `{"name":"John","age":30}`

	var data map[string]interface{}
	err := UnmarshalString(jsonStr, &data)
	assert.NoError(t, err)
	assert.Equal(t, "John", data["name"])
	assert.Equal(t, float64(30), data["age"])

	// 测试空字符串
	emptyStr := ""
	var emptyData map[string]interface{}
	err = UnmarshalString(emptyStr, &emptyData)
	assert.NoError(t, err)
}

// 测试文件操作相关函数
func TestFileOperations(t *testing.T) {
	// 创建临时文件路径
	fileName := filepath.Join(t.TempDir(), "test.json")
	data := map[string]interface{}{
		"name": "John",
		"age":  30,
	}

	// 测试写入文件
	err := WriteToFile(fileName, data)
	assert.NoError(t, err)

	// 验证文件存在
	fileInfo, err := os.Stat(fileName)
	assert.NoError(t, err)
	assert.False(t, fileInfo.IsDir())

	// 测试从文件读取
	var readData map[string]interface{}
	err = ReadFromFile(fileName, &readData)
	assert.NoError(t, err)
	assert.Equal(t, data["name"], readData["name"])
	// JSON解析时数字会变成float64，需要进行类型转换比较
	assert.Equal(t, float64(data["age"].(int)), readData["age"])

	// 测试读取不存在的文件
	var nonExistentData map[string]interface{}
	err = ReadFromFile("non-existent-file.json", &nonExistentData)
	assert.Error(t, err)
	assert.True(t, errors.Is(err, ErrFileNotFound))

	// 测试读取无效JSON文件
	invalidFileName := filepath.Join(t.TempDir(), "invalid.json")
	os.WriteFile(invalidFileName, []byte("this is not JSON"), 0644)
	var invalidData map[string]interface{}
	err = ReadFromFile(invalidFileName, &invalidData)
	assert.Error(t, err)
	assert.True(t, errors.Is(err, ErrInvalidJSON))
}

// 测试Validate函数
func TestValidate(t *testing.T) {
	tests := []struct {
		name     string
		jsonStr  string
		expected bool
	}{{
		name:     "valid JSON",
		jsonStr:  `{"name":"John","age":30}`,
		expected: true,
	}, {
		name:     "invalid JSON",
		jsonStr:  `{"name":"John",}`, // 尾部有逗号
		expected: false,
	}, {
		name:     "empty string",
		jsonStr:  "",
		expected: false, // jsoniter.Valid返回false表示空字符串
	}, {
		name:     "null value",
		jsonStr:  `null`,
		expected: true,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Validate(tt.jsonStr)
			assert.Equal(t, tt.expected, result)
		})
	}
}

// 测试Compact函数
func TestCompact(t *testing.T) {
	prettyJSON := `{
  "name": "John",
  "age": 30,
  "address": {
    "street": "123 Main St",
    "city": "Anytown"
  }
}`

	// 预期的压缩结果
	expectedCompact := `{"name":"John","age":30,"address":{"street":"123 Main St","city":"Anytown"}}`

	// 测试压缩
	compactJSON, err := Compact(prettyJSON)
	assert.NoError(t, err)
	assert.Equal(t, expectedCompact, compactJSON)
	assert.NotContains(t, compactJSON, "\n")
	assert.NotContains(t, compactJSON, "  ")

	// 测试无效JSON的压缩
	invalidJSON := `{"name":"John",}`
	_, err = Compact(invalidJSON)
	assert.Error(t, err)

	// 测试压缩已压缩的JSON
	alreadyCompact := expectedCompact
	result, err := Compact(alreadyCompact)
	assert.NoError(t, err)
	assert.Equal(t, alreadyCompact, result) // 应该保持不变
}

// 示例函数，展示如何使用这些API
func ExampleMarshal() {
	data := map[string]interface{}{
		"name": "John",
		"age":  30,
	}
	bytes, _ := Marshal(data)
	fmt.Println(string(bytes))
}

func ExampleMarshalPretty() {
	data := map[string]interface{}{
		"name": "John",
		"age":  30,
	}
	bytes, _ := MarshalPretty(data)
	fmt.Println(string(bytes))
}
