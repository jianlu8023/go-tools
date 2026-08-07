package base62

import (
	"errors"
	"strings"
)

// 定义base62字符集，按照标准顺序：数字、小写字母、大写字母
const chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// IntToBase62 将非负整数转换为 base62 字符串
// @param num 输入整数（必须 >= 0）
// @return string base62 编码后的字符串
func IntToBase62(num int64) string {
	if num < 0 {
		return ""
	}
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

// Encode 将字节数组编码为 base62 字符串。
//
// 实现说明：采用与标准库 base32/base64 类似的大端字节流分组算法。
// 每次从输入读取 7 字节（56 位）作为一个单元，将其视为一个大整数并转为 10 位 base62 数字
// （7 字节 = 56 位，最大值 2^56-1 = 72057594037927935，62^10 = 839299365868340224 > 2^56-1）。
// 不足 7 字节的末尾组按实际字节数计算所需 base62 位数，并用前导 '0' 补齐到该位数，
// 解码时按位数还原原始字节数，保证严格对称、不丢失前导零字节。
//
// @param data 输入字节数组
// @return string base62 编码后的字符串
func Encode(data []byte) string {
	if len(data) == 0 {
		return ""
	}

	// 每组处理 7 字节，对应 10 个 base62 字符
	const groupBytes = 7
	const groupDigits = 10

	var result []byte
	// 用前导 '0' 补齐每组到固定 10 位，保证解码时能按 10 位切分
	for i := 0; i < len(data); i += groupBytes {
		end := i + groupBytes
		if end > len(data) {
			end = len(data)
		}
		chunk := data[i:end]

		// 将 chunk 视为大端无符号整数
		var num int64
		for j := 0; j < len(chunk); j++ {
			num = num<<8 | int64(chunk[j])
		}

		// 计算该 chunk 对应的 base62 数字位数
		digits := base62DigitLen(len(chunk))
		// 转为 base62 字符串（低位在前）
		var group []byte
		if num == 0 {
			// 全为 '0' 字符（注意是字符 '0'，不是零值字节）
			group = make([]byte, digits)
			for k := range group {
				group[k] = '0'
			}
		} else {
			for num > 0 {
				group = append(group, chars[num%62])
				num = num / 62
			}
		}
		// 反转得到大端序
		for k, m := 0, len(group)-1; k < m; k, m = k+1, m-1 {
			group[k], group[m] = group[m], group[k]
		}
		// 前导补 '0' 至 digits 位
		padded := make([]byte, digits)
		for k := range padded {
			padded[k] = '0'
		}
		copy(padded[digits-len(group):], group)
		result = append(result, padded...)
	}

	return string(result)
}

// Decode 将 base62 字符串解码为字节数组。
//
// 与 Encode 严格对称：每 10 个 base62 字符对应 7 字节原始数据，
// 末尾组按实际位数还原对应字节数，保证前导零字节不丢失。
//
// @param s base62 编码的字符串
// @return []byte 解码后的字节数组
// @return error 解码过程中的错误
func Decode(s string) ([]byte, error) {
	if s == "" {
		return []byte{}, nil
	}

	const groupBytes = 7
	const groupDigits = 10

	var result []byte
	for i := 0; i < len(s); i += groupDigits {
		end := i + groupDigits
		isLast := end > len(s)
		if isLast {
			end = len(s)
		}
		group := s[i:end]

		// 该组对应的字节数
		var byteLen int
		if isLast {
			byteLen = base62ByteLen(len(group))
		} else {
			byteLen = groupBytes
		}

		// 解析 base62 字符串为整数
		var num int64
		for k := 0; k < len(group); k++ {
			idx := strings.IndexByte(chars, group[k])
			if idx == -1 {
				return nil, errors.New("invalid base62 character")
			}
			// 检查溢出
			if num > (9223372036854775807-int64(idx))/62 {
				return nil, errors.New("integer overflow")
			}
			num = num*62 + int64(idx)
		}

		// 将整数转为大端字节数组，长度为 byteLen（保留前导零）
		chunk := make([]byte, byteLen)
		for k := byteLen - 1; k >= 0; k-- {
			chunk[k] = byte(num & 0xFF)
			num = num >> 8
		}
		result = append(result, chunk...)
	}

	return result, nil
}

// base62DigitLen 返回 n 字节大端无符号整数所需的 base62 字符位数。
// n ∈ [1, 7]，对应 8n 位的整数，需要最小的 d 使 62^d > 2^(8n)-1。
// 预计算结果：1→2, 2→3, 3→5, 4→6, 5→7, 6→9, 7→10
func base62DigitLen(n int) int {
	switch n {
	case 1:
		return 2
	case 2:
		return 3
	case 3:
		return 5
	case 4:
		return 6
	case 5:
		return 7
	case 6:
		return 9
	case 7:
		return 10
	default:
		return 10
	}
}

// base62ByteLen 是 base62DigitLen 的反函数，返回 d 个 base62 字符对应的字节数。
// 对于 Encode 产生的有效位数 {2,3,5,6,7,9,10}，严格对应 {1,2,3,4,5,6,7} 字节。
// 其他位数（1,4,8）不会由 Encode 产生，此处按可表示的最大字节数向下取整以做容错。
func base62ByteLen(d int) int {
	switch d {
	case 1:
		return 1
	case 2:
		return 1
	case 3:
		return 2
	case 4:
		return 2
	case 5:
		return 3
	case 6:
		return 4
	case 7:
		return 5
	case 8:
		return 5
	case 9:
		return 6
	default:
		return 7
	}
}
