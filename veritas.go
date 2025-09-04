// Package veritas provides a comprehensive collection of input validation functions
// for various data formats including strings, numbers, Brazilian documents, and more.
package veritas

import (
	"fmt"
	"regexp"
	"strings"
)

// Validator represents the main validation struct that provides access to all validation methods.
type Validator struct {
	// Future: could include configuration options, custom validators, etc.
}

// New creates a new instance of the Validator.
func New() *Validator {
	return &Validator{}
}

// Validate performs validation on a value using the provided validation function.
// It returns a ValidationError if validation fails, nil otherwise.
func (v *Validator) Validate(field string, value interface{}, validator func(interface{}) error) *ValidationError {
	if err := validator(value); err != nil {
		return NewValidationError(field, ErrorTypeInvalid, err.Error(), value)
	}
	return nil
}

// ValidateMultiple performs multiple validations and returns all validation errors.
func (v *Validator) ValidateMultiple(validations ...func() *ValidationError) []*ValidationError {
	var errors []*ValidationError
	for _, validation := range validations {
		if err := validation(); err != nil {
			errors = append(errors, err)
		}
	}
	return errors
}

// HasErrors checks if there are any validation errors.
func (v *Validator) HasErrors(errors []*ValidationError) bool {
	return len(errors) > 0
}

// CleanString removes leading/trailing whitespace and converts to lowercase if specified.
func (v *Validator) CleanString(s string, toLower bool) string {
	cleaned := strings.TrimSpace(s)
	if toLower {
		cleaned = strings.ToLower(cleaned)
	}
	return cleaned
}

// IsEmpty checks if a string is empty or contains only whitespace.
func (v *Validator) IsEmpty(s string) bool {
	return strings.TrimSpace(s) == ""
}

// IsNotEmpty checks if a string is not empty and contains non-whitespace characters.
func (v *Validator) IsNotEmpty(s string) bool {
	return !v.IsEmpty(s)
}

// CompileRegex compiles a regular expression pattern and returns an error if invalid.
func (v *Validator) CompileRegex(pattern string) (*regexp.Regexp, error) {
	return regexp.Compile(pattern)
}

// MatchRegex checks if a string matches the given regular expression pattern.
func (v *Validator) MatchRegex(s, pattern string) (bool, error) {
	regex, err := v.CompileRegex(pattern)
	if err != nil {
		return false, fmt.Errorf("invalid regex pattern: %w", err)
	}
	return regex.MatchString(s), nil
}
