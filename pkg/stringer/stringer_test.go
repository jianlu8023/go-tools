package stringer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsBlank(t *testing.T) {
	assert.True(t, IsBlank(""), "Empty string should be considered blank")
	assert.True(t, IsBlank("   "), "Whitespace string should be considered blank")
	assert.False(t, IsBlank("not blank"), "Non-empty string should not be considered blank")
}

func TestIsTrue(t *testing.T) {
	assert.True(t, IsTrue("true"), "String 'true' should be considered true")
	assert.True(t, IsTrue("TRUE"), "String 'TRUE' should be considered true")
	assert.True(t, IsTrue(" True "), "String ' True ' with whitespace should be considered true")
	assert.False(t, IsTrue("false"), "String 'false' should not be considered true")
	assert.False(t, IsTrue("not true"), "String 'not true' should not be considered true")
}

func TestIsFalse(t *testing.T) {
	assert.True(t, IsFalse("false"), "String 'false' should be considered false")
	assert.True(t, IsFalse("FALSE"), "String 'FALSE' should be considered false")
	assert.True(t, IsFalse(" False "), "String ' False ' with whitespace should be considered false")
	assert.False(t, IsFalse("true"), "String 'true' should not be considered false")
	assert.False(t, IsFalse("not false"), "String 'not false' should not be considered false")
}

func TestCompareIgnoreCase(t *testing.T) {
	assert.True(t, CompareIgnoreCase("hello", "HELLO"), "Strings 'hello' and 'HELLO' should be considered equal")
	assert.True(t, CompareIgnoreCase("  hello  ", "HELLO"), "Strings with whitespace should be considered equal")
	assert.True(t, CompareIgnoreCase("hello world", "HELLO  WORLD"), "Strings with internal whitespace should be considered equal")
	assert.False(t, CompareIgnoreCase("hello", "world"), "Different strings should not be considered equal")
}

func TestDeduplicateString(t *testing.T) {
	// 测试基本去重功能
	assert.Equal(t, "a,b,c", DeduplicateString("a,b,a,c,b", ",", false), "Basic deduplication failed")

	// 测试忽略空元素
	assert.Equal(t, "a,b,c", DeduplicateString("a,,,b,,a,c,b", ",", true), "Deduplication with empty elements failed")

	// 测试空字符串输入
	assert.Equal(t, "", DeduplicateString("", ",", false), "Empty string input should return empty string")

	// 测试保持顺序
	assert.Equal(t, "a,b,c", DeduplicateString("a,b,c,a,b,c", ",", false), "Order preservation failed")
}

func TestReplace(t *testing.T) {
	template := "https://api.github.com/repo/{{owner}}/{{repo}}"

	str := Replace(template, map[string]string{
		"{{owner}}": "ceshi",
		"{{repo}}":  "ceshi",
	})
	assert.Equal(t, "https://api.github.com/repo/ceshi/ceshi", str, "Replace function failed")

	// 测试多个替换模式
	complexTemplate := "{{greeting}}, {{name}}! Welcome to {{place}}."
	complexResult := Replace(complexTemplate, map[string]string{
		"{{greeting}}": "Hello",
		"{{name}}":     "User",
		"{{place}}":    "Go Tools",
	})
	assert.Equal(t, "Hello, User! Welcome to Go Tools.", complexResult, "Multiple pattern replacement failed")
}

func TestSubstring(t *testing.T) {
	// 测试基本截取
	assert.Equal(t, "llo", Substring("hello", 2, 5), "Basic substring failed")
	// 测试边界条件
	assert.Equal(t, "hello", Substring("hello", 0, 10), "Substring with end beyond length failed")
	assert.Equal(t, "hello", Substring("hello", -5, 5), "Substring with negative start failed")
	assert.Equal(t, "", Substring("hello", 3, 2), "Substring with start > end failed")
}

func TestCountOccurrences(t *testing.T) {
	// 测试基本计数
	assert.Equal(t, 2, CountOccurrences("hello world, hello golang", "hello"), "Count occurrences failed")
	// 测试不存在的子字符串
	assert.Equal(t, 0, CountOccurrences("hello world", "test"), "Count non-existent substring failed")
	// 测试空字符串
	assert.Equal(t, 0, CountOccurrences("hello", ""), "Count empty substring failed")
}

