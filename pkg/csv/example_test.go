package csv_test

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jianlu8023/go-tools/v2/pkg/csv"
)

// User 用户结构体，用于演示CSV处理
type User struct {
	ID    int    `csv:"id"`
	Name  string `csv:"name"`
	Email string `csv:"email"`
	Age   int    `csv:"age"`
}

// ExampleMarshal 演示如何将结构体切片转换为CSV格式
func ExampleMarshal() {
	users := []User{
		{ID: 1, Name: "Alice", Email: "alice@example.com", Age: 25},
		{ID: 2, Name: "Bob", Email: "bob@example.com", Age: 30},
		{ID: 3, Name: "Charlie", Email: "charlie@example.com", Age: 35},
	}

	// 将结构体切片转换为CSV字节数据
	data, err := csv.Marshal(users)
	if err != nil {
		fmt.Printf("Error marshaling CSV: %v\n", err)
		return
	}

	fmt.Printf("CSV data:\n%s\n", string(data))

	// Output:
	// CSV data:
	// id,name,email,age
	// 1,Alice,alice@example.com,25
	// 2,Bob,bob@example.com,30
	// 3,Charlie,charlie@example.com,35
}

// ExampleMarshalString 演示如何将结构体切片转换为CSV字符串
func ExampleMarshalString() {
	users := []User{
		{ID: 1, Name: "Alice", Email: "alice@example.com", Age: 25},
		{ID: 2, Name: "Bob", Email: "bob@example.com", Age: 30},
	}

	// 将结构体切片转换为CSV字符串
	csvStr, err := csv.MarshalString(users)
	if err != nil {
		fmt.Printf("Error marshaling CSV: %v\n", err)
		return
	}

	fmt.Printf("CSV string:\n%s\n", csvStr)

	// Output:
	// CSV string:
	// id,name,email,age
	// 1,Alice,alice@example.com,25
	// 2,Bob,bob@example.com,30
}

// ExampleMarshalFile 演示如何将结构体切片写入CSV文件
func ExampleMarshalFile() {
	users := []User{
		{ID: 1, Name: "Alice", Email: "alice@example.com", Age: 25},
		{ID: 2, Name: "Bob", Email: "bob@example.com", Age: 30},
	}

	// 使用临时目录
	tempDir, err := os.MkdirTemp("", "csv_example")
	if err != nil {
		fmt.Printf("Error creating temp dir: %v\n", err)
		return
	}
	defer os.RemoveAll(tempDir) // 清理临时目录

	filename := filepath.Join(tempDir, "example_users.csv")

	// 将结构体切片写入CSV文件
	err = csv.MarshalFile(users, filename)
	if err != nil {
		fmt.Printf("Error marshaling CSV to file: %v\n", err)
		return
	}

	fmt.Println("CSV data written to file")

	// 读取并显示文件内容
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	fmt.Printf("File content:\n%s\n", string(data))

	// Output:
	// CSV data written to file
	// File content:
	// id,name,email,age
	// 1,Alice,alice@example.com,25
	// 2,Bob,bob@example.com,30
}

// ExampleUnmarshal 演示如何将CSV数据解析为结构体切片
func ExampleUnmarshal() {
	// CSV数据
	csvData := []byte("id,name,email,age\n1,Alice,alice@example.com,25\n2,Bob,bob@example.com,30\n")

	var users []User
	err := csv.Unmarshal(csvData, &users)
	if err != nil {
		fmt.Printf("Error unmarshaling CSV: %v\n", err)
		return
	}

	fmt.Printf("Parsed users:\n")
	for _, user := range users {
		fmt.Printf("ID: %d, Name: %s, Email: %s, Age: %d\n", user.ID, user.Name, user.Email, user.Age)
	}

	// Output:
	// Parsed users:
	// ID: 1, Name: Alice, Email: alice@example.com, Age: 25
	// ID: 2, Name: Bob, Email: bob@example.com, Age: 30
}

