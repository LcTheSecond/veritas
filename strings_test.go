// Package veritas provides comprehensive unit tests for string validation functions.
package veritas

import (
	"testing"
)

// TestValidateString_ValidCases tests valid string lengths
func TestValidateString_ValidCases(t *testing.T) {
	tests := []struct {
		name      string
		str       string
		minLength int
		maxLength int
		expected  error
	}{
		{
			name:      "String within bounds",
			str:       "Hello World",
			minLength: 5,
			maxLength: 20,
			expected:  nil,
		},
		{
			name:      "String at minimum length",
			str:       "Hi",
			minLength: 2,
			maxLength: 10,
			expected:  nil,
		},
		{
			name:      "String at maximum length",
			str:       "Hello",
			minLength: 2,
			maxLength: 5,
			expected:  nil,
		},
		{
			name:      "String with UTF-8 characters",
			str:       "Café",
			minLength: 1,
			maxLength: 10,
			expected:  nil,
		},
		{
			name:      "String with emojis",
			str:       "Hello 😀",
			minLength: 1,
			maxLength: 10,
			expected:  nil,
		},
		{
			name:      "String with Chinese characters",
			str:       "你好世界",
			minLength: 1,
			maxLength: 10,
			expected:  nil,
		},
		{
			name:      "String with Arabic characters",
			str:       "مرحبا",
			minLength: 1,
			maxLength: 10,
			expected:  nil,
		},
		{
			name:      "String with Cyrillic characters",
			str:       "Привет",
			minLength: 1,
			maxLength: 10,
			expected:  nil,
		},
		{
			name:      "String with numbers and symbols",
			str:       "Test123!@#",
			minLength: 5,
			maxLength: 15,
			expected:  nil,
		},
		{
			name:      "String with whitespace",
			str:       "  Hello World  ",
			minLength: 5,
			maxLength: 20,
			expected:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateString(tt.str, tt.minLength, tt.maxLength)
			if err != tt.expected {
				t.Errorf("ValidateString() error = %v, expected %v", err, tt.expected)
			}
		})
	}
}

// TestValidateString_InvalidLengths tests invalid string lengths
func TestValidateString_InvalidLengths(t *testing.T) {
	tests := []struct {
		name      string
		str       string
		minLength int
		maxLength int
		expected  string
	}{
		{
			name:      "String too short",
			str:       "Hi",
			minLength: 5,
			maxLength: 20,
			expected:  "string must be at least 5 characters long",
		},
		{
			name:      "String too long",
			str:       "This is a very long string that exceeds the maximum length",
			minLength: 5,
			maxLength: 20,
			expected:  "string must be at most 20 characters long",
		},
		{
			name:      "Empty string with minimum length",
			str:       "",
			minLength: 1,
			maxLength: 10,
			expected:  "string must be at least 1 characters long",
		},
		{
			name:      "Single character with minimum length 2",
			str:       "A",
			minLength: 2,
			maxLength: 10,
			expected:  "string must be at least 2 characters long",
		},
		{
			name:      "String exactly at max length + 1",
			str:       "Hello",
			minLength: 2,
			maxLength: 4,
			expected:  "string must be at most 4 characters long",
		},
		{
			name:      "UTF-8 string too short",
			str:       "你",
			minLength: 2,
			maxLength: 10,
			expected:  "string must be at least 2 characters long",
		},
		{
			name:      "UTF-8 string too long",
			str:       "你好世界测试",
			minLength: 1,
			maxLength: 4,
			expected:  "string must be at most 4 characters long",
		},
		{
			name:      "Emoji string too short",
			str:       "😀",
			minLength: 2,
			maxLength: 10,
			expected:  "string must be at least 2 characters long",
		},
		{
			name:      "Emoji string too long",
			str:       "😀😀😀😀😀",
			minLength: 1,
			maxLength: 4,
			expected:  "string must be at most 4 characters long",
		},
		{
			name:      "String with only whitespace",
			str:       "   ",
			minLength: 1,
			maxLength: 10,
			expected:  "string must be at least 1 characters long",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateString(tt.str, tt.minLength, tt.maxLength)
			if err == nil {
				t.Errorf("ValidateString() expected error, got nil")
			} else if err.Error() != tt.expected {
				t.Errorf("ValidateString() error = %v, expected %v", err.Error(), tt.expected)
			}
		})
	}
}

