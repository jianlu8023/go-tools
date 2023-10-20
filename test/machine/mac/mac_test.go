package mac

import (
	"fmt"
	"testing"

	"github.com/jianlu8023/go-tools/pkg/machine/mac"
)

func TestGetMacAddr(t *testing.T) {
	addr := mac.GetMacAddr()
	fmt.Println(addr)
}
