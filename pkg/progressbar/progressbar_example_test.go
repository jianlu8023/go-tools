package progressbar

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"time"
)

func ExampleNew() {
	// 创建一个简单的进度条
	bar := New(100)

	// 模拟一些工作
	for i := 0; i < 100; i++ {
		bar.Add(1)
		// 模拟工作延迟
		time.Sleep(10 * time.Millisecond)
	}

	fmt.Println("进度条完成!")
	// Output: 进度条完成!
}

func ExampleDefault() {
	// 创建一个默认进度条
	bar := Default(100, "处理中")

	// 模拟一些工作
	for i := 0; i < 100; i++ {
		bar.Add(1)
		time.Sleep(5 * time.Millisecond)
	}

	// Output:
}

func ExampleDefaultBytes() {
	// 创建一个字节进度条
	bar := DefaultBytes(1024*1024, "下载文件")

	// 模拟下载过程
	for i := 0; i < 100; i++ {
		bar.Add(1024 * 10) // 每次增加10KB
		time.Sleep(50 * time.Millisecond)
	}

	// Output:
}

func ExampleNewOptions() {
	// 创建一个自定义进度条
	bar := NewOptions(100,
		OptionSetWidth(20),
		OptionSetDescription("自定义进度条"),
		OptionShowCount(),
		OptionSetTheme(Theme{
			Saucer:        "=",
			SaucerHead:    ">",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}),
	)

	// 模拟一些工作
	for i := 0; i < 100; i++ {
		bar.Add(1)
		time.Sleep(10 * time.Millisecond)
	}

	// Output:
}

func ExampleProgressBar_Write() {
	// 使用Write方法自动跟踪进度
	bar := DefaultBytes(100, "写入数据")

	// 创建一些数据
	data := bytes.Repeat([]byte("data"), 25) // 100字节数据

	// 使用Write方法写入数据，进度条会自动更新
	_, err := bar.Write(data)
	if err != nil {
		fmt.Printf("写入错误: %v\n", err)
		return
	}

	// Output:
}

func ExampleNewReader() {
	// 创建一个Reader来跟踪读取进度
	data := strings.Repeat("example data ", 100)
	reader := strings.NewReader(data)

	// 创建进度条
	bar := Default(int64(len(data)), "读取数据")

	// 创建包装的Reader
	wrappedReader := NewReader(reader, bar)

	// 读取数据
	buffer := make([]byte, len(data))
	_, err := io.ReadFull(&wrappedReader.Reader, buffer)
	if err != nil {
		fmt.Printf("读取错误: %v\n", err)
		return
	}

	// Output:
}

func ExampleOptionShowIts() {
	// 创建显示迭代速度的进度条
	bar := NewOptions(100,
		OptionSetDescription("处理数据"),
		OptionShowIts(),
	)

	// 模拟一些工作
	for i := 0; i < 100; i++ {
		bar.Add(1)
		time.Sleep(50 * time.Millisecond)
	}

	// Output:
}

func ExampleOptionThrottle() {
	// 创建节流的进度条（限制更新频率）
	bar := NewOptions(100,
		OptionSetDescription("节流进度条"),
		OptionThrottle(100*time.Millisecond), // 每100ms最多更新一次
	)

	// 快速增加进度，但由于节流，实际更新会较少
	for i := 0; i < 100; i++ {
		bar.Add(1)
		// 没有延迟，会触发节流
	}

	// Output:
}

func ExampleProgressBar_ChangeMax() {
	// 创建进度条并动态更改最大值
	bar := New(50)
	bar.Describe("动态调整")

	// 处理前50个单位
	for i := 0; i < 50; i++ {
		bar.Add(1)
		time.Sleep(20 * time.Millisecond)
	}

	// 更改最大值为100
	bar.ChangeMax(100)

	// 处理剩余的50个单位
	for i := 0; i < 50; i++ {
		bar.Add(1)
		time.Sleep(20 * time.Millisecond)
	}

	// Output:
}

func ExampleOptionSpinnerType() {
	// 创建旋转动画进度条（未知长度）
	bar := NewOptions(-1, // -1表示未知长度，会自动使用旋转动画
		OptionSetDescription("处理中"),
		OptionSpinnerType(9), // 使用特定的旋转动画类型
	)

	// 模拟一些工作
	for i := 0; i < 100; i++ {
		bar.Add(1)
		time.Sleep(50 * time.Millisecond)
	}

	bar.Finish()
	// Output:
}

func ExampleOptionClearOnFinish() {
	// 创建完成时清除的进度条
	bar := NewOptions(100,
		OptionSetDescription("完成后清除"),
		OptionClearOnFinish(), // 完成后清除进度条显示
	)

	// 模拟一些工作
	for i := 0; i < 100; i++ {
		bar.Add(1)
		time.Sleep(10 * time.Millisecond)
	}

	// Output:
}

func ExampleOptionOnCompletion() {
	// 创建带完成回调的进度条
	bar := NewOptions(100,
		OptionSetDescription("带回调"),
		OptionOnCompletion(func() {
			fmt.Println("进度条完成!")
		}),
	)

	// 模拟一些工作
	for i := 0; i < 100; i++ {
		bar.Add(1)
		time.Sleep(10 * time.Millisecond)
	}

	// Output: 进度条完成!
}

func ExampleProgressBar_AddDetail() {
	// 创建带详细信息的进度条
	bar := NewOptions(100,
		OptionSetDescription("详细信息"),
		OptionSetMaxDetailRow(3), // 最多显示3行详细信息
	)

	// 模拟一些工作并添加详细信息
	for i := 0; i < 100; i++ {
		bar.Add(1)

		// 每20步添加一条详细信息
		if i%20 == 0 {
			bar.AddDetail(fmt.Sprintf("步骤 %d 完成", i))
		}

		time.Sleep(20 * time.Millisecond)
	}

	// Output:
}
