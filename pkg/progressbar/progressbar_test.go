package progressbar

import (
	"io"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	// 测试创建新的进度条
	pb := New(100)
	assert.NotNil(t, pb)
	assert.Equal(t, 100, pb.GetMax())
}

func TestNew64(t *testing.T) {
	// 测试创建新的进度条(64位)
	pb := New64(1000)
	assert.NotNil(t, pb)
	assert.Equal(t, int64(1000), pb.GetMax64())
}

func TestDefault(t *testing.T) {
	// 测试创建默认进度条
	pb := Default(100, "test")
	assert.NotNil(t, pb)
	assert.Equal(t, int64(100), pb.GetMax64())
}

func TestDefaultBytes(t *testing.T) {
	// 测试创建默认字节进度条
	pb := DefaultBytes(1024, "download")
	assert.NotNil(t, pb)
	assert.Equal(t, int64(1024), pb.GetMax64())
}

func TestAdd(t *testing.T) {
	// 测试增加进度
	pb := New(100)
	err := pb.Add(10)
	assert.NoError(t, err)
	// 注意：由于进度条的内部状态无法直接访问，我们只能测试不返回错误
}

func TestAdd64(t *testing.T) {
	// 测试增加进度(64位)
	pb := New64(1000)
	err := pb.Add64(100)
	assert.NoError(t, err)
}

func TestSet(t *testing.T) {
	// 测试设置进度
	pb := New(100)
	err := pb.Set(50)
	assert.NoError(t, err)
}

func TestSet64(t *testing.T) {
	// 测试设置进度(64位)
	pb := New64(1000)
	err := pb.Set64(500)
	assert.NoError(t, err)
}

func TestFinish(t *testing.T) {
	// 测试完成进度条
	pb := New(100)
	err := pb.Finish()
	assert.NoError(t, err)
	assert.True(t, pb.IsFinished())
}

func TestReset(t *testing.T) {
	// 测试重置进度条
	pb := New(100)
	pb.Add(50)
	pb.Reset()
	// 重置后进度应该为0，但由于封装限制无法直接验证
}

func TestDescribe(t *testing.T) {
	// 测试设置描述
	pb := New(100)
	pb.Describe("test description")
	// 由于封装限制无法直接验证描述是否设置成功
}

func TestChangeMax(t *testing.T) {
	// 测试更改最大值
	pb := New(100)
	pb.ChangeMax(200)
	assert.Equal(t, 200, pb.GetMax())
}

func TestChangeMax64(t *testing.T) {
	// 测试更改最大值(64位)
	pb := New64(1000)
	pb.ChangeMax64(2000)
	assert.Equal(t, int64(2000), pb.GetMax64())
}

func TestAddMax(t *testing.T) {
	// 测试增加最大值
	pb := New(100)
	pb.AddMax(50)
	assert.Equal(t, 150, pb.GetMax())
}

func TestAddMax64(t *testing.T) {
	// 测试增加最大值(64位)
	pb := New64(1000)
	pb.AddMax64(500)
	assert.Equal(t, int64(1500), pb.GetMax64())
}

func TestIsFinished(t *testing.T) {
	// 测试检查是否完成
	pb := New(100)
	assert.False(t, pb.IsFinished())
	pb.Finish()
	assert.True(t, pb.IsFinished())
}

func TestState(t *testing.T) {
	// 测试获取状态
	pb := New(100)
	state := pb.State()
	assert.NotNil(t, state)
}

func TestClear(t *testing.T) {
	// 测试清除进度条
	pb := New(100)
	err := pb.Clear()
	assert.NoError(t, err)
}

func TestExit(t *testing.T) {
	// 测试退出进度条
	pb := New(100)
	err := pb.Exit()
	assert.NoError(t, err)
}

func TestWrite(t *testing.T) {
	// 测试Write方法
	pb := DefaultBytes(100, "test")
	data := []byte("test data")
	n, err := pb.Write(data)
	assert.NoError(t, err)
	assert.Equal(t, len(data), n)
}

func TestString(t *testing.T) {
	// 测试String方法
	pb := New(100)
	str := pb.String()
	// 字符串可能为空，这取决于进度条的实现，我们只测试不panic
	assert.NotNil(t, str)
}

func TestNewOptions(t *testing.T) {
	// 测试创建带选项的进度条
	pb := NewOptions(100,
		OptionSetWidth(50),
		OptionSetDescription("test"),
		OptionShowCount(),
	)
	assert.NotNil(t, pb)
}

