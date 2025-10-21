package sm3

import (
	"encoding/hex"
	"testing"
)

func TestSM3(t *testing.T) {
	bytes, err := HashSm3("hello world")
	if err != nil {
		t.Error(err)
	}

	// 验证结果
	expected := "44f0061e69fa6fdfc290c494654a05dc0c053da7e5c52b84ef93a9d67d3fff88"
	actual := hex.EncodeToString(bytes)

	if actual != expected {
		t.Errorf("Hash mismatch. Expected: %s, Got: %s", expected, actual)
	}
}
