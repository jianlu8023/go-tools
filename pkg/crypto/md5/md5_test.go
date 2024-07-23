package md5

import (
	"fmt"
	"testing"
)

func TestMd5(t *testing.T) {

	str := MD5("hello world")
	fmt.Println(str)
}
