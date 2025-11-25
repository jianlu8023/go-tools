package graphics

func ExampleResize() {
	// 创建一个示例图像
	dc := NewContext(200, 100)
	dc.SetRGBA255(255, 0, 0, 255) // 红色背景
	dc.DrawRectangle(0, 0, 200, 100)
	dc.Fill()

	// 缩放图像到100x50
	resizedImg := Resize(100, 50, dc.Image(), Lanczos3)

	// 保存缩放后的图像
	resizedDC := NewContextFromImage(resizedImg)
	resizedDC.SavePNG("resized_example.png")
}

func ExampleThumbnail() {
	// 创建一个示例图像
	dc := NewContext(200, 100)
	dc.SetRGBA255(0, 255, 0, 255) // 绿色背景
	dc.DrawRectangle(0, 0, 200, 100)
	dc.Fill()

	// 生成缩略图，最大尺寸为50x50
	thumbnailImg := Thumbnail(50, 50, dc.Image(), MitchellNetravali)

	// 保存缩略图
	thumbnailDC := NewContextFromImage(thumbnailImg)
	thumbnailDC.SavePNG("thumbnail_example.png")
}

func ExampleContext_Resize() {
	// 创建一个上下文
	dc := NewContext(200, 100)
	dc.SetRGBA255(0, 0, 255, 255) // 蓝色背景
	dc.DrawRectangle(0, 0, 200, 100)
	dc.Fill()

	// 使用上下文方法缩放图像
	resizedDC := dc.Resize(100, 50, Lanczos3)
	resizedDC.SavePNG("context_resize_example.png")
}

func ExampleContext_Thumbnail() {
	// 创建一个上下文
	dc := NewContext(200, 100)
	dc.SetRGBA255(255, 255, 0, 255) // 黄色背景
	dc.DrawRectangle(0, 0, 200, 100)
	dc.Fill()

	// 使用上下文方法生成缩略图
	thumbnailDC := dc.Thumbnail(50, 50, MitchellNetravali)
	thumbnailDC.SavePNG("context_thumbnail_example.png")
}

func ExampleResizeWithDefault() {
	// 创建一个示例图像
	dc := NewContext(200, 100)
	dc.SetRGBA255(255, 0, 255, 255) // 紫色背景
	dc.DrawRectangle(0, 0, 200, 100)
	dc.Fill()

	// 使用默认插值函数缩放图像
	resizedImg := ResizeWithDefault(100, 50, dc.Image())

	// 保存缩放后的图像
	resizedDC := NewContextFromImage(resizedImg)
	resizedDC.SavePNG("default_resize_example.png")
}

func ExampleSetDefaultInterpolation() {
	// 设置默认插值函数
	SetDefaultInterpolation(Lanczos3)

	// 创建示例图像
	dc := NewContext(200, 100)
	dc.SetRGBA255(0, 255, 255, 255) // 青色背景
	dc.DrawRectangle(0, 0, 200, 100)
	dc.Fill()

	// 使用默认插值函数缩放图像
	resizedImg := ResizeWithDefault(100, 50, dc.Image())

	// 保存结果
	resizedDC := NewContextFromImage(resizedImg)
	resizedDC.SavePNG("default_interpolation_example.png")
}
