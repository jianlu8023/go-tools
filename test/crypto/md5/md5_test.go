package md5

import (
	"fmt"
	"testing"

	"github.com/jianlu8023/go-tools/pkg/crypto/md5"
)

func TestMd5(t *testing.T) {

	str := md5.MD5("hello world")
	fmt.Println(str)
}
