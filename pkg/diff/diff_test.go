package diff

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiff(t *testing.T) {
	// 测试基本差异功能
	seq1 := []string{"apple", "banana", "cherry"}
	seq2 := []string{"apple", "grape", "cherry", "date"}

	result := Diff(seq1, seq2)

	// 验证结果数量
	assert.Equal(t, 5, len(result))

	// 验证具体差异
	assert.Equal(t, "apple", result[0].Payload)
	assert.Equal(t, Common, result[0].Type)

	assert.Equal(t, "banana", result[1].Payload)
	assert.Equal(t, LeftOnly, result[1].Type)

	assert.Equal(t, "grape", result[2].Payload)
	assert.Equal(t, RightOnly, result[2].Type)

	assert.Equal(t, "cherry", result[3].Payload)
	assert.Equal(t, Common, result[3].Type)

	assert.Equal(t, "date", result[4].Payload)
	assert.Equal(t, RightOnly, result[4].Type)
}

func TestHTMLDiff(t *testing.T) {
	// 测试HTML差异功能
	seq1 := []string{"line1", "line2"}
	seq2 := []string{"line1", "line2 modified"}

	result := HTMLDiff(seq1, seq2)

	// 验证返回的HTML不为空
	assert.NotEmpty(t, result)

	// 验证包含HTML标签
	assert.True(t, strings.Contains(result, "<tr"))
	assert.True(t, strings.Contains(result, "<td"))
	assert.True(t, strings.Contains(result, "line1"))
}

func TestSplitLines(t *testing.T) {
	// 测试文本行分割功能
	text := "line1\nline2\nline3\n"
	result := SplitLines(text)

	expected := []string{"line1", "line2", "line3"}
	assert.Equal(t, expected, result)

	// 测试不以换行符结尾的文本
	text2 := "line1\nline2\nline3"
	result2 := SplitLines(text2)
	assert.Equal(t, expected, result2)

	// 测试空文本
	text3 := ""
	result3 := SplitLines(text3)
	assert.Equal(t, []string{}, result3)
}

func TestFormatDiff(t *testing.T) {
	// 测试差异格式化功能
	records := []Record{
		{Payload: "line1", Type: Common},
		{Payload: "line2", Type: LeftOnly},
		{Payload: "line3", Type: RightOnly},
	}

	result := FormatDiff(records)

	expected := " line1\n-line2\n+line3\n"
	assert.Equal(t, expected, result)
}

func TestDiffRecordString(t *testing.T) {
	// 测试DiffRecord的String方法
	record1 := Record{Payload: "test", Type: Common}
	assert.Equal(t, " test", record1.String())

	record2 := Record{Payload: "test", Type: LeftOnly}
	assert.Equal(t, "-test", record2.String())

	record3 := Record{Payload: "test", Type: RightOnly}
	assert.Equal(t, "+test", record3.String())
}

func TestEmptySequences(t *testing.T) {
	// 测试空序列的差异
	seq1 := []string{}
	seq2 := []string{}

	result := Diff(seq1, seq2)
	assert.Equal(t, 0, len(result))

	// 测试一个空序列和一个非空序列
	seq3 := []string{"item1", "item2"}
	seq4 := []string{}

	result2 := Diff(seq3, seq4)
	assert.Equal(t, 2, len(result2))
	assert.Equal(t, LeftOnly, result2[0].Type)
	assert.Equal(t, LeftOnly, result2[1].Type)
}
