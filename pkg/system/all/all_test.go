package all

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSystemInformation(t *testing.T) {
	information, err := SystemInformation()
	assert.NoError(t, err, "get system information failed")
	assert.NotEmpty(t, information, "system information is empty")
	fmt.Printf("System Information: %+v\n", information)
}
