package graphics

import (
	"image"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createTestImage(width, height int) image.Image {
	// 创建一个简单的测试图像
	ctx := NewContext(width, height)

	// 设置背景为蓝色
	ctx.SetRGBA255(0, 0, 255, 255)
	ctx.DrawRectangle(0, 0, float64(width), float64(height))
	ctx.Fill()

	// 在中心绘制一个红色圆圈
	ctx.SetRGBA255(255, 0, 0, 255)
	ctx.DrawCircle(float64(width)/2, float64(height)/2, float64(width)/4)
	ctx.Fill()

	return ctx.Image()
}

func TestResize(t *testing.T) {
	// 创建一个测试图像
	img := createTestImage(100, 100)

	// 测试缩放功能
	resized := Resize(50, 50, img, Lanczos3)
	assert.NotNil(t, resized)
	assert.Equal(t, 50, resized.Bounds().Dx())
	assert.Equal(t, 50, resized.Bounds().Dy())
}

func TestResizeWithAspectRatio(t *testing.T) {
	// 创建一个测试图像
	img := createTestImage(100, 50)

	// 测试保持宽高比的缩放（宽度为0）
	resized := Resize(0, 25, img, Bilinear)
	assert.NotNil(t, resized)
	assert.Equal(t, 50, resized.Bounds().Dx()) // 宽度应该按比例缩放
	assert.Equal(t, 25, resized.Bounds().Dy())

	// 测试保持宽高比的缩放（高度为0）
	resized = Resize(200, 0, img, Bilinear)
	assert.NotNil(t, resized)
	assert.Equal(t, 200, resized.Bounds().Dx())
	assert.Equal(t, 100, resized.Bounds().Dy()) // 高度应该按比例缩放
}

func TestThumbnail(t *testing.T) {
	// 创建一个测试图像
	img := createTestImage(100, 50)

	// 测试缩略图功能
	thumbnail := Thumbnail(30, 30, img, MitchellNetravali)
	assert.NotNil(t, thumbnail)
	// 缩略图应该保持宽高比，且不超过指定尺寸
	assert.True(t, thumbnail.Bounds().Dx() <= 30)
	assert.True(t, thumbnail.Bounds().Dy() <= 30)
}

func TestThumbnailWithSmallerImage(t *testing.T) {
	// 创建一个比限制尺寸小的图像
	img := createTestImage(20, 10)

	// 测试缩略图功能（图像比限制尺寸小）
	thumbnail := Thumbnail(30, 30, img, NearestNeighbor)
	assert.NotNil(t, thumbnail)
	// 应该返回原始图像
	assert.Equal(t, img, thumbnail)
}

func TestInterpolationFunctions(t *testing.T) {
	img := createTestImage(50, 50)

	// 测试所有插值函数
	interpolations := []InterpolationFunction{
		NearestNeighbor,
		Bilinear,
		Bicubic,
		MitchellNetravali,
		Lanczos2,
		Lanczos3,
	}

	for _, interp := range interpolations {
		resized := Resize(30, 30, img, interp)
		assert.NotNil(t, resized)
		assert.Equal(t, 30, resized.Bounds().Dx())
		assert.Equal(t, 30, resized.Bounds().Dy())
	}
}

func TestContextResize(t *testing.T) {
	// 创建一个上下文
	ctx := NewContext(100, 100)
	ctx.SetRGBA255(255, 0, 0, 255)
	ctx.DrawRectangle(0, 0, 100, 100)
	ctx.Fill()

	// 测试上下文缩放功能
	resizedCtx := ctx.Resize(50, 50, Lanczos3)
	assert.NotNil(t, resizedCtx)
	assert.Equal(t, 50, resizedCtx.Width())
	assert.Equal(t, 50, resizedCtx.Height())
}

func TestContextThumbnail(t *testing.T) {
	// 创建一个上下文
	ctx := NewContext(100, 50)
	ctx.SetRGBA255(0, 255, 0, 255)
	ctx.DrawRectangle(0, 0, 100, 50)
	ctx.Fill()

	// 测试上下文缩略图功能
	thumbnailCtx := ctx.Thumbnail(30, 30, MitchellNetravali)
	assert.NotNil(t, thumbnailCtx)
	// 缩略图应该保持宽高比，且不超过指定尺寸
	assert.True(t, thumbnailCtx.Width() <= 30)
	assert.True(t, thumbnailCtx.Height() <= 30)
}

func TestDefaultInterpolation(t *testing.T) {
	// 测试默认插值函数设置和获取
	original := GetDefaultInterpolation()

	// 设置新的默认插值函数
	SetDefaultInterpolation(Bilinear)
	assert.Equal(t, Bilinear, GetDefaultInterpolation())

	// 恢复原始值
	SetDefaultInterpolation(original)
	assert.Equal(t, original, GetDefaultInterpolation())
}

func TestResizeWithDefault(t *testing.T) {
	img := createTestImage(100, 100)

	// 测试使用默认插值函数的缩放
	resized := ResizeWithDefault(50, 50, img)
	assert.NotNil(t, resized)
	assert.Equal(t, 50, resized.Bounds().Dx())
	assert.Equal(t, 50, resized.Bounds().Dy())
}

func TestContextResizeWithDefault(t *testing.T) {
	// 创建一个上下文
	ctx := NewContext(100, 100)
	ctx.SetRGBA255(0, 0, 255, 255)
	ctx.DrawRectangle(0, 0, 100, 100)
	ctx.Fill()

	// 测试使用默认插值函数的上下文缩放
	resizedCtx := ctx.ResizeWithDefault(50, 50)
	assert.NotNil(t, resizedCtx)
	assert.Equal(t, 50, resizedCtx.Width())
	assert.Equal(t, 50, resizedCtx.Height())
}
