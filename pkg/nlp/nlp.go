package nlp

import (
	"strings"

	"github.com/go-ego/gse"
)

// Segmenter 中文分词器
type Segmenter struct {
	seg gse.Segmenter
	log Log
}

// New 创建一个新的分词器实例
// dicts: 字典文件路径，可以为空
func New(dicts ...string) (*Segmenter, error) {
	return newSegmenter(Default, dicts...)
}

func NewWithLog(log Log, dicts ...string) (*Segmenter, error) {
	return newSegmenter(log, dicts...)
}

func newSegmenter(log Log, dicts ...string) (*Segmenter, error) {
	var seg gse.Segmenter
	var err error
	switch log {
	case Off:
		seg.SkipLog = true
		seg.MoreLog = false
	case Less:
		seg.SkipLog = false
		seg.MoreLog = false
	case More:
		seg.SkipLog = false
		seg.MoreLog = true
	case Default:
		// no-op
	default:
		// np-op
	}
	if len(dicts) > 0 {
		err = seg.LoadDict(dicts...)
	} else {
		// 加载默认字典
		err = seg.LoadDict()
		if err != nil {
			return nil, err
		}
	}

	// 加载停用词字典
	err = seg.LoadStop()
	if err != nil {
		return nil, err
	}

	return &Segmenter{seg: seg, log: log}, nil
}

// NewWithEmbed 创建一个新的分词器实例，使用内嵌字典
// dicts: 内嵌字典内容
func NewWithEmbed(dicts ...string) (*Segmenter, error) {
	return newSegmenterWithEmbed(Default, dicts...)
}

func NewWithEmbedWithLog(log Log, dicts ...string) (*Segmenter, error) {
	return newSegmenterWithEmbed(log, dicts...)
}

func newSegmenterWithEmbed(log Log, dicts ...string) (*Segmenter, error) {
	var seg gse.Segmenter
	var err error
	switch log {
	case Off:
		seg.SkipLog = true
		seg.MoreLog = false
	case Less:
		seg.SkipLog = false
		seg.MoreLog = false
	case More:
		seg.SkipLog = false
		seg.MoreLog = true
	case Default:
		// no-op
	default:
		// np-op
	}
	if len(dicts) > 0 {
		err = seg.LoadDictEmbed(dicts...)
	} else {
		// 加载默认内嵌字典
		err = seg.LoadDictEmbed()
		if err != nil {
			return nil, err
		}
	}

	// 加载停用词字典
	err = seg.LoadStopEmbed()
	if err != nil {
		return nil, err
	}

	return &Segmenter{seg: seg, log: log}, nil
}

// Cut 对文本进行分词
// text: 要分词的文本
// hmm: 是否使用HMM算法
func (s *Segmenter) Cut(text string, hmm ...bool) []string {
	return s.seg.Cut(text, hmm...)
}

// CutAll 全模式分词
// text: 要分词的文本
func (s *Segmenter) CutAll(text string) []string {
	return s.seg.CutAll(text)
}

// CutSearch 搜索引擎模式分词
// text: 要分词的文本
// hmm: 是否使用HMM算法
func (s *Segmenter) CutSearch(text string, hmm ...bool) []string {
	return s.seg.CutSearch(text, hmm...)
}

// CutDAG DAG分词
// text: 要分词的文本
func (s *Segmenter) CutDAG(text string) []string {
	return s.seg.CutDAG(text)
}

// Pos 词性标注
// text: 要标注的文本
// hmm: 是否使用HMM算法
func (s *Segmenter) Pos(text string, hmm ...bool) []gse.SegPos {
	return s.seg.Pos(text, hmm...)
}

// AddToken 添加自定义词汇
// text: 词汇
// freq: 频率
// pos: 词性
func (s *Segmenter) AddToken(text string, freq float64, pos ...string) error {
	return s.seg.AddToken(text, freq, pos...)
}

// AddStop 添加停用词
// text: 停用词
func (s *Segmenter) AddStop(text string) {
	s.seg.AddStop(text)
}

// IsStop 判断是否为停用词
// text: 词汇
func (s *Segmenter) IsStop(text string) bool {
	return s.seg.IsStop(text)
}

// Trim 过滤停用词
// text: 分词结果
func (s *Segmenter) Trim(text []string) []string {
	return s.seg.Trim(text)
}

// TrimPos 过滤停用词（词性）
// pos: 词性标注结果
func (s *Segmenter) TrimPos(pos []gse.SegPos) []gse.SegPos {
	return s.seg.TrimPos(pos)
}

// Analyze 分析分词结果
// text: 分词结果
// original: 原始文本
func (s *Segmenter) Analyze(text []string, original string) []gse.AnalyzeToken {
	return s.seg.Analyze(text, original)
}

// HMMCut HMM分词
// text: 要分词的文本
func (s *Segmenter) HMMCut(text string) []string {
	return s.seg.HMMCut(text)
}

// CutStop 使用分词器对文本进行分词并过滤停用词
// text: 要分词的文本
// hmm: 是否使用HMM算法
func (s *Segmenter) CutStop(text string, hmm ...bool) []string {
	return s.seg.CutStop(text, hmm...)
}

// defaultSeg 默认分词器实例
var defaultSeg *Segmenter

func init() {
	// 初始化默认分词器
	defaultSeg, _ = New()
}

// Cut 使用默认分词器对文本进行分词
// text: 要分词的文本
// hmm: 是否使用HMM算法
func Cut(text string, hmm ...bool) []string {
	if defaultSeg == nil {
		return []string{text}
	}
	return defaultSeg.Cut(text, hmm...)
}

// CutAll 使用默认分词器进行全模式分词
// text: 要分词的文本
func CutAll(text string) []string {
	if defaultSeg == nil {
		return []string{text}
	}
	return defaultSeg.CutAll(text)
}

// CutSearch 使用默认分词器进行搜索引擎模式分词
// text: 要分词的文本
// hmm: 是否使用HMM算法
func CutSearch(text string, hmm ...bool) []string {
	if defaultSeg == nil {
		return []string{text}
	}
	return defaultSeg.CutSearch(text, hmm...)
}

// CutStop 使用默认分词器对文本进行分词并过滤停用词
// text: 要分词的文本
// hmm: 是否使用HMM算法
func CutStop(text string, hmm ...bool) []string {
	if defaultSeg == nil {
		return []string{text}
	}
	return defaultSeg.CutStop(text, hmm...)
}

// Pos 使用默认分词器进行词性标注
// text: 要标注的文本
// hmm: 是否使用HMM算法
func Pos(text string, hmm ...bool) []gse.SegPos {
	if defaultSeg == nil {
		return []gse.SegPos{}
	}
	return defaultSeg.Pos(text, hmm...)
}

// Join 将分词结果连接成字符串
// segments: 分词结果
// separator: 分隔符
func Join(segments []string, separator string) string {
	return strings.Join(segments, separator)
}
