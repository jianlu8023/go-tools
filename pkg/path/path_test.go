package path

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriteToFileAndFileExists(t *testing.T) {
	// 准备测试文件路径
	testFilePath := "test_write.txt"
	defer func() {
		// 测试结束后删除测试文件
		_ = RemoveFile(testFilePath)
	}()

	// 测试写入新文件
	content := "Hello, World!"
	err := WriteToFile(testFilePath, content, false)
	assert.NoError(t, err)

	// 测试文件是否存在
	exists, err := FileExists(testFilePath)
	assert.NoError(t, err)
	assert.True(t, exists)

	// 测试文件非空
	notEmpty, err := FileNotEmpty(testFilePath)
	assert.NoError(t, err)
	assert.True(t, notEmpty)

	// 测试不覆盖写入已存在的文件
	err = WriteToFile(testFilePath, "New content", false)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "文件已存在")

	// 测试强制覆盖写入已存在的文件
	err = WriteToFile(testFilePath, "New content", true)
	assert.NoError(t, err)
}

func TestAppendWrite(t *testing.T) {
	// 准备测试文件路径
	testFilePath := "test_append.txt"
	defer func() {
		// 测试结束后删除测试文件
		_ = RemoveFile(testFilePath)
	}()

	// 先写入初始内容
	err := WriteToFile(testFilePath, "Initial content", false)
	assert.NoError(t, err)

	// 测试追加写入
	err = appendWrite(testFilePath, "\nAppended content")
	assert.NoError(t, err)

	// 读取文件内容验证
	lines, allContent, err := ReadFileToLinesAndAll(testFilePath)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(lines))
	assert.Contains(t, allContent, "Initial content")
	assert.Contains(t, allContent, "Appended content")
}

func TestCreateDir(t *testing.T) {
	// 准备测试目录路径
	testDirPath := "test_dir"
	defer func() {
		// 测试结束后删除测试目录
		_ = os.RemoveAll(testDirPath)
	}()

	// 测试创建新目录
	ok, err := CreateDir(testDirPath)
	assert.NoError(t, err)
	assert.True(t, ok)

	// 检查目录是否存在
	dirInfo, err := os.Stat(testDirPath)
	assert.NoError(t, err)
	assert.True(t, dirInfo.IsDir())

	// 测试创建已存在的目录
	ok, err = CreateDir(testDirPath)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "该目录已经存在")
	assert.True(t, ok) // 即使有错误，返回值仍然为true，表示目录存在
}

func TestClearDir(t *testing.T) {
	// 准备测试目录路径
	testDirPath := "test_clear_dir"
	defer func() {
		// 测试结束后删除测试目录
		_ = os.RemoveAll(testDirPath)
	}()

	// 创建测试目录
	_, err := CreateDir(testDirPath)
	assert.NoError(t, err)

	// 在目录中创建一些文件
	testFilePath1 := filepath.Join(testDirPath, "file1.txt")
	testFilePath2 := filepath.Join(testDirPath, "file2.txt")
	err = WriteToFile(testFilePath1, "Content 1", false)
	assert.NoError(t, err)
	err = WriteToFile(testFilePath2, "Content 2", false)
	assert.NoError(t, err)

	// 创建子目录并在其中创建文件
	subDirPath := filepath.Join(testDirPath, "subdir")
	_, err = CreateDir(subDirPath)
	assert.NoError(t, err)
	subFilePath := filepath.Join(subDirPath, "subfile.txt")
	err = WriteToFile(subFilePath, "Sub content", false)
	assert.NoError(t, err)

	// 测试清空目录
	err = ClearDir(testDirPath)
	assert.NoError(t, err)

	// 验证目录为空但仍然存在
	isEmpty, err := IsDirEmpty(testDirPath)
	assert.NoError(t, err)
	assert.True(t, isEmpty)
}

