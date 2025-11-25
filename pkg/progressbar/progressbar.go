package progressbar

import (
	"io"
	"time"

	"github.com/schollz/progressbar/v3"
)

// ProgressBar 封装进度条功能
type ProgressBar struct {
	*progressbar.ProgressBar
}

// New 创建一个新的进度条
func New(max int) *ProgressBar {
	return &ProgressBar{ProgressBar: progressbar.New(max)}
}

// New64 创建一个新的进度条(64位)
func New64(max int64) *ProgressBar {
	return &ProgressBar{ProgressBar: progressbar.New64(max)}
}

// Default 创建默认进度条
func Default(max int64, description ...string) *ProgressBar {
	return &ProgressBar{ProgressBar: progressbar.Default(max, description...)}
}

// DefaultBytes 创建默认字节进度条
func DefaultBytes(maxBytes int64, description ...string) *ProgressBar {
	return &ProgressBar{ProgressBar: progressbar.DefaultBytes(maxBytes, description...)}
}

// DefaultSilent 创建静默默认进度条
func DefaultSilent(max int64, description ...string) *ProgressBar {
	return &ProgressBar{ProgressBar: progressbar.DefaultSilent(max, description...)}
}

// DefaultBytesSilent 创建静默默认字节进度条
func DefaultBytesSilent(maxBytes int64, description ...string) *ProgressBar {
	return &ProgressBar{ProgressBar: progressbar.DefaultBytesSilent(maxBytes, description...)}
}

// NewOptions 创建带选项的进度条
func NewOptions(max int, options ...Option) *ProgressBar {
	// 转换选项
	pbOptions := make([]progressbar.Option, len(options))
	for i, opt := range options {
		pbOptions[i] = progressbar.Option(opt)
	}
	return &ProgressBar{ProgressBar: progressbar.NewOptions(max, pbOptions...)}
}

// NewOptions64 创建带选项的进度条(64位)
func NewOptions64(max int64, options ...Option) *ProgressBar {
	// 转换选项
	pbOptions := make([]progressbar.Option, len(options))
	for i, opt := range options {
		pbOptions[i] = progressbar.Option(opt)
	}
	return &ProgressBar{ProgressBar: progressbar.NewOptions64(max, pbOptions...)}
}

// Add 增加进度
func (pb *ProgressBar) Add(num int) error {
	return pb.ProgressBar.Add(num)
}

// Add64 增加进度(64位)
func (pb *ProgressBar) Add64(num int64) error {
	return pb.ProgressBar.Add64(num)
}

// Set 设置进度
func (pb *ProgressBar) Set(num int) error {
	return pb.ProgressBar.Set(num)
}

// Set64 设置进度(64位)
func (pb *ProgressBar) Set64(num int64) error {
	return pb.ProgressBar.Set64(num)
}

// Finish 完成进度条
func (pb *ProgressBar) Finish() error {
	return pb.ProgressBar.Finish()
}

// Close 关闭进度条
func (pb *ProgressBar) Close() error {
	return pb.ProgressBar.Close()
}

// Reset 重置进度条
func (pb *ProgressBar) Reset() {
	pb.ProgressBar.Reset()
}

// Describe 设置描述
func (pb *ProgressBar) Describe(description string) {
	pb.ProgressBar.Describe(description)
}

// ChangeMax 更改最大值
func (pb *ProgressBar) ChangeMax(newMax int) {
	pb.ProgressBar.ChangeMax(newMax)
}

// ChangeMax64 更改最大值(64位)
func (pb *ProgressBar) ChangeMax64(newMax int64) {
	pb.ProgressBar.ChangeMax64(newMax)
}

// AddMax 增加最大值
func (pb *ProgressBar) AddMax(added int) {
	pb.ProgressBar.AddMax(added)
}

// AddMax64 增加最大值(64位)
func (pb *ProgressBar) AddMax64(added int64) {
	pb.ProgressBar.AddMax64(added)
}

// GetMax 获取最大值
func (pb *ProgressBar) GetMax() int {
	return pb.ProgressBar.GetMax()
}

