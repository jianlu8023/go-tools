package file

import (
	"fmt"
	"io"
	"os"

	"github.com/jianlu8023/go-tools/internal/logger"
	p "github.com/jianlu8023/go-tools/pkg/path"
)

var log = logger.GetLogger()

// WriteToFile 写入文件
// @param path 文件路径
// @param context 写入内容
// @return error 错误信息
func WriteToFile(path string, context string) error {
	exists := p.Exists(path)
	if exists {

		log.Infof("File %s already exists", path)
		return fmt.Errorf("File %s already exists", path)
	}
	// TODO 需要判断是绝对路经还是相对路经 然后创建文件，添加参数 force 强制覆盖
	file, err := os.Create(path)
	if err != nil {
		// fmt.Println("Create File Error ", err)
		log.Errorf("Create File Error %v", err)
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
