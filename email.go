// Package veritas provides email validation functions.
package veritas

import (
	"fmt"
	"regexp"
)

// ValidateEmail validates an email address format.
func ValidateEmail(email interface{}) error {
	emailStr, ok := email.(string)
	if !ok {
		return fmt.Errorf("email must be a string")
	}

	emailStr = cleanString(emailStr, true)
	if isEmpty(emailStr) {
		return fmt.Errorf("email cannot be empty")
	}

	// Simple email regex
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	matched, err := regexp.MatchString(emailRegex, emailStr)
	if err != nil {
		return fmt.Errorf("email validation error: %w", err)
	}

	if !matched {
		return fmt.Errorf("invalid email format")
	}

	return nil
}
