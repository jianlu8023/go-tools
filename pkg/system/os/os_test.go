package os

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSystemOs(t *testing.T) {
	os := SystemOsInfo()
	assert.NotEmpty(t, os, "system os info is empty")
	fmt.Printf("System os: %s\n", os)
}
