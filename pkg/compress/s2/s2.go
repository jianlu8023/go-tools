package s2

import (
	"bytes"
	"io"
	"os"
	"path/filepath"

	"github.com/klauspost/compress/s2"
)

// Compress 使用s2算法压缩字节数据
func Compress(data []byte) ([]byte, error) {
	var buf bytes.Buffer
	w := s2.NewWriter(&buf)

	if _, err := w.Write(data); err != nil {
		return nil, err
	}

	if err := w.Close(); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// Decompress 使用s2算法解压缩字节数据
func Decompress(data []byte) ([]byte, error) {
	return DecompressWithLimit(data, 0)
}

// DecompressWithLimit 使用s2算法解压缩字节数据，并限制最大解压大小
// 如果maxSize为0，则不限制大小
func DecompressWithLimit(data []byte, maxSize int64) ([]byte, error) {
	r := s2.NewReader(bytes.NewReader(data))

	var buf bytes.Buffer

	if maxSize > 0 {
		// 使用io.LimitReader限制读取大小
		limitedReader := io.LimitReader(r, maxSize)
		if _, err := io.Copy(&buf, limitedReader); err != nil {
			return nil, err
		}

		// 检查是否还有更多数据可读，如果有则说明超出了限制
		if _, err := r.Read(make([]byte, 1)); err == nil {
			return nil, io.ErrUnexpectedEOF
		}
	} else {
		if _, err := io.Copy(&buf, r); err != nil {
			return nil, err
		}
	}

	return buf.Bytes(), nil
}

// CompressFile 压缩文件
func CompressFile(srcPath, destPath string) error {
	// 打开源文件
	srcFile, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	// 确保目标目录存在
	destDir := filepath.Dir(destPath)
	if err := os.MkdirAll(destDir, 0755); err != nil {
		return err
	}

	// 创建目标文件
	destFile, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer destFile.Close()

	// 创建s2写入器
	w := s2.NewWriter(destFile)
	defer w.Close()

	// 复制数据进行压缩
	if _, err := io.Copy(w, srcFile); err != nil {
		return err
	}

	return nil
}

// DecompressFile 解压文件
func DecompressFile(srcPath, destPath string) error {
	return DecompressFileWithLimit(srcPath, destPath, 0)
}

// DecompressFileWithLimit 解压文件，并限制最大解压大小
// 如果maxSize为0，则不限制大小
func DecompressFileWithLimit(srcPath, destPath string, maxSize int64) error {
	// 打开源文件
	srcFile, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	// 创建s2读取器
	r := s2.NewReader(srcFile)

	// 确保目标目录存在
	destDir := filepath.Dir(destPath)
	if err := os.MkdirAll(destDir, 0755); err != nil {
		return err
	}

	// 创建目标文件
	destFile, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer destFile.Close()

	// 复制数据进行解压
	if maxSize > 0 {
		limitedReader := io.LimitReader(r, maxSize)
		bytesRead, err := io.Copy(destFile, limitedReader)
		if err != nil {
			return err
		}
		if bytesRead >= maxSize {
			// 检查是否还有更多数据可读
			if _, err := r.Read(make([]byte, 1)); err == nil {
				return io.ErrUnexpectedEOF
			}
		}
	} else {
		if _, err := io.Copy(destFile, r); err != nil {
			return err
		}
	}

	return nil
}
