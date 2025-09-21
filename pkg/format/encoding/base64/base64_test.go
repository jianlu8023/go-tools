package base64

import (
	"fmt"
	"testing"
)

func TestBase64(t *testing.T) {
	base64 := ToBase64([]byte("hello world"))
	fmt.Println(base64)
}

func TestBase64Decode(t *testing.T) {
	toByte, err := ToByte("aGVsbG8gd29ybGQ=")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println(string(toByte)) // hello world

}
