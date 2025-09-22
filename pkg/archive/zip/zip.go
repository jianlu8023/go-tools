package zip

import (
	"archive/zip"
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// Zip 压缩文件或目录到zip文件
// @param sourcePath 源文件或目录路径
// @param targetZipPath 目标zip文件路径
// @param includeRootDir 是否包含根目录（仅当sourcePath是目录时有效）
// @return error 错误信息
func Zip(sourcePath, targetZipPath string, includeRootDir bool) error {
	// 创建目标zip文件
	zipFile, err := os.Create(targetZipPath)
	if err != nil {
		return err
	}
	defer func(zipFile *os.File) {
		err := zipFile.Close()
		if err != nil {
			panic(err)
		}
	}(zipFile)

	// 创建zip写入器
	zipWriter := zip.NewWriter(zipFile)
	defer func(zipWriter *zip.Writer) {
		err := zipWriter.Close()
		if err != nil {
			panic(err)
		}
	}(zipWriter)

	// 获取源路径信息
	sourceInfo, err := os.Stat(sourcePath)
	if err != nil {
		return err
	}

	// 根据源路径类型执行不同的压缩逻辑
	if sourceInfo.IsDir() {
		// 压缩目录
		rootDir := filepath.Base(sourcePath)
		return filepath.Walk(sourcePath, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			// 计算在zip文件中的相对路径
			relPath, err := filepath.Rel(sourcePath, path)
			if err != nil {
				return err
			}

			// 构建zip内部路径
			var zipPath string
			if includeRootDir {
				zipPath = filepath.Join(rootDir, relPath)
			} else {
				zipPath = relPath
			}

			// 如果路径为空，跳过（处理根目录情况）
			if zipPath == "." {
				return nil
			}

			return addFileToZip(zipWriter, path, zipPath, info)
		})
	} else {
		// 压缩单个文件
		zipPath := filepath.Base(sourcePath)
		return addFileToZip(zipWriter, sourcePath, zipPath, sourceInfo)
	}
}

// addFileToZip 将文件添加到zip归档中
func addFileToZip(zipWriter *zip.Writer, filePath, zipPath string, info os.FileInfo) error {
	// 创建zip文件头
	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}

	// 设置文件头信息
	header.Name = zipPath
	// 对于可执行文件，需要特别设置Mode
	if info.Mode()&0111 != 0 {
		header.SetMode(info.Mode())
	}

	// 为目录添加尾部斜杠
	if info.IsDir() {
		header.Name += string(filepath.Separator)
	}

	// 创建zip文件写入器
	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}

	// 如果是目录，不需要写入内容
	if info.IsDir() {
		return nil
	}

	// 打开源文件
	sourceFile, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer func(sourceFile *os.File) {
		err := sourceFile.Close()
		if err != nil {
			panic(err)
		}
	}(sourceFile)

	// 使用带缓冲的读取器提高效率
	bufReader := bufio.NewReader(sourceFile)

	// 将文件内容写入zip归档
	_, err = io.Copy(writer, bufReader)
	return err
}

// ZipFiles 压缩多个文件到zip文件
// @param files 要压缩的文件路径列表
// @param targetZipPath 目标zip文件路径
// @param baseDir 在zip文件中的基础目录（可选，为空则无基础目录）
// @return error 错误信息
func ZipFiles(files []string, targetZipPath string, baseDir string) error {
	// 创建目标zip文件
	zipFile, err := os.Create(targetZipPath)
	if err != nil {
		return err
	}
	defer func(zipFile *os.File) {
		err := zipFile.Close()
		if err != nil {
			panic(err)
		}
	}(zipFile)

	// 创建zip写入器
	zipWriter := zip.NewWriter(zipFile)
	defer func(zipWriter *zip.Writer) {
		err := zipWriter.Close()
		if err != nil {
			panic(err)
		}
	}(zipWriter)

	// 处理每个文件
	for _, filePath := range files {
		info, err := os.Stat(filePath)
		if err != nil {
			return err
		}

		// 构建zip内部路径
		zipPath := filepath.Base(filePath)
		if baseDir != "" {
			zipPath = filepath.Join(baseDir, zipPath)
		}

		// 添加文件到zip
		if err := addFileToZip(zipWriter, filePath, zipPath, info); err != nil {
			return err
		}
	}

	return nil
}

