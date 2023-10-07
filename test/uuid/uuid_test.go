package uuid

import (
	"fmt"
	"testing"

	"github.com/jianlu8023/go-tools/pkg/uuid"
)

// TestUUID
// @Description: TestUUID
// @author ght
// @date 2023-10-07 17:19:36
// @param t:
func TestUUID(t *testing.T) {
	genUUID := uuid.GenUUID()
	fmt.Println(genUUID)
}
