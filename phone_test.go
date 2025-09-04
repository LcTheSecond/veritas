// Package veritas provides comprehensive unit tests for phone validation functions.
package veritas

import (
	"testing"
)

// TestValidatePhone_ValidMobileCases tests valid Brazilian mobile phone numbers
func TestValidatePhone_ValidMobileCases(t *testing.T) {
	tests := []struct {
		name     string
		phone    string
		expected error
	}{
		{
			name:     "Valid mobile with country code and formatting",
			phone:    "+55 (41) 99504-8710",
			expected: nil,
		},
		{
			name:     "Valid mobile without country code",
			phone:    "(41) 99504-8710",
			expected: nil,
		},
		{
			name:     "Valid mobile with dots",
			phone:    "+55 41.99504.8710",
			expected: nil,
		},
		{
			name:     "Valid mobile with spaces",
			phone:    "+55 41 99504 8710",
			expected: nil,
		},
		{
			name:     "Valid mobile without formatting",
			phone:    "+5541995048710",
			expected: nil,
		},
		{
			name:     "Valid mobile from São Paulo",
			phone:    "+55 11 98765-4321",
			expected: nil,
		},
		{
			name:     "Valid mobile from Rio de Janeiro",
			phone:    "+55 21 99999-8888",
			expected: nil,
		},
		{
			name:     "Valid mobile from Minas Gerais",
			phone:    "+55 31 91234-5678",
			expected: nil,
		},
		{
			name:     "Valid mobile from Paraná",
			phone:    "+55 41 99876-5432",
			expected: nil,
		},
		{
			name:     "Valid mobile from Santa Catarina",
			phone:    "+55 47 98765-4321",
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidatePhone(tt.phone)
			if err != tt.expected {
				t.Errorf("ValidatePhone() error = %v, expected %v", err, tt.expected)
			}
		})
	}
}

// TestValidatePhone_ValidLandlineCases tests valid Brazilian landline phone numbers
func TestValidatePhone_ValidLandlineCases(t *testing.T) {
	tests := []struct {
		name     string
		phone    string
		expected error
	}{
		{
			name:     "Valid landline with country code and formatting",
			phone:    "+55 (41) 3346-4468",
			expected: nil,
		},
		{
			name:     "Valid landline without country code",
			phone:    "(41) 3346-4468",
			expected: nil,
		},
		{
			name:     "Valid landline with dots",
			phone:    "+55 41.3346.4468",
			expected: nil,
		},
		{
			name:     "Valid landline with spaces",
			phone:    "+55 41 3346 4468",
			expected: nil,
		},
		{
			name:     "Valid landline without formatting",
			phone:    "+554133464468",
			expected: nil,
		},
		{
			name:     "Valid landline from São Paulo",
			phone:    "+55 11 3333-4444",
			expected: nil,
		},
		{
			name:     "Valid landline from Rio de Janeiro",
			phone:    "+55 21 2222-3333",
			expected: nil,
		},
		{
			name:     "Valid landline from Minas Gerais",
			phone:    "+55 31 3333-2222",
			expected: nil,
		},
		{
			name:     "Valid landline from Paraná",
			phone:    "+55 41 3333-1111",
			expected: nil,
		},
		{
			name:     "Valid landline from Santa Catarina",
			phone:    "+55 47 3333-0000",
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidatePhone(tt.phone)
			if err != tt.expected {
				t.Errorf("ValidatePhone() error = %v, expected %v", err, tt.phone)
			}
		})
	}
}

