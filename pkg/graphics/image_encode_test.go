package graphics

import (
	"image"
	"image/color"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatePNGImage(t *testing.T) {
	// 创建一个简单的RGBA图像用于测试
	img := image.NewRGBA(image.Rect(0, 0, 10, 10))
	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			img.Set(x, y, color.RGBA{255, 0, 0, 255}) // 红色
		}
	}

	data, err := CreatePNGImage(img)
	assert.NoError(t, err)
	assert.NotNil(t, data)
	assert.Greater(t, len(data), 0)
}

func TestCreateJPGImage(t *testing.T) {
	// 创建一个简单的RGBA图像用于测试
	img := image.NewRGBA(image.Rect(0, 0, 10, 10))
	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			img.Set(x, y, color.RGBA{0, 255, 0, 255}) // 绿色
		}
	}

	data, err := CreateJPGImage(img, 80)
	assert.NoError(t, err)
	assert.NotNil(t, data)
	assert.Greater(t, len(data), 0)
}

func TestCreatePNGImageURI(t *testing.T) {
	// 创建一个简单的RGBA图像用于测试
	img := image.NewRGBA(image.Rect(0, 0, 5, 5))
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			img.Set(x, y, color.RGBA{0, 0, 255, 255}) // 蓝色
		}
	}

	uri, err := CreatePNGImageURI(img)
	assert.NoError(t, err)
	assert.NotEmpty(t, uri)
	assert.Contains(t, uri, "data:image/png;base64,")
}

func TestCreateJPGImageURI(t *testing.T) {
	// 创建一个简单的RGBA图像用于测试
	img := image.NewRGBA(image.Rect(0, 0, 5, 5))
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			img.Set(x, y, color.RGBA{255, 255, 0, 255}) // 青色
		}
	}

	uri, err := CreateJPGImageURI(img, 90)
	assert.NoError(t, err)
	assert.NotEmpty(t, uri)
	assert.Contains(t, uri, "data:image/jpeg;base64,")
}
