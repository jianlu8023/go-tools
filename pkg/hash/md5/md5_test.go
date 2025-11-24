package md5

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	data := []byte("hello world")
	expected := []byte{0x5e, 0xb6, 0x3b, 0xbb, 0xe0, 0x1e, 0xee, 0xd0, 0x93, 0xcb, 0x22, 0xbb, 0x8f, 0x5a, 0xcd, 0xc3}
	result := Sum(data)
	assert.Equal(t, expected, result)
}

func TestSumHex(t *testing.T) {
	data := []byte("hello world")
	expected := "5eb63bbbe01eeed093cb22bb8f5acdc3"
	result := SumHex(data)
	assert.Equal(t, expected, result)
}

func TestSumString(t *testing.T) {
	s := "hello world"
	expected := []byte{0x5e, 0xb6, 0x3b, 0xbb, 0xe0, 0x1e, 0xee, 0xd0, 0x93, 0xcb, 0x22, 0xbb, 0x8f, 0x5a, 0xcd, 0xc3}
	result := SumString(s)
	assert.Equal(t, expected, result)
}

func TestSumStringHex(t *testing.T) {
	s := "hello world"
	expected := "5eb63bbbe01eeed093cb22bb8f5acdc3"
	result := SumStringHex(s)
	assert.Equal(t, expected, result)
}

func TestNew(t *testing.T) {
	h := New()
	_, ok := h.(Digest)
	assert.True(t, ok)

	data := []byte("hello world")
	h.Write(data)
	result := h.Sum(nil)
	expected := []byte{0x5e, 0xb6, 0x3b, 0xbb, 0xe0, 0x1e, 0xee, 0xd0, 0x93, 0xcb, 0x22, 0xbb, 0x8f, 0x5a, 0xcd, 0xc3}
	assert.Equal(t, expected, result)
}

func TestEmptyString(t *testing.T) {
	// MD5 of empty string is d41d8cd98f00b204e9800998ecf8427e
	expected := "d41d8cd98f00b204e9800998ecf8427e"
	result := SumHex([]byte(""))
	assert.Equal(t, expected, result)

	result2 := SumStringHex("")
	assert.Equal(t, expected, result2)
}