// GetMax64 获取最大值(64位)
func (pb *ProgressBar) GetMax64() int64 {
	return pb.ProgressBar.GetMax64()
}

// IsFinished 检查是否完成
func (pb *ProgressBar) IsFinished() bool {
	return pb.ProgressBar.IsFinished()
}

// IsStarted 检查是否已开始
func (pb *ProgressBar) IsStarted() bool {
	return pb.ProgressBar.IsStarted()
}

// State 获取状态
func (pb *ProgressBar) State() State {
	state := pb.ProgressBar.State()
	return State{
		CurrentPercent: state.CurrentPercent,
		CurrentBytes:   state.CurrentBytes,
		SecondsSince:   state.SecondsSince,
		SecondsLeft:    state.SecondsLeft,
		KBsPerSecond:   state.KBsPerSecond,
		Description:    state.Description,
	}
}

// RenderBlank 渲染空白进度条
func (pb *ProgressBar) RenderBlank() error {
	return pb.ProgressBar.RenderBlank()
}

// Clear 清除进度条
func (pb *ProgressBar) Clear() error {
	return pb.ProgressBar.Clear()
}

// Exit 退出进度条
func (pb *ProgressBar) Exit() error {
	return pb.ProgressBar.Exit()
}

// StartWithoutRender 开始但不渲染
func (pb *ProgressBar) StartWithoutRender() {
	pb.ProgressBar.StartWithoutRender()
}

// StartHTTPServer 启动HTTP服务器
func (pb *ProgressBar) StartHTTPServer(hostPort string) {
	pb.ProgressBar.StartHTTPServer(hostPort)
}

// Write 实现io.Writer接口
func (pb *ProgressBar) Write(b []byte) (n int, err error) {
	return pb.ProgressBar.Write(b)
}

// Read 实现io.Reader接口
func (pb *ProgressBar) Read(b []byte) (n int, err error) {
	return pb.ProgressBar.Read(b)
}

// String 获取进度条字符串表示
func (pb *ProgressBar) String() string {
	return pb.ProgressBar.String()
}

// Reader 封装Reader功能
type Reader struct {
	progressbar.Reader
}

// NewReader 创建新的Reader
func NewReader(r io.Reader, bar *ProgressBar) Reader {
	return Reader{Reader: progressbar.NewReader(r, bar.ProgressBar)}
}

// Read 实现io.Reader接口
func (r *Reader) Read(p []byte) (n int, err error) {
	return r.Reader.Read(p)
}

// Close 实现io.Closer接口
func (r *Reader) Close() error {
	return r.Reader.Close()
}

// State 进度条状态
type State struct {
	CurrentPercent float64
	CurrentBytes   float64
	SecondsSince   float64
	SecondsLeft    float64
	KBsPerSecond   float64
	Description    string
}

// Theme 主题
type Theme struct {
	Saucer        string
	SaucerHead    string
	SaucerPadding string
	BarStart      string
	BarEnd        string
}

// Option 选项类型
type Option progressbar.Option

// OptionSetWriter 设置写入器
func OptionSetWriter(w io.Writer) Option {
	return Option(progressbar.OptionSetWriter(w))
}

// OptionSetWidth 设置宽度
func OptionSetWidth(s int) Option {
	return Option(progressbar.OptionSetWidth(s))
}

// OptionSetDescription 设置描述
func OptionSetDescription(description string) Option {
	return Option(progressbar.OptionSetDescription(description))
}

// OptionEnableColorCodes 启用颜色代码
func OptionEnableColorCodes(colorCodes bool) Option {
	return Option(progressbar.OptionEnableColorCodes(colorCodes))
}

// OptionShowBytes 显示字节
func OptionShowBytes(val bool) Option {
	return Option(progressbar.OptionShowBytes(val))
}

// OptionShowCount 显示计数
func OptionShowCount() Option {
	return Option(progressbar.OptionShowCount())
}

