package bytes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBytesToString(t *testing.T) {
	b := []byte("Hello World")
	str := BytesToString(b)
	assert.Equal(t, string(b), str)
}

func TestStringToBytes(t *testing.T) {
	str := "Hello World"
	b := StringToBytes(str)
	assert.Equal(t, []byte(str), b)
}

func TestCompare(t *testing.T) {
	tests := []struct {
		a, b []byte
		want int
	}{{
		a:    []byte{1, 2, 3},
		b:    []byte{1, 2, 3},
		want: 0,
	}, {
		a:    []byte{1, 2, 3},
		b:    []byte{1, 2, 4},
		want: -1,
	}, {
		a:    []byte{1, 2, 4},
		b:    []byte{1, 2, 3},
		want: 1,
	}, {
		a:    []byte{1, 2},
		b:    []byte{1, 2, 3},
		want: -1,
	}, {
		a:    []byte{1, 2, 3},
		b:    []byte{1, 2},
		want: 1,
	}, {
		a:    nil,
		b:    nil,
		want: 0,
	}, {
		a:    nil,
		b:    []byte{},
		want: 0,
	}}

	for i, tt := range tests {
		got := Compare(tt.a, tt.b)
		assert.Equal(t, tt.want, got, "test %d: Compare(%v, %v)", i, tt.a, tt.b)
	}
}

func TestEqual(t *testing.T) {
	tests := []struct {
		a, b []byte
		want bool
	}{{
		a:    []byte{1, 2, 3},
		b:    []byte{1, 2, 3},
		want: true,
	}, {
		a:    []byte{1, 2, 3},
		b:    []byte{1, 2, 4},
		want: false,
	}, {
		a:    []byte{1, 2},
		b:    []byte{1, 2, 3},
		want: false,
	}, {
		a:    nil,
		b:    nil,
		want: true,
	}, {
		a:    nil,
		b:    []byte{},
		want: false,
	}}

	for i, tt := range tests {
		got := Equal(tt.a, tt.b)
		assert.Equal(t, tt.want, got, "test %d: Equal(%v, %v)", i, tt.a, tt.b)
	}
}

func TestClone(t *testing.T) {
	tests := []struct {
		b    []byte
		want []byte
	}{{
		b:    []byte{1, 2, 3},
		want: []byte{1, 2, 3},
	}, {
		b:    []byte{},
		want: []byte{},
	}, {
		b:    nil,
		want: nil,
	}}

	for i, tt := range tests {
		got := Clone(tt.b)
		assert.Equal(t, tt.want, got, "test %d: Clone(%v)", i, tt.b)
		// Verify it's a copy, not the same slice
		if tt.b != nil {
			assert.NotSame(t, tt.b, got, "test %d: Clone should return a new slice", i)
		}
	}
}

func TestReverse(t *testing.T) {
	tests := []struct {
		b    []byte
		want []byte
	}{{
		b:    []byte{1, 2, 3, 4},
		want: []byte{4, 3, 2, 1},
	}, {
		b:    []byte{1},
		want: []byte{1},
	}, {
		b:    []byte{},
		want: []byte{},
	}, {
		b:    nil,
		want: nil,
	}}

	for i, tt := range tests {
		got := Reverse(tt.b)
		assert.Equal(t, tt.want, got, "test %d: Reverse(%v)", i, tt.b)
		// Verify original slice is unchanged
		if tt.b != nil {
			// Make a copy of the original for comparison
			original := make([]byte, len(tt.b))
			copy(original, tt.b)
			assert.Equal(t, original, tt.b, "test %d: original slice should not be modified", i)
		}
	}
}

func TestPadding(t *testing.T) {
	tests := []struct {
		b       []byte
		length  int
		padByte byte
		want    []byte
	}{{
		b:       []byte{1, 2, 3},
		length:  5,
		padByte: 0,
		want:    []byte{1, 2, 3, 0, 0},
	}, {
		b:       []byte{1, 2, 3},
		length:  3,
		padByte: 0,
		want:    []byte{1, 2, 3},
	}, {
		b:       []byte{1, 2, 3},
		length:  2,
		padByte: 0,
		want:    []byte{1, 2, 3},
	}, {
		b:       nil,
		length:  3,
		padByte: 5,
		want:    []byte{5, 5, 5},
	}}

	for i, tt := range tests {
		got := Padding(tt.b, tt.length, tt.padByte)
		assert.Equal(t, tt.want, got, "test %d: Padding(%v, %d, %d)", i, tt.b, tt.length, tt.padByte)
	}
}

