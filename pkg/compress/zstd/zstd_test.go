package zstd

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"testing"
)

// 测试简单字符串的压缩和解压缩
func TestCompressAndDecompress(t *testing.T) {
	// 测试数据
	testData := []byte("Hello, this is a test for zstd compression!")

	// 压缩数据
	compressed, err := Compress(testData)
	if err != nil {
		t.Fatalf("Compress failed: %v", err)
	}

	// 对于小数据，压缩后可能会变大，这是正常的
	if len(compressed) > len(testData) {
		t.Logf("Compressed data is larger than original: %d > %d (this is normal for small data)", len(compressed), len(testData))
	}

	// 解压缩数据
	decompressed, err := Decompress(compressed)
	if err != nil {
		t.Fatalf("Decompress failed: %v", err)
	}

	// 验证解压后的数据是否与原始数据相同
	if !bytes.Equal(decompressed, testData) {
		t.Fatalf("Decompressed data does not match original data")
	}
}

// 测试大小限制功能
func TestDecompressWithSizeLimit(t *testing.T) {
	// 测试数据
	testData := []byte("This is a test data for size limit decompression")

	// 压缩数据
	compressed, err := Compress(testData)
	if err != nil {
		t.Fatalf("Compress failed: %v", err)
	}

	// 使用足够大的限制，应该成功解压
	decompressed, err := DecompressWithLimit(compressed, int64(len(testData)*2))
	if err != nil {
		t.Fatalf("Decompress with sufficient limit failed: %v", err)
	}
	if !bytes.Equal(decompressed, testData) {
		t.Fatalf("Decompressed data does not match original data")
	}

	// 使用过小的限制，应该失败
	_, err = DecompressWithLimit(compressed, int64(len(testData)/2))
	if err == nil || err != io.ErrUnexpectedEOF {
		t.Fatalf("Decompress with insufficient limit should have failed with io.ErrUnexpectedEOF, got: %v", err)
	}
}

// 测试文件压缩和解压缩
func TestCompressFileAndDecompressFile(t *testing.T) {
	// 创建临时目录
	tempDir, err := os.MkdirTemp("", "zstd-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// 创建测试文件
	srcPath := filepath.Join(tempDir, "test.txt")
	testContent := []byte("This is a test file content for zstd compression and decompression")
	if err := os.WriteFile(srcPath, testContent, 0644); err != nil {
		t.Fatalf("Failed to write test file: %v", err)
	}

	// 压缩文件
	compressedPath := filepath.Join(tempDir, "test.txt.zst")
	if err := CompressFile(srcPath, compressedPath); err != nil {
		t.Fatalf("CompressFile failed: %v", err)
	}

	// 验证压缩文件是否存在
	if _, err := os.Stat(compressedPath); os.IsNotExist(err) {
		t.Fatalf("Compressed file does not exist")
	}

	// 解压缩文件
	decompressedPath := filepath.Join(tempDir, "test-decompressed.txt")
	if err := DecompressFile(compressedPath, decompressedPath); err != nil {
		t.Fatalf("DecompressFile failed: %v", err)
	}

	// 验证解压后的数据是否与原始数据相同
	decompressedContent, err := os.ReadFile(decompressedPath)
	if err != nil {
		t.Fatalf("Failed to read decompressed file: %v", err)
	}
	if !bytes.Equal(decompressedContent, testContent) {
		t.Fatalf("Decompressed file content does not match original content")
	}
}

// 测试大文件压缩（这里使用相对较大的数据来测试性能和正确性）
func TestLargeDataCompression(t *testing.T) {
	// 创建一个较大的测试数据（约1MB）
	largeData := make([]byte, 1024*1024)
	for i := range largeData {
		largeData[i] = byte(i % 256)
	}

	// 压缩数据
	compressed, err := Compress(largeData)
	if err != nil {
		t.Fatalf("Compress large data failed: %v", err)
	}

	// 验证压缩效果（对于这种重复模式的数据，应该有较好的压缩率）
	if float64(len(compressed))/float64(len(largeData)) > 0.5 {
		t.Logf("Compression ratio is not optimal: %.2f%%", float64(len(compressed))/float64(len(largeData))*100)
	}

	// 解压缩数据
	decompressed, err := Decompress(compressed)
	if err != nil {
		t.Fatalf("Decompress large data failed: %v", err)
	}

	// 验证解压后的数据是否与原始数据相同
	if !bytes.Equal(decompressed, largeData) {
		t.Fatalf("Decompressed large data does not match original data")
	}
}
