package lz4

import (
	"archive/tar"
	"bytes"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"

	"github.com/pierrec/lz4/v4"
)

// Compress 压缩字节数据
// @param data 待压缩的数据
// @return []byte 压缩后的数据
// @return error 错误信息
func Compress(data []byte) ([]byte, error) {
	var buf bytes.Buffer
	writer := lz4.NewWriter(&buf)

	_, err := writer.Write(data)
	if err != nil {
		return nil, fmt.Errorf("failed to write data to lz4 writer: %w", err)
	}

	// 确保所有数据都被写入
	if err := writer.Close(); err != nil {
		return nil, fmt.Errorf("failed to close lz4 writer: %w", err)
	}

	return buf.Bytes(), nil
}

// Decompress 解压字节数据
// @param data 待解压的数据
// @param maxDecompressedSize 最大解压大小（防止解压炸弹），0表示不限制
// @return []byte 解压后的数据
// @return error 错误信息
func Decompress(data []byte, maxDecompressedSize int64) ([]byte, error) {
	reader := lz4.NewReader(bytes.NewReader(data))

	var buf bytes.Buffer
	if maxDecompressedSize > 0 {
		// 使用LimitReader并允许读取超过限制1个字节以便检测
		limitedReader := io.LimitReader(reader, maxDecompressedSize+1)
		bytesRead, err := io.Copy(&buf, limitedReader)
		if err != nil {
			return nil, fmt.Errorf("failed to decompress data: %w", err)
		}
		// 检查是否超过了最大大小限制
		if bytesRead > maxDecompressedSize {
			return nil, fmt.Errorf("decompressed data size exceeds maximum allowed size of %d bytes", maxDecompressedSize)
		}
	} else {
		_, err := io.Copy(&buf, reader)
		if err != nil {
			return nil, fmt.Errorf("failed to decompress data: %w", err)
		}
	}

	return buf.Bytes(), nil
}

// CompressFile 压缩单个文件
// @param sourceFile 源文件路径
// @param targetFile 目标压缩文件路径
// @return error 错误信息
func CompressFile(sourceFile, targetFile string) error {
	// 打开源文件
	src, err := os.Open(sourceFile)
	if err != nil {
		return fmt.Errorf("failed to open source file: %w", err)
	}
	defer src.Close()

	// 创建目标文件
	dst, err := os.Create(targetFile)
	if err != nil {
		return fmt.Errorf("failed to create target file: %w", err)
	}
	defer dst.Close()

	// 创建lz4写入器
	writer := lz4.NewWriter(dst)
	defer writer.Close()

	// 复制文件内容
	_, err = io.Copy(writer, src)
	if err != nil {
		return fmt.Errorf("failed to compress file: %w", err)
	}

	// 确保所有数据都被写入
	if err := writer.Close(); err != nil {
		return fmt.Errorf("failed to flush compressed data: %w", err)
	}

	return nil
}

// DecompressFile 解压单个文件
// @param sourceFile 源压缩文件路径
// @param targetFile 目标文件路径
// @param maxDecompressedSize 最大解压大小（防止解压炸弹），0表示不限制
// @return error 错误信息
func DecompressFile(sourceFile, targetFile string, maxDecompressedSize int64) error {
	// 打开源压缩文件
	src, err := os.Open(sourceFile)
	if err != nil {
		return fmt.Errorf("failed to open source file: %w", err)
	}
	defer src.Close()

	// 创建目标文件
	dst, err := os.Create(targetFile)
	if err != nil {
		return fmt.Errorf("failed to create target file: %w", err)
	}
	defer dst.Close()

	// 创建lz4读取器
	reader := lz4.NewReader(src)

	// 复制解压后的内容
	if maxDecompressedSize > 0 {
		// 设置读取限制
		limitedReader := &io.LimitedReader{R: reader, N: maxDecompressedSize}
		_, err = io.Copy(dst, limitedReader)
	} else {
		_, err = io.Copy(dst, reader)
	}

	if err != nil {
		return fmt.Errorf("failed to decompress file: %w", err)
	}

	return nil
}

