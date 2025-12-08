package sqlnull

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestStringPtr(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "non-empty string",
			input:    "hello",
			expected: "hello",
		},
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := StringPtr(tt.input)
			assert.NotNil(t, result)
			assert.Equal(t, tt.expected, *result)
		})
	}
}

func TestStringDefaultPtr(t *testing.T) {
	result := StringDefaultPtr()
	assert.NotNil(t, result)
	assert.Equal(t, "", *result)
}

func TestInt64Ptr(t *testing.T) {
	tests := []struct {
		name     string
		input    int64
		expected int64
	}{
		{
			name:     "positive number",
			input:    42,
			expected: 42,
		},
		{
			name:     "negative number",
			input:    -10,
			expected: -10,
		},
		{
			name:     "zero",
			input:    0,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Int64Ptr(tt.input)
			assert.NotNil(t, result)
			assert.Equal(t, tt.expected, *result)
		})
	}
}

func TestInt64DefaultPtr(t *testing.T) {
	result := Int64DefaultPtr()
	assert.NotNil(t, result)
	assert.Equal(t, int64(0), *result)
}

func TestTruePtr(t *testing.T) {
	result := TruePtr()
	assert.NotNil(t, result)
	assert.True(t, *result)
}

func TestFalsePtr(t *testing.T) {
	result := FalsePtr()
	assert.NotNil(t, result)
	assert.False(t, *result)
}

func TestBoolPtr(t *testing.T) {
	tests := []struct {
		name     string
		input    bool
		expected bool
	}{
		{
			name:     "true value",
			input:    true,
			expected: true,
		},
		{
			name:     "false value",
			input:    false,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := BoolPtr(tt.input)
			assert.NotNil(t, result)
			assert.Equal(t, tt.expected, *result)
		})
	}
}

func TestBoolDefaultPtr(t *testing.T) {
	result := BoolDefaultPtr()
	assert.NotNil(t, result)
	assert.False(t, *result)
}

func TestFloat64Ptr(t *testing.T) {
	tests := []struct {
		name     string
		input    float64
		expected float64
	}{
		{
			name:     "positive float",
			input:    3.14,
			expected: 3.14,
		},
		{
			name:     "negative float",
			input:    -2.5,
			expected: -2.5,
		},
		{
			name:     "zero float",
			input:    0.0,
			expected: 0.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Float64Ptr(tt.input)
			assert.NotNil(t, result)
			assert.Equal(t, tt.expected, *result)
		})
	}
}

func TestFloat64DefaultPtr(t *testing.T) {
	result := Float64DefaultPtr()
	assert.NotNil(t, result)
	assert.Equal(t, 0.0, *result)
}

func TestInt32Ptr(t *testing.T) {
	tests := []struct {
		name     string
		input    int32
		expected int32
	}{
		{
			name:     "positive int32",
			input:    42,
			expected: 42,
		},
		{
			name:     "negative int32",
			input:    -10,
			expected: -10,
		},
		{
			name:     "zero int32",
			input:    0,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Int32Ptr(tt.input)
			assert.NotNil(t, result)
			assert.Equal(t, tt.expected, *result)
		})
	}
}

func TestInt32DefaultPtr(t *testing.T) {
	result := Int32DefaultPtr()
	assert.NotNil(t, result)
	assert.Equal(t, int32(0), *result)
}

func TestInt16Ptr(t *testing.T) {
	tests := []struct {
		name     string
		input    int16
		expected int16
	}{
		{
			name:     "positive int16",
			input:    42,
			expected: 42,
		},
		{
			name:     "negative int16",
			input:    -10,
			expected: -10,
		},
		{
			name:     "zero int16",
			input:    0,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Int16Ptr(tt.input)
			assert.NotNil(t, result)
			assert.Equal(t, tt.expected, *result)
		})
	}
}

func TestInt16DefaultPtr(t *testing.T) {
	result := Int16DefaultPtr()
	assert.NotNil(t, result)
	assert.Equal(t, int16(0), *result)
}

func TestTimePtr(t *testing.T) {
	now := time.Now()
	result := TimePtr(now)
	assert.NotNil(t, result)
	assert.Equal(t, now, *result)
}

func TestTimeDefaultPtr(t *testing.T) {
	result := TimeDefaultPtr()
	assert.NotNil(t, result)
	assert.Equal(t, time.Time{}, *result)
}

func TestIntPtr(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{
			name:     "positive int",
			input:    42,
			expected: 42,
		},
		{
			name:     "negative int",
			input:    -10,
			expected: -10,
		},
		{
			name:     "zero int",
			input:    0,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IntPtr(tt.input)
			assert.NotNil(t, result)
			assert.Equal(t, tt.expected, *result)
		})
	}
}

