package csv

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/gocarina/gocsv"
	"github.com/jianlu8023/go-tools/v2/pkg/path"
)

// Client CSV客户端封装
type Client struct {
	// 可以添加一些自定义配置
}

// NewClient 创建一个新的CSV客户端
func NewClient() *Client {
	return &Client{}
}

// Marshal 将结构体切片转换为CSV格式的字节数据
// data: 要转换的结构体切片
// 返回CSV格式的字节数据和可能的错误
func (c *Client) Marshal(data interface{}) ([]byte, error) {
	buffer := &bytes.Buffer{}
	err := gocsv.Marshal(data, buffer)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

// MarshalString 将结构体切片转换为CSV格式的字符串
// data: 要转换的结构体切片
// 返回CSV格式的字符串和可能的错误
func (c *Client) MarshalString(data interface{}) (string, error) {
	buffer := &bytes.Buffer{}
	err := gocsv.Marshal(data, buffer)
	if err != nil {
		return "", err
	}
	return buffer.String(), nil
}

// MarshalFile 将结构体切片写入CSV文件
// data: 要写入的结构体切片
// filename: 目标文件名
// 返回可能的错误
func (c *Client) MarshalFile(data interface{}, filename string) error {
	// 使用path包确保目录存在
	dir := filepath.Dir(filename)

	// 使用path包创建目录
	created, err := path.CreateDir(dir)
	// 只有在创建失败且不是因为目录已存在的情况下才返回错误
	if !created && err != nil {
		// 检查是否是因为目录已存在导致的错误
		if fmt.Sprintf("%s", err) != fmt.Sprintf("该目录已经存在: %s", dir) {
			return fmt.Errorf("创建目录失败: %w", err)
		}
	}

	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("创建文件失败: %w", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Printf("关闭文件失败: %v", err)
		}
	}(file)

	return gocsv.MarshalFile(data, file)
}

// MarshalWithoutHeaders 将结构体切片转换为不包含标题行的CSV格式字节数据
// data: 要转换的结构体切片
// 返回CSV格式的字节数据和可能的错误
func (c *Client) MarshalWithoutHeaders(data interface{}) ([]byte, error) {
	buffer := &bytes.Buffer{}
	err := gocsv.MarshalWithoutHeaders(data, buffer)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

// Unmarshal 将CSV格式的数据解析为结构体切片
// data: CSV格式的数据
// out: 用于接收解析结果的结构体切片指针
// 返回可能的错误
func (c *Client) Unmarshal(data []byte, out interface{}) error {
	return gocsv.UnmarshalBytes(data, out)
}

// UnmarshalString 将CSV格式的字符串解析为结构体切片
// data: CSV格式的字符串
// out: 用于接收解析结果的结构体切片指针
// 返回可能的错误
func (c *Client) UnmarshalString(data string, out interface{}) error {
	return gocsv.UnmarshalString(data, out)
}

// UnmarshalFile 从CSV文件中读取数据并解析为结构体切片
// filename: CSV文件名
// out: 用于接收解析结果的结构体切片指针
// 返回可能的错误
func (c *Client) UnmarshalFile(filename string, out interface{}) error {
	// 使用path包检查文件是否存在
	exists, err := path.FileExists(filename)
	if err != nil {
		return fmt.Errorf("检查文件存在性失败: %w", err)
	}

	if !exists {
		return fmt.Errorf("文件不存在: %s", filename)
	}

	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("打开文件失败: %w", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

			fmt.Printf("关闭文件失败: %v", err)
		}
	}(file)

	return gocsv.UnmarshalFile(file, out)
}

// UnmarshalWithoutHeaders 将不包含标题行的CSV格式数据解析为结构体切片
// data: 不包含标题行的CSV格式数据
// out: 用于接收解析结果的结构体切片指针
// 返回可能的错误
func (c *Client) UnmarshalWithoutHeaders(data []byte, out interface{}) error {
	return gocsv.UnmarshalWithoutHeaders(bytes.NewReader(data), out)
}

// SetCSVReader 设置自定义CSV读取器
// csvReader: CSV读取器工厂函数
func (c *Client) SetCSVReader(csvReader func(io.Reader) gocsv.CSVReader) {
	gocsv.SetCSVReader(csvReader)
}

// SetCSVWriter 设置自定义CSV写入器
// csvWriter: CSV写入器工厂函数
func (c *Client) SetCSVWriter(csvWriter func(io.Writer) *gocsv.SafeCSVWriter) {
	gocsv.SetCSVWriter(csvWriter)
}

// SetHeaderNormalizer 设置标题标准化函数
// normalizer: 标题标准化函数
func (c *Client) SetHeaderNormalizer(normalizer gocsv.Normalizer) {
	gocsv.SetHeaderNormalizer(normalizer)
}

// defaultClient 默认CSV客户端实例（包内私有）
var defaultClient = NewClient()

// Marshal 使用默认客户端将结构体切片转换为CSV格式的字节数据
// data: 要转换的结构体切片
// 返回CSV格式的字节数据和可能的错误
func Marshal(data interface{}) ([]byte, error) {
	return defaultClient.Marshal(data)
}

// MarshalString 使用默认客户端将结构体切片转换为CSV格式的字符串
// data: 要转换的结构体切片
// 返回CSV格式的字符串和可能的错误
func MarshalString(data interface{}) (string, error) {
	return defaultClient.MarshalString(data)
}

// MarshalFile 使用默认客户端将结构体切片写入CSV文件
// data: 要写入的结构体切片
// filename: 目标文件名
// 返回可能的错误
func MarshalFile(data interface{}, filename string) error {
	return defaultClient.MarshalFile(data, filename)
}

// MarshalWithoutHeaders 使用默认客户端将结构体切片转换为不包含标题行的CSV格式字节数据
// data: 要转换的结构体切片
// 返回CSV格式的字节数据和可能的错误
func MarshalWithoutHeaders(data interface{}) ([]byte, error) {
	return defaultClient.MarshalWithoutHeaders(data)
}

// Unmarshal 使用默认客户端将CSV格式的数据解析为结构体切片
// data: CSV格式的数据
// out: 用于接收解析结果的结构体切片指针
// 返回可能的错误
func Unmarshal(data []byte, out interface{}) error {
	return defaultClient.Unmarshal(data, out)
}

// UnmarshalString 使用默认客户端将CSV格式的字符串解析为结构体切片
// data: CSV格式的字符串
// out: 用于接收解析结果的结构体切片指针
// 返回可能的错误
func UnmarshalString(data string, out interface{}) error {
	return defaultClient.UnmarshalString(data, out)
}

// UnmarshalFile 使用默认客户端从CSV文件中读取数据并解析为结构体切片
// filename: CSV文件名
// out: 用于接收解析结果的结构体切片指针
// 返回可能的错误
func UnmarshalFile(filename string, out interface{}) error {
	return defaultClient.UnmarshalFile(filename, out)
}

// UnmarshalWithoutHeaders 使用默认客户端将不包含标题行的CSV格式数据解析为结构体切片
// data: 不包含标题行的CSV格式数据
// out: 用于接收解析结果的结构体切片指针
// 返回可能的错误
func UnmarshalWithoutHeaders(data []byte, out interface{}) error {
	return defaultClient.UnmarshalWithoutHeaders(data, out)
}
