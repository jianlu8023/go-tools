package bytes

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"math"
	"unsafe"
)

const (
	Byte = 1

	KiloByte = Byte * 1000
	MegaByte = KiloByte * 1000
	GigaByte = MegaByte * 1000
	TeraByte = GigaByte * 1000

	KibiByte = Byte * 1024
	MebiByte = KibiByte * 1024
	GibiByte = MebiByte * 1024
)

func HumanDecimal(b int64) string {
	var value float64
	var unit string

	switch {
	case b >= TeraByte:
		value = float64(b) / TeraByte
		unit = "TB"
	case b >= GigaByte:
		value = float64(b) / GigaByte
		unit = "GB"
	case b >= MegaByte:
		value = float64(b) / MegaByte
		unit = "MB"
	case b >= KiloByte:
		value = float64(b) / KiloByte
		unit = "KB"
	default:
		return fmt.Sprintf("%d B", b)
	}

	switch {
	case value >= 10:
		return fmt.Sprintf("%d %s", int(value), unit)
	case value != math.Trunc(value):
		return fmt.Sprintf("%.1f %s", value, unit)
	default:
		return fmt.Sprintf("%d %s", int(value), unit)
	}
}

func HumanBinary(b uint64) string {
	switch {
	case b >= GibiByte:
		return fmt.Sprintf("%.1f GiB", float64(b)/GibiByte)
	case b >= MebiByte:
		return fmt.Sprintf("%.1f MiB", float64(b)/MebiByte)
	case b >= KibiByte:
		return fmt.Sprintf("%.1f KiB", float64(b)/KibiByte)
	default:
		return fmt.Sprintf("%d B", b)
	}
}

// BytesToInt le bytehelper to int32, little endian
func BytesToInt(b []byte) (int32, error) {
	bytesBuffer := bytes.NewBuffer(b)

	var x int32
	err := binary.Read(bytesBuffer, binary.LittleEndian, &x)
	if err != nil {
		return -1, err
	}
	return x, nil
}

// BytesToInt64 le bytehelper to int64, little endian
func BytesToInt64(b []byte) (int64, error) {
	bytesBuffer := bytes.NewBuffer(b)
	var x int64
	err := binary.Read(bytesBuffer, binary.LittleEndian, &x)
	if err != nil {
		return -1, err
	}
	return x, nil
}

// IntToBytes int32 to le bytehelper, little endian
func IntToBytes(x int32) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	err := binary.Write(bytesBuffer, binary.LittleEndian, x)
	if err != nil {
		return nil
	}
	return bytesBuffer.Bytes()
}

// Int64ToBytes int64 to le bytehelper, little endian
func Int64ToBytes(x int64) ([]byte, error) {
	bytesBuffer := bytes.NewBuffer([]byte{})
	err := binary.Write(bytesBuffer, binary.LittleEndian, x)
	if err != nil {
		return nil, err
	}
	return bytesBuffer.Bytes(), nil
}

// BytesToUint64 le bytehelper to uint64, little endian
func BytesToUint64(b []byte) (uint64, error) {
	bytesBuffer := bytes.NewBuffer(b)
	var x uint64
	err := binary.Read(bytesBuffer, binary.LittleEndian, &x)
	if err != nil {
		return 0, err
	}
	return x, nil
}

// Uint64ToBytes uint64 to le bytehelper, little endian
func Uint64ToBytes(x uint64) ([]byte, error) {
	bytesBuffer := bytes.NewBuffer([]byte{})
	err := binary.Write(bytesBuffer, binary.LittleEndian, x)
	if err != nil {
		return nil, err
	}
	return bytesBuffer.Bytes(), nil
}

// BytesPrefix returns key range that satisfy the given prefix.
// This only applicable for the standard 'bytehelper comparer'.
func BytesPrefix(prefix []byte) ([]byte, []byte) {
	var limit []byte
	for i := len(prefix) - 1; i >= 0; i-- {
		c := prefix[i]
		if c < 0xff {
			limit = make([]byte, i+1)
			copy(limit, prefix)
			limit[i] = c + 1
			break
		}
	}
	return prefix, limit
}

