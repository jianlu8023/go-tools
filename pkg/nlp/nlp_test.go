package nlp

import (
	"testing"

	"github.com/go-ego/gse"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	// 测试创建分词器实例
	seg, err := New()
	assert.NoError(t, err)
	assert.NotNil(t, seg)

	// 测试创建分词器实例（带字典参数）
	seg2, err := New("zh")
	assert.NoError(t, err)
	assert.NotNil(t, seg2)
}

func TestNewWithEmbed(t *testing.T) {
	// 测试创建分词器实例（内嵌字典）
	seg, err := NewWithEmbed()
	assert.NoError(t, err)
	assert.NotNil(t, seg)

	// 测试创建分词器实例（带内嵌字典参数）
	seg2, err := NewWithEmbed("zh")
	assert.NoError(t, err)
	assert.NotNil(t, seg2)
}

func TestCut(t *testing.T) {
	seg, err := New()
	assert.NoError(t, err)

	// 测试基本分词功能
	text := "我爱编程"
	result := seg.Cut(text)
	assert.NotEmpty(t, result)

	// 测试带HMM的分词功能
	result2 := seg.Cut(text, true)
	assert.NotEmpty(t, result2)
}

func TestCutAll(t *testing.T) {
	seg, err := New()
	assert.NoError(t, err)

	// 测试全模式分词
	text := "我爱编程"
	result := seg.CutAll(text)
	assert.NotEmpty(t, result)
}

func TestCutSearch(t *testing.T) {
	seg, err := New()
	assert.NoError(t, err)

	// 测试搜索引擎模式分词
	text := "我爱编程"
	result := seg.CutSearch(text)
	assert.NotEmpty(t, result)

	// 测试带HMM的搜索引擎模式分词
	result2 := seg.CutSearch(text, true)
	assert.NotEmpty(t, result2)
}

func TestPos(t *testing.T) {
	seg, err := New()
	assert.NoError(t, err)

	// 测试词性标注
	text := "我爱编程"
	result := seg.Pos(text)
	assert.NotEmpty(t, result)

	// 测试带HMM的词性标注
	result2 := seg.Pos(text, true)
	assert.NotEmpty(t, result2)
}

func TestAddToken(t *testing.T) {
	seg, err := New()
	assert.NoError(t, err)

	// 测试添加自定义词汇
	err = seg.AddToken("自定义词汇", 100, "n")
	assert.NoError(t, err)

	// 验证自定义词汇是否生效
	result := seg.Cut("这是自定义词汇")
	assert.Contains(t, result, "自定义词汇")
}

func TestStopWords(t *testing.T) {
	seg, err := New()
	assert.NoError(t, err)

	// 添加停用词
	seg.AddStop("的")

	// 测试停用词判断
	assert.True(t, seg.IsStop("的"))
	assert.False(t, seg.IsStop("编程"))

	// 测试过滤停用词
	segments := []string{"我", "的", "编程"}
	result := seg.Trim(segments)
	assert.NotContains(t, result, "的")
	assert.Contains(t, result, "我")
	assert.Contains(t, result, "编程")
}

func TestTrimPos(t *testing.T) {
	seg, err := New()
	assert.NoError(t, err)

	// 添加停用词
	seg.AddStop("的")

	// 测试过滤停用词（词性）
	pos := []gse.SegPos{
		{Text: "我", Pos: "r"},
		{Text: "的", Pos: "u"},
		{Text: "编程", Pos: "v"},
	}
	result := seg.TrimPos(pos)

	// 验证"的"被过滤掉
	found := false
	for _, item := range result {
		if item.Text == "的" {
			found = true
			break
		}
	}
	assert.False(t, found)
}

func TestAnalyze(t *testing.T) {
	seg, err := New()
	assert.NoError(t, err)

	// 测试分析分词结果
	text := "我爱编程"
	segments := seg.Cut(text)
	result := seg.Analyze(segments, text)
	assert.NotEmpty(t, result)
}

func TestHMMCut(t *testing.T) {
	seg, err := New()
	assert.NoError(t, err)

	// 测试HMM分词
	text := "我爱编程"
	result := seg.HMMCut(text)
	assert.NotEmpty(t, result)
}

func TestCutStop(t *testing.T) {
	seg, err := New()
	assert.NoError(t, err)

	// 测试分词并过滤停用词
	text := "我爱编程的技术"
	result := seg.CutStop(text)
	assert.NotEmpty(t, result)
}

func TestDefaultFunctions(t *testing.T) {
	// 测试默认分词器函数
	text := "我爱编程"

	// 测试Cut函数
	result := Cut(text)
	assert.NotEmpty(t, result)

	// 测试CutAll函数
	result2 := CutAll(text)
	assert.NotEmpty(t, result2)

	// 测试CutSearch函数
	result3 := CutSearch(text)
	assert.NotEmpty(t, result3)

	// 测试CutStop函数
	result4 := CutStop(text)
	assert.NotEmpty(t, result4)

	// 测试Pos函数
	result5 := Pos(text)
	assert.NotEmpty(t, result5)
}

func TestJoin(t *testing.T) {
	// 测试连接分词结果
	segments := []string{"我", "爱", "编程"}
	result := Join(segments, "/")
	expected := "我/爱/编程"
	assert.Equal(t, expected, result)
}
