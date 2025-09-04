// Package veritas provides string validation functions.
package veritas

import (
	"fmt"
	"unicode/utf8"
)

// String validates that a string is not empty and within length bounds.
func (v *Validator) String(str interface{}, minLength, maxLength int) error {
	strValue, ok := str.(string)
	if !ok {
		return fmt.Errorf("value must be a string")
	}

	length := utf8.RuneCountInString(strValue)

	if length < minLength {
		return fmt.Errorf("string must be at least %d characters long", minLength)
	}

	if length > maxLength {
		return fmt.Errorf("string must be at most %d characters long", maxLength)
	}

	return nil
}
