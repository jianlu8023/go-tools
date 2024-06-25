package file

import (
	"fmt"
	"io"
	"os"
)

// WriteToFile 写入文件
// @param path 文件路径
// @param context 写入内容
// @return error 错误信息
func WriteToFile(path string, context string) error {
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("Create File Error ", err)
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Close File Error ", err)
		}
	}(file)

	_, err = io.WriteString(file, context)
	if err != nil {
		fmt.Println("Write File Error ", err)
		return err
	}
	return nil
}
