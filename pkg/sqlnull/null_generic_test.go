//go:build go1.18
// +build go1.18

package sqlnull

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNullPtr(t *testing.T) {
	type testCase[T any] struct {
		name     string
		input    sql.Null[T]
		expected *T
	}

	// Test with string
	stringTests := []testCase[string]{
		{
			name:     "valid string",
			input:    sql.Null[string]{V: "hello", Valid: true},
			expected: stringPtr("hello"),
		},
		{
			name:     "invalid string",
			input:    sql.Null[string]{V: "", Valid: false},
			expected: nil,
		},
	}

	for _, tt := range stringTests {
		t.Run(tt.name, func(t *testing.T) {
			result := NullPtr(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}

	// Test with int
	intTests := []testCase[int]{
		{
			name:     "valid int",
			input:    sql.Null[int]{V: 42, Valid: true},
			expected: intPtr(42),
		},
		{
			name:     "invalid int",
			input:    sql.Null[int]{V: 0, Valid: false},
			expected: nil,
		},
	}

	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			result := NullPtr(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}

	// Test with bool
	boolTests := []testCase[bool]{
		{
			name:     "valid bool true",
			input:    sql.Null[bool]{V: true, Valid: true},
			expected: boolPtr(true),
		},
		{
			name:     "valid bool false",
			input:    sql.Null[bool]{V: false, Valid: true},
			expected: boolPtr(false),
		},
		{
			name:     "invalid bool",
			input:    sql.Null[bool]{V: false, Valid: false},
			expected: nil,
		},
	}

	for _, tt := range boolTests {
		t.Run(tt.name, func(t *testing.T) {
			result := NullPtr(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestToNull(t *testing.T) {
	type testCase[T any] struct {
		name     string
		input    *T
		expected sql.Null[T]
	}

	// Test with string
	stringTests := []testCase[string]{
		{
			name:     "non-nil string",
			input:    stringPtr("hello"),
			expected: sql.Null[string]{V: "hello", Valid: true},
		},
		{
			name:     "nil string",
			input:    nil,
			expected: sql.Null[string]{V: "", Valid: false},
		},
	}

	for _, tt := range stringTests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToNull(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}

	// Test with int
	intTests := []testCase[int]{
		{
			name:     "non-nil int",
			input:    intPtr(42),
			expected: sql.Null[int]{V: 42, Valid: true},
		},
		{
			name:     "nil int",
			input:    nil,
			expected: sql.Null[int]{V: 0, Valid: false},
		},
	}

	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToNull(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}

	// Test with bool
	boolTests := []testCase[bool]{
		{
			name:     "non-nil bool true",
			input:    boolPtr(true),
			expected: sql.Null[bool]{V: true, Valid: true},
		},
		{
			name:     "non-nil bool false",
			input:    boolPtr(false),
			expected: sql.Null[bool]{V: false, Valid: true},
		},
		{
			name:     "nil bool",
			input:    nil,
			expected: sql.Null[bool]{V: false, Valid: false},
		},
	}

	for _, tt := range boolTests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToNull(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

// Helper functions to create pointers
func intPtr(i int) *int {
	return &i
}
