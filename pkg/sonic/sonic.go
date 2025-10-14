package sonic

import (
	"errors"
	"io"
	"os"
	"strings"

	"github.com/bytedance/sonic"
)

var (
	// ErrInvalidJSON 表示JSON格式无效
	ErrInvalidJSON = errors.New("invalid JSON format")
	// ErrFileNotFound 表示文件未找到
	ErrFileNotFound = errors.New("JSON file not found")
	// ErrReadingFile 表示读取文件失败
	ErrReadingFile = errors.New("failed to read JSON file")
	// ErrWritingFile 表示写入文件失败
	ErrWritingFile = errors.New("failed to write JSON file")
)

type Sonic struct {
	defaultSonic sonic.API
}

func NewStandardSonic() *Sonic {
	return &Sonic{
		defaultSonic: sonic.ConfigStd,
	}
}

func NewDefaultSonic() *Sonic {
	return &Sonic{
		defaultSonic: sonic.ConfigDefault,
	}
}

func NewFastestSonic() *Sonic {
	return &Sonic{
		defaultSonic: sonic.ConfigFastest,
	}
}

// Marshal 将Go对象转换为JSON字节数组
// 这是对jsoniter.Marshal的封装，提供更好的性能
func (s *Sonic) Marshal(v interface{}) ([]byte, error) {
	return s.defaultSonic.Marshal(v)
}

// MarshalIndent 生成格式化的JSON字节数组
// prefix: 每行的前缀
// indent: 缩进字符串，通常是空格或制表符
func (s *Sonic) MarshalIndent(v interface{}, prefix, indent string) ([]byte, error) {
	return s.defaultSonic.MarshalIndent(v, prefix, indent)
}

// NewEncoder 创建一个新的JSON编码器
func (s *Sonic) NewEncoder(w io.Writer) sonic.Encoder {
	return s.defaultSonic.NewEncoder(w)
}

// Unmarshal 将JSON字节数组解析为Go对象
func (s *Sonic) Unmarshal(data []byte, v interface{}) error {
	if len(data) == 0 {
		return nil // 空数据视为有效输入
	}
	return s.defaultSonic.Unmarshal(data, v)
}

// NewDecoder 创建一个新的JSON解码器
func (s *Sonic) NewDecoder(r io.Reader) sonic.Decoder {
	return s.defaultSonic.NewDecoder(r)
}

// MarshalPretty 生成美化的JSON字符串，带有标准缩进
func (s *Sonic) MarshalPretty(v interface{}) ([]byte, error) {
	return s.MarshalIndent(v, "", "  ")
}

// MarshalString 将Go对象转换为JSON字符串
func (s *Sonic) MarshalString(v interface{}) (string, error) {
	bytes, err := s.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// UnmarshalString 将JSON字符串解析为Go对象
func (s *Sonic) UnmarshalString(str string, v interface{}) error {
	return s.Unmarshal([]byte(str), v)
}

// ReadFromFile 从文件中读取JSON数据并解析为Go对象
func (s *Sonic) ReadFromFile(filePath string, v interface{}) error {
	// 检查文件是否存在
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return ErrFileNotFound
	}

	// 读取文件内容
	data, err := os.ReadFile(filePath)
	if err != nil {
		return errors.Join(ErrReadingFile, err)
	}

	// 解析JSON数据
	if err := s.Unmarshal(data, v); err != nil {
		return errors.Join(ErrInvalidJSON, err)
	}

	return nil
}

// WriteToFile 将Go对象序列化为JSON并写入文件
func (s *Sonic) WriteToFile(filePath string, v interface{}) error {
	// 序列化Go对象为JSON
	data, err := s.MarshalPretty(v)
	if err != nil {
		return err
	}

	// 写入文件
	if err := os.WriteFile(filePath, data, 0644); err != nil {
		return errors.Join(ErrWritingFile, err)
	}

	return nil
}

// Validate 验证JSON字符串是否有效
func (s *Sonic) Validate(jsonStr string) bool {
	return s.defaultSonic.Valid([]byte(jsonStr))
}

