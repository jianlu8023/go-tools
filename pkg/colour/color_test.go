package colour

import (
	"testing"
)

func TestBlue(t *testing.T) {
	t.Log(Blue("hello world"))
}

func TestGrey(t *testing.T) {
	t.Log(Grey("hello world"))
}
