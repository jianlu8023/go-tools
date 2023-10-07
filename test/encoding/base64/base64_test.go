package base64

import (
	"fmt"
	"testing"

	"github.com/jianlu8023/go-tools/pkg/encoding/base64"
)

// TestBase64
// @Description: TestBase64
// @author ght
// @date 2023-10-07 17:08:51
// @param t:
func TestBase64(t *testing.T) {
	encode := base64.ByteToBase64([]byte("hello world"))
	fmt.Println(encode)
	decode, err := base64.Base64ToByte(encode)
	if err != nil {
		fmt.Println(fmt.Sprintf("err:%s", err))
	}
	fmt.Println(decode)
}
