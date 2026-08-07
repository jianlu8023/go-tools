package lz4

import (
	"archive/tar"
	"bytes"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"

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
func CompressFile(sourceFile, targetFile string) (err error) {
	// 打开源文件
	src, err := os.Open(sourceFile)
	if err != nil {
		return fmt.Errorf("failed to open source file: %w", err)
	}
	defer func() {
		if closeErr := src.Close(); closeErr != nil && err == nil {
			err = closeErr
		}
	}()

	// 创建目标文件
	dst, err := os.Create(targetFile)
	if err != nil {
		return fmt.Errorf("failed to create target file: %w", err)
	}

	// 创建lz4写入器
	writer := lz4.NewWriter(dst)

	// 复制文件内容
	if _, err = io.Copy(writer, src); err != nil {
		_ = writer.Close()
		_ = dst.Close()
		return fmt.Errorf("failed to compress file: %w", err)
	}

	// 关闭 lz4 writer 刷出尾部（仅 Close 一次，避免 P002 重复 Close 问题）
	if err := writer.Close(); err != nil {
		_ = dst.Close()
		return fmt.Errorf("failed to flush compressed data: %w", err)
	}

	// 关闭目标文件
	if closeErr := dst.Close(); closeErr != nil {
		return fmt.Errorf("failed to close target file: %w", closeErr)
	}

	return nil
}

// DecompressFile 解压单个文件
// @param sourceFile 源压缩文件路径
// @param targetFile 目标文件路径
// @param maxDecompressedSize 最大解压大小（防止解压炸弹），0表示不限制
// @return error 错误信息
func DecompressFile(sourceFile, targetFile string, maxDecompressedSize int64) (err error) {
	// 打开源压缩文件
	src, err := os.Open(sourceFile)
	if err != nil {
		return fmt.Errorf("failed to open source file: %w", err)
	}
	defer func() {
		if closeErr := src.Close(); closeErr != nil && err == nil {
			err = closeErr
		}
	}()

	// 创建目标文件
	dst, err := os.Create(targetFile)
	if err != nil {
		return fmt.Errorf("failed to create target file: %w", err)
	}
	defer func() {
		if closeErr := dst.Close(); closeErr != nil && err == nil {
			err = closeErr
		}
	}()

	// 创建lz4读取器
	reader := lz4.NewReader(src)

	// 复制解压后的内容
	if maxDecompressedSize > 0 {
		// 使用 LimitReader 读取 maxDecompressedSize+1 字节以检测是否超限
		limitedReader := io.LimitReader(reader, maxDecompressedSize+1)
		bytesRead, copyErr := io.Copy(dst, limitedReader)
		if copyErr != nil {
			return fmt.Errorf("failed to decompress file: %w", copyErr)
		}
		// 读取字节数超过限制说明数据不完整（被静默截断），需返回错误而非误报成功
		if bytesRead > maxDecompressedSize {
			return fmt.Errorf("decompressed data size exceeds maximum allowed size of %d bytes", maxDecompressedSize)
		}
	} else {
		if _, err := io.Copy(dst, reader); err != nil {
			return fmt.Errorf("failed to decompress file: %w", err)
		}
	}

	return nil
}

// CompressDirToTarlz4 压缩目录到tar.lz4文件
// @param dirPath 源目录路径
// @param targetTarlz4Path 目标tar.lz4文件路径
// @return error 错误信息
func CompressDirToTarlz4(dirPath, targetTarlz4Path string) (err error) {
	// 创建目标文件
	tarFile, err := os.Create(targetTarlz4Path)
	if err != nil {
		return fmt.Errorf("failed to create target file: %w", err)
	}
	defer func() {
		if closeErr := tarFile.Close(); closeErr != nil && err == nil {
			err = closeErr
		}
	}()

	// 创建lz4写入器
	lz4Writer := lz4.NewWriter(tarFile)

	// 创建tar写入器
	tarWriter := tar.NewWriter(lz4Writer)

	// 遍历目录并添加文件
	walkErr := filepath.Walk(dirPath, func(filePath string, info os.FileInfo, err error) error {
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

		// 显式关闭文件，避免在 Walk 回调中 defer 累积导致大量句柄同时打开（N-A09）
		_, copyErr := io.Copy(tarWriter, file)
		closeErr := file.Close()
		if copyErr != nil {
			return copyErr
		}
		return closeErr
	})

	// 无论 Walk 是否成功，都需要按顺序关闭 tarWriter 和 lz4Writer 以刷出缓冲数据
	// 关闭顺序：先 tarWriter（依赖 lz4Writer），再 lz4Writer（依赖 tarFile）
	closeTarErr := tarWriter.Close()
	closeLz4Err := lz4Writer.Close()

	if walkErr != nil {
		return fmt.Errorf("failed to walk directory: %w", walkErr)
	}
	if closeTarErr != nil {
		return fmt.Errorf("failed to close tar writer: %w", closeTarErr)
	}
	if closeLz4Err != nil {
		return fmt.Errorf("failed to close lz4 writer: %w", closeLz4Err)
	}

	return nil
}

// UnCompressTarlz4ToDir 解压tar.lz4文件到目录
// @param sourceTarlz4Path 源tar.lz4文件路径
// @param dirPath 目标目录路径
// @return error 错误信息
func UnCompressTarlz4ToDir(sourceTarlz4Path, dirPath string) (err error) {
	// 规范化目标目录路径，确保后续路径遍历检查可靠
	absDirPath, err := filepath.Abs(filepath.Clean(dirPath))
	if err != nil {
		return fmt.Errorf("failed to get absolute directory: %w", err)
	}

	// 打开源压缩文件
	sourceFile, err := os.Open(sourceTarlz4Path)
	if err != nil {
		return fmt.Errorf("failed to open source file: %w", err)
	}
	defer func() {
		if closeErr := sourceFile.Close(); closeErr != nil && err == nil {
			err = closeErr
		}
	}()

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
		targetPath := filepath.Join(absDirPath, header.Name)

		// 路径遍历检查：防止恶意归档通过 ../ 等写入目标目录之外的路径（N-A02）
		if !isWithinDir(absDirPath, targetPath) {
			return fmt.Errorf("path traversal detected, refuse to extract: %s", header.Name)
		}

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

			// 显式关闭文件，避免在 for 循环中 defer 累积导致大量句柄同时打开
			_, copyErr := io.Copy(file, tarReader)
			closeErr := file.Close()
			if copyErr != nil {
				return fmt.Errorf("failed to write file content: %w", copyErr)
			}
			if closeErr != nil {
				return fmt.Errorf("failed to close file %s: %w", targetPath, closeErr)
			}
		}
	}

	return nil
}

// isWithinDir 判断 targetPath 是否位于 baseDir 之内（含 baseDir 自身）
// 通过比较清理后的相对路径，防止 `..`、绝对路径、UNC 路径等跳出 baseDir
func isWithinDir(baseDir, targetPath string) bool {
	rel, err := filepath.Rel(baseDir, targetPath)
	if err != nil {
		return false
	}
	// rel 以 ".." 开头表示 targetPath 在 baseDir 之外
	return rel != ".." && !strings.HasPrefix(rel, ".."+string(filepath.Separator))
}
