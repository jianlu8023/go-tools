package diff_test

import (
	"fmt"

	"github.com/jianlu8023/go-tools/v2/pkg/diff"
)

// ExampleDiff 演示如何使用Diff函数比较两个字符串序列
func ExampleDiff() {
	seq1 := []string{"apple", "banana", "cherry"}
	seq2 := []string{"apple", "grape", "cherry", "date"}

	result := diff.Diff(seq1, seq2)

	for _, record := range result {
		fmt.Println(record.String())
	}

	// Output:
	//  apple
	// -banana
	// +grape
	//  cherry
	// +date
}

// ExampleHTMLDiff 演示如何使用HTMLDiff函数生成HTML格式的差异
func ExampleHTMLDiff() {
	seq1 := []string{"Hello", "World"}
	seq2 := []string{"Hello", "Go", "World"}

	html := diff.HTMLDiff(seq1, seq2)
	fmt.Println(html)

	// Output:
	// <tr><td class="line-num">1</td><td><pre>Hello</pre></td><td><pre>Hello</pre></td><td class="line-num">1</td></tr>
	// <tr><td class="line-num"></td><td></td><td class="added"><pre>Go</pre></td><td class="line-num">2</td></tr>
	// <tr><td class="line-num">2</td><td><pre>World</pre></td><td><pre>World</pre></td><td class="line-num">3</td></tr>
}

// ExampleSplitLines 演示如何使用SplitLines函数分割文本
func ExampleSplitLines() {
	text := "line1\nline2\nline3"
	lines := diff.SplitLines(text)

	for i, line := range lines {
		fmt.Printf("Line %d: %s\n", i+1, line)
	}

	// Output:
	// Line 1: line1
	// Line 2: line2
	// Line 3: line3
}

// ExampleFormatDiff 演示如何使用FormatDiff函数格式化差异结果
func ExampleFormatDiff() {
	records := []diff.Record{
		{Payload: "common line", Type: diff.Common},
		{Payload: "removed line", Type: diff.LeftOnly},
		{Payload: "added line", Type: diff.RightOnly},
	}

	result := diff.FormatDiff(records)
	fmt.Print(result)

	// Output:
	//  common line
	// -removed line
	// +added line
}