// ExampleUnmarshalString 演示如何将CSV字符串解析为结构体切片
func ExampleUnmarshalString() {
	// CSV字符串
	csvStr := "id,name,email,age\n1,Alice,alice@example.com,25\n2,Bob,bob@example.com,30\n"

	var users []User
	err := csv.UnmarshalString(csvStr, &users)
	if err != nil {
		fmt.Printf("Error unmarshaling CSV string: %v\n", err)
		return
	}

	fmt.Printf("Parsed users:\n")
	for _, user := range users {
		fmt.Printf("ID: %d, Name: %s, Email: %s, Age: %d\n", user.ID, user.Name, user.Email, user.Age)
	}

	// Output:
	// Parsed users:
	// ID: 1, Name: Alice, Email: alice@example.com, Age: 25
	// ID: 2, Name: Bob, Email: bob@example.com, Age: 30
}

// ExampleUnmarshalFile 演示如何从CSV文件中读取数据并解析为结构体切片
func ExampleUnmarshalFile() {
	// 使用临时目录
	tempDir, err := os.MkdirTemp("", "csv_example_read")
	if err != nil {
		fmt.Printf("Error creating temp dir: %v\n", err)
		return
	}
	defer os.RemoveAll(tempDir) // 清理临时目录

	filename := filepath.Join(tempDir, "example_read_users.csv")

	// 创建测试文件
	csvContent := "id,name,email,age\n1,Alice,alice@example.com,25\n2,Bob,bob@example.com,30\n"
	err = os.WriteFile(filename, []byte(csvContent), 0644)
	if err != nil {
		fmt.Printf("Error creating test file: %v\n", err)
		return
	}

	var users []User
	err = csv.UnmarshalFile(filename, &users)
	if err != nil {
		fmt.Printf("Error unmarshaling CSV file: %v\n", err)
		return
	}

	fmt.Println("Parsed users from file:")
	for _, user := range users {
		fmt.Printf("ID: %d, Name: %s, Email: %s, Age: %d\n", user.ID, user.Name, user.Email, user.Age)
	}

	// Output:
	// Parsed users from file:
	// ID: 1, Name: Alice, Email: alice@example.com, Age: 25
	// ID: 2, Name: Bob, Email: bob@example.com, Age: 30
}

// ExampleMarshalWithoutHeaders 演示如何将结构体切片转换为不包含标题行的CSV格式
func ExampleMarshalWithoutHeaders() {
	users := []User{
		{ID: 1, Name: "Alice", Email: "alice@example.com", Age: 25},
		{ID: 2, Name: "Bob", Email: "bob@example.com", Age: 30},
	}

	// 将结构体切片转换为不包含标题行的CSV字节数据
	data, err := csv.MarshalWithoutHeaders(users)
	if err != nil {
		fmt.Printf("Error marshaling CSV without headers: %v\n", err)
		return
	}

	fmt.Printf("CSV data without headers:\n%s\n", string(data))

	// Output:
	// CSV data without headers:
	// 1,Alice,alice@example.com,25
	// 2,Bob,bob@example.com,30
}

// ExampleUnmarshalWithoutHeaders 演示如何将不包含标题行的CSV数据解析为结构体切片
func ExampleUnmarshalWithoutHeaders() {
	// 不包含标题行的CSV数据
	csvData := []byte("1,Alice,alice@example.com,25\n2,Bob,bob@example.com,30\n")

	var users []User
	err := csv.UnmarshalWithoutHeaders(csvData, &users)
	if err != nil {
		fmt.Printf("Error unmarshaling CSV without headers: %v\n", err)
		return
	}

	fmt.Println("Parsed users from CSV without headers:")
	for _, user := range users {
		fmt.Printf("ID: %d, Name: %s, Email: %s, Age: %d\n", user.ID, user.Name, user.Email, user.Age)
	}

	// Output:
	// Parsed users from CSV without headers:
	// ID: 1, Name: Alice, Email: alice@example.com, Age: 25
	// ID: 2, Name: Bob, Email: bob@example.com, Age: 30
}
