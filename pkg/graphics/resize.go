package graphics

import (
	"image"

	"github.com/nfnt/resize"
)

// InterpolationFunction 插值函数类型
type InterpolationFunction int

const (
	// NearestNeighbor 最近邻插值
	NearestNeighbor InterpolationFunction = iota
	// Bilinear 双线性插值
	Bilinear
	// Bicubic 双三次插值
	Bicubic
	// MitchellNetravali Mitchell-Netravali插值
	MitchellNetravali
	// Lanczos2 Lanczos插值(a=2)
	Lanczos2
	// Lanczos3 Lanczos插值(a=3)
	Lanczos3
)

// Resize 缩放图像到指定尺寸
// width: 目标宽度，如果为0则按比例计算
// height: 目标高度，如果为0则按比例计算
// img: 原始图像
// interp: 插值函数
func Resize(width, height uint, img image.Image, interp InterpolationFunction) image.Image {
	return resize.Resize(width, height, img, resize.InterpolationFunction(interp))
}

// Thumbnail 生成缩略图，保持宽高比
// maxWidth: 最大宽度
// maxHeight: 最大高度
// img: 原始图像
// interp: 插值函数
func Thumbnail(maxWidth, maxHeight uint, img image.Image, interp InterpolationFunction) image.Image {
	return resize.Thumbnail(maxWidth, maxHeight, img, resize.InterpolationFunction(interp))
}

// Resize 缩放当前上下文中的图像
func (c *Context) Resize(width, height uint, interp InterpolationFunction) *Context {
	resizedImg := Resize(width, height, c.Image(), interp)
	return NewContextFromImage(resizedImg)
}

// Thumbnail 生成当前上下文图像的缩略图
func (c *Context) Thumbnail(maxWidth, maxHeight uint, interp InterpolationFunction) *Context {
	thumbnailImg := Thumbnail(maxWidth, maxHeight, c.Image(), interp)
	return NewContextFromImage(thumbnailImg)
}

// 默认插值函数
var defaultInterpolation = Lanczos3

// SetDefaultInterpolation 设置默认插值函数
func SetDefaultInterpolation(interp InterpolationFunction) {
	defaultInterpolation = interp
}

// GetDefaultInterpolation 获取默认插值函数
func GetDefaultInterpolation() InterpolationFunction {
	return defaultInterpolation
}

// ResizeWithDefault 使用默认插值函数缩放图像
func ResizeWithDefault(width, height uint, img image.Image) image.Image {
	return Resize(width, height, img, defaultInterpolation)
}

// ThumbnailWithDefault 使用默认插值函数生成缩略图
func ThumbnailWithDefault(maxWidth, maxHeight uint, img image.Image) image.Image {
	return Thumbnail(maxWidth, maxHeight, img, defaultInterpolation)
}

// ResizeContextWithDefault 使用默认插值函数缩放当前上下文中的图像
func (c *Context) ResizeWithDefault(width, height uint) *Context {
	return c.Resize(width, height, defaultInterpolation)
}

// ThumbnailContextWithDefault 使用默认插值函数生成当前上下文图像的缩略图
func (c *Context) ThumbnailWithDefault(maxWidth, maxHeight uint) *Context {
	return c.Thumbnail(maxWidth, maxHeight, defaultInterpolation)
}
