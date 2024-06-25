package path

import (
	"fmt"
	"os"
)

// Ensure 确保目录存在
// @param path 目录路径
// @return error 错误信息
func Ensure(path string) error {
	err := os.MkdirAll(path, os.ModeDir)
	if err != nil {
		fmt.Println("Mkdir Error ", err)
		return err
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
