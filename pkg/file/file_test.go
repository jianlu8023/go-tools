package file

import (
	"testing"
)

func TestWriteToFile(t *testing.T) {
	err := WriteToFile("./test.txt", "Hello, World!", true)
	if err != nil {
		t.Errorf("writeToFile error %v", err)
		return
	}
}
