package nlp_test

import (
	"fmt"

	"github.com/jianlu8023/go-tools/v2/pkg/nlp"
)

// ExampleCut 演示如何使用基本分词功能
func ExampleCut() {
	text := "我爱编程和开源技术"

	// 使用默认分词器进行分词
	segments := nlp.Cut(text)
	fmt.Println(segments)

	// 使用默认分词器进行搜索引擎模式分词
	segments2 := nlp.CutSearch(text)
	fmt.Println(segments2)

	// Output:
	// [我 爱 编程 和 开源 技术]
	// [我 爱 编程 和 开源 技术]
}

// ExampleNew 演示如何创建自定义分词器
func ExampleNew() {
	// 创建自定义分词器
	seg, err := nlp.New()
	if err != nil {
		fmt.Printf("创建分词器失败: %v\n", err)
		return
	}

	text := "Go语言是一门优秀的编程语言"
	segments := seg.Cut(text)
	fmt.Println(segments)

	// Output:
	// [go 语言 是 一门 优秀 的 编程语言]
}

// ExamplePos 演示如何进行词性标注
func ExamplePos() {
	text := "我爱编程"

	// 进行词性标注
	pos := nlp.Pos(text)
	for _, item := range pos {
		fmt.Printf("%s/%s ", item.Text, item.Pos)
	}
	fmt.Println()

	// Output:
	// 我/r 爱/v 编程/n
}

// ExampleJoin 演示如何连接分词结果
func ExampleJoin() {
	segments := []string{"我", "爱", "编程"}
	result := nlp.Join(segments, "-")
	fmt.Println(result)

	// Output:
	// 我-爱-编程
}

// ExampleCutAll 演示全模式分词
func ExampleCutAll() {
	text := "我爱编程"

	// 全模式分词
	segments := nlp.CutAll(text)
	fmt.Println(segments)

	// Output:
	// [我 爱 编程]
}
