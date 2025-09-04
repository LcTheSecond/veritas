// Package veritas provides comprehensive unit tests for email validation functions.
package veritas

import (
	"testing"
)

// TestValidateEmail_ValidCases tests valid email addresses
func TestValidateEmail_ValidCases(t *testing.T) {
	tests := []struct {
		name     string
		email    string
		expected error
	}{
		{
			name:     "Simple valid email",
			email:    "user@example.com",
			expected: nil,
		},
		{
			name:     "Email with subdomain",
			email:    "user@mail.example.com",
			expected: nil,
		},
		{
			name:     "Email with numbers",
			email:    "user123@example123.com",
			expected: nil,
		},
		{
			name:     "Email with dots in local part",
			email:    "user.name@example.com",
			expected: nil,
		},
		{
			name:     "Email with plus sign",
			email:    "user+tag@example.com",
			expected: nil,
		},
		{
			name:     "Email with hyphens",
			email:    "user-name@example-domain.com",
			expected: nil,
		},
		{
			name:     "Email with underscores",
			email:    "user_name@example.com",
			expected: nil,
		},
		{
			name:     "Email with percent sign",
			email:    "user%name@example.com",
			expected: nil,
		},
		{
			name:     "Email with mixed case",
			email:    "User.Name@Example.COM",
			expected: nil,
		},
		{
			name:     "Email with long domain",
			email:    "user@very-long-domain-name.example.com",
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateEmail(tt.email)
			if err != tt.expected {
				t.Errorf("ValidateEmail() error = %v, expected %v", err, tt.expected)
			}
		})
	}
}

// TestValidateEmail_InvalidFormats tests invalid email formats
func TestValidateEmail_InvalidFormats(t *testing.T) {
	tests := []struct {
		name     string
		email    string
		expected string
	}{
		{
			name:     "Missing @ symbol",
			email:    "userexample.com",
			expected: "invalid email format",
		},
		{
			name:     "Multiple @ symbols",
			email:    "user@@example.com",
			expected: "invalid email format",
		},
		{
			name:     "Missing domain",
			email:    "user@",
			expected: "invalid email format",
		},
		{
			name:     "Missing local part",
			email:    "@example.com",
			expected: "invalid email format",
		},
		{
			name:     "Missing TLD",
			email:    "user@example",
			expected: "invalid email format",
		},
		{
			name:     "Invalid characters in domain",
			email:    "user@example..com",
			expected: "invalid email format",
		},
		{
			name:     "Space in email",
			email:    "user name@example.com",
			expected: "invalid email format",
		},
		{
			name:     "Invalid TLD length",
			email:    "user@example.c",
			expected: "invalid email format",
		},
		{
			name:     "Special characters in domain",
			email:    "user@example#.com",
			expected: "invalid email format",
		},
		{
			name:     "Consecutive dots in local part",
			email:    "user..name@example.com",
			expected: "invalid email format",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateEmail(tt.email)
			if err == nil {
				t.Errorf("ValidateEmail() expected error, got nil")
			} else if err.Error() != tt.expected {
				t.Errorf("ValidateEmail() error = %v, expected %v", err.Error(), tt.expected)
			}
		})
	}
}

// TestValidateEmail_EdgeCases tests edge cases for email validation
func TestValidateEmail_EdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		email    string
		expected string
	}{
		{
			name:     "Empty string",
			email:    "",
			expected: "email cannot be empty",
		},
		{
			name:     "Only whitespace",
			email:    "   ",
			expected: "email cannot be empty",
		},
		{
			name:     "Email with leading whitespace",
			email:    " user@example.com",
			expected: "", // Should be cleaned and valid
		},
		{
			name:     "Email with trailing whitespace",
			email:    "user@example.com ",
			expected: "", // Should be cleaned and valid
		},
		{
			name:     "Email with mixed case and whitespace",
			email:    " User.Name@Example.COM ",
			expected: "", // Should be cleaned and valid
		},
		{
			name:     "Very long local part",
			email:    "verylonglocalpartthatmightexceedlimits@example.com",
			expected: "", // Should still be valid
		},
		{
			name:     "Email with international domain",
			email:    "user@example.co.uk",
			expected: "",
		},
		{
			name:     "Email with numbers in TLD",
			email:    "user@example.com123",
			expected: "invalid email format",
		},
		{
			name:     "Email with special characters in TLD",
			email:    "user@example.c-m",
			expected: "invalid email format",
		},
		{
			name:     "Email with single character local part",
			email:    "a@example.com",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateEmail(tt.email)
			if tt.expected == "" {
				if err != nil {
					t.Errorf("ValidateEmail() error = %v, expected nil", err)
				}
			} else {
				if err == nil {
					t.Errorf("ValidateEmail() expected error, got nil")
				} else if err.Error() != tt.expected {
					t.Errorf("ValidateEmail() error = %v, expected %v", err.Error(), tt.expected)
				}
			}
		})
	}
}

// TestValidateEmail_TypeValidation tests type validation for email
func TestValidateEmail_TypeValidation(t *testing.T) {
	tests := []struct {
		name     string
		email    interface{}
		expected string
	}{
		{
			name:     "Integer input",
			email:    123,
			expected: "email must be a string",
		},
		{
			name:     "Float input",
			email:    123.45,
			expected: "email must be a string",
		},
		{
			name:     "Boolean input",
			email:    true,
			expected: "email must be a string",
		},
		{
			name:     "Nil input",
			email:    nil,
			expected: "email must be a string",
		},
		{
			name:     "Slice input",
			email:    []string{"user", "@", "example.com"},
			expected: "email must be a string",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateEmail(tt.email)
			if err == nil {
				t.Errorf("ValidateEmail() expected error, got nil")
			} else if err.Error() != tt.expected {
				t.Errorf("ValidateEmail() error = %v, expected %v", err.Error(), tt.expected)
			}
		})
	}
}
