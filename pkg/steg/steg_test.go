package steg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasic(t *testing.T) {

	tests := []struct {
		name     string
		plain    []byte
		password []byte
		expected []byte
	}{
		{
			name:     "no-password-test",
			plain:    []byte("hello world"),
			password: nil,
			expected: []byte("hello world"),
		},
		{
			name:     "password-test",
			plain:    []byte("hello world"),
			password: []byte("password"),
			expected: []byte("hello world"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encode, err := Encode(tt.plain, tt.password)
			assert.NoError(t, err, "encode error")
			decode, err := Decode(encode, tt.password)
			assert.NoError(t, err, "decode error")
			assert.Equal(t, tt.expected, decode, "decode error")
		})
	}
}
