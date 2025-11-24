package graphics

import (
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"io"
	"os"

	"github.com/fogleman/gg"
)

// Context 封装gg.Context，提供图形绘制上下文
type Context struct {
	*gg.Context
}

// NewContext 创建一个新的图形绘制上下文
func NewContext(width, height int) *Context {
	return &Context{Context: gg.NewContext(width, height)}
}

// NewContextFromImage 从现有图像创建图形绘制上下文
func NewContextFromImage(img image.Image) *Context {
	return &Context{Context: gg.NewContextForImage(img)}
}

// NewContextFromReader 从图像读取器创建图形绘制上下文
func NewContextFromReader(r io.Reader) (*Context, error) {
	img, _, err := image.Decode(r)
	if err != nil {
		return nil, err
	}
	return &Context{Context: gg.NewContextForImage(img)}, nil
}

// NewContextFromFile 从图像文件创建图形绘制上下文
func NewContextFromFile(filename string) (*Context, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return &Context{Context: gg.NewContextForImage(img)}, nil
}

// SetLineWidth 设置线条宽度
func (c *Context) SetLineWidth(width float64) {
	c.Context.SetLineWidth(width)
}

// SetLineCap 设置线条端点样式
func (c *Context) SetLineCap(cap gg.LineCap) {
	c.Context.SetLineCap(cap)
}

// SetLineJoin 设置线条连接样式
func (c *Context) SetLineJoin(join gg.LineJoin) {
	c.Context.SetLineJoin(join)
}

// SetDash 设置虚线样式
func (c *Context) SetDash(dashes ...float64) {
	c.Context.SetDash(dashes...)
}

// SetDashOffset 设置虚线偏移
func (c *Context) SetDashOffset(offset float64) {
	c.Context.SetDashOffset(offset)
}

// SetFillStyle 设置填充样式
func (c *Context) SetFillStyle(pattern gg.Pattern) {
	c.Context.SetFillStyle(pattern)
}

// SetStrokeStyle 设置描边样式
func (c *Context) SetStrokeStyle(pattern gg.Pattern) {
	c.Context.SetStrokeStyle(pattern)
}

// SetFillRule 设置填充规则
func (c *Context) SetFillRule(rule gg.FillRule) {
	c.Context.SetFillRule(rule)
}

// SetColor 设置颜色
func (c *Context) SetColor(clr color.Color) {
	c.Context.SetColor(clr)
}

// SetHexColor 设置十六进制颜色
func (c *Context) SetHexColor(x string) {
	c.Context.SetHexColor(x)
}

// SetRGB 设置RGB颜色
func (c *Context) SetRGB(r, g, b float64) {
	c.Context.SetRGB(r, g, b)
}

// SetRGBA 设置RGBA颜色
func (c *Context) SetRGBA(r, g, b, a float64) {
	c.Context.SetRGBA(r, g, b, a)
}

// SetRGB255 设置RGB颜色(0-255)
func (c *Context) SetRGB255(r, g, b int) {
	c.Context.SetRGB255(r, g, b)
}

// SetRGBA255 设置RGBA颜色(0-255)
func (c *Context) SetRGBA255(r, g, b, a int) {
	c.Context.SetRGBA255(r, g, b, a)
}

// DrawRectangle 绘制矩形
func (c *Context) DrawRectangle(x, y, w, h float64) {
	c.Context.DrawRectangle(x, y, w, h)
}

// DrawRoundedRectangle 绘制圆角矩形
func (c *Context) DrawRoundedRectangle(x, y, w, h, r float64) {
	c.Context.DrawRoundedRectangle(x, y, w, h, r)
}

// DrawCircle 绘制圆形
func (c *Context) DrawCircle(x, y, r float64) {
	c.Context.DrawCircle(x, y, r)
}

// DrawArc 绘制弧线
func (c *Context) DrawArc(x, y, r, a1, a2 float64) {
	c.Context.DrawArc(x, y, r, a1, a2)
}

// DrawEllipse 绘制椭圆
func (c *Context) DrawEllipse(x, y, rx, ry float64) {
	c.Context.DrawEllipse(x, y, rx, ry)
}

// DrawRegularPolygon 绘制正多边形
func (c *Context) DrawRegularPolygon(n int, x, y, r, rotation float64) {
	c.Context.DrawRegularPolygon(n, x, y, r, rotation)
}

// DrawImage 绘制图像
func (c *Context) DrawImage(img image.Image, x, y int) {
	c.Context.DrawImage(img, x, y)
}

// DrawImageAnchored 绘制锚定图像
func (c *Context) DrawImageAnchored(img image.Image, x, y int, ax, ay float64) {
	c.Context.DrawImageAnchored(img, x, y, ax, ay)
}

// DrawString 绘制字符串
func (c *Context) DrawString(s string, x, y float64) {
	c.Context.DrawString(s, x, y)
}

// DrawStringAnchored 绘制锚定字符串
func (c *Context) DrawStringAnchored(s string, x, y, ax, ay float64) {
	c.Context.DrawStringAnchored(s, x, y, ax, ay)
}

// DrawStringWrapped 绘制包装字符串
func (c *Context) DrawStringWrapped(s string, x, y, ax, ay, width, lineSpacing float64, align gg.Align) {
	c.Context.DrawStringWrapped(s, x, y, ax, ay, width, lineSpacing, align)
}

