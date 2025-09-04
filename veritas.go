// Package veritas provides a comprehensive collection of input validation functions
// for various data formats including strings, numbers, Brazilian documents, and more.
package veritas

import (
	"fmt"
	"regexp"
	"strings"
)

// cleanString removes leading/trailing whitespace and converts to lowercase if specified.
func cleanString(s string, toLower bool) string {
	cleaned := strings.TrimSpace(s)
	if toLower {
		cleaned = strings.ToLower(cleaned)
	}
	return cleaned
}

// isEmpty checks if a string is empty or contains only whitespace.
func isEmpty(s string) bool {
	return strings.TrimSpace(s) == ""
}

// isNotEmpty checks if a string is not empty and contains non-whitespace characters.
func isNotEmpty(s string) bool {
	return !isEmpty(s)
}

// compileRegex compiles a regular expression pattern and returns an error if invalid.
func compileRegex(pattern string) (*regexp.Regexp, error) {
	return regexp.Compile(pattern)
}

// matchRegex checks if a string matches the given regular expression pattern.
func matchRegex(s, pattern string) (bool, error) {
	regex, err := compileRegex(pattern)
	if err != nil {
		return false, fmt.Errorf("invalid regex pattern: %w", err)
	}
	return regex.MatchString(s), nil
}
