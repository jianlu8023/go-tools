package blake2

import (
	"os"
	"path/filepath"
	"testing"
)

func TestHash(t *testing.T) {
	data := []byte("Hello, Blake2!")

	// 测试Blake2b-256
	hashValue, err := Hash(data, Blake2b_256)
	if err != nil {
		t.Fatalf("Blake2b-256哈希计算失败: %v", err)
	}

	if len(hashValue) == 0 {
		t.Error("Blake2b-256哈希值不能为空")
	}

	// 测试Blake2b-512
	hashValue, err = Hash(data, Blake2b_512)
	if err != nil {
		t.Fatalf("Blake2b-512哈希计算失败: %v", err)
	}

	if len(hashValue) == 0 {
		t.Error("Blake2b-512哈希值不能为空")
	}

	// 测试Blake2s-256
	hashValue, err = Hash(data, Blake2s_256)
	if err != nil {
		t.Fatalf("Blake2s-256哈希计算失败: %v", err)
	}

	if len(hashValue) == 0 {
		t.Error("Blake2s-256哈希值不能为空")
	}
}

func TestHashText(t *testing.T) {
	text := "Hello, Blake2 text hashing!"

	// 测试Blake2b-256文本哈希
	hashValue, err := HashText(text, Blake2b_256)
	if err != nil {
		t.Fatalf("Blake2b-256文本哈希计算失败: %v", err)
	}

	if len(hashValue) == 0 {
		t.Error("Blake2b-256文本哈希值不能为空")
	}
}

func TestHashFile(t *testing.T) {
	// 创建临时目录
	tempDir := filepath.Join(os.TempDir(), "blake2_test")
	if err := os.MkdirAll(tempDir, 0755); err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	testFile := filepath.Join(tempDir, "test.txt")
	content := "Hello, Blake2 file hashing!"

	// 创建测试文件
	err := os.WriteFile(testFile, []byte(content), 0644)
	if err != nil {
		t.Fatalf("创建测试文件失败: %v", err)
	}

	// 测试Blake2b-256文件哈希
	hashValue, err := HashFile(testFile, Blake2b_256)
	if err != nil {
		t.Fatalf("Blake2b-256文件哈希计算失败: %v", err)
	}

	if len(hashValue) == 0 {
		t.Error("Blake2b-256文件哈希值不能为空")
	}
}

func TestConvenienceFunctions(t *testing.T) {
	data := []byte("Hello, Blake2 convenience functions!")
	text := "Hello, Blake2 convenience functions text!"

	// 测试Blake2b-256便利函数
	hashValue, err := HashBlake2b256(data)
	if err != nil {
		t.Fatalf("Blake2b-256便利函数哈希计算失败: %v", err)
	}

	if len(hashValue) == 0 {
		t.Error("Blake2b-256便利函数哈希值不能为空")
	}

	// 测试Blake2b-256文本便利函数
	hashValue, err = HashTextBlake2b256(text)
	if err != nil {
		t.Fatalf("Blake2b-256文本便利函数哈希计算失败: %v", err)
	}

	if len(hashValue) == 0 {
		t.Error("Blake2b-256文本便利函数哈希值不能为空")
	}

	// 测试Blake2b-512便利函数
	hashValue, err = HashBlake2b512(data)
	if err != nil {
		t.Fatalf("Blake2b-512便利函数哈希计算失败: %v", err)
	}

	if len(hashValue) == 0 {
		t.Error("Blake2b-512便利函数哈希值不能为空")
	}

	// 测试Blake2b-512文本便利函数
	hashValue, err = HashTextBlake2b512(text)
	if err != nil {
		t.Fatalf("Blake2b-512文本便利函数哈希计算失败: %v", err)
	}

	if len(hashValue) == 0 {
		t.Error("Blake2b-512文本便利函数哈希值不能为空")
	}
}

func TestHashFileConvenience(t *testing.T) {
	// 创建临时目录
	tempDir := filepath.Join(os.TempDir(), "blake2_test")
	if err := os.MkdirAll(tempDir, 0755); err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	testFile := filepath.Join(tempDir, "test.txt")
	content := "Hello, Blake2 file convenience functions!"

	// 创建测试文件
	err := os.WriteFile(testFile, []byte(content), 0644)
	if err != nil {
		t.Fatalf("创建测试文件失败: %v", err)
	}

	// 测试Blake2b-256文件便利函数
	hashValue, err := HashFileBlake2b256(testFile)
	if err != nil {
		t.Fatalf("Blake2b-256文件便利函数哈希计算失败: %v", err)
	}

	if len(hashValue) == 0 {
		t.Error("Blake2b-256文件便利函数哈希值不能为空")
	}

	// 测试Blake2b-512文件便利函数
	hashValue, err = HashFileBlake2b512(testFile)
	if err != nil {
		t.Fatalf("Blake2b-512文件便利函数哈希计算失败: %v", err)
	}

	if len(hashValue) == 0 {
		t.Error("Blake2b-512文件便利函数哈希值不能为空")
	}
}

func TestNilHashFuncErrors(t *testing.T) {
	data := []byte("test data")

	// 测试使用nil哈希函数
	_, err := Hash(data, nil)
	if err == nil {
		t.Error("使用nil哈希函数应该返回错误")
	}
}

func TestEmptyFileNameErrors(t *testing.T) {
	// 测试空文件名
	_, err := HashFile("", Blake2b_256)
	if err == nil {
		t.Error("使用空文件名应该返回错误")
	}
}
