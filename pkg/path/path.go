package path

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

var (
	ErrIsNull = errIsNull()
)

func errIsNull() error { return errors.New("path is null") }

// Ensure 确保目录存在
// @param path 目录路径
// @param isDir 是否文件夹
// @return error 错误信息
func Ensure(path string, isDir bool) error {
	abs, err := IsAbs(path)
	if err != nil {
		return err
	}

	if isDir {
		if _, err := os.Stat(abs); os.IsNotExist(err) {
			return os.MkdirAll(abs, os.ModePerm)
		}
	} else {
		dir := filepath.Dir(abs)
		// 如果路径不存在，创建目录
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			if err := os.MkdirAll(dir, os.ModePerm); err != nil {
				return err
			}
		}
		if _, err := os.Stat(abs); os.IsNotExist(err) {
			f, err := os.Create(abs)
			defer func(f *os.File) {
				err := f.Close()
				if err != nil {
					return
				}
			}(f)
			if err != nil {
				return err
			}
		}
		// fileInfo, err := os.Stat(path)
		// if os.IsNotExist(err) {
		// 	file, err := os.Create(path)
		// 	if err != nil {
		// 		return err
		// 	}
		// 	defer func(file *os.File) {
		// 		_ = file.Close()
		// 	}(file)
		// } else if err != nil {
		// 	return err
		// } else if !fileInfo.IsDir() {
		// 	return nil
		// }
	}
	return nil
}

// Exists 判断文件或目录是否存在
// @param path 文件或目录路径
// @return bool 是否存在
func Exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// IsAbs 判断路径是否为绝对路径 若不是绝对路经则将其转换为绝对路径
// @param path 路径
// @return string 绝对路径
// @return error 错误信息
func IsAbs(path string) (string, error) {
	// 判断路径是否为绝对路径
	if len(path) == 0 {
		return "", ErrIsNull
	}
	abs := strings.HasPrefix(path, "/")
	if abs {
		return path, nil
	} else {
		s, err := filepath.Abs(path)
		if err != nil {
			return "", err
		} else {
			return s, nil
		}
	}
}
