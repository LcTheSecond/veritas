// Package veritas provides comprehensive unit tests for CPF validation functions.
package veritas

import (
	"testing"
)

// TestValidateCPF_ValidCases tests valid CPF numbers with various formats
func TestValidateCPF_ValidCases(t *testing.T) {
	tests := []struct {
		name     string
		cpf      string
		expected error
	}{
		{
			name:     "Valid CPF with dots and hyphens",
			cpf:      "111.444.777-35",
			expected: nil,
		},
		{
			name:     "Valid CPF without formatting",
			cpf:      "11144477735",
			expected: nil,
		},
		{
			name:     "Valid CPF with spaces",
			cpf:      "111 444 777 35",
			expected: nil,
		},
		{
			name:     "Another valid CPF",
			cpf:      "123.456.789-09",
			expected: nil,
		},
		{
			name:     "Valid CPF with mixed formatting",
			cpf:      "123.456.789.09",
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateCPF(tt.cpf)
			if err != tt.expected {
				t.Errorf("ValidateCPF() error = %v, expected %v", err, tt.expected)
			}
		})
	}
}

// TestValidateCPF_InvalidFormats tests invalid CPF formats
func TestValidateCPF_InvalidFormats(t *testing.T) {
	tests := []struct {
		name     string
		cpf      string
		expected string
	}{
		{
			name:     "CPF too short",
			cpf:      "1234567890",
			expected: "CPF must have exactly 11 digits",
		},
		{
			name:     "CPF too long",
			cpf:      "123456789012",
			expected: "CPF must have exactly 11 digits",
		},
		{
			name:     "CPF with letters",
			cpf:      "1234567890a",
			expected: "CPF must have exactly 11 digits",
		},
		{
			name:     "CPF with special characters only",
			cpf:      "abc.def.ghi-jk",
			expected: "CPF must have exactly 11 digits",
		},
		{
			name:     "Empty string",
			cpf:      "",
			expected: "CPF must have exactly 11 digits",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateCPF(tt.cpf)
			if err == nil {
				t.Errorf("ValidateCPF() expected error, got nil")
			} else if err.Error() != tt.expected {
				t.Errorf("ValidateCPF() error = %v, expected %v", err.Error(), tt.expected)
			}
		})
	}
}

// TestValidateCPF_InvalidCheckDigits tests CPFs with invalid check digits
func TestValidateCPF_InvalidCheckDigits(t *testing.T) {
	tests := []struct {
		name     string
		cpf      string
		expected string
	}{
		{
			name:     "Invalid first check digit",
			cpf:      "111.444.777-45",
			expected: "invalid CPF check digits",
		},
		{
			name:     "Invalid second check digit",
			cpf:      "111.444.777-36",
			expected: "invalid CPF check digits",
		},
		{
			name:     "Both check digits invalid",
			cpf:      "111.444.777-99",
			expected: "invalid CPF check digits",
		},
		{
			name:     "Wrong check digits for another CPF",
			cpf:      "123.456.789-00",
			expected: "invalid CPF check digits",
		},
		{
			name:     "Check digits swapped",
			cpf:      "111.444.777-53",
			expected: "invalid CPF check digits",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateCPF(tt.cpf)
			if err == nil {
				t.Errorf("ValidateCPF() expected error, got nil")
			} else if err.Error() != tt.expected {
				t.Errorf("ValidateCPF() error = %v, expected %v", err.Error(), tt.expected)
			}
		})
	}
}

// TestValidateCPF_EdgeCases tests edge cases for CPF validation
func TestValidateCPF_EdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		cpf      string
		expected string
	}{
		{
			name:     "All zeros",
			cpf:      "000.000.000-00",
			expected: "CPF cannot be a sequence of identical digits",
		},
		{
			name:     "All ones",
			cpf:      "111.111.111-11",
			expected: "CPF cannot be a sequence of identical digits",
		},
		{
			name:     "All nines",
			cpf:      "999.999.999-99",
			expected: "CPF cannot be a sequence of identical digits",
		},
		{
			name:     "CPF with only spaces",
			cpf:      "   ",
			expected: "CPF must have exactly 11 digits",
		},
		{
			name:     "CPF with mixed separators",
			cpf:      "111.444-777.35",
			expected: "", // Should still be valid after cleaning
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateCPF(tt.cpf)
			if tt.expected == "" {
				if err != nil {
					t.Errorf("ValidateCPF() error = %v, expected nil", err)
				}
			} else {
				if err == nil {
					t.Errorf("ValidateCPF() expected error, got nil")
				} else if err.Error() != tt.expected {
					t.Errorf("ValidateCPF() error = %v, expected %v", err.Error(), tt.expected)
				}
			}
		})
	}
}

// TestValidateCPF_TypeValidation tests type validation for CPF
func TestValidateCPF_TypeValidation(t *testing.T) {
	tests := []struct {
		name     string
		cpf      interface{}
		expected string
	}{
		{
			name:     "Integer input",
			cpf:      11144477735,
			expected: "CPF must be a string",
		},
		{
			name:     "Float input",
			cpf:      11144477735.0,
			expected: "CPF must be a string",
		},
		{
			name:     "Boolean input",
			cpf:      true,
			expected: "CPF must be a string",
		},
		{
			name:     "Nil input",
			cpf:      nil,
			expected: "CPF must be a string",
		},
		{
			name:     "Slice input",
			cpf:      []string{"111", "444", "777", "35"},
			expected: "CPF must be a string",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateCPF(tt.cpf)
			if err == nil {
				t.Errorf("ValidateCPF() expected error, got nil")
			} else if err.Error() != tt.expected {
				t.Errorf("ValidateCPF() error = %v, expected %v", err.Error(), tt.expected)
			}
		})
	}
}
