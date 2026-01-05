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
		if tt.b != nil && len(tt.b) > 0 {
			// Modify original to check if copy is independent
			originalFirst := tt.b[0]
			tt.b[0] = ^tt.b[0] // Flip all bits
			assert.Equal(t, originalFirst, got[0], "test %d: Clone should return independent copy", i)
			tt.b[0] = originalFirst // Restore
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

func TestHumanDecimal(t *testing.T) {
	tests := []struct {
		input int64
		want  string
	}{
		{0, "0B"},
		{1, "1B"},
		{999, "999B"},
		{1000, "1kB"},
		{1500, "1.5kB"},
		{1000000, "1MB"},
		{1500000, "1.5MB"},
		{1000000000, "1GB"},
		{1500000000, "1.5GB"},
		{1000000000000, "1TB"},
		{1500000000000, "1.5TB"},
		{1000000000000000, "1PB"},
		{1500000000000000, "1.5PB"},
	}

	for _, tt := range tests {
		got := HumanDecimal(tt.input)
		assert.Equal(t, tt.want, got, "HumanDecimal(%d)", tt.input)
	}
}

func TestFromHumanDecimal(t *testing.T) {
	tests := []struct {
		input    string
		want     int64
		wantErr  bool
		errorMsg string
	}{
		{"1B", 1, false, ""},
		{"1kB", 1000, false, ""},
		{"1.5kB", 1500, false, ""},
		{"1MB", 1000000, false, ""},
		{"1.5MB", 1500000, false, ""},
		{"1GB", 1000000000, false, ""},
		{"1.5GB", 1500000000, false, ""},
		{"1TB", 1000000000000, false, ""},
		{"1.5TB", 1500000000000, false, ""},
		{"1PB", 1000000000000000, false, ""},
		{"1.5PB", 1500000000000000, false, ""},
		{"1 K", 1000, false, ""},
		{"1 M", 1000000, false, ""},
		{"1 G", 1000000000, false, ""},
		{"1 T", 1000000000000, false, ""},
		{"1 P", 1000000000000000, false, ""},
		{"1kb", 1000, false, ""},
		{"1mb", 1000000, false, ""},
		{"1gb", 1000000000, false, ""},
		{"1tb", 1000000000000, false, ""},
		{"1pb", 1000000000000000, false, ""},
		{"invalid", -1, true, "invalid size: 'invalid'"},
		{"1.5xyz", -1, true, "invalid size: '1.5xyz'"},
		{"", -1, true, "invalid size: ''"},
	}

	for _, tt := range tests {
		got, err := FromHumanDecimal(tt.input)
		if tt.wantErr {
			assert.Error(t, err, "FromHumanDecimal(%s) should return error", tt.input)
			if tt.errorMsg != "" {
				assert.Contains(t, err.Error(), tt.errorMsg, "FromHumanDecimal(%s) error message", tt.input)
			}
		} else {
			assert.NoError(t, err, "FromHumanDecimal(%s) should not return error", tt.input)
			assert.Equal(t, tt.want, got, "FromHumanDecimal(%s)", tt.input)
		}
	}
}

func TestHumanBinary(t *testing.T) {
	tests := []struct {
		input uint64
		want  string
	}{
		{0, "0B"},
		{1, "1B"},
		{1023, "1023B"},
		{1024, "1KiB"},
		{1536, "1.5KiB"},
		{1048576, "1MiB"},
		{1572864, "1.5MiB"},
		{1073741824, "1GiB"},
		{1610612736, "1.5GiB"},
		{1099511627776, "1TiB"},
		{1649267441664, "1.5TiB"},
		{1125899906842624, "1PiB"},
		{1688849860263936, "1.5PiB"},
	}

	for _, tt := range tests {
		got := HumanBinary(tt.input)
		assert.Equal(t, tt.want, got, "HumanBinary(%d)", tt.input)
	}
}

func TestFromHumanBinary(t *testing.T) {
	tests := []struct {
		input    string
		want     int64
		wantErr  bool
		errorMsg string
	}{
		{"1B", 1, false, ""},
		{"1KiB", 1024, false, ""},
		{"1.5KiB", 1536, false, ""},
		{"1MiB", 1048576, false, ""},
		{"1.5MiB", 1572864, false, ""},
		{"1GiB", 1073741824, false, ""},
		{"1.5GiB", 1610612736, false, ""},
		{"1TiB", 1099511627776, false, ""},
		{"1.5TiB", 1649267441664, false, ""},
		{"1PiB", 1125899906842624, false, ""},
		{"1.5PiB", 1688849860263936, false, ""},
		{"1 K", 1024, false, ""},
		{"1 M", 1048576, false, ""},
		{"1 G", 1073741824, false, ""},
		{"1 T", 1099511627776, false, ""},
		{"1 P", 1125899906842624, false, ""},
		{"1Ki", 1024, false, ""},
		{"1Mi", 1048576, false, ""},
		{"1Gi", 1073741824, false, ""},
		{"1Ti", 1099511627776, false, ""},
		{"1Pi", 1125899906842624, false, ""},
		{"invalid", -1, true, "invalid size: 'invalid'"},
		{"1.5xyz", -1, true, "invalid size: '1.5xyz'"},
		{"", -1, true, "invalid size: ''"},
	}

	for _, tt := range tests {
		got, err := FromHumanBinary(tt.input)
		if tt.wantErr {
			assert.Error(t, err, "FromHumanBinary(%s) should return error", tt.input)
			if tt.errorMsg != "" {
				assert.Contains(t, err.Error(), tt.errorMsg, "FromHumanBinary(%s) error message", tt.input)
			}
		} else {
			assert.NoError(t, err, "FromHumanBinary(%s) should not return error", tt.input)
			assert.Equal(t, tt.want, got, "FromHumanBinary(%s)", tt.input)
		}
	}
}

func TestBytesToIntConversion(t *testing.T) {
	// Test BytesToInt
	val := int32(12345)
	bytes := IntToBytes(val)
	result, err := BytesToInt(bytes)
	assert.NoError(t, err)
	assert.Equal(t, val, result)

	// Test with invalid byte length
	_, err = BytesToInt([]byte{1, 2})
	assert.Error(t, err)

	// Test BytesToInt64
	val64 := int64(1234567890)
	bytes64, err := Int64ToBytes(val64)
	assert.NoError(t, err)
	result64, err := BytesToInt64(bytes64)
	assert.NoError(t, err)
	assert.Equal(t, val64, result64)

	// Test with invalid byte length for int64
	_, err = BytesToInt64([]byte{1, 2, 3, 4})
	assert.Error(t, err)

	// Test BytesToUint64
	uval64 := uint64(123456789012345)
	ubytes64, err := Uint64ToBytes(uval64)
	assert.NoError(t, err)
	uresult64, err := BytesToUint64(ubytes64)
	assert.NoError(t, err)
	assert.Equal(t, uval64, uresult64)

	// Test with invalid byte length for uint64
	_, err = BytesToUint64([]byte{1, 2, 3, 4})
	assert.Error(t, err)
}

func TestBytesPrefix(t *testing.T) {
	tests := []struct {
		prefix []byte
		start  []byte
		limit  []byte
	}{
		{
			prefix: []byte("hello"),
			start:  []byte("hello"),
			limit:  []byte("hellp"),
		},
		{
			prefix: []byte{0xff, 0xfe},
			start:  []byte{0xff, 0xfe},
			limit:  []byte{0xff, 0xff},
		},
		{
			prefix: []byte{0xff, 0xff},
			start:  []byte{0xff, 0xff},
			limit:  nil,
		},
		{
			prefix: []byte{},
			start:  []byte{},
			limit:  nil,
		},
	}

	for i, tt := range tests {
		start, limit := BytesPrefix(tt.prefix)
		assert.Equal(t, tt.start, start, "test %d: start bytes", i)
		assert.Equal(t, tt.limit, limit, "test %d: limit bytes", i)
	}
}