// Unzip 解压zip文件到指定目录
// @param zipFile 压缩文件路径
// @param destDir 解压路径
// @return []string 解压文件路径列表
// @return error 错误信息
func Unzip(zipFile string, destDir string) ([]string, error) {
	// 确保目标目录存在
	if err := os.MkdirAll(destDir, os.ModePerm); err != nil {
		return nil, fmt.Errorf("failed to create destination directory: %w", err)
	}

	// 打开zip文件
	zipReader, err := zip.OpenReader(zipFile)
	if err != nil {
		return nil, fmt.Errorf("failed to open zip file %s: %w", zipFile, err)
	}
	defer func(zipReader *zip.ReadCloser) {
		err := zipReader.Close()
		if err != nil {
			panic(err)
		}
	}(zipReader)

	var extractedPaths []string

	// 遍历zip文件中的所有文件
	for _, f := range zipReader.File {
		// 安全检查：防止路径遍历攻击
		if containsPathTraversal(f.Name) {
			return nil, fmt.Errorf("security violation: path traversal detected in file name '%s'", f.Name)
		}

		// 构建目标文件路径
		destPath := filepath.Join(destDir, f.Name)

		// 确保目标文件所在目录存在
		destDir := filepath.Dir(destPath)
		if err := os.MkdirAll(destDir, os.ModePerm); err != nil {
			return nil, fmt.Errorf("failed to create directory for %s: %w", f.Name, err)
		}

		if f.FileInfo().IsDir() {
			// 创建目录
			if err := os.MkdirAll(destPath, os.ModePerm); err != nil {
				return nil, fmt.Errorf("failed to create directory %s: %w", destPath, err)
			}
			extractedPaths = append(extractedPaths, destPath)
		} else {
			// 解压文件
			if err := extractFile(f, destPath); err != nil {
				return nil, fmt.Errorf("failed to extract file %s: %w", f.Name, err)
			}
			extractedPaths = append(extractedPaths, destPath)
		}
	}

	return extractedPaths, nil
}

// containsPathTraversal 检查路径是否包含路径遍历字符
func containsPathTraversal(path string) bool {
	// 检查Unix/Linux和Windows的路径遍历模式
	if strings.Contains(path, "../") || strings.Contains(path, "..\\") ||
		strings.HasPrefix(path, "..") || strings.HasPrefix(path, "/") ||
		strings.HasPrefix(path, "\\") || strings.HasPrefix(path, "C:") {
		return true
	}

	// 检查规范化后的路径是否改变，这可以检测更复杂的路径遍历尝试
	cleanPath := filepath.Clean(path)
	return cleanPath != path || strings.HasPrefix(cleanPath, "..")
}

// extractFile 从zip文件中解压单个文件
func extractFile(f *zip.File, destPath string) error {
	// 打开zip中的文件
	inFile, err := f.Open()
	if err != nil {
		return err
	}
	// 注意：这里不使用defer，而是在函数结束前直接关闭，避免循环中积累大量未关闭的文件描述符

	// 创建目标文件
	outFile, err := os.OpenFile(destPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
	if err != nil {
		// 确保在错误情况下也关闭inFile
		err := inFile.Close()
		if err != nil {
			return err
		}
		return err
	}

	// 使用带缓冲的读写器提高效率
	bufReader := bufio.NewReader(inFile)
	bufWriter := bufio.NewWriter(outFile)

	// 复制文件内容
	_, err = io.Copy(bufWriter, bufReader)
	// 确保刷新缓冲区
	if flushErr := bufWriter.Flush(); flushErr != nil && err == nil {
		err = flushErr
	}

	// 关闭文件
	if err := inFile.Close(); err != nil {
		return err
	}

	if err := outFile.Close(); err != nil {
		return err
	}

	return err
}

// UnzipSingleFile 从zip文件中解压单个文件
// @param zipFile zip文件路径
// @param fileName zip中的文件名
// @param destPath 目标文件路径
// @return error 错误信息
func UnzipSingleFile(zipFile, fileName, destPath string) error {
	// 打开zip文件
	zipReader, err := zip.OpenReader(zipFile)
	if err != nil {
		return fmt.Errorf("failed to open zip file %s: %w", zipFile, err)
	}
	defer func(zipReader *zip.ReadCloser) {
		err := zipReader.Close()
		if err != nil {
			panic(err)
		}
	}(zipReader)

	// 查找指定的文件
	for _, f := range zipReader.File {
		if f.Name == fileName {
			// 确保目标目录存在
			destDir := filepath.Dir(destPath)
			if err := os.MkdirAll(destDir, os.ModePerm); err != nil {
				return fmt.Errorf("failed to create destination directory: %w", err)
			}

			// 解压文件
			return extractFile(f, destPath)
		}
	}

	return fmt.Errorf("file '%s' not found in zip file '%s'", fileName, zipFile)
}
