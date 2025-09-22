package stringer

import (
	"regexp"
	"strings"
)

const (
	// True 字符串true
	True = "true"
	// False 字符串false
	False = "false"
	// Blank 空字符串
	Blank = ""
)

// IsBlank 判断字符串是否为空
// @param str 字符串
// @return 是否为空
func IsBlank(str string) bool {
	return CompareIgnoreCase(Blank, str)
}

// IsTrue 判断字符串是否为true
// @param str 字符串
// @return 是否为true
func IsTrue(str string) bool {
	return CompareIgnoreCase(True, str)
}

// IsFalse 判断字符串是否为false
// @param str 字符串
// @return 是否为false
func IsFalse(str string) bool {
	return CompareIgnoreCase(False, str)
}

// CompareIgnoreCase 比较两个字符串是否一样
// @param str1: 字符串1
// @param str2: 字符串2
// @return bool: 两个字符串是否一致
func CompareIgnoreCase(str1, str2 string) bool {
	// 去除空格
	str1 = strings.TrimSpace(str1)
	str2 = strings.TrimSpace(str2)

	// 去除str中间的空格
	str1 = strings.Join(strings.Fields(str1), "")
	str2 = strings.Join(strings.Fields(str2), "")

	// 全部转小写
	str1 = strings.ToLower(str1)
	str2 = strings.ToLower(str2)

	// 比较字符串
	return str1 == str2
}

// DeduplicateString 对指定分隔符的字符串进行去重，保留元素首次出现的顺序
// @param str 要去重的字符串
// @param separator 用于分割和拼接字符串的分隔符
// @param ignoreEmpty 是否忽略因连续分隔符或开头/结尾分隔符产生的空字符串元素
// @return string 去重后的字符串
func DeduplicateString(str string, separator string, ignoreEmpty bool) string {
	// 如果输入是空字符串，直接返回空字符串
	if str == "" {
		return ""
	}

	// 如果分隔符是空字符串， Split 会将每个字符作为元素，Join 会直接拼接
	// 这可能是预期行为，所以不做特殊处理
	// if separator == "" { ... }

	// 1. 分割字符串
	elements := strings.Split(str, separator)

	// 2. 去重并保持顺序
	// 使用 map 来记录已经遇到的元素，key 是元素值，value 是空结构体（节省内存）
	seen := make(map[string]struct{})
	// 使用一个切片来存储去重后的元素，按照它们首次出现的顺序
	var uniqueElements []string

	for _, element := range elements {
		// 根据 ignoreEmpty 参数，判断是否跳过空字符串元素
		if ignoreEmpty && IsBlank(element) {
			continue
		}

		// 检查当前元素是否已经在 seen map 中
		if _, ok := seen[element]; !ok {
			// 如果不在，说明是第一次遇到这个元素
			// 将元素添加到 seen map 中
			seen[element] = struct{}{}
			// 将元素添加到 uniqueElements 切片中，保持顺序
			uniqueElements = append(uniqueElements, element)
		}
	}

	// 3. 拼接字符串
	// 使用指定的分隔符将去重后的元素切片拼接回字符串
	return strings.Join(uniqueElements, separator)
}

// Replace 替换字符串中的模式
// @param template: 模板
// @param replacements: 替换的信息
// @return string: 替换后
func Replace(template string, replacements map[string]string) string {
	for pattern, replacement := range replacements {
		re := regexp.MustCompile(pattern)
		template = re.ReplaceAllString(template, replacement)

	}
	return template
}

// Substring 截取子字符串
// @param str: 原始字符串
// @param start: 起始索引
// @param end: 结束索引
// @return string: 截取的子字符串
func Substring(str string, start, end int) string {
	// 处理边界条件
	if start < 0 {
		start = 0
	}
	if end > len(str) {
		end = len(str)
	}
	if start >= end {
		return ""
	}
	return str[start:end]
}

// CountOccurrences 计算子字符串出现的次数
// @param str: 原始字符串
// @param substr: 子字符串
// @return int: 出现次数
func CountOccurrences(str, substr string) int {
	if len(substr) == 0 {
		return 0
	}
	count := 0
	start := 0
	for {
		index := strings.Index(str[start:], substr)
		if index == -1 {
			break
		}
		count++
		start += index + len(substr)
	}
	return count
}

// ToCamelCase 转换为驼峰命名
// @param str: 原始字符串
// @return string: 驼峰命名的字符串
func ToCamelCase(str string) string {
	// 分割字符串
	split := regexp.MustCompile(`[\s_-]+`).Split(str, -1)
	var result strings.Builder

	for i, word := range split {
		if i == 0 {
			// 第一个单词保持小写
			result.WriteString(strings.ToLower(word))
		} else {
			// 后续单词首字母大写
			if len(word) > 0 {
				result.WriteString(strings.ToUpper(string(word[0])))
				if len(word) > 1 {
					result.WriteString(strings.ToLower(word[1:]))
				}
			}
		}
	}

	return result.String()
}

// ToSnakeCase 转换为下划线命名
// @param str: 原始字符串
// @return string: 下划线命名的字符串
func ToSnakeCase(str string) string {
	// 处理驼峰命名
	snake := regexp.MustCompile(`([a-z0-9])([A-Z])`).ReplaceAllString(str, `${1}_${2}`)
	// 处理空格和连字符
	snake = regexp.MustCompile(`[\s-]+`).ReplaceAllString(snake, "_")
	// 全部转为小写
	return strings.ToLower(snake)
}

// Truncate 截断字符串到指定长度
// @param str: 原始字符串
// @param maxLength: 最大长度
// @param suffix: 截断后的后缀
// @return string: 截断后的字符串
func Truncate(str string, maxLength int, suffix string) string {
	if len(str) <= maxLength {
		return str
	}
	if maxLength <= len(suffix) {
		return suffix[:maxLength]
	}
	return str[:maxLength-len(suffix)] + suffix
}

// PadLeft 在字符串左侧填充字符
// @param str: 原始字符串
// @param length: 目标长度
// @param padChar: 填充字符
// @return string: 填充后的字符串
func PadLeft(str string, length int, padChar string) string {
	if len(str) >= length {
		return str
	}
	padding := strings.Repeat(padChar, length-len(str))
	return padding + str
}

// PadRight 在字符串右侧填充字符
// @param str: 原始字符串
// @param length: 目标长度
// @param padChar: 填充字符
// @return string: 填充后的字符串
func PadRight(str string, length int, padChar string) string {
	if len(str) >= length {
		return str
	}
	padding := strings.Repeat(padChar, length-len(str))
	return str + padding
}

// Capitalize 将字符串首字母大写
// @param str: 原始字符串
// @return string: 首字母大写的字符串
func Capitalize(str string) string {
	if len(str) == 0 {
		return str
	}
	return strings.ToUpper(string(str[0])) + strings.ToLower(str[1:])
}

// IsNumeric 判断字符串是否为纯数字
// @param str: 待判断的字符串
// @return bool: 是否为纯数字
func IsNumeric(str string) bool {
	match, _ := regexp.MatchString(`^\d+$`, str)
	return match
}

// Reverse 反转字符串
// @param str: 原始字符串
// @return string: 反转后的字符串
func Reverse(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// Join 将字符串切片连接为一个字符串
// @param elements: 字符串切片
// @param separator: 分隔符
// @return string: 连接后的字符串
func Join(elements []string, separator string) string {
	return strings.Join(elements, separator)
}