// StringToBytes converts string to byte slice without a memory allocation.
func StringToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{s, len(s)},
	))
}

// BytesToString converts byte slice to string without a memory allocation.
func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// Compare compares two byte slices lexicographically.
// It returns 0 if a == b, -1 if a < b, and 1 if a > b.
func Compare(a, b []byte) int {
	minLen := len(a)
	if len(b) < minLen {
		minLen = len(b)
	}

	for i := 0; i < minLen; i++ {
		if a[i] != b[i] {
			if a[i] < b[i] {
				return -1
			}
			return 1
		}
	}

	// All compared bytes are equal, so the shorter slice is considered smaller
	if len(a) == len(b) {
		return 0
	} else if len(a) < len(b) {
		return -1
	}
	return 1
}

// Equal checks if two byte slices are equal.
// nil and empty slice are considered not equal.
func Equal(a, b []byte) bool {
	// 特殊处理nil和空切片的情况
	if (a == nil && b != nil) || (a != nil && b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

// Clone creates a copy of the given byte slice.
func Clone(b []byte) []byte {
	if b == nil {
		return nil
	}

	clone := make([]byte, len(b))
	copy(clone, b)
	return clone
}

// Reverse reverses the order of elements in the byte slice.
// It returns a new slice with the reversed elements.
func Reverse(b []byte) []byte {
	if b == nil {
		return nil
	}

	result := make([]byte, len(b))
	for i, j := 0, len(b)-1; i < len(b); i, j = i+1, j-1 {
		result[i] = b[j]
	}

	return result
}

// Padding pads the byte slice to the specified length with the given pad byte.
// If the slice is already longer than the specified length, it is returned unchanged.
func Padding(b []byte, length int, padByte byte) []byte {
	if b == nil {
		b = []byte{}
	}

	if len(b) >= length {
		return b
	}

	padded := make([]byte, length)
	copy(padded, b)

	// Fill the remaining bytes with padByte
	for i := len(b); i < length; i++ {
		padded[i] = padByte
	}

	return padded
}

// Slice safely slices the byte slice from start to end indices.
// It returns an error if the indices are out of bounds or if start > end.
func Slice(b []byte, start, end int) ([]byte, error) {
	if start < 0 || end > len(b) || start > end {
		return nil, errors.New("invalid slice indices")
	}

	result := make([]byte, end-start)
	copy(result, b[start:end])
	return result, nil
}

// Contains checks if the subslice exists within the byte slice.
func Contains(b, sub []byte) bool {
	if len(sub) == 0 {
		return true
	}

	if len(b) < len(sub) {
		return false
	}

	for i := 0; i <= len(b)-len(sub); i++ {
		found := true
		for j := 0; j < len(sub); j++ {
			if b[i+j] != sub[j] {
				found = false
				break
			}
		}
		if found {
			return true
		}
	}

	return false
}

// Index returns the index of the first occurrence of the subslice within the byte slice.
// It returns -1 if the subslice is not found.
func Index(b, sub []byte) int {
	if len(sub) == 0 {
		return 0
	}

	if len(b) < len(sub) {
		return -1
	}

	for i := 0; i <= len(b)-len(sub); i++ {
		found := true
		for j := 0; j < len(sub); j++ {
			if b[i+j] != sub[j] {
				found = false
				break
			}
		}
		if found {
			return i
		}
	}

	return -1
}

// HexEncode converts a byte slice to a hexadecimal string.
func HexEncode(b []byte) string {
	return hex.EncodeToString(b)
}

// HexDecode converts a hexadecimal string to a byte slice.
func HexDecode(s string) ([]byte, error) {
	return hex.DecodeString(s)
}
