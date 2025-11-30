package lz4

import (
	"bytes"
	"crypto/rand"
	"os"
	"path/filepath"
	"testing"
)

// TestCompressAndDecompress 测试基本的压缩和解压功能
func TestCompressAndDecompress(t *testing.T) {
	tests := []struct {
		name string
		data []byte
	}{{
		name: "空数据",
		data: []byte{},
	}, {
		name: "简单字符串",
		data: []byte("hello world"),
	}, {
		name: "重复模式数据",
		data: bytes.Repeat([]byte("abcdefgh"), 100),
	}, {
		name: "随机数据",
		data: func() []byte {
			data := make([]byte, 1024)
			_, _ = rand.Read(data)
			return data
		}(),
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 压缩数据
			compressed, err := Compress(tt.data)
			if err != nil {
				t.Fatalf("Compress failed: %v", err)
			}

			// 对于非空数据，压缩后的数据通常应该更小（除了一些特殊情况）
			if len(tt.data) > 0 && len(compressed) >= len(tt.data)*2 {
				t.Logf("Warning: Compressed data is not smaller: original=%d, compressed=%d", len(tt.data), len(compressed))
			}

			// 解压数据
			decompressed, err := Decompress(compressed, 0)
			if err != nil {
				t.Fatalf("Decompress failed: %v", err)
			}

			// 验证解压后的数据与原始数据一致
			if !bytes.Equal(decompressed, tt.data) {
				t.Fatalf("Decompressed data does not match original: original=%d bytes, decompressed=%d bytes", len(tt.data), len(decompressed))
			}
		})
	}
}

// TestDecompressWithSizeLimit 测试带大小限制的解压功能
func TestDecompressWithSizeLimit(t *testing.T) {
	// 创建一个可压缩的数据
	original := bytes.Repeat([]byte("hello"), 1000)

	// 压缩数据
	compressed, err := Compress(original)
	if err != nil {
		t.Fatalf("Compress failed: %v", err)
	}

	// 测试成功的情况：设置足够大的限制
	decompressed, err := Decompress(compressed, int64(len(original)))
	if err != nil {
		t.Fatalf("Decompress with sufficient limit failed: %v", err)
	}
	if !bytes.Equal(decompressed, original) {
		t.Fatalf("Decompressed data does not match original")
	}

	// 测试失败的情况：设置太小的限制
	decompressed, err = Decompress(compressed, int64(len(original)/2))
	if err == nil {
		t.Fatal("Decompress with insufficient limit should have failed")
	}
}