func TestIsDirEmpty(t *testing.T) {
	// 准备测试目录路径
	testDirPath1 := "test_empty_dir"
	testDirPath2 := "test_non_empty_dir"
	defer func() {
		// 测试结束后删除测试目录
		_ = os.RemoveAll(testDirPath1)
		_ = os.RemoveAll(testDirPath2)
	}()

	// 创建测试目录
	_, err := CreateDir(testDirPath1)
	assert.NoError(t, err)
	_, err = CreateDir(testDirPath2)
	assert.NoError(t, err)

	// 在第二个目录中创建文件
	testFilePath := filepath.Join(testDirPath2, "file.txt")
	err = WriteToFile(testFilePath, "Content", false)
	assert.NoError(t, err)

	// 测试空目录
	isEmpty, err := IsDirEmpty(testDirPath1)
	assert.NoError(t, err)
	assert.True(t, isEmpty)

	// 测试非空目录
	isEmpty, err = IsDirEmpty(testDirPath2)
	assert.NoError(t, err)
	assert.False(t, isEmpty)
}

func TestFileCopy(t *testing.T) {
	// 准备测试文件路径
	srcFilePath := "test_source.txt"
	dstFilePath := "test_destination.txt"
	defer func() {
		// 测试结束后删除测试文件
		_ = RemoveFile(srcFilePath)
		_ = RemoveFile(dstFilePath)
	}()

	// 创建源文件
	srcContent := "Content to be copied"
	err := WriteToFile(srcFilePath, srcContent, false)
	assert.NoError(t, err)

	// 测试文件拷贝
	err = FileCopy(srcFilePath, dstFilePath)
	assert.NoError(t, err)

	// 验证目标文件内容
	_, dstContent, err := ReadFileToLinesAndAll(dstFilePath)
	assert.NoError(t, err)
	assert.Equal(t, srcContent, dstContent)
}

func TestFileCopyToDir(t *testing.T) {
	// 准备测试文件和目录路径
	srcFilePath := "test_source_to_dir.txt"
	dstDirPath := "test_destination_dir"
	defer func() {
		// 测试结束后删除测试文件和目录
		_ = RemoveFile(srcFilePath)
		_ = os.RemoveAll(dstDirPath)
	}()

	// 创建源文件
	srcContent := "Content to be copied to directory"
	err := WriteToFile(srcFilePath, srcContent, false)
	assert.NoError(t, err)

	// 测试文件拷贝到目录
	err = FileCopyToDir(srcFilePath, dstDirPath)
	assert.NoError(t, err)

	// 验证目标文件内容
	dstFilePath := filepath.Join(dstDirPath, filepath.Base(srcFilePath))
	_, dstContent, err := ReadFileToLinesAndAll(dstFilePath)
	assert.NoError(t, err)
	assert.Equal(t, srcContent, dstContent)
}

func TestSplitPath(t *testing.T) {
	// 测试各种路径格式的切分
	tests := []struct {
		path     string
		expected []string
	}{{
		path:     "/a/b/c/d.txt",
		expected: []string{"a", "b", "c", "d.txt"},
	}, {
		path:     "a/b/c/d.txt",
		expected: []string{"a", "b", "c", "d.txt"},
	}, {
		path:     "./a/b/c/d.txt",
		expected: []string{".", "a", "b", "c", "d.txt"},
	}, {
		path:     "../a/b/c/d.txt",
		expected: []string{"..", "a", "b", "c", "d.txt"},
	}, {
		path:     "C:\\a\\b\\c\\d.txt",
		expected: []string{"C:", "a", "b", "c", "d.txt"},
	}, {
		path:     "\\a\\b\\c\\d.txt",
		expected: []string{"a", "b", "c", "d.txt"},
	}}

	for _, tt := range tests {
		result := SplitPath(tt.path)
		assert.Equal(t, tt.expected, result, "SplitPath(%s)", tt.path)
	}
}

