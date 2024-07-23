package file

import (
	"testing"
)

func TestWriteToFile(t *testing.T) {

	err := WriteToFile("/root/go/src/github.com/jianlu8023/go-tools/testdata/test.txt", "Hello, World!")
	if err != nil {
		t.Errorf("writeToFile error %v", err)
		return
	}
}
