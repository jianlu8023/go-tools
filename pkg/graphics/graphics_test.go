package graphics

import (
	"image/color"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewContext(t *testing.T) {
	// 测试创建新的图形上下文
	ctx := NewContext(100, 100)
	assert.NotNil(t, ctx)
	assert.Equal(t, 100, ctx.Width())
	assert.Equal(t, 100, ctx.Height())
}

func TestNewContextFromImage(t *testing.T) {
	// 创建一个简单的图像用于测试
	ctx := NewContext(50, 50)
	ctx.SetRGBA255(255, 0, 0, 255)
	ctx.DrawRectangle(0, 0, 50, 50)
	ctx.Fill()

	// 从图像创建新的上下文
	newCtx := NewContextFromImage(ctx.Image())
	assert.NotNil(t, newCtx)
	assert.Equal(t, 50, newCtx.Width())
	assert.Equal(t, 50, newCtx.Height())
}

func TestColorMethods(t *testing.T) {
	ctx := NewContext(100, 100)

	// 测试设置颜色的方法
	ctx.SetColor(color.RGBA{255, 0, 0, 255})
	ctx.SetHexColor("#00FF00")
	ctx.SetRGB(0, 0, 1)
	ctx.SetRGBA(0, 1, 0, 0.5)
	ctx.SetRGB255(128, 128, 128)
	ctx.SetRGBA255(64, 64, 64, 128)

	// 验证没有panic
	assert.True(t, true)
}

func TestDrawingMethods(t *testing.T) {
	ctx := NewContext(100, 100)

	// 测试各种绘制方法
	ctx.DrawRectangle(10, 10, 50, 30)
	ctx.DrawRoundedRectangle(10, 50, 50, 30, 5)
	ctx.DrawCircle(75, 25, 15)
	ctx.DrawEllipse(75, 75, 20, 10)
	ctx.DrawRegularPolygon(6, 25, 25, 15, 0)

	// 测试路径操作
	ctx.MoveTo(0, 0)
	ctx.LineTo(100, 100)
	ctx.QuadraticTo(50, 0, 100, 50)
	ctx.CubicTo(0, 50, 50, 100, 100, 0)
	ctx.ClosePath()

	// 验证没有panic
	assert.True(t, true)
}

func TestLineMethods(t *testing.T) {
	ctx := NewContext(100, 100)

	// 测试线条相关方法
	ctx.SetLineWidth(2)
	ctx.SetLineCap(1)
	ctx.SetLineJoin(1)
	ctx.SetDash(5, 3)
	ctx.SetDashOffset(2)

	// 测试绘制线条
	ctx.DrawLine(0, 0, 100, 100)

	// 验证没有panic
	assert.True(t, true)
}

func TestTextMethods(t *testing.T) {
	ctx := NewContext(100, 100)

	// 测试文本绘制方法
	ctx.DrawString("Hello", 10, 10)
	ctx.DrawStringAnchored("World", 50, 50, 0.5, 0.5)
	ctx.DrawStringWrapped("This is a long text that should be wrapped", 10, 80, 0, 0, 80, 1.2, 1)

	// 验证没有panic
	assert.True(t, true)
}

func TestTransformMethods(t *testing.T) {
	ctx := NewContext(100, 100)

	// 测试变换方法
	ctx.Rotate(0.5)
	ctx.RotateAbout(0.5, 50, 50)
	ctx.Translate(10, 10)
	ctx.Scale(1.5, 1.5)
	ctx.ScaleAbout(1.5, 1.5, 50, 50)
	ctx.Shear(0.2, 0.2)
	ctx.ShearAbout(0.2, 0.2, 50, 50)
	ctx.TransformPoint(10, 10)
	ctx.InvertY()

	// 验证没有panic
	assert.True(t, true)
}

func TestFillAndStroke(t *testing.T) {
	ctx := NewContext(100, 100)

	// 绘制一个形状
	ctx.DrawRectangle(10, 10, 50, 30)

	// 测试填充和描边
	ctx.Fill()
	ctx.ClearPath()

	ctx.DrawRectangle(10, 10, 50, 30)
	ctx.Stroke()
	ctx.ClearPath()

	ctx.DrawRectangle(10, 10, 50, 30)
	ctx.FillStroke()

	// 验证没有panic
	assert.True(t, true)
}

func TestClearMethods(t *testing.T) {
	ctx := NewContext(100, 100)

	// 绘制一些内容
	ctx.SetRGBA255(255, 0, 0, 255)
	ctx.DrawRectangle(0, 0, 100, 100)
	ctx.Fill()

	// 测试清除方法
	ctx.Clear()
	ctx.DrawRectangle(0, 0, 50, 50)
	ctx.ClearPath()

	// 验证没有panic
	assert.True(t, true)
}

func TestSaveAndLoad(t *testing.T) {
	ctx := NewContext(100, 100)

	// 绘制一些内容
	ctx.SetRGBA255(255, 0, 0, 255)
	ctx.DrawRectangle(10, 10, 50, 30)
	ctx.Fill()

	// 测试保存PNG
	tempFile := "/tmp/test_graphics.png"
	err := ctx.SavePNG(tempFile)
	assert.NoError(t, err)

	// 验证文件存在
	_, err = os.Stat(tempFile)
	assert.NoError(t, err)

	// 清理临时文件
	os.Remove(tempFile)
}

func TestGradientCreation(t *testing.T) {
	// 测试渐变创建方法
	linear := CreateLinearGradient(0, 0, 100, 100)
	assert.NotNil(t, linear)

	radial := CreateRadialGradient(50, 50, 25)
	assert.NotNil(t, radial)
}

func TestDefaultContext(t *testing.T) {
	// 测试默认上下文功能
	SetDefaultSize(200, 200)
	ctx := GetDefaultContext()
	assert.NotNil(t, ctx)
	assert.Equal(t, 200, ctx.Width())
	assert.Equal(t, 200, ctx.Height())

	// 测试默认上下文绘制方法
	SetHexColor("#FF0000")
	DrawRectangle(10, 10, 50, 30)
	Fill()

	// 测试保存
	tempFile := "/tmp/test_default_graphics.png"
	err := SavePNG(tempFile)
	assert.NoError(t, err)

	// 验证文件存在
	_, err = os.Stat(tempFile)
	assert.NoError(t, err)

	// 清理临时文件
	os.Remove(tempFile)
}

func TestNewContextFromFile(t *testing.T) {
	// 创建一个测试图像文件
	ctx := NewContext(50, 50)
	ctx.SetRGBA255(0, 255, 0, 255)
	ctx.DrawRectangle(0, 0, 50, 50)
	ctx.Fill()

	tempFile := "/tmp/test_input.png"
	err := ctx.SavePNG(tempFile)
	assert.NoError(t, err)

	// 从文件创建上下文
	newCtx, err := NewContextFromFile(tempFile)
	assert.NoError(t, err)
	assert.NotNil(t, newCtx)
	assert.Equal(t, 50, newCtx.Width())
	assert.Equal(t, 50, newCtx.Height())

	// 清理临时文件
	os.Remove(tempFile)
}