func TestToCamelCase(t *testing.T) {
	// 测试空格分隔的字符串
	assert.Equal(t, "helloWorld", ToCamelCase("hello world"), "Convert space-separated string to camel case failed")
	// 测试下划线分隔的字符串
	assert.Equal(t, "helloWorld", ToCamelCase("hello_world"), "Convert underscore-separated string to camel case failed")
	// 测试连字符分隔的字符串
	assert.Equal(t, "helloWorld", ToCamelCase("hello-world"), "Convert hyphen-separated string to camel case failed")
	// 测试混合分隔符的字符串
	assert.Equal(t, "helloWorldGolang", ToCamelCase("hello_world-golang"), "Convert mixed separators string to camel case failed")
}

func TestToSnakeCase(t *testing.T) {
	// 测试驼峰命名的字符串
	assert.Equal(t, "hello_world", ToSnakeCase("helloWorld"), "Convert camel case string to snake case failed")
	// 测试空格分隔的字符串
	assert.Equal(t, "hello_world", ToSnakeCase("hello world"), "Convert space-separated string to snake case failed")
	// 测试连字符分隔的字符串
	assert.Equal(t, "hello_world", ToSnakeCase("hello-world"), "Convert hyphen-separated string to snake case failed")
	// 测试混合格式的字符串
	assert.Equal(t, "hello_world_golang", ToSnakeCase("helloWorld-golang"), "Convert mixed format string to snake case failed")
}

func TestTruncate(t *testing.T) {
	// 测试不需要截断的情况
	assert.Equal(t, "hello", Truncate("hello", 10, "..."), "Truncate with sufficient length failed")
	// 测试需要截断的情况
	assert.Equal(t, "hello...", Truncate("hello world", 8, "..."), "Truncate with suffix failed")
	// 测试空字符串
	assert.Equal(t, "", Truncate("", 5, "..."), "Truncate empty string failed")
	// 测试后缀长度大于最大长度的情况
	assert.Equal(t, "...", Truncate("hello world", 3, "..."), "Truncate with suffix longer than max length failed")
}

func TestPadLeft(t *testing.T) {
	// 测试基本填充
	assert.Equal(t, "000hello", PadLeft("hello", 8, "0"), "Left pad with zeros failed")
	// 测试不需要填充的情况
	assert.Equal(t, "hello", PadLeft("hello", 3, "0"), "Left pad with sufficient length failed")
	// 测试空字符串
	assert.Equal(t, "0000", PadLeft("", 4, "0"), "Left pad empty string failed")
}

func TestPadRight(t *testing.T) {
	// 测试基本填充
	assert.Equal(t, "hello000", PadRight("hello", 8, "0"), "Right pad with zeros failed")
	// 测试不需要填充的情况
	assert.Equal(t, "hello", PadRight("hello", 3, "0"), "Right pad with sufficient length failed")
	// 测试空字符串
	assert.Equal(t, "0000", PadRight("", 4, "0"), "Right pad empty string failed")
}

func TestCapitalize(t *testing.T) {
	// 测试基本首字母大写
	assert.Equal(t, "Hello", Capitalize("hello"), "Capitalize first letter failed")
	// 测试全大写字符串
	assert.Equal(t, "Hello", Capitalize("HELLO"), "Capitalize all uppercase string failed")
	// 测试空字符串
	assert.Equal(t, "", Capitalize(""), "Capitalize empty string failed")
}

func TestIsNumeric(t *testing.T) {
	// 测试纯数字字符串
	assert.True(t, IsNumeric("12345"), "Pure numeric string should be considered numeric")
	// 测试非数字字符串
	assert.False(t, IsNumeric("123abc"), "String with letters should not be considered numeric")
	assert.False(t, IsNumeric("abc123"), "String starting with letters should not be considered numeric")
	// 测试空字符串
	assert.False(t, IsNumeric(""), "Empty string should not be considered numeric")
}

func TestReverse(t *testing.T) {
	// 测试基本反转
	assert.Equal(t, "dlrow olleh", Reverse("hello world"), "Basic string reverse failed")
	// 测试空字符串
	assert.Equal(t, "", Reverse(""), "Reverse empty string failed")
	// 测试单字符字符串
	assert.Equal(t, "a", Reverse("a"), "Reverse single character string failed")
}

func TestJoin(t *testing.T) {
	// 测试基本连接
	assert.Equal(t, "a,b,c", Join([]string{"a", "b", "c"}, ","), "Basic string join failed")
	// 测试空切片
	assert.Equal(t, "", Join([]string{}, ","), "Join empty slice failed")
	// 测试空分隔符
	assert.Equal(t, "abc", Join([]string{"a", "b", "c"}, ""), "Join with empty separator failed")
}
