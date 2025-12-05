package sqlnull

import (
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNullStringPtr(t *testing.T) {
	tests := []struct {
		name     string
		input    sql.NullString
		expected *string
	}{
		{
			name:     "valid string",
			input:    sql.NullString{String: "hello", Valid: true},
			expected: stringPtr("hello"),
		},
		{
			name:     "invalid string",
			input:    sql.NullString{String: "", Valid: false},
			expected: nil,
		},
		{
			name:     "empty valid string",
			input:    sql.NullString{String: "", Valid: true},
			expected: stringPtr(""),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NullStringPtr(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestStringToNull(t *testing.T) {
	tests := []struct {
		name     string
		input    *string
		expected sql.NullString
	}{
		{
			name:     "non-nil string",
			input:    stringPtr("hello"),
			expected: sql.NullString{String: "hello", Valid: true},
		},
		{
			name:     "nil string",
			input:    nil,
			expected: sql.NullString{String: "", Valid: false},
		},
		{
			name:     "empty string",
			input:    stringPtr(""),
			expected: sql.NullString{String: "", Valid: true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := StringToNull(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestNullInt64Ptr(t *testing.T) {
	tests := []struct {
		name     string
		input    sql.NullInt64
		expected *int64
	}{
		{
			name:     "valid int64",
			input:    sql.NullInt64{Int64: 42, Valid: true},
			expected: int64Ptr(42),
		},
		{
			name:     "invalid int64",
			input:    sql.NullInt64{Int64: 0, Valid: false},
			expected: nil,
		},
		{
			name:     "zero valid int64",
			input:    sql.NullInt64{Int64: 0, Valid: true},
			expected: int64Ptr(0),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NullInt64Ptr(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestInt64ToNull(t *testing.T) {
	tests := []struct {
		name     string
		input    *int64
		expected sql.NullInt64
	}{
		{
			name:     "non-nil int64",
			input:    int64Ptr(42),
			expected: sql.NullInt64{Int64: 42, Valid: true},
		},
		{
			name:     "nil int64",
			input:    nil,
			expected: sql.NullInt64{Int64: 0, Valid: false},
		},
		{
			name:     "zero int64",
			input:    int64Ptr(0),
			expected: sql.NullInt64{Int64: 0, Valid: true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Int64ToNull(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestNullBoolPtr(t *testing.T) {
	tests := []struct {
		name     string
		input    sql.NullBool
		expected *bool
	}{
		{
			name:     "valid true bool",
			input:    sql.NullBool{Bool: true, Valid: true},
			expected: boolPtr(true),
		},
		{
			name:     "valid false bool",
			input:    sql.NullBool{Bool: false, Valid: true},
			expected: boolPtr(false),
		},
		{
			name:     "invalid bool",
			input:    sql.NullBool{Bool: false, Valid: false},
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NullBoolPtr(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestBoolToNull(t *testing.T) {
	tests := []struct {
		name     string
		input    *bool
		expected sql.NullBool
	}{
		{
			name:     "non-nil true bool",
			input:    boolPtr(true),
			expected: sql.NullBool{Bool: true, Valid: true},
		},
		{
			name:     "non-nil false bool",
			input:    boolPtr(false),
			expected: sql.NullBool{Bool: false, Valid: true},
		},
		{
			name:     "nil bool",
			input:    nil,
			expected: sql.NullBool{Bool: false, Valid: false},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := BoolToNull(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestNullFloat64Ptr(t *testing.T) {
	tests := []struct {
		name     string
		input    sql.NullFloat64
		expected *float64
	}{
		{
			name:     "valid float64",
			input:    sql.NullFloat64{Float64: 3.14, Valid: true},
			expected: float64Ptr(3.14),
		},
		{
			name:     "invalid float64",
			input:    sql.NullFloat64{Float64: 0.0, Valid: false},
			expected: nil,
		},
		{
			name:     "zero valid float64",
			input:    sql.NullFloat64{Float64: 0.0, Valid: true},
			expected: float64Ptr(0.0),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NullFloat64Ptr(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestFloat64ToNull(t *testing.T) {
	tests := []struct {
		name     string
		input    *float64
		expected sql.NullFloat64
	}{
		{
			name:     "non-nil float64",
			input:    float64Ptr(3.14),
			expected: sql.NullFloat64{Float64: 3.14, Valid: true},
		},
		{
			name:     "nil float64",
			input:    nil,
			expected: sql.NullFloat64{Float64: 0.0, Valid: false},
		},
		{
			name:     "zero float64",
			input:    float64Ptr(0.0),
			expected: sql.NullFloat64{Float64: 0.0, Valid: true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Float64ToNull(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestNullTimePtr(t *testing.T) {
	now := time.Now()
	tests := []struct {
		name     string
		input    sql.NullTime
		expected *time.Time
	}{
		{
			name:     "valid time",
			input:    sql.NullTime{Time: now, Valid: true},
			expected: &now,
		},
		{
			name:     "invalid time",
			input:    sql.NullTime{Time: time.Time{}, Valid: false},
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NullTimePtr(tt.input)
			if tt.expected == nil {
				assert.Nil(t, result)
			} else {
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

func TestTimeToNull(t *testing.T) {
	now := time.Now()
	tests := []struct {
		name     string
		input    *time.Time
		expected sql.NullTime
	}{
		{
			name:     "non-nil time",
			input:    &now,
			expected: sql.NullTime{Time: now, Valid: true},
		},
		{
			name:     "nil time",
			input:    nil,
			expected: sql.NullTime{Time: time.Time{}, Valid: false},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := TimeToNull(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

// Helper functions to create pointers
func stringPtr(s string) *string {
	return &s
}

func int64Ptr(i int64) *int64 {
	return &i
}

func boolPtr(b bool) *bool {
	return &b
}

func float64Ptr(f float64) *float64 {
	return &f
}