// TestValidateString_EdgeCases tests edge cases for string validation
func TestValidateString_EdgeCases(t *testing.T) {
	tests := []struct {
		name      string
		str       string
		minLength int
		maxLength int
		expected  error
	}{
		{
			name:      "Zero length bounds",
			str:       "",
			minLength: 0,
			maxLength: 0,
			expected:  nil,
		},
		{
			name:      "Same min and max length",
			str:       "Hi",
			minLength: 2,
			maxLength: 2,
			expected:  nil,
		},
		{
			name:      "Very large bounds",
			str:       "Hello",
			minLength: 1,
			maxLength: 1000,
			expected:  nil,
		},
		{
			name:      "String with newlines",
			str:       "Hello\nWorld",
			minLength: 5,
			maxLength: 20,
			expected:  nil,
		},
		{
			name:      "String with tabs",
			str:       "Hello\tWorld",
			minLength: 5,
			maxLength: 20,
			expected:  nil,
		},
		{
			name:      "String with carriage returns",
			str:       "Hello\r\nWorld",
			minLength: 5,
			maxLength: 20,
			expected:  nil,
		},
		{
			name:      "String with mixed whitespace",
			str:       "  Hello\t\nWorld  ",
			minLength: 5,
			maxLength: 20,
			expected:  nil,
		},
		{
			name:      "String with special characters",
			str:       "!@#$%^&*()",
			minLength: 5,
			maxLength: 15,
			expected:  nil,
		},
		{
			name:      "String with unicode spaces",
			str:       "Hello\u00A0World", // Non-breaking space
			minLength: 5,
			maxLength: 20,
			expected:  nil,
		},
		{
			name:      "String with zero-width characters",
			str:       "Hello\u200BWorld", // Zero-width space
			minLength: 5,
			maxLength: 20,
			expected:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateString(tt.str, tt.minLength, tt.maxLength)
			if err != tt.expected {
				t.Errorf("ValidateString() error = %v, expected %v", err, tt.expected)
			}
		})
	}
}

// TestValidateString_TypeValidation tests type validation for string validation
func TestValidateString_TypeValidation(t *testing.T) {
	tests := []struct {
		name      string
		str       interface{}
		minLength int
		maxLength int
		expected  string
	}{
		{
			name:      "Integer input",
			str:       123,
			minLength: 1,
			maxLength: 10,
			expected:  "value must be a string",
		},
		{
			name:      "Float input",
			str:       123.45,
			minLength: 1,
			maxLength: 10,
			expected:  "value must be a string",
		},
		{
			name:      "Boolean input",
			str:       true,
			minLength: 1,
			maxLength: 10,
			expected:  "value must be a string",
		},
		{
			name:      "Nil input",
			str:       nil,
			minLength: 1,
			maxLength: 10,
			expected:  "value must be a string",
		},
		{
			name:      "Slice input",
			str:       []string{"Hello", "World"},
			minLength: 1,
			maxLength: 10,
			expected:  "value must be a string",
		},
		{
			name:      "Map input",
			str:       map[string]string{"key": "value"},
			minLength: 1,
			maxLength: 10,
			expected:  "value must be a string",
		},
		{
			name:      "Struct input",
			str:       struct{ Name string }{"Hello"},
			minLength: 1,
			maxLength: 10,
			expected:  "value must be a string",
		},
		{
			name:      "Channel input",
			str:       make(chan string),
			minLength: 1,
			maxLength: 10,
			expected:  "value must be a string",
		},
		{
			name:      "Function input",
			str:       func() string { return "Hello" },
			minLength: 1,
			maxLength: 10,
			expected:  "value must be a string",
		},
		{
			name:      "Pointer input",
			str:       &[]string{"Hello"}[0],
			minLength: 1,
			maxLength: 10,
			expected:  "value must be a string",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateString(tt.str, tt.minLength, tt.maxLength)
			if err == nil {
				t.Errorf("ValidateString() expected error, got nil")
			} else if err.Error() != tt.expected {
				t.Errorf("ValidateString() error = %v, expected %v", err.Error(), tt.expected)
			}
		})
	}
}

// TestValidateString_BoundaryConditions tests boundary conditions
func TestValidateString_BoundaryConditions(t *testing.T) {
	tests := []struct {
		name      string
		str       string
		minLength int
		maxLength int
		expected  string
	}{
		{
			name:      "String exactly at minimum length",
			str:       "Hi",
			minLength: 2,
			maxLength: 10,
			expected:  "",
		},
		{
			name:      "String exactly at maximum length",
			str:       "Hello",
			minLength: 2,
			maxLength: 5,
			expected:  "",
		},
		{
			name:      "String one character below minimum",
			str:       "H",
			minLength: 2,
			maxLength: 10,
			expected:  "string must be at least 2 characters long", // This should fail
		},
		{
			name:      "String one character above maximum",
			str:       "Hello!",
			minLength: 2,
			maxLength: 5,
			expected:  "string must be at most 5 characters long", // This should fail
		},
		{
			name:      "UTF-8 string at boundary",
			str:       "你好",
			minLength: 2,
			maxLength: 2,
			expected:  "",
		},
		{
			name:      "Emoji string at boundary",
			str:       "😀😀",
			minLength: 2,
			maxLength: 2,
			expected:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateString(tt.str, tt.minLength, tt.maxLength)
			if tt.expected == "" {
				if err != nil {
					t.Errorf("ValidateString() error = %v, expected nil", err)
				}
			} else {
				if err == nil {
					t.Errorf("ValidateString() expected error, got nil")
				} else if err.Error() != tt.expected {
					t.Errorf("ValidateString() error = %v, expected %v", err.Error(), tt.expected)
				}
			}
		})
	}
}
