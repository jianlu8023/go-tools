package graphics

import (
	"image/color"
)

func Example() {
	// 创建一个400x400的画布
	dc := NewContext(400, 400)

	// 设置背景色为白色
	dc.SetColor(color.White)
	dc.DrawRectangle(0, 0, 400, 400)
	dc.Fill()

	// 绘制一个红色的圆
	dc.SetRGBA255(255, 0, 0, 255)
	dc.DrawCircle(200, 200, 100)
	dc.Fill()

	// 绘制一个蓝色的矩形
	dc.SetRGBA255(0, 0, 255, 255)
	dc.DrawRectangle(100, 100, 200, 100)
	dc.Stroke()

	// 保存为PNG文件
	dc.SavePNG("/tmp/example1.png")
}

func ExampleContext_DrawRectangle() {
	// 创建一个300x300的画布
	dc := NewContext(300, 300)

	// 设置背景色
	dc.SetRGB(0.8, 0.8, 1)
	dc.DrawRectangle(0, 0, 300, 300)
	dc.Fill()

	// 绘制绿色圆角矩形
	dc.SetRGBA255(0, 255, 0, 255)
	dc.DrawRoundedRectangle(50, 50, 100, 80, 10)
	dc.FillStroke()

	// 保存为PNG文件
	dc.SavePNG("/tmp/example2.png")
}

func ExampleContext_DrawString() {
	// 创建一个400x200的画布
	dc := NewContext(400, 200)

	// 设置背景色
	dc.SetRGBA255(240, 240, 240, 255)
	dc.DrawRectangle(0, 0, 400, 200)
	dc.Fill()

	// 绘制文本
	dc.SetRGBA255(0, 0, 0, 255)

	// 简单文本
	dc.DrawString("Hello, World!", 50, 50)

	// 居中文本
	dc.DrawStringAnchored("Centered Text", 200, 100, 0.5, 0.5)

	// 保存为PNG文件
	dc.SavePNG("/tmp/example3.png")
}

func ExampleContext_DrawImage() {
	// 创建一个画布
	dc := NewContext(300, 300)

	// 设置背景色
	dc.SetRGBA255(255, 255, 200, 255)
	dc.DrawRectangle(0, 0, 300, 300)
	dc.Fill()

	// 从默认上下文创建一个简单图像
	SetDefaultSize(100, 100)
	SetHexColor("#FF0000")
	DrawRectangle(10, 10, 80, 80)
	Fill()

	// 获取图像并绘制到主画布上
	img := GetDefaultContext().Image()
	dc.DrawImage(img, 100, 100)

	// 保存为PNG文件
	dc.SavePNG("/tmp/example4.png")
}

func ExampleCreateLinearGradient() {
	// 创建一个400x400的画布
	dc := NewContext(400, 400)

	// 创建线性渐变
	linear := CreateLinearGradient(0, 0, 400, 400)
	// 注意：gg库的渐变使用方式可能需要根据具体版本调整
	// 这里仅展示API调用方式
	dc.SetFillStyle(linear)

	// 绘制填充矩形
	dc.DrawRectangle(0, 0, 400, 400)
	dc.Fill()

	// 保存为PNG文件
	dc.SavePNG("/tmp/example5.png")
}

func ExampleContext_Translate() {
	// 创建一个300x300的画布
	dc := NewContext(300, 300)

	// 设置背景色
	dc.SetRGBA255(240, 240, 240, 255)
	dc.DrawRectangle(0, 0, 300, 300)
	dc.Fill()

	// 平移并旋转
	dc.Translate(150, 150)
	dc.Rotate(0.5)

	// 绘制红色矩形
	dc.SetRGBA255(255, 0, 0, 255)
	dc.DrawRectangle(-25, -25, 50, 50)
	dc.Fill()

	// 保存为PNG文件
	dc.SavePNG("/tmp/example6.png")
}

func ExampleNewContextFromImage() {
	// 首先创建一个图像
	dc := NewContext(100, 100)
	dc.SetRGBA255(255, 0, 0, 255)
	dc.DrawCircle(50, 50, 40)
	dc.Fill()

	// 从现有图像创建新的上下文
	newDC := NewContextFromImage(dc.Image())

	// 在现有图像上添加内容
	newDC.SetRGBA255(0, 255, 0, 128)
	newDC.DrawRectangle(25, 25, 50, 50)
	newDC.Fill()

	// 保存修改后的图像
	newDC.SavePNG("/tmp/modified.png")
}
