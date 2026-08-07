package zlib

import (
	"bytes"
	"io"
	"os"
	"path/filepath"

	"github.com/klauspost/compress/zlib"
)

// Compress 使用zlib算法压缩字节数据
func Compress(data []byte) ([]byte, error) {
	var buf bytes.Buffer
	w := zlib.NewWriter(&buf)

	if _, err := w.Write(data); err != nil {
		return nil, err
	}

	if err := w.Close(); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// Decompress 使用zlib算法解压缩字节数据
func Decompress(data []byte) ([]byte, error) {
	return DecompressWithLimit(data, 0)
}

// DecompressWithLimit 使用zlib算法解压缩字节数据，并限制最大解压大小
// 如果maxSize为0，则不限制大小
func DecompressWithLimit(data []byte, maxSize int64) ([]byte, error) {
	r, err := zlib.NewReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	defer r.Close()

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
func CompressFile(srcPath, destPath string) (err error) {
	// 打开源文件
	srcFile, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer func() {
		if closeErr := srcFile.Close(); closeErr != nil && err == nil {
			err = closeErr
		}
	}()

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
	// zlib.Writer.Close 负责刷出缓冲数据并写入尾部校验和，必须显式关闭并检查错误；
	// 为确保 Close 错误能正确返回，先 Close writer 再 Close 文件。
	w := zlib.NewWriter(destFile)

	// 复制数据进行压缩
	if _, err := io.Copy(w, srcFile); err != nil {
		_ = w.Close()
		_ = destFile.Close()
		return err
	}

	// 关闭 zlib writer 刷出尾部校验和
	if err := w.Close(); err != nil {
		_ = destFile.Close()
		return err
	}

	// 关闭目标文件
	if closeErr := destFile.Close(); closeErr != nil {
		return closeErr
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

	// 创建zlib读取器
	r, err := zlib.NewReader(srcFile)
	if err != nil {
		return err
	}
	defer r.Close()

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
