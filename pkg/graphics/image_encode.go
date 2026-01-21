package graphics

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/jpeg"
	"image/png"
)

// CreatePNGImage 将RGBA图像编码为PNG格式的字节切片
func CreatePNGImage(img *image.RGBA) ([]byte, error) {
	out := new(bytes.Buffer)
	err := png.Encode(out, img)
	if err != nil {
		return nil, err
	}

	return out.Bytes(), nil
}

// CreateJPGImage 将RGBA图像编码为JPG格式的字节切片
func CreateJPGImage(img *image.RGBA, quality int) ([]byte, error) {
	out := new(bytes.Buffer)
	err := jpeg.Encode(out, img, &jpeg.Options{Quality: quality})
	if err != nil {
		return nil, err
	}

	return out.Bytes(), nil
}

// CreatePNGImageURI 将RGBA图像编码为PNG格式的数据URI
func CreatePNGImageURI(img *image.RGBA) (string, error) {
	data, err := CreatePNGImage(img)
	if err != nil {
		return "", err
	}

	uri := "data:image/png;base64," + base64.StdEncoding.EncodeToString(data)
	return uri, nil
}

// CreateJPGImageURI 将RGBA图像编码为JPG格式的数据URI
func CreateJPGImageURI(img *image.RGBA, quality int) (string, error) {
	data, err := CreateJPGImage(img, quality)
	if err != nil {
		return "", err
	}

	uri := "data:image/jpeg;base64," + base64.StdEncoding.EncodeToString(data)
	return uri, nil
}
