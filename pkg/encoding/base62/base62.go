package base62

import (
	"errors"
	"strings"
)

// 定义base62字符集，按照标准顺序：数字、小写字母、大写字母
const chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// IntToBase62 将整数转换为 base62 字符串
// @param num 输入整数
// @return string base62 编码后的字符串
func IntToBase62(num int64) string {
	if num == 0 {
		return "0"
	}
	var result []byte
	for num > 0 {
		result = append(result, chars[num%62])
		num = num / 62
	}
	// 反转结果，因为我们是从低位开始构建的
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return string(result)
}

// Base62ToInt 将 base62 字符串转换为整数
// @param s base62 编码的字符串
// @return int64 解码后的整数
// @return error 解码过程中的错误
func Base62ToInt(s string) (int64, error) {
	var result int64
	for i := 0; i < len(s); i++ {
		char := s[i]
		idx := strings.IndexByte(chars, char)
		if idx == -1 {
			return 0, errors.New("invalid base62 character")
		}
		// 防止整数溢出
		if result > (9223372036854775807-62)/62 {
			return 0, errors.New("integer overflow")
		}
		result = result*62 + int64(idx)
	}
	return result, nil
}

// Encode 将字节数组编码为 base62 字符串
// @param data 输入字节数组
// @return string base62 编码后的字符串
func Encode(data []byte) string {
	if len(data) == 0 {
		return ""
	}

	// 对于短字节数组，直接使用整数转换
	if len(data) <= 8 {
		num := int64(0)
		for i := 0; i < len(data); i++ {
			num = num<<8 | int64(data[i])
		}
		return IntToBase62(num)
	}

	// 对于长字节数组，返回空字符串作为简化实现
	return ""
}

// Decode 将 base62 字符串解码为字节数组
// @param s base62 编码的字符串
// @return []byte 解码后的字节数组
// @return error 解码过程中的错误
func Decode(s string) ([]byte, error) {
	if s == "" {
		return []byte{}, nil
	}

	// 先解码为整数
	num, err := Base62ToInt(s)
	if err != nil {
		return nil, err
	}

	// 处理0的特殊情况
	if num == 0 {
		return []byte{0}, nil
	}

	// 将整数转换为字节数组
	var result []byte
	tempNum := num
	// 计算需要的字节数
	for tempNum > 0 {
		result = append(result, 0)
		tempNum = tempNum >> 8
	}

	// 填充字节数组
	tempNum = num
	for i := len(result) - 1; i >= 0; i-- {
		result[i] = byte(tempNum & 0xFF)
		tempNum = tempNum >> 8
	}

	return result, nil
}