func TestIntDefaultPtr(t *testing.T) {
	result := IntDefaultPtr()
	assert.NotNil(t, result)
	assert.Equal(t, 0, *result)
}

func TestInt8Ptr(t *testing.T) {
	tests := []struct {
		name     string
		input    int8
		expected int8
	}{
		{
			name:     "positive int8",
			input:    42,
			expected: 42,
		},
		{
			name:     "negative int8",
			input:    -10,
			expected: -10,
		},
		{
			name:     "zero int8",
			input:    0,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Int8Ptr(tt.input)
			assert.NotNil(t, result)
			assert.Equal(t, tt.expected, *result)
		})
	}
}

func TestInt8DefaultPtr(t *testing.T) {
	result := Int8DefaultPtr()
	assert.NotNil(t, result)
	assert.Equal(t, int8(0), *result)
}

func TestUintPtr(t *testing.T) {
	tests := []struct {
		name     string
		input    uint
		expected uint
	}{
		{
			name:     "positive uint",
			input:    42,
			expected: 42,
		},
		{
			name:     "zero uint",
			input:    0,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := UintPtr(tt.input)
			assert.NotNil(t, result)
			assert.Equal(t, tt.expected, *result)
		})
	}
}

func TestUintDefaultPtr(t *testing.T) {
	result := UintDefaultPtr()
	assert.NotNil(t, result)
	assert.Equal(t, uint(0), *result)
}

func TestUint8Ptr(t *testing.T) {
	tests := []struct {
		name     string
		input    uint8
		expected uint8
	}{
		{
			name:     "positive uint8",
			input:    42,
			expected: 42,
		},
		{
			name:     "zero uint8",
			input:    0,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Uint8Ptr(tt.input)
			assert.NotNil(t, result)
			assert.Equal(t, tt.expected, *result)
		})
	}
}

func TestUint8DefaultPtr(t *testing.T) {
	result := Uint8DefaultPtr()
	assert.NotNil(t, result)
	assert.Equal(t, uint8(0), *result)
}

func TestUint16Ptr(t *testing.T) {
	tests := []struct {
		name     string
		input    uint16
		expected uint16
	}{
		{
			name:     "positive uint16",
			input:    42,
			expected: 42,
		},
		{
			name:     "zero uint16",
			input:    0,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Uint16Ptr(tt.input)
			assert.NotNil(t, result)
			assert.Equal(t, tt.expected, *result)
		})
	}
}

func TestUint16DefaultPtr(t *testing.T) {
	result := Uint16DefaultPtr()
	assert.NotNil(t, result)
	assert.Equal(t, uint16(0), *result)
}

func TestUint32Ptr(t *testing.T) {
	tests := []struct {
		name     string
		input    uint32
		expected uint32
	}{
		{
			name:     "positive uint32",
			input:    42,
			expected: 42,
		},
		{
			name:     "zero uint32",
			input:    0,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Uint32Ptr(tt.input)
			assert.NotNil(t, result)
			assert.Equal(t, tt.expected, *result)
		})
	}
}

func TestUint32DefaultPtr(t *testing.T) {
	result := Uint32DefaultPtr()
	assert.NotNil(t, result)
	assert.Equal(t, uint32(0), *result)
}

func TestUint64Ptr(t *testing.T) {
	tests := []struct {
		name     string
		input    uint64
		expected uint64
	}{
		{
			name:     "positive uint64",
			input:    42,
			expected: 42,
		},
		{
			name:     "zero uint64",
			input:    0,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Uint64Ptr(tt.input)
			assert.NotNil(t, result)
			assert.Equal(t, tt.expected, *result)
		})
	}
}

func TestUint64DefaultPtr(t *testing.T) {
	result := Uint64DefaultPtr()
	assert.NotNil(t, result)
	assert.Equal(t, uint64(0), *result)
}

func TestFloat32Ptr(t *testing.T) {
	tests := []struct {
		name     string
		input    float32
		expected float32
	}{
		{
			name:     "positive float32",
			input:    3.14,
			expected: 3.14,
		},
		{
			name:     "negative float32",
			input:    -2.5,
			expected: -2.5,
		},
		{
			name:     "zero float32",
			input:    0.0,
			expected: 0.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Float32Ptr(tt.input)
			assert.NotNil(t, result)
			assert.Equal(t, tt.expected, *result)
		})
	}
}

func TestFloat32DefaultPtr(t *testing.T) {
	result := Float32DefaultPtr()
	assert.NotNil(t, result)
	assert.Equal(t, float32(0.0), *result)
}
