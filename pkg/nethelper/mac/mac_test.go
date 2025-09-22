package mac

import (
	"fmt"
	"testing"
)

func TestGetMacAddr(t *testing.T) {
	addr := GetMacAddr()
	fmt.Println(addr)
}
