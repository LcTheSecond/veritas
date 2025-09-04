// Package veritas provides comprehensive unit tests for utility functions.
package veritas

import (
	"testing"
)

// TestCleanString tests the cleanString utility function
func TestCleanString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		toLower  bool
		expected string
	}{
		{
			name:     "String with leading/trailing spaces, no lowercase",
			input:    "  Hello World  ",
			toLower:  false,
			expected: "Hello World",
		},
		{
			name:     "String with leading/trailing spaces, with lowercase",
			input:    "  Hello World  ",
			toLower:  true,
			expected: "hello world",
		},
		{
			name:     "String with tabs and newlines, no lowercase",
			input:    "\t\nHello\tWorld\n",
			toLower:  false,
			expected: "Hello\tWorld",
		},
		{
			name:     "String with tabs and newlines, with lowercase",
			input:    "\t\nHello\tWorld\n",
			toLower:  true,
			expected: "hello\tworld",
		},
		{
			name:     "Empty string, no lowercase",
			input:    "",
			toLower:  false,
			expected: "",
		},
		{
			name:     "Empty string, with lowercase",
			input:    "",
			toLower:  true,
			expected: "",
		},
		{
			name:     "String with only spaces, no lowercase",
			input:    "   ",
			toLower:  false,
			expected: "",
		},
		{
			name:     "String with only spaces, with lowercase",
			input:    "   ",
			toLower:  true,
			expected: "",
		},
		{
			name:     "String with mixed case, no lowercase",
			input:    "  Hello WORLD  ",
			toLower:  false,
			expected: "Hello WORLD",
		},
		{
			name:     "String with mixed case, with lowercase",
			input:    "  Hello WORLD  ",
			toLower:  true,
			expected: "hello world",
		},
		{
			name:     "String with special characters, no lowercase",
			input:    "  Hello@World#123  ",
			toLower:  false,
			expected: "Hello@World#123",
		},
		{
			name:     "String with special characters, with lowercase",
			input:    "  Hello@World#123  ",
			toLower:  true,
			expected: "hello@world#123",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := cleanString(tt.input, tt.toLower)
			if result != tt.expected {
				t.Errorf("cleanString() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

// TestIsEmpty tests the isEmpty utility function
func TestIsEmpty(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Empty string",
			input:    "",
			expected: true,
		},
		{
			name:     "String with only spaces",
			input:    "   ",
			expected: true,
		},
		{
			name:     "String with only tabs",
			input:    "\t\t",
			expected: true,
		},
		{
			name:     "String with only newlines",
			input:    "\n\n",
			expected: true,
		},
		{
			name:     "String with mixed whitespace",
			input:    " \t\n ",
			expected: true,
		},
		{
			name:     "String with content",
			input:    "Hello",
			expected: false,
		},
		{
			name:     "String with content and leading spaces",
			input:    "  Hello",
			expected: false,
		},
		{
			name:     "String with content and trailing spaces",
			input:    "Hello  ",
			expected: false,
		},
		{
			name:     "String with content and mixed whitespace",
			input:    " \tHello\n ",
			expected: false,
		},
		{
			name:     "String with special characters",
			input:    "!@#",
			expected: false,
		},
		{
			name:     "String with numbers",
			input:    "123",
			expected: false,
		},
		{
			name:     "String with unicode characters",
			input:    "你好",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isEmpty(tt.input)
			if result != tt.expected {
				t.Errorf("isEmpty() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

// TestIsNotEmpty tests the isNotEmpty utility function
func TestIsNotEmpty(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Empty string",
			input:    "",
			expected: false,
		},
		{
			name:     "String with only spaces",
			input:    "   ",
			expected: false,
		},
		{
			name:     "String with only tabs",
			input:    "\t\t",
			expected: false,
		},
		{
			name:     "String with only newlines",
			input:    "\n\n",
			expected: false,
		},
		{
			name:     "String with mixed whitespace",
			input:    " \t\n ",
			expected: false,
		},
		{
			name:     "String with content",
			input:    "Hello",
			expected: true,
		},
		{
			name:     "String with content and leading spaces",
			input:    "  Hello",
			expected: true,
		},
		{
			name:     "String with content and trailing spaces",
			input:    "Hello  ",
			expected: true,
		},
		{
			name:     "String with content and mixed whitespace",
			input:    " \tHello\n ",
			expected: true,
		},
		{
			name:     "String with special characters",
			input:    "!@#",
			expected: true,
		},
		{
			name:     "String with numbers",
			input:    "123",
			expected: true,
		},
		{
			name:     "String with unicode characters",
			input:    "你好",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isNotEmpty(tt.input)
			if result != tt.expected {
				t.Errorf("isNotEmpty() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

// TestCompileRegex tests the compileRegex utility function
func TestCompileRegex(t *testing.T) {
	tests := []struct {
		name        string
		pattern     string
		expectError bool
	}{
		{
			name:        "Valid simple pattern",
			pattern:     "hello",
			expectError: false,
		},
		{
			name:        "Valid pattern with anchors",
			pattern:     "^hello$",
			expectError: false,
		},
		{
			name:        "Valid pattern with character classes",
			pattern:     "[a-zA-Z0-9]",
			expectError: false,
		},
		{
			name:        "Valid pattern with quantifiers",
			pattern:     "a+",
			expectError: false,
		},
		{
			name:        "Valid pattern with groups",
			pattern:     "(hello|world)",
			expectError: false,
		},
		{
			name:        "Valid pattern with escape sequences",
			pattern:     "\\d+",
			expectError: false,
		},
		{
			name:        "Valid pattern with unicode",
			pattern:     "\\p{L}+",
			expectError: false,
		},
		{
			name:        "Valid complex pattern",
			pattern:     "^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$",
			expectError: false,
		},
		{
			name:        "Invalid pattern with unclosed bracket",
			pattern:     "[a-z",
			expectError: true,
		},
		{
			name:        "Invalid pattern with unclosed parenthesis",
			pattern:     "(hello",
			expectError: true,
		},
		{
			name:        "Invalid pattern with invalid escape",
			pattern:     "\\",
			expectError: true,
		},
		{
			name:        "Invalid pattern with invalid quantifier",
			pattern:     "a{",
			expectError: true,
		},
		{
			name:        "Empty pattern",
			pattern:     "",
			expectError: false,
		},
		{
			name:        "Pattern with special regex characters",
			pattern:     ".*+?^${}[]|()",
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			regex, err := compileRegex(tt.pattern)
			if tt.expectError {
				if err == nil {
					t.Errorf("compileRegex() expected error, got nil")
				}
			} else {
				if err != nil {
					t.Errorf("compileRegex() unexpected error: %v", err)
				}
				if regex == nil {
					t.Errorf("compileRegex() returned nil regex")
				}
			}
		})
	}
}

// TestMatchRegex tests the matchRegex utility function
func TestMatchRegex(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		pattern     string
		expected    bool
		expectError bool
	}{
		{
			name:        "String matches simple pattern",
			input:       "hello",
			pattern:     "hello",
			expected:    true,
			expectError: false,
		},
		{
			name:        "String does not match simple pattern",
			input:       "world",
			pattern:     "hello",
			expected:    false,
			expectError: false,
		},
		{
			name:        "String matches pattern with anchors",
			input:       "hello",
			pattern:     "^hello$",
			expected:    true,
			expectError: false,
		},
		{
			name:        "String does not match pattern with anchors",
			input:       "hello world",
			pattern:     "^hello$",
			expected:    false,
			expectError: false,
		},
		{
			name:        "String matches pattern with character classes",
			input:       "a",
			pattern:     "[a-zA-Z]",
			expected:    true,
			expectError: false,
		},
		{
			name:        "String does not match pattern with character classes",
			input:       "1",
			pattern:     "[a-zA-Z]",
			expected:    false,
			expectError: false,
		},
		{
			name:        "String matches pattern with quantifiers",
			input:       "aaa",
			pattern:     "a+",
			expected:    true,
			expectError: false,
		},
		{
			name:        "String does not match pattern with quantifiers",
			input:       "b",
			pattern:     "a+",
			expected:    false,
			expectError: false,
		},
		{
			name:        "String matches pattern with groups",
			input:       "hello",
			pattern:     "(hello|world)",
			expected:    true,
			expectError: false,
		},
		{
			name:        "String matches pattern with groups (second option)",
			input:       "world",
			pattern:     "(hello|world)",
			expected:    true,
			expectError: false,
		},
		{
			name:        "String does not match pattern with groups",
			input:       "test",
			pattern:     "(hello|world)",
			expected:    false,
			expectError: false,
		},
		{
			name:        "String matches pattern with escape sequences",
			input:       "123",
			pattern:     "\\d+",
			expected:    true,
			expectError: false,
		},
		{
			name:        "String does not match pattern with escape sequences",
			input:       "abc",
			pattern:     "\\d+",
			expected:    false,
			expectError: false,
		},
		{
			name:        "String matches email pattern",
			input:       "user@example.com",
			pattern:     "^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$",
			expected:    true,
			expectError: false,
		},
		{
			name:        "String does not match email pattern",
			input:       "invalid-email",
			pattern:     "^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$",
			expected:    false,
			expectError: false,
		},
		{
			name:        "Empty string matches empty pattern",
			input:       "",
			pattern:     "",
			expected:    true,
			expectError: false,
		},
		{
			name:        "Empty string does not match non-empty pattern",
			input:       "",
			pattern:     "hello",
			expected:    false,
			expectError: false,
		},
		{
			name:        "Non-empty string does not match empty pattern",
			input:       "hello",
			pattern:     "",
			expected:    true,
			expectError: false,
		},
		{
			name:        "Invalid regex pattern",
			input:       "hello",
			pattern:     "[a-z",
			expected:    false,
			expectError: true,
		},
		{
			name:        "String with unicode characters",
			input:       "你好",
			pattern:     "\\p{L}+",
			expected:    true,
			expectError: false,
		},
		{
			name:        "String with special characters",
			input:       "hello@world#123",
			pattern:     ".*",
			expected:    true,
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := matchRegex(tt.input, tt.pattern)
			if tt.expectError {
				if err == nil {
					t.Errorf("matchRegex() expected error, got nil")
				}
			} else {
				if err != nil {
					t.Errorf("matchRegex() unexpected error: %v", err)
				}
				if result != tt.expected {
					t.Errorf("matchRegex() = %v, expected %v", result, tt.expected)
				}
			}
		})
	}
}

// TestCompileRegex_Integration tests compileRegex with actual regex operations
func TestCompileRegex_Integration(t *testing.T) {
	tests := []struct {
		name     string
		pattern  string
		testStr  string
		expected bool
	}{
		{
			name:     "Email validation pattern",
			pattern:  "^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$",
			testStr:  "user@example.com",
			expected: true,
		},
		{
			name:     "Phone number pattern",
			pattern:  "^\\+?[1-9]\\d{1,14}$",
			testStr:  "+1234567890",
			expected: true,
		},
		{
			name:     "Alphanumeric pattern",
			pattern:  "^[a-zA-Z0-9]+$",
			testStr:  "Hello123",
			expected: true,
		},
		{
			name:     "Date pattern",
			pattern:  "^\\d{4}-\\d{2}-\\d{2}$",
			testStr:  "2023-12-25",
			expected: true,
		},
		{
			name:     "URL pattern",
			pattern:  "^https?://[^\\s/$.?#].[^\\s]*$",
			testStr:  "https://example.com",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			regex, err := compileRegex(tt.pattern)
			if err != nil {
				t.Errorf("compileRegex() error: %v", err)
				return
			}

			result := regex.MatchString(tt.testStr)
			if result != tt.expected {
				t.Errorf("regex.MatchString() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

// TestUtilityFunctions_EdgeCases tests edge cases for utility functions
func TestUtilityFunctions_EdgeCases(t *testing.T) {
	t.Run("cleanString with unicode whitespace", func(t *testing.T) {
		input := "\u00A0\u2000\u2001\u2002\u2003\u2004\u2005\u2006\u2007\u2008\u2009\u200A\u200B\u200C\u200D\u2028\u2029\u202F\u205F\u3000Hello\u00A0\u2000\u2001\u2002\u2003\u2004\u2005\u2006\u2007\u2008\u2009\u200A\u200B\u200C\u200D\u2028\u2029\u202F\u205F\u3000"
		result := cleanString(input, false)
		expected := "Hello"
		if result != expected {
			t.Errorf("cleanString() with unicode whitespace = %v, expected %v", result, expected)
		}
	})

	t.Run("isEmpty with unicode whitespace", func(t *testing.T) {
		input := "\u00A0\u2000\u2001\u2002\u2003\u2004\u2005\u2006\u2007\u2008\u2009\u200A\u200B\u200C\u200D\u2028\u2029\u202F\u205F\u3000"
		result := isEmpty(input)
		expected := true
		if result != expected {
			t.Errorf("isEmpty() with unicode whitespace = %v, expected %v", result, expected)
		}
	})

	t.Run("matchRegex with unicode characters", func(t *testing.T) {
		input := "Hello 世界"
		pattern := "\\p{L}+"
		result, err := matchRegex(input, pattern)
		if err != nil {
			t.Errorf("matchRegex() error: %v", err)
		}
		expected := true
		if result != expected {
			t.Errorf("matchRegex() with unicode = %v, expected %v", result, expected)
		}
	})

	t.Run("compileRegex with complex unicode pattern", func(t *testing.T) {
		pattern := "\\p{L}+\\s+\\p{L}+"
		regex, err := compileRegex(pattern)
		if err != nil {
			t.Errorf("compileRegex() error: %v", err)
		}
		if regex == nil {
			t.Errorf("compileRegex() returned nil regex")
		}
	})
}
