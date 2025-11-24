package base64x

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToBase64(t *testing.T) {
	data := []byte("hello world")
	expected := "aGVsbG8gd29ybGQ="
	result := ToBase64(data)
	assert.Equal(t, expected, result)
}

func TestToByte(t *testing.T) {
	data := "aGVsbG8gd29ybGQ="
	expected := []byte("hello world")
	result, err := ToByte(data)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestToBase64URL(t *testing.T) {
	data := []byte("hello world?")
	expected := "aGVsbG8gd29ybGQ_"
	result := ToBase64URL(data)
	assert.Equal(t, expected, result)
}

func TestToByteURL(t *testing.T) {
	data := "aGVsbG8gd29ybGQ_"
	expected := []byte("hello world?")
	result, err := ToByteURL(data)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestToBase64NoPadding(t *testing.T) {
	data := []byte("hello")
	expected := "aGVsbG8"
	result := ToBase64NoPadding(data)
	assert.Equal(t, expected, result)
}

func TestToBase64URLNoPadding(t *testing.T) {
	data := []byte("hello")
	expected := "aGVsbG8"
	result := ToBase64URLNoPadding(data)
	assert.Equal(t, expected, result)
}