// TestCompressFileAndDecompressFile 测试文件压缩和解压功能
func TestCompressFileAndDecompressFile(t *testing.T) {
	// 创建临时目录
	tempDir, err := os.MkdirTemp("", "lz4-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// 创建测试文件
	sourceFile := filepath.Join(tempDir, "source.txt")
	content := "This is a test file for lz4 compression and decompression."
	if err := os.WriteFile(sourceFile, []byte(content), 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// 压缩文件
	compressedFile := filepath.Join(tempDir, "compressed.lz4")
	if err := CompressFile(sourceFile, compressedFile); err != nil {
		t.Fatalf("CompressFile failed: %v", err)
	}

	// 验证压缩文件存在且大小合理
	if info, err := os.Stat(compressedFile); err != nil {
		t.Fatalf("Compressed file does not exist: %v", err)
	} else if info.Size() == 0 {
		t.Fatal("Compressed file is empty")
	}

	// 解压文件
	decompressedFile := filepath.Join(tempDir, "decompressed.txt")
	if err := DecompressFile(compressedFile, decompressedFile, 0); err != nil {
		t.Fatalf("DecompressFile failed: %v", err)
	}

	// 验证解压后的文件内容
	decompressedContent, err := os.ReadFile(decompressedFile)
	if err != nil {
		t.Fatalf("Failed to read decompressed file: %v", err)
	}

	if string(decompressedContent) != content {
		t.Fatalf("Decompressed content does not match original: got=%s, want=%s", string(decompressedContent), content)
	}
}

// TestCompressDirToTarlz4AndUnCompressTarlz4ToDir 测试目录压缩和解压功能
func TestCompressDirToTarlz4AndUnCompressTarlz4ToDir(t *testing.T) {
	// 创建临时目录
	tempDir, err := os.MkdirTemp("", "lz4-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// 创建测试目录结构
	sourceDir := filepath.Join(tempDir, "source")
	if err := os.Mkdir(sourceDir, 0755); err != nil {
		t.Fatalf("Failed to create source dir: %v", err)
	}

	// 创建测试文件
	if err := os.WriteFile(filepath.Join(sourceDir, "file1.txt"), []byte("content1"), 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// 创建子目录和文件
	subDir := filepath.Join(sourceDir, "subdir")
	if err := os.Mkdir(subDir, 0755); err != nil {
		t.Fatalf("Failed to create subdir: %v", err)
	}

	if err := os.WriteFile(filepath.Join(subDir, "file2.txt"), []byte("content2"), 0644); err != nil {
		t.Fatalf("Failed to create test file in subdir: %v", err)
	}

	// 压缩目录
	compressedFile := filepath.Join(tempDir, "archive.tar.lz4")
	if err := CompressDirToTarlz4(sourceDir, compressedFile); err != nil {
		t.Fatalf("CompressDirToTarlz4 failed: %v", err)
	}

	// 验证压缩文件存在
	if _, err := os.Stat(compressedFile); err != nil {
		t.Fatalf("Compressed file does not exist: %v", err)
	}

	// 解压目录
	decompressedDir := filepath.Join(tempDir, "decompressed")
	if err := UnCompressTarlz4ToDir(compressedFile, decompressedDir); err != nil {
		t.Fatalf("UnCompressTarlz4ToDir failed: %v", err)
	}

	// 验证解压后的目录结构
	expectedFiles := []struct {
		path    string
		content string
	}{{
		path:    filepath.Join(decompressedDir, "source", "file1.txt"),
		content: "content1",
	}, {
		path:    filepath.Join(decompressedDir, "source", "subdir", "file2.txt"),
		content: "content2",
	}}

	for _, ef := range expectedFiles {
		content, err := os.ReadFile(ef.path)
		if err != nil {
			t.Fatalf("Failed to read decompressed file %s: %v", ef.path, err)
		}
		if string(content) != ef.content {
			t.Fatalf("Decompressed content of %s does not match: got=%s, want=%s", ef.path, string(content), ef.content)
		}
	}
}

// BenchmarkCompress 基准测试压缩性能
func BenchmarkCompress(b *testing.B) {
	// 创建测试数据
	data := make([]byte, 1024*1024) // 1MB
	_, _ = rand.Read(data)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := Compress(data)
		if err != nil {
			b.Fatalf("Compress failed: %v", err)
		}
	}
}

// BenchmarkDecompress 基准测试解压性能
func BenchmarkDecompress(b *testing.B) {
	// 创建测试数据
	data := make([]byte, 1024*1024) // 1MB
	_, _ = rand.Read(data)

	// 先压缩数据
	compressed, err := Compress(data)
	if err != nil {
		b.Fatalf("Compress failed: %v", err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := Decompress(compressed, 0)
		if err != nil {
			b.Fatalf("Decompress failed: %v", err)
		}
	}
}

// TestErrorHandling 测试错误处理
func TestErrorHandling(t *testing.T) {
	// 测试不存在的源文件
	if err := CompressFile("/path/to/nonexistent/file", "output.lz4"); err == nil {
		t.Fatal("CompressFile should fail for nonexistent source file")
	}

	// 测试不存在的源目录
	if err := DecompressFile("/path/to/nonexistent/file.lz4", "output.txt", 0); err == nil {
		t.Fatal("DecompressFile should fail for nonexistent source file")
	}

	// 测试无效的压缩数据
	invalidData := []byte("not lz4 compressed data")
	if _, err := Decompress(invalidData, 0); err == nil {
		t.Fatal("Decompress should fail for invalid compressed data")
	}

	// 测试无效的tar.lz4文件
	tempDir, err := os.MkdirTemp("", "lz4-test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	invalidFile := filepath.Join(tempDir, "invalid.tar.lz4")
	if err := os.WriteFile(invalidFile, []byte("not a tar.lz4 file"), 0644); err != nil {
		t.Fatalf("Failed to create invalid file: %v", err)
	}

	decompressDir := filepath.Join(tempDir, "decompress")
	if err := UnCompressTarlz4ToDir(invalidFile, decompressDir); err == nil {
		t.Fatal("UnCompressTarlz4ToDir should fail for invalid tar.lz4 file")
	}
}