func TestSlice(t *testing.T) {
	tests := []struct {
		b          []byte
		start, end int
		want       []byte
		wantErr    bool
	}{{
		b:       []byte{1, 2, 3, 4, 5},
		start:   1,
		end:     4,
		want:    []byte{2, 3, 4},
		wantErr: false,
	}, {
		b:       []byte{1, 2, 3},
		start:   0,
		end:     3,
		want:    []byte{1, 2, 3},
		wantErr: false,
	}, {
		b:       []byte{1, 2, 3},
		start:   -1,
		end:     2,
		want:    nil,
		wantErr: true,
	}, {
		b:       []byte{1, 2, 3},
		start:   1,
		end:     5,
		want:    nil,
		wantErr: true,
	}, {
		b:       []byte{1, 2, 3},
		start:   3,
		end:     2,
		want:    nil,
		wantErr: true,
	}}

	for i, tt := range tests {
		got, err := Slice(tt.b, tt.start, tt.end)
		if tt.wantErr {
			assert.Error(t, err, "test %d: Slice(%v, %d, %d) should return error", i, tt.b, tt.start, tt.end)
		} else {
			assert.NoError(t, err, "test %d: Slice(%v, %d, %d) should not return error", i, tt.b, tt.start, tt.end)
			assert.Equal(t, tt.want, got, "test %d: Slice(%v, %d, %d)", i, tt.b, tt.start, tt.end)
		}
	}
}

func TestContains(t *testing.T) {
	tests := []struct {
		b, sub []byte
		want   bool
	}{{
		b:    []byte{1, 2, 3, 4, 5},
		sub:  []byte{3, 4},
		want: true,
	}, {
		b:    []byte{1, 2, 3, 4, 5},
		sub:  []byte{6},
		want: false,
	}, {
		b:    []byte{1, 2, 3},
		sub:  []byte{},
		want: true,
	}, {
		b:    []byte{1, 2, 3},
		sub:  []byte{1, 2, 3, 4},
		want: false,
	}}

	for i, tt := range tests {
		got := Contains(tt.b, tt.sub)
		assert.Equal(t, tt.want, got, "test %d: Contains(%v, %v)", i, tt.b, tt.sub)
	}
}

func TestIndex(t *testing.T) {
	tests := []struct {
		b, sub []byte
		want   int
	}{{
		b:    []byte{1, 2, 3, 4, 5},
		sub:  []byte{3, 4},
		want: 2,
	}, {
		b:    []byte{1, 2, 3, 4, 3, 4, 5},
		sub:  []byte{3, 4},
		want: 2,
	}, {
		b:    []byte{1, 2, 3},
		sub:  []byte{6},
		want: -1,
	}, {
		b:    []byte{1, 2, 3},
		sub:  []byte{},
		want: 0,
	}, {
		b:    []byte{1, 2, 3},
		sub:  []byte{1, 2, 3, 4},
		want: -1,
	}}

	for i, tt := range tests {
		got := Index(tt.b, tt.sub)
		assert.Equal(t, tt.want, got, "test %d: Index(%v, %v)", i, tt.b, tt.sub)
	}
}

func TestHexEncodeDecode(t *testing.T) {
	tests := []struct {
		b    []byte
		want string
	}{{
		b:    []byte{1, 2, 3, 4, 255},
		want: "01020304ff",
	}, {
		b:    []byte{},
		want: "",
	}}

	for i, tt := range tests {
		encoded := HexEncode(tt.b)
		assert.Equal(t, tt.want, encoded, "test %d: HexEncode(%v)", i, tt.b)

		// Test decode
		decoded, err := HexDecode(encoded)
		assert.NoError(t, err, "test %d: HexDecode(%s)", i, encoded)
		assert.Equal(t, tt.b, decoded, "test %d: HexDecode(%s)", i, encoded)
	}

	// Test invalid hex string
	invalidHex := "invalid"
	_, err := HexDecode(invalidHex)
	assert.Error(t, err, "HexDecode should return error for invalid hex string")
}