// TestValidatePhone_InvalidFormats tests invalid phone number formats
func TestValidatePhone_InvalidFormats(t *testing.T) {
	tests := []struct {
		name     string
		phone    string
		expected string
	}{
		{
			name:     "Phone too short",
			phone:    "+55 41 123",
			expected: "invalid Brazilian phone number format",
		},
		{
			name:     "Phone too long",
			phone:    "+55 41 123456789012",
			expected: "invalid Brazilian phone number format",
		},
		{
			name:     "Invalid country code",
			phone:    "+56 41 99504-8710",
			expected: "invalid Brazilian phone number format",
		},
		{
			name:     "Missing country code for international format",
			phone:    "41 99504-8710",
			expected: "invalid Brazilian phone number format",
		},
		{
			name:     "Invalid DDD",
			phone:    "+55 00 99504-8710",
			expected: "invalid area code (DDD)",
		},
		{
			name:     "Non-existent DDD",
			phone:    "+55 99 99504-8710",
			expected: "invalid area code (DDD)",
		},
		{
			name:     "Mobile without 9",
			phone:    "+55 41 8504-8710",
			expected: "mobile number must start with 9 after area code",
		},
		{
			name:     "Landline with 9",
			phone:    "+55 41 93346-4468",
			expected: "invalid Brazilian phone number format",
		},
		{
			name:     "Phone with letters",
			phone:    "+55 41 99504-871a",
			expected: "invalid phone number digits",
		},
		{
			name:     "Empty string",
			phone:    "",
			expected: "phone cannot be empty",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidatePhone(tt.phone)
			if err == nil {
				t.Errorf("ValidatePhone() expected error, got nil")
			} else if err.Error() != tt.expected {
				t.Errorf("ValidatePhone() error = %v, expected %v", err.Error(), tt.expected)
			}
		})
	}
}

// TestValidatePhone_EdgeCases tests edge cases for phone validation
func TestValidatePhone_EdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		phone    string
		expected string
	}{
		{
			name:     "Phone with only spaces",
			phone:    "   ",
			expected: "phone cannot be empty",
		},
		{
			name:     "Phone with mixed separators",
			phone:    "+55 (41) 99504.8710",
			expected: "", // Should still be valid after cleaning
		},
		{
			name:     "Phone with extra parentheses",
			phone:    "+55 ((41)) 99504-8710",
			expected: "", // Should still be valid after cleaning
		},
		{
			name:     "Phone with multiple spaces",
			phone:    "+55   41   99504   8710",
			expected: "", // Should still be valid after cleaning
		},
		{
			name:     "Phone with leading/trailing spaces",
			phone:    " +55 41 99504-8710 ",
			expected: "", // Should still be valid after cleaning
		},
		{
			name:     "Phone with special characters",
			phone:    "+55-41-99504-8710",
			expected: "", // Should still be valid after cleaning
		},
		{
			name:     "Phone with dots and hyphens",
			phone:    "+55.41.99504-8710",
			expected: "", // Should still be valid after cleaning
		},
		{
			name:     "Phone with only digits",
			phone:    "5541995048710",
			expected: "", // Should be valid as mobile
		},
		{
			name:     "Phone with only 10 digits",
			phone:    "4133464468",
			expected: "", // Should be valid as landline
		},
		{
			name:     "Phone with invalid DDD format",
			phone:    "+55 4 99504-8710",
			expected: "invalid Brazilian phone number format",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidatePhone(tt.phone)
			if tt.expected == "" {
				if err != nil {
					t.Errorf("ValidatePhone() error = %v, expected nil", err)
				}
			} else {
				if err == nil {
					t.Errorf("ValidatePhone() expected error, got nil")
				} else if err.Error() != tt.expected {
					t.Errorf("ValidatePhone() error = %v, expected %v", err.Error(), tt.expected)
				}
			}
		})
	}
}

// TestValidatePhone_TypeValidation tests type validation for phone
func TestValidatePhone_TypeValidation(t *testing.T) {
	tests := []struct {
		name     string
		phone    interface{}
		expected string
	}{
		{
			name:     "Integer input",
			phone:    5541995048710,
			expected: "phone must be a string",
		},
		{
			name:     "Float input",
			phone:    5541995048710.0,
			expected: "phone must be a string",
		},
		{
			name:     "Boolean input",
			phone:    true,
			expected: "phone must be a string",
		},
		{
			name:     "Nil input",
			phone:    nil,
			expected: "phone must be a string",
		},
		{
			name:     "Slice input",
			phone:    []string{"+55", "41", "99504-8710"},
			expected: "phone must be a string",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidatePhone(tt.phone)
			if err == nil {
				t.Errorf("ValidatePhone() expected error, got nil")
			} else if err.Error() != tt.expected {
				t.Errorf("ValidatePhone() error = %v, expected %v", err.Error(), tt.expected)
			}
		})
	}
}