func TestNewOptions64(t *testing.T) {
	// 测试创建带选项的进度条(64位)
	pb := NewOptions64(1000,
		OptionSetWidth(50),
		OptionSetDescription("test"),
		OptionShowCount(),
	)
	assert.NotNil(t, pb)
}

func TestOptionFunctions(t *testing.T) {
	// 测试各种选项函数
	options := []Option{
		OptionSetWidth(50),
		OptionSetDescription("test"),
		OptionEnableColorCodes(true),
		OptionShowBytes(true),
		OptionShowCount(),
		OptionSetVisibility(true),
		OptionFullWidth(),
		OptionThrottle(100 * time.Millisecond),
		OptionClearOnFinish(),
		OptionShowIts(),
		OptionUseANSICodes(true),
		OptionShowTotalBytes(true),
		OptionSetRenderBlankState(true),
		OptionShowDescriptionAtLineEnd(),
		OptionSetPredictTime(true),
		OptionUseIECUnits(true),
		OptionSetSpinnerChangeInterval(100 * time.Millisecond),
		OptionSpinnerType(9),
		OptionSetElapsedTime(true),
	}

	// 确保所有选项函数都能正常创建选项
	for _, opt := range options {
		assert.NotNil(t, opt)
	}
}

func TestTheme(t *testing.T) {
	// 测试主题
	assert.NotNil(t, ThemeDefault)
	assert.NotNil(t, ThemeASCII)
}

func TestOptionSetTheme(t *testing.T) {
	// 测试设置主题选项
	theme := Theme{
		Saucer:        "=",
		SaucerHead:    ">",
		SaucerPadding: " ",
		BarStart:      "[",
		BarEnd:        "]",
	}

	option := OptionSetTheme(theme)
	assert.NotNil(t, option)
}

func TestNewReader(t *testing.T) {
	// 测试创建Reader
	data := "test data for reader"
	reader := strings.NewReader(data)
	pb := Default(int64(len(data)), "reading")

	wrappedReader := NewReader(reader, pb)
	assert.NotNil(t, wrappedReader)

	// 测试读取
	buffer := make([]byte, len(data))
	n, err := wrappedReader.Read(buffer)
	assert.NoError(t, err)
	assert.Equal(t, len(data), n)
	assert.Equal(t, data, string(buffer[:n]))
}

func TestReaderClose(t *testing.T) {
	// 测试Reader关闭
	reader := &mockReadCloser{}
	pb := Default(100, "test")
	wrappedReader := NewReader(reader, pb)

	err := wrappedReader.Close()
	assert.NoError(t, err)
}

// mockReadCloser 用于测试的模拟ReadCloser
type mockReadCloser struct {
	closed bool
}

func (m *mockReadCloser) Read(p []byte) (n int, err error) {
	if m.closed {
		return 0, io.EOF
	}
	return 0, nil
}

func (m *mockReadCloser) Close() error {
	m.closed = true
	return nil
}

func TestDefaultSilent(t *testing.T) {
	// 测试创建静默默认进度条
	pb := DefaultSilent(100, "silent test")
	assert.NotNil(t, pb)
}

func TestDefaultBytesSilent(t *testing.T) {
	// 测试创建静默默认字节进度条
	pb := DefaultBytesSilent(1024, "silent download")
	assert.NotNil(t, pb)
}

func TestOptionSetItsString(t *testing.T) {
	// 测试设置迭代速度字符串选项
	option := OptionSetItsString("it/s")
	assert.NotNil(t, option)
}

func TestOptionSpinnerCustom(t *testing.T) {
	// 测试自定义旋转动画选项
	spinner := []string{"|", "/", "-", "\\"}
	option := OptionSpinnerCustom(spinner)
	assert.NotNil(t, option)
}

func TestOptionSetMaxDetailRow(t *testing.T) {
	// 测试设置最大详细行选项
	option := OptionSetMaxDetailRow(5)
	assert.NotNil(t, option)
}

func TestRenderBlank(t *testing.T) {
	// 测试渲染空白进度条
	pb := New(100)
	err := pb.RenderBlank()
	assert.NoError(t, err)
}

func TestStartWithoutRender(t *testing.T) {
	// 测试开始但不渲染
	pb := New(100)
	pb.StartWithoutRender()
	// 由于是异步操作，只能测试不panic
	assert.True(t, pb.IsStarted())
}
