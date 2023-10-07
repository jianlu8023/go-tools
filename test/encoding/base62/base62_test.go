package base62

import (
	"fmt"
	"testing"

	"github.com/jianlu8023/go-tools/pkg/encoding/base62"
)

func TestBase62(t *testing.T) {
	toBase62 := base62.IntToBase62(10)
	fmt.Println(toBase62)
}