// CompressDirToTarlz4 压缩目录到tar.lz4文件
// @param dirPath 源目录路径
// @param targetTarlz4Path 目标tar.lz4文件路径
// @return error 错误信息
func CompressDirToTarlz4(dirPath, targetTarlz4Path string) error {
	// 创建目标文件
	tarFile, err := os.Create(targetTarlz4Path)
	if err != nil {
		return fmt.Errorf("failed to create target file: %w", err)
	}
	defer tarFile.Close()

	// 创建lz4写入器
	lz4Writer := lz4.NewWriter(tarFile)
	defer lz4Writer.Close()

	// 创建tar写入器
	tarWriter := tar.NewWriter(lz4Writer)
	defer tarWriter.Close()

	// 遍历目录并添加文件
	err = filepath.Walk(dirPath, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 计算相对路径
		baseDir := filepath.Base(dirPath)
		relPath, err := filepath.Rel(dirPath, filePath)
		if err != nil {
			return err
		}

		// 创建tar头
		header, err := tar.FileInfoHeader(info, info.Name())
		if err != nil {
			return err
		}

		// 设置正确的路径
		header.Name = path.Join(baseDir, relPath)

		// 写入tar头
		if err = tarWriter.WriteHeader(header); err != nil {
			return err
		}

		// 如果是目录，不需要写入内容
		if info.IsDir() {
			return nil
		}

		// 打开源文件
		file, err := os.Open(filePath)
		if err != nil {
			return err
		}
		defer file.Close()

		// 复制文件内容到tar
		_, err = io.Copy(tarWriter, file)
		return err
	})

	if err != nil {
		return fmt.Errorf("failed to walk directory: %w", err)
	}

	// 确保所有数据都被写入
	if err := tarWriter.Close(); err != nil {
		return fmt.Errorf("failed to close tar writer: %w", err)
	}

	if err := lz4Writer.Close(); err != nil {
		return fmt.Errorf("failed to close lz4 writer: %w", err)
	}

	return nil
}

// UnCompressTarlz4ToDir 解压tar.lz4文件到目录
// @param sourceTarlz4Path 源tar.lz4文件路径
// @param dirPath 目标目录路径
// @return error 错误信息
func UnCompressTarlz4ToDir(sourceTarlz4Path, dirPath string) error {
	// 打开源压缩文件
	sourceFile, err := os.Open(sourceTarlz4Path)
	if err != nil {
		return fmt.Errorf("failed to open source file: %w", err)
	}
	defer sourceFile.Close()

	// 创建lz4读取器
	lz4Reader := lz4.NewReader(sourceFile)

	// 创建tar读取器
	tarReader := tar.NewReader(lz4Reader)

	// 遍历tar文件中的所有条目
	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			// 到达文件末尾
			break
		} else if err != nil {
			return fmt.Errorf("failed to read tar entry: %w", err)
		}

		// 构建目标路径
		targetPath := filepath.Join(dirPath, header.Name)

		// 根据条目类型执行不同操作
		switch header.Typeflag {
		case tar.TypeDir:
			// 创建目录
			if err := os.MkdirAll(targetPath, os.FileMode(header.Mode)); err != nil {
				return fmt.Errorf("failed to create directory: %w", err)
			}
		case tar.TypeReg:
			// 创建文件所在目录
			if err := os.MkdirAll(filepath.Dir(targetPath), os.ModePerm); err != nil {
				return fmt.Errorf("failed to create directory for file: %w", err)
			}

			// 创建目标文件
			file, err := os.OpenFile(targetPath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, os.FileMode(header.Mode))
			if err != nil {
				return fmt.Errorf("failed to create file: %w", err)
			}

			// 复制文件内容
			_, err = io.Copy(file, tarReader)
			file.Close()
			if err != nil {
				return fmt.Errorf("failed to write file content: %w", err)
			}
		}
	}

	return nil
}
