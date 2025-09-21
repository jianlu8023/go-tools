package base62

import (
	"fmt"
	"testing"
)

func TestBase62(t *testing.T) {
	toBase62 := IntToBase62(100)
	fmt.Println(toBase62)
}
