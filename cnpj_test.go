// Package veritas provides comprehensive unit tests for CNPJ validation functions.
package veritas

import (
	"testing"
)

// TestValidateCNPJ_ValidCases tests valid CNPJ numbers with various formats
func TestValidateCNPJ_ValidCases(t *testing.T) {
	tests := []struct {
		name     string
		cnpj     string
		expected error
	}{
		{
			name:     "Valid CNPJ with dots, slashes and hyphens",
			cnpj:     "11.222.333/0001-81",
			expected: nil,
		},
		{
			name:     "Valid CNPJ without formatting",
			cnpj:     "11222333000181",
			expected: nil,
		},
		{
			name:     "Valid CNPJ with spaces",
			cnpj:     "11 222 333 0001 81",
			expected: nil,
		},
		{
			name:     "Another valid CNPJ",
			cnpj:     "12.345.678/0001-95",
			expected: nil,
		},
		{
			name:     "Valid CNPJ with mixed formatting",
			cnpj:     "12.345.678.0001.95",
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateCNPJ(tt.cnpj)
			if err != tt.expected {
				t.Errorf("ValidateCNPJ() error = %v, expected %v", err, tt.expected)
			}
		})
	}
}

// TestValidateCNPJ_InvalidFormats tests invalid CNPJ formats
func TestValidateCNPJ_InvalidFormats(t *testing.T) {
	tests := []struct {
		name     string
		cnpj     string
		expected string
	}{
		{
			name:     "CNPJ too short",
			cnpj:     "1234567890123",
			expected: "CNPJ must have exactly 14 digits",
		},
		{
			name:     "CNPJ too long",
			cnpj:     "123456789012345",
			expected: "CNPJ must have exactly 14 digits",
		},
		{
			name:     "CNPJ with letters",
			cnpj:     "1234567890123a",
			expected: "CNPJ must have exactly 14 digits",
		},
		{
			name:     "CNPJ with special characters only",
			cnpj:     "ab.cde.fgh/ijkl-mn",
			expected: "CNPJ must have exactly 14 digits",
		},
		{
			name:     "Empty string",
			cnpj:     "",
			expected: "CNPJ must have exactly 14 digits",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateCNPJ(tt.cnpj)
			if err == nil {
				t.Errorf("ValidateCNPJ() expected error, got nil")
			} else if err.Error() != tt.expected {
				t.Errorf("ValidateCNPJ() error = %v, expected %v", err.Error(), tt.expected)
			}
		})
	}
}

// TestValidateCNPJ_InvalidCheckDigits tests CNPJs with invalid check digits
func TestValidateCNPJ_InvalidCheckDigits(t *testing.T) {
	tests := []struct {
		name     string
		cnpj     string
		expected string
	}{
		{
			name:     "Invalid first check digit",
			cnpj:     "11.222.333/0001-91",
			expected: "invalid CNPJ check digits",
		},
		{
			name:     "Invalid second check digit",
			cnpj:     "11.222.333/0001-82",
			expected: "invalid CNPJ check digits",
		},
		{
			name:     "Both check digits invalid",
			cnpj:     "11.222.333/0001-99",
			expected: "invalid CNPJ check digits",
		},
		{
			name:     "Wrong check digits for another CNPJ",
			cnpj:     "12.345.678/0001-00",
			expected: "invalid CNPJ check digits",
		},
		{
			name:     "Check digits swapped",
			cnpj:     "11.222.333/0001-18",
			expected: "invalid CNPJ check digits",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateCNPJ(tt.cnpj)
			if err == nil {
				t.Errorf("ValidateCNPJ() expected error, got nil")
			} else if err.Error() != tt.expected {
				t.Errorf("ValidateCNPJ() error = %v, expected %v", err.Error(), tt.expected)
			}
		})
	}
}

// TestValidateCNPJ_EdgeCases tests edge cases for CNPJ validation
func TestValidateCNPJ_EdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		cnpj     string
		expected string
	}{
		{
			name:     "All zeros",
			cnpj:     "00.000.000/0000-00",
			expected: "CNPJ cannot be a sequence of identical digits",
		},
		{
			name:     "All ones",
			cnpj:     "11.111.111/1111-11",
			expected: "CNPJ cannot be a sequence of identical digits",
		},
		{
			name:     "All nines",
			cnpj:     "99.999.999/9999-99",
			expected: "CNPJ cannot be a sequence of identical digits",
		},
		{
			name:     "CNPJ with only spaces",
			cnpj:     "   ",
			expected: "CNPJ must have exactly 14 digits",
		},
		{
			name:     "CNPJ with mixed separators",
			cnpj:     "11.222-333/0001.81",
			expected: "", // Should still be valid after cleaning
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateCNPJ(tt.cnpj)
			if tt.expected == "" {
				if err != nil {
					t.Errorf("ValidateCNPJ() error = %v, expected nil", err)
				}
			} else {
				if err == nil {
					t.Errorf("ValidateCNPJ() expected error, got nil")
				} else if err.Error() != tt.expected {
					t.Errorf("ValidateCNPJ() error = %v, expected %v", err.Error(), tt.expected)
				}
			}
		})
	}
}

// TestValidateCNPJ_TypeValidation tests type validation for CNPJ
func TestValidateCNPJ_TypeValidation(t *testing.T) {
	tests := []struct {
		name     string
		cnpj     interface{}
		expected string
	}{
		{
			name:     "Integer input",
			cnpj:     11222333000181,
			expected: "CNPJ must be a string",
		},
		{
			name:     "Float input",
			cnpj:     11222333000181.0,
			expected: "CNPJ must be a string",
		},
		{
			name:     "Boolean input",
			cnpj:     true,
			expected: "CNPJ must be a string",
		},
		{
			name:     "Nil input",
			cnpj:     nil,
			expected: "CNPJ must be a string",
		},
		{
			name:     "Slice input",
			cnpj:     []string{"11", "222", "333", "0001", "81"},
			expected: "CNPJ must be a string",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateCNPJ(tt.cnpj)
			if err == nil {
				t.Errorf("ValidateCNPJ() expected error, got nil")
			} else if err.Error() != tt.expected {
				t.Errorf("ValidateCNPJ() error = %v, expected %v", err.Error(), tt.expected)
			}
		})
	}
}
