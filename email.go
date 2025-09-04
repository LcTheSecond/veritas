// Package veritas provides email validation functions.
package veritas

import (
	"fmt"
	"regexp"
)

// Email validates an email address format.
func (v *Validator) Email(email interface{}) error {
	emailStr, ok := email.(string)
	if !ok {
		return fmt.Errorf("email must be a string")
	}

	emailStr = v.CleanString(emailStr, true)
	if v.IsEmpty(emailStr) {
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
