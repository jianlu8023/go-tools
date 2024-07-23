package file

import (
	"io"
	"os"

	p "github.com/jianlu8023/go-tools/pkg/path"
)

// WriteToFile 写入文件
// @param path 文件路径
// @param content 写入内容
// @return error 错误信息
// @return bool 是否复盖
func WriteToFile(path string, content string, force bool) error {
	if abs, err := p.IsAbs(path); err == nil {
		exists := p.Exists(abs)
		if exists && force {
			_ = os.RemoveAll(abs)
			return overrideWrite(abs, content)
		} else if exists && !force {
			return appendWrite(abs, content)
		} else {
			return overrideWrite(abs, content)
		}
	} else {
		return err
	}
}

// appendWrite 追加写入文件
// @param path 文件路径
// @param content 写入内容
// @return error 错误信息
func appendWrite(path, content string) error {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}

	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	_, err = io.WriteString(file, content)
	if err != nil {
		return err
	}

	return nil
}

// overrideWrite 复盖写入文件
// @param path 文件路径
// @param content 写入内容
// @return error 错误信息
func overrideWrite(path, content string) error {
	err := p.Ensure(path)
	if err != nil {
		return err
	}
	file, err := os.OpenFile(path, os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	_, err = io.WriteString(file, content)
	if err != nil {
		return err
	}
	return nil
}