func TestFirstDir(t *testing.T) {
	// 测试获取第一个目录
	tests := []struct {
		path     string
		expected string
	}{{
		path:     "/a/b/c/d.txt",
		expected: "a",
	}, {
		path:     "a/b/c/d.txt",
		expected: "a",
	}, {
		path:     "./a/b/c/d.txt",
		expected: ".",
	}, {
		path:     "../a/b/c/d.txt",
		expected: "..",
	}, {
		path:     "C:\\a\\b\\c\\d.txt",
		expected: "C:",
	}, {
		path:     "\\a\\b\\c\\d.txt",
		expected: "a",
	}}

	for _, tt := range tests {
		result := FirstDir(tt.path)
		assert.Equal(t, tt.expected, result, "FirstDir(%s)", tt.path)
	}
}

func TestRemoveFile(t *testing.T) {
	// 准备测试文件路径
	testFilePath := "test_remove.txt"

	// 创建测试文件
	err := WriteToFile(testFilePath, "Content", false)
	assert.NoError(t, err)

	// 测试删除存在的文件
	err = RemoveFile(testFilePath)
	assert.NoError(t, err)

	// 验证文件已删除
	exists, err := FileExists(testFilePath)
	assert.NoError(t, err)
	assert.False(t, exists)

	// 测试删除不存在的文件（不应该报错）
	err = RemoveFile(testFilePath)
	assert.NoError(t, err)
}

func TestGetWorkDir(t *testing.T) {
	// 测试获取当前工作目录
	wd, err := GetWorkDir()
	assert.NoError(t, err)
	assert.NotEmpty(t, wd)

	// 验证返回的路径与os.Getwd()一致
	osWd, err := os.Getwd()
	assert.NoError(t, err)
	assert.Equal(t, osWd, wd)
}

func TestFilterFileByCondition(t *testing.T) {
	// 准备测试目录路径
	testDirPath := "test_filter_dir"
	defer func() {
		// 测试结束后删除测试目录
		_ = os.RemoveAll(testDirPath)
	}()

	// 创建测试目录
	_, err := CreateDir(testDirPath)
	assert.NoError(t, err)

	// 在目录中创建一些文件
	testFilePath1 := filepath.Join(testDirPath, "file1.txt")
	testFilePath2 := filepath.Join(testDirPath, "file2.json")
	testFilePath3 := filepath.Join(testDirPath, ".hidden.txt")
	testFilePath4 := filepath.Join(testDirPath, "config.ini")
	err = WriteToFile(testFilePath1, "Content 1", false)
	assert.NoError(t, err)
	err = WriteToFile(testFilePath2, "Content 2", false)
	assert.NoError(t, err)
	err = WriteToFile(testFilePath3, "Content 3", false)
	assert.NoError(t, err)
	err = WriteToFile(testFilePath4, "Content 4", false)
	assert.NoError(t, err)

	// 创建子目录
	subDirPath := filepath.Join(testDirPath, "subdir")
	_, err = CreateDir(subDirPath)
	assert.NoError(t, err)

	// 测试按文件后缀过滤
	condition := FileFilterCondition{
		FileNameSuffix: ".txt",
		ContainsHidden: false,
		ContainsDir:    false,
	}
	files := FilterFileByCondition(testDirPath, condition)
	assert.Contains(t, files, testFilePath1)
	assert.NotContains(t, files, testFilePath2)
	assert.NotContains(t, files, testFilePath3) // 默认不包含隐藏文件
	assert.NotContains(t, files, subDirPath)    // 默认不包含目录

	// 测试包含隐藏文件
	condition.ContainsHidden = true
	files = FilterFileByCondition(testDirPath, condition)
	assert.Contains(t, files, testFilePath1)
	assert.Contains(t, files, testFilePath3)

	// 测试包含目录
	condition.ContainsDir = true
	files = FilterFileByCondition(testDirPath, condition)
	assert.Contains(t, files, subDirPath)

	// 测试按正则表达式过滤
	condition = FileFilterCondition{
		FileNameRegex: "^file\\d+",
	}
	files = FilterFileByCondition(testDirPath, condition)
	assert.Contains(t, files, testFilePath1)
	assert.Contains(t, files, testFilePath2)
	assert.NotContains(t, files, testFilePath4)
}