// Fill 填充当前路径
func (c *Context) Fill() {
	c.Context.Fill()
}

// Stroke 描边当前路径
func (c *Context) Stroke() {
	c.Context.Stroke()
}

// FillStroke 填充并描边当前路径
func (c *Context) FillStroke() {
	c.Context.Fill()
	c.Context.Stroke()
}

// Clear 清除画布
func (c *Context) Clear() {
	c.Context.Clear()
}

// ClearPath 清除路径
func (c *Context) ClearPath() {
	c.Context.ClearPath()
}

// Save 保存图像到文件(PNG格式)
func (c *Context) SavePNG(filename string) error {
	return c.Context.SavePNG(filename)
}

// EncodePNG 编码为PNG格式
func (c *Context) EncodePNG(w io.Writer) error {
	return png.Encode(w, c.Context.Image())
}

// EncodeJPG 编码为JPG格式
func (c *Context) EncodeJPG(w io.Writer, options *jpeg.Options) error {
	return jpeg.Encode(w, c.Context.Image(), options)
}

// Image 返回当前图像
func (c *Context) Image() image.Image {
	return c.Context.Image()
}

// Width 返回图像宽度
func (c *Context) Width() int {
	return c.Context.Width()
}

// Height 返回图像高度
func (c *Context) Height() int {
	return c.Context.Height()
}

// LoadFontFace 加载字体
func (c *Context) LoadFontFace(path string, points float64) error {
	return c.Context.LoadFontFace(path, points)
}

// DrawLine 绘制线条
func (c *Context) DrawLine(x1, y1, x2, y2 float64) {
	c.Context.DrawLine(x1, y1, x2, y2)
}

// MoveTo 移动到指定点
func (c *Context) MoveTo(x, y float64) {
	c.Context.MoveTo(x, y)
}

// LineTo 画线到指定点
func (c *Context) LineTo(x, y float64) {
	c.Context.LineTo(x, y)
}

// QuadraticTo 绘制二次贝塞尔曲线
func (c *Context) QuadraticTo(x1, y1, x2, y2 float64) {
	c.Context.QuadraticTo(x1, y1, x2, y2)
}

// CubicTo 绘制三次贝塞尔曲线
func (c *Context) CubicTo(x1, y1, x2, y2, x3, y3 float64) {
	c.Context.CubicTo(x1, y1, x2, y2, x3, y3)
}

// ClosePath 闭合路径
func (c *Context) ClosePath() {
	c.Context.ClosePath()
}

// Rotate 旋转画布
func (c *Context) Rotate(angle float64) {
	c.Context.Rotate(angle)
}

// RotateAbout 围绕指定点旋转
func (c *Context) RotateAbout(angle, x, y float64) {
	c.Context.RotateAbout(angle, x, y)
}

// Translate 平移画布
func (c *Context) Translate(x, y float64) {
	c.Context.Translate(x, y)
}

// Scale 缩放画布
func (c *Context) Scale(x, y float64) {
	c.Context.Scale(x, y)
}

// ScaleAbout 围绕指定点缩放
func (c *Context) ScaleAbout(sx, sy, x, y float64) {
	c.Context.ScaleAbout(sx, sy, x, y)
}

// Shear 剪切画布
func (c *Context) Shear(x, y float64) {
	c.Context.Shear(x, y)
}

// ShearAbout 围绕指定点剪切
func (c *Context) ShearAbout(sx, sy, x, y float64) {
	c.Context.ShearAbout(sx, sy, x, y)
}

// TransformPoint 变换点坐标
func (c *Context) TransformPoint(x, y float64) (float64, float64) {
	return c.Context.TransformPoint(x, y)
}

// InvertY 翻转Y轴
func (c *Context) InvertY() {
	c.Context.InvertY()
}

// CreateLinearGradient 创建线性渐变
func CreateLinearGradient(x1, y1, x2, y2 float64) gg.Gradient {
	return gg.NewLinearGradient(x1, y1, x2, y2)
}

// CreateRadialGradient 创建径向渐变
func CreateRadialGradient(x, y, r float64) gg.Gradient {
	return gg.NewRadialGradient(x, y, x, y, r, r)
}

// 默认客户端函数
var defaultContext *Context

// SetDefaultSize 设置默认画布大小
func SetDefaultSize(width, height int) {
	defaultContext = NewContext(width, height)
}

// GetDefaultContext 获取默认上下文
func GetDefaultContext() *Context {
	if defaultContext == nil {
		SetDefaultSize(500, 500)
	}
	return defaultContext
}

// DrawRectangle 使用默认上下文绘制矩形
func DrawRectangle(x, y, w, h float64) {
	GetDefaultContext().DrawRectangle(x, y, w, h)
}

// DrawCircle 使用默认上下文绘制圆形
func DrawCircle(x, y, r float64) {
	GetDefaultContext().DrawCircle(x, y, r)
}

// SetHexColor 使用默认上下文设置十六进制颜色
func SetHexColor(x string) {
	GetDefaultContext().SetHexColor(x)
}

// Fill 使用默认上下文填充
func Fill() {
	GetDefaultContext().Fill()
}

// SavePNG 使用默认上下文保存PNG文件
func SavePNG(filename string) error {
	return GetDefaultContext().SavePNG(filename)
}
