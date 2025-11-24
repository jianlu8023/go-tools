package csv

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestUser 用于测试的结构体
type TestUser struct {
	ID    int    `csv:"id"`
	Name  string `csv:"name"`
	Email string `csv:"email"`
	Age   int    `csv:"age"`
}

func TestNewClient(t *testing.T) {
	client := NewClient()
	assert.NotNil(t, client)
}

func TestMarshal(t *testing.T) {
	client := NewClient()

	users := []TestUser{
		{ID: 1, Name: "Alice", Email: "alice@example.com", Age: 25},
		{ID: 2, Name: "Bob", Email: "bob@example.com", Age: 30},
	}

	// 测试Marshal
	data, err := client.Marshal(users)
	assert.NoError(t, err)
	assert.NotEmpty(t, data)

	// 验证包含标题行和数据
	csvStr := string(data)
	assert.Contains(t, csvStr, "id,name,email,age")
	assert.Contains(t, csvStr, "Alice")
	assert.Contains(t, csvStr, "Bob")
}

func TestMarshalString(t *testing.T) {
	client := NewClient()

	users := []TestUser{
		{ID: 1, Name: "Alice", Email: "alice@example.com", Age: 25},
		{ID: 2, Name: "Bob", Email: "bob@example.com", Age: 30},
	}

	// 测试MarshalString
	csvStr, err := client.MarshalString(users)
	assert.NoError(t, err)
	assert.NotEmpty(t, csvStr)

	// 验证包含标题行和数据
	assert.Contains(t, csvStr, "id,name,email,age")
	assert.Contains(t, csvStr, "Alice")
	assert.Contains(t, csvStr, "Bob")
}

func TestMarshalFile(t *testing.T) {
	client := NewClient()

	users := []TestUser{
		{ID: 1, Name: "Alice", Email: "alice@example.com", Age: 25},
		{ID: 2, Name: "Bob", Email: "bob@example.com", Age: 30},
	}

	// 使用临时目录
	tempDir := t.TempDir()
	filename := filepath.Join(tempDir, "test_users.csv")

	// 测试MarshalFile
	err := client.MarshalFile(users, filename)
	assert.NoError(t, err)

	// 验证文件存在且内容正确
	data, err := os.ReadFile(filename)
	assert.NoError(t, err)
	assert.NotEmpty(t, data)

	csvStr := string(data)
	assert.Contains(t, csvStr, "id,name,email,age")
	assert.Contains(t, csvStr, "Alice")
	assert.Contains(t, csvStr, "Bob")
}

func TestMarshalWithoutHeaders(t *testing.T) {
	client := NewClient()

	users := []TestUser{
		{ID: 1, Name: "Alice", Email: "alice@example.com", Age: 25},
		{ID: 2, Name: "Bob", Email: "bob@example.com", Age: 30},
	}

	// 测试MarshalWithoutHeaders
	data, err := client.MarshalWithoutHeaders(users)
	assert.NoError(t, err)
	assert.NotEmpty(t, data)

	// 验证不包含标题行
	csvStr := string(data)
	assert.NotContains(t, csvStr, "id,name,email,age")
	assert.Contains(t, csvStr, "Alice")
	assert.Contains(t, csvStr, "Bob")
}

func TestUnmarshal(t *testing.T) {
	client := NewClient()

	// 准备CSV数据
	csvData := []byte("id,name,email,age\n1,Alice,alice@example.com,25\n2,Bob,bob@example.com,30\n")

	var users []TestUser
	err := client.Unmarshal(csvData, &users)
	assert.NoError(t, err)
	assert.Len(t, users, 2)

	// 验证数据正确性
	assert.Equal(t, 1, users[0].ID)
	assert.Equal(t, "Alice", users[0].Name)
	assert.Equal(t, "alice@example.com", users[0].Email)
	assert.Equal(t, 25, users[0].Age)

	assert.Equal(t, 2, users[1].ID)
	assert.Equal(t, "Bob", users[1].Name)
	assert.Equal(t, "bob@example.com", users[1].Email)
	assert.Equal(t, 30, users[1].Age)
}

func TestUnmarshalString(t *testing.T) {
	client := NewClient()

	// 准备CSV字符串
	csvStr := "id,name,email,age\n1,Alice,alice@example.com,25\n2,Bob,bob@example.com,30\n"

	var users []TestUser
	err := client.UnmarshalString(csvStr, &users)
	assert.NoError(t, err)
	assert.Len(t, users, 2)

	// 验证数据正确性
	assert.Equal(t, 1, users[0].ID)
	assert.Equal(t, "Alice", users[0].Name)
	assert.Equal(t, "alice@example.com", users[0].Email)
	assert.Equal(t, 25, users[0].Age)

	assert.Equal(t, 2, users[1].ID)
	assert.Equal(t, "Bob", users[1].Name)
	assert.Equal(t, "bob@example.com", users[1].Email)
	assert.Equal(t, 30, users[1].Age)
}

func TestUnmarshalFile(t *testing.T) {
	client := NewClient()

	// 使用临时目录
	tempDir := t.TempDir()
	filename := filepath.Join(tempDir, "test_users_read.csv")

	// 创建测试文件
	csvContent := "id,name,email,age\n1,Alice,alice@example.com,25\n2,Bob,bob@example.com,30\n"
	err := os.WriteFile(filename, []byte(csvContent), 0644)
	assert.NoError(t, err)

	var users []TestUser
	err = client.UnmarshalFile(filename, &users)
	assert.NoError(t, err)
	assert.Len(t, users, 2)

	// 验证数据正确性
	assert.Equal(t, 1, users[0].ID)
	assert.Equal(t, "Alice", users[0].Name)
	assert.Equal(t, "alice@example.com", users[0].Email)
	assert.Equal(t, 25, users[0].Age)

	assert.Equal(t, 2, users[1].ID)
	assert.Equal(t, "Bob", users[1].Name)
	assert.Equal(t, "bob@example.com", users[1].Email)
	assert.Equal(t, 30, users[1].Age)
}

func TestUnmarshalWithoutHeaders(t *testing.T) {
	client := NewClient()

	// 准备不包含标题行的CSV数据
	csvData := []byte("1,Alice,alice@example.com,25\n2,Bob,bob@example.com,30\n")

	var users []TestUser
	err := client.UnmarshalWithoutHeaders(csvData, &users)
	assert.NoError(t, err)
	assert.Len(t, users, 2)

	// 验证数据正确性
	assert.Equal(t, 1, users[0].ID)
	assert.Equal(t, "Alice", users[0].Name)
	assert.Equal(t, "alice@example.com", users[0].Email)
	assert.Equal(t, 25, users[0].Age)

	assert.Equal(t, 2, users[1].ID)
	assert.Equal(t, "Bob", users[1].Name)
	assert.Equal(t, "bob@example.com", users[1].Email)
	assert.Equal(t, 30, users[1].Age)
}

func TestDefaultClientFunctions(t *testing.T) {
	users := []TestUser{
		{ID: 1, Name: "Alice", Email: "alice@example.com", Age: 25},
		{ID: 2, Name: "Bob", Email: "bob@example.com", Age: 30},
	}

	// 测试默认客户端的Marshal函数
	data, err := Marshal(users)
	assert.NoError(t, err)
	assert.NotEmpty(t, data)

	// 测试默认客户端的MarshalString函数
	csvStr, err := MarshalString(users)
	assert.NoError(t, err)
	assert.NotEmpty(t, csvStr)

	// 测试默认客户端的Unmarshal函数
	var users2 []TestUser
	err = Unmarshal(data, &users2)
	assert.NoError(t, err)
	assert.Len(t, users2, 2)
}
