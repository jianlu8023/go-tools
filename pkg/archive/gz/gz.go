package gz

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// CompressDirToTargz 压缩目录到tar.gz文件
//
//	例如: CompressDirToTargz("testdata/testDir", "testdata/XXX.tar.gz")
//	会将testDir目录完整压缩到testdata/XXX.tar.gz文件中, 且压缩文件中的目录结构根目录为testDir，与testDir目录结构一致。
//	如果testDir目录结构如下:
//	testDir
//	├── subdir1
//	│   └── testFile1.txt
//	├── subdir2
//	└── testFile.txt
//	则压缩文件中的目录结构为:
//	testDir
//	├── subdir1
//	│   └── testFile1.txt
//	├── subdir2
//	└── testFile.txt
func CompressDirToTargz(dirPath, targetTarGzPath string) (err error) {
	tarFile, err := os.Create(targetTarGzPath)
	if err != nil {
		return err
	}
	defer func() {
		if closeErr := tarFile.Close(); closeErr != nil && err == nil {
			err = closeErr
		}
	}()

	gzWriter := gzip.NewWriter(tarFile)
	defer func() {
		if closeErr := gzWriter.Close(); closeErr != nil && err == nil {
			err = closeErr
		}
	}()

	tarWriter := tar.NewWriter(gzWriter)
	defer func() {
		if closeErr := tarWriter.Close(); closeErr != nil && err == nil {
			err = closeErr
		}
	}()

	err = filepath.Walk(dirPath, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		baseDir := filepath.Base(dirPath)
		relPath, err := filepath.Rel(dirPath, filePath)
		if err != nil {
			return err
		}

		header, err := tar.FileInfoHeader(info, info.Name())
		if err != nil {
			return err
		}
		header.Name = path.Join(baseDir, relPath)

		if err = tarWriter.WriteHeader(header); err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		file, err := os.Open(filePath)
		if err != nil {
			return err
		}

		// 显式关闭文件，避免在 Walk 回调中 defer 累积导致大量句柄同时打开
		_, copyErr := io.Copy(tarWriter, file)
		closeErr := file.Close()
		if copyErr != nil {
			return copyErr
		}
		return closeErr
	})
	if err != nil {
		return err
	}
	return nil
}

// UnCompressTargzToDir 解压tar.gz文件到目录
//
//	例如: UnCompressTargzToDir("testdata/XXX.tar.gz", "testdata/YYY")
//	会将XXX.tar.gz文件完整解压到testdata/YYY目录中。
//	如果XXX.tar.gz文件中的目录结构为:
//	testDir
//	├── subdir1
//	│   └── testFile1.txt
//	├── subdir2
//	└── testFile.txt
//	则解压后的目录结构为:
//	YYY
//	└── testDir
//	    ├── subdir1
//	    │   └── testFile1.txt
//	    ├── subdir2
//	    └── testFile.txt
func UnCompressTargzToDir(sourceTarGzPath, dirPath string) (err error) {
	// 规范化目标目录路径，确保后续路径遍历检查可靠
	absDirPath, err := filepath.Abs(filepath.Clean(dirPath))
	if err != nil {
		return err
	}

	sourceTarGz, err := os.Open(sourceTarGzPath)
	if err != nil {
		return err
	}
	defer func() {
		if closeErr := sourceTarGz.Close(); closeErr != nil && err == nil {
			err = closeErr
		}
	}()

	gzReader, err := gzip.NewReader(sourceTarGz)
	if err != nil {
		return err
	}
	defer func() {
		if closeErr := gzReader.Close(); closeErr != nil && err == nil {
			err = closeErr
		}
	}()

	tarReader := tar.NewReader(gzReader)

	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		targetPath := filepath.Join(absDirPath, header.Name)

		// 路径遍历检查：防止恶意归档通过 ../../../ 等写入目标目录之外的路径
		if !isWithinDir(absDirPath, targetPath) {
			return fmt.Errorf("path traversal detected, refuse to extract: %s", header.Name)
		}

		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(targetPath, os.FileMode(header.Mode)); err != nil {
				return err
			}
		case tar.TypeReg:
			file, err := os.OpenFile(targetPath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, os.FileMode(header.Mode))
			if err != nil {
				return err
			}

			// 显式关闭文件，避免在 for 循环中 defer 累积导致大量句柄同时打开
			_, copyErr := io.Copy(file, tarReader)
			closeErr := file.Close()
			if copyErr != nil {
				return copyErr
			}
			if closeErr != nil {
				return closeErr
			}
		}
	}
	return nil
}

// isWithinDir 判断 targetPath 是否位于 baseDir 之内（含 baseDir 自身）
// 通过比较清理后的绝对路径前缀，防止符号链接或 `..` 跳出 baseDir
func isWithinDir(baseDir, targetPath string) bool {
	rel, err := filepath.Rel(baseDir, targetPath)
	if err != nil {
		return false
	}
	// rel 以 ".." 开头表示 targetPath 在 baseDir 之外
	return rel != ".." && !strings.HasPrefix(rel, ".."+string(filepath.Separator))
}
