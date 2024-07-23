package uuid

import (
	"fmt"
	"testing"
)

func TestUUID(t *testing.T) {
	genUUID := GenUUID()
	fmt.Println(genUUID)
}