// OptionSetTheme 设置主题
func OptionSetTheme(t Theme) Option {
	pbTheme := progressbar.Theme{
		Saucer:        t.Saucer,
		SaucerHead:    t.SaucerHead,
		SaucerPadding: t.SaucerPadding,
		BarStart:      t.BarStart,
		BarEnd:        t.BarEnd,
	}
	return Option(progressbar.OptionSetTheme(pbTheme))
}

// OptionSetVisibility 设置可见性
func OptionSetVisibility(visibility bool) Option {
	return Option(progressbar.OptionSetVisibility(visibility))
}

// OptionFullWidth 全宽度
func OptionFullWidth() Option {
	return Option(progressbar.OptionFullWidth())
}

// OptionThrottle 节流
func OptionThrottle(duration time.Duration) Option {
	return Option(progressbar.OptionThrottle(duration))
}

// OptionClearOnFinish 完成时清除
func OptionClearOnFinish() Option {
	return Option(progressbar.OptionClearOnFinish())
}

// OptionOnCompletion 完成时回调
func OptionOnCompletion(cmpl func()) Option {
	return Option(progressbar.OptionOnCompletion(cmpl))
}

// OptionShowElapsedTimeOnFinish 完成时显示已用时间
func OptionShowElapsedTimeOnFinish() Option {
	return Option(progressbar.OptionShowElapsedTimeOnFinish())
}

// OptionShowIts 显示迭代速度
func OptionShowIts() Option {
	return Option(progressbar.OptionShowIts())
}

// OptionSetItsString 设置迭代速度字符串
func OptionSetItsString(iterationString string) Option {
	return Option(progressbar.OptionSetItsString(iterationString))
}

// OptionUseANSICodes 使用ANSI代码
func OptionUseANSICodes(val bool) Option {
	return Option(progressbar.OptionUseANSICodes(val))
}

// OptionShowTotalBytes 显示总字节数
func OptionShowTotalBytes(flag bool) Option {
	return Option(progressbar.OptionShowTotalBytes(flag))
}

// OptionSetRenderBlankState 设置渲染空白状态
func OptionSetRenderBlankState(r bool) Option {
	return Option(progressbar.OptionSetRenderBlankState(r))
}

// OptionShowDescriptionAtLineEnd 在行尾显示描述
func OptionShowDescriptionAtLineEnd() Option {
	return Option(progressbar.OptionShowDescriptionAtLineEnd())
}

// OptionSetPredictTime 设置预测时间
func OptionSetPredictTime(predictTime bool) Option {
	return Option(progressbar.OptionSetPredictTime(predictTime))
}

// OptionUseIECUnits 使用IEC单位
func OptionUseIECUnits(val bool) Option {
	return Option(progressbar.OptionUseIECUnits(val))
}

// OptionSetSpinnerChangeInterval 设置旋转动画变化间隔
func OptionSetSpinnerChangeInterval(interval time.Duration) Option {
	return Option(progressbar.OptionSetSpinnerChangeInterval(interval))
}

// OptionSpinnerType 旋转动画类型
func OptionSpinnerType(spinnerType int) Option {
	return Option(progressbar.OptionSpinnerType(spinnerType))
}

// OptionSpinnerCustom 自定义旋转动画
func OptionSpinnerCustom(spinner []string) Option {
	return Option(progressbar.OptionSpinnerCustom(spinner))
}

// OptionSetMaxDetailRow 设置最大详细行
func OptionSetMaxDetailRow(row int) Option {
	return Option(progressbar.OptionSetMaxDetailRow(row))
}

// OptionSetElapsedTime 设置已用时间
func OptionSetElapsedTime(elapsedTime bool) Option {
	return Option(progressbar.OptionSetElapsedTime(elapsedTime))
}

// 默认主题
var (
	ThemeDefault = Theme{
		Saucer:        "█",
		SaucerPadding: " ",
		BarStart:      "|",
		BarEnd:        "|",
	}
	ThemeASCII = Theme{
		Saucer:        "=",
		SaucerPadding: " ",
		BarStart:      "|",
		BarEnd:        "|",
	}
)
