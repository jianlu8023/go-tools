package diff

import (
	"strings"

	"github.com/aryann/difflib"
)

// Type 表示差异类型
type Type int

const (
	// Common 表示两个序列中都存在的元素
	Common Type = iota
	// LeftOnly 表示仅在第一个序列中存在的元素
	LeftOnly
	// RightOnly 表示仅在第二个序列中存在的元素
	RightOnly
)

// Record 表示一个差异记录
type Record struct {
	// Payload 差异内容
	Payload string
	// Type 差异类型
	Type Type
}

// String 返回差异记录的字符串表示
func (d Record) String() string {
	var prefix string
	switch d.Type {
	case Common:
		prefix = " "
	case LeftOnly:
		prefix = "-"
	case RightOnly:
		prefix = "+"
	}
	return prefix + d.Payload
}

// Diff 计算两个字符串序列之间的差异
// seq1: 第一个字符串序列
// seq2: 第二个字符串序列
// 返回差异记录列表
func Diff(seq1, seq2 []string) []Record {
	delta := difflib.Diff(seq1, seq2)
	result := make([]Record, len(delta))

	for i, d := range delta {
		record := Record{
			Payload: d.Payload,
		}

		switch d.Delta {
		case difflib.Common:
			record.Type = Common
		case difflib.LeftOnly:
			record.Type = LeftOnly
		case difflib.RightOnly:
			record.Type = RightOnly
		}

		result[i] = record
	}

	return result
}

// HTMLDiff 返回两个字符串序列差异的HTML表示
// seq1: 第一个字符串序列
// seq2: 第二个字符串序列
// 返回HTML格式的差异表示
func HTMLDiff(seq1, seq2 []string) string {
	return difflib.HTMLDiff(seq1, seq2)
}

// SplitLines 将文本按行分割
// text: 要分割的文本
// 返回分割后的字符串切片
// 兼容 Windows(\r\n)、Unix(\n) 和旧版 Mac(\r) 风格的换行符
func SplitLines(text string) []string {
	// 统一换行符：先将 \r\n 转为 \n，再处理孤立的 \r
	normalized := strings.ReplaceAll(text, "\r\n", "\n")
	normalized = strings.ReplaceAll(normalized, "\r", "\n")
	lines := strings.Split(normalized, "\n")
	// 如果最后一行为空，则移除它（对应文本末尾的换行）
	if len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}
	return lines
}

// FormatDiff 格式化显示差异结果
// records: 差异记录列表
// 返回格式化后的差异字符串
func FormatDiff(records []Record) string {
	var builder strings.Builder
	for _, record := range records {
		builder.WriteString(record.String())
		builder.WriteString("\n")
	}
	return builder.String()
}