// Compact 压缩JSON字符串，移除所有空白字符
func (s *Sonic) Compact(jsonStr string) (string, error) {
	// 检查JSON是否有效
	if !s.Validate(jsonStr) {
		return "", ErrInvalidJSON
	}

	// 手动实现压缩逻辑
	var buf strings.Builder
	buf.Grow(len(jsonStr)) // 预分配容量
	inString := false
	escapeNext := false

	for _, char := range jsonStr {
		// 处理转义字符
		if escapeNext {
			buf.WriteRune(char)
			escapeNext = false
			continue
		}

		// 处理双引号
		if char == '"' {
			buf.WriteRune(char)
			inString = !inString
			continue
		}

		// 处理转义字符
		if char == '\\' {
			buf.WriteRune(char)
			escapeNext = true
			continue
		}

		// 如果不在字符串中，跳过空白字符
		if !inString && (char == ' ' || char == '\t' || char == '\n' || char == '\r') {
			continue
		}

		// 其他字符正常写入
		buf.WriteRune(char)
	}

	return buf.String(), nil
}

var defaultSonic = NewStandardSonic()

// Marshal 将Go对象转换为JSON字节数组
// 这是对jsoniter.Marshal的封装，提供更好的性能
func Marshal(v interface{}) ([]byte, error) {
	return defaultSonic.Marshal(v)
}

// MarshalIndent 生成格式化的JSON字节数组
// prefix: 每行的前缀
// indent: 缩进字符串，通常是空格或制表符
func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error) {
	return defaultSonic.MarshalIndent(v, prefix, indent)
}

// NewEncoder 创建一个新的JSON编码器
func NewEncoder(w io.Writer) sonic.Encoder {
	return defaultSonic.NewEncoder(w)
}

// Unmarshal 将JSON字节数组解析为Go对象
func Unmarshal(data []byte, v interface{}) error {
	if len(data) == 0 {
		return nil // 空数据视为有效输入
	}
	return defaultSonic.Unmarshal(data, v)
}

// NewDecoder 创建一个新的JSON解码器
func NewDecoder(r io.Reader) sonic.Decoder {
	return defaultSonic.NewDecoder(r)
}

// MarshalPretty 生成美化的JSON字符串，带有标准缩进
func MarshalPretty(v interface{}) ([]byte, error) {
	return MarshalIndent(v, "", "  ")
}

// MarshalString 将Go对象转换为JSON字符串
func MarshalString(v interface{}) (string, error) {
	bytes, err := Marshal(v)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// UnmarshalString 将JSON字符串解析为Go对象
func UnmarshalString(s string, v interface{}) error {
	return Unmarshal([]byte(s), v)
}

// ReadFromFile 从文件中读取JSON数据并解析为Go对象
func ReadFromFile(filePath string, v interface{}) error {
	// 检查文件是否存在
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return ErrFileNotFound
	}

	// 读取文件内容
	data, err := os.ReadFile(filePath)
	if err != nil {
		return errors.Join(ErrReadingFile, err)
	}

	// 解析JSON数据
	if err := Unmarshal(data, v); err != nil {
		return errors.Join(ErrInvalidJSON, err)
	}

	return nil
}

// WriteToFile 将Go对象序列化为JSON并写入文件
func WriteToFile(filePath string, v interface{}) error {
	// 序列化Go对象为JSON
	data, err := MarshalPretty(v)
	if err != nil {
		return err
	}

	// 写入文件
	if err := os.WriteFile(filePath, data, 0644); err != nil {
		return errors.Join(ErrWritingFile, err)
	}

	return nil
}

// Validate 验证JSON字符串是否有效
func Validate(jsonStr string) bool {
	return defaultSonic.Validate(jsonStr)
}

// Compact 压缩JSON字符串，移除所有空白字符
func Compact(jsonStr string) (string, error) {
	// 检查JSON是否有效
	if !Validate(jsonStr) {
		return "", ErrInvalidJSON
	}

	// 手动实现压缩逻辑
	var buf strings.Builder
	buf.Grow(len(jsonStr)) // 预分配容量
	inString := false
	escapeNext := false

	for _, char := range jsonStr {
		// 处理转义字符
		if escapeNext {
			buf.WriteRune(char)
			escapeNext = false
			continue
		}

		// 处理双引号
		if char == '"' {
			buf.WriteRune(char)
			inString = !inString
			continue
		}

		// 处理转义字符
		if char == '\\' {
			buf.WriteRune(char)
			escapeNext = true
			continue
		}

		// 如果不在字符串中，跳过空白字符
		if !inString && (char == ' ' || char == '\t' || char == '\n' || char == '\r') {
			continue
		}

		// 其他字符正常写入
		buf.WriteRune(char)
	}

	return buf.String(), nil
}
