// Package veritas provides numeric validation functions.
package veritas

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// ValidateNumber validates that a value is a valid number.
func ValidateNumber(num interface{}) error {
	_, err := parseNumber(num)
	if err != nil {
		return err
	}
	return nil
}

// ValidatePositive validates that a number is positive (> 0).
func ValidatePositive(num interface{}) error {
	numValue, err := parseNumber(num)
	if err != nil {
		return err
	}
	if numValue <= 0 {
		return fmt.Errorf("number must be positive")
	}
	return nil
}

// ValidateNegative validates that a number is negative (< 0).
func ValidateNegative(num interface{}) error {
	numValue, err := parseNumber(num)
	if err != nil {
		return err
	}
	if numValue >= 0 {
		return fmt.Errorf("number must be negative")
	}
	return nil
}

// ValidateEven validates that a number is even.
func ValidateEven(num interface{}) error {
	numValue, err := parseNumber(num)
	if err != nil {
		return err
	}
	if int(numValue)%2 != 0 {
		return fmt.Errorf("number must be even")
	}
	return nil
}

// ValidateBiggerThan validates that a number is bigger than the given value.
func ValidateBiggerThan(num interface{}, than float64) error {
	numValue, err := parseNumber(num)
	if err != nil {
		return err
	}
	if numValue <= than {
		return fmt.Errorf("number must be bigger than %v", than)
	}
	return nil
}

// ValidateSmallerThan validates that a number is smaller than the given value.
func ValidateSmallerThan(num interface{}, than float64) error {
	numValue, err := parseNumber(num)
	if err != nil {
		return err
	}
	if numValue >= than {
		return fmt.Errorf("number must be smaller than %v", than)
	}
	return nil
}

// ValidateBetween validates that a number is between min and max (inclusive).
func ValidateBetween(num interface{}, min, max float64) error {
	numValue, err := parseNumber(num)
	if err != nil {
		return err
	}
	if numValue < min || numValue > max {
		return fmt.Errorf("number must be between %v and %v", min, max)
	}
	return nil
}

// ValidatePrime validates that a number is a prime number.
func ValidatePrime(num interface{}) error {
	numValue, err := parseNumber(num)
	if err != nil {
		return err
	}

	// Convert to integer
	intValue := int(numValue)
	if float64(intValue) != numValue {
		return fmt.Errorf("prime number must be an integer")
	}

	if intValue < 2 {
		return fmt.Errorf("number must be at least 2 to be prime")
	}

	// Check if prime
	for i := 2; i <= int(math.Sqrt(float64(intValue))); i++ {
		if intValue%i == 0 {
			return fmt.Errorf("number is not prime")
		}
	}

	return nil
}

// parseNumber converts various number types to float64.
func parseNumber(number interface{}) (float64, error) {
	switch n := number.(type) {
	case string:
		n = strings.TrimSpace(n)
		if isEmpty(n) {
			return 0, fmt.Errorf("number cannot be empty")
		}
		return strconv.ParseFloat(n, 64)
	case int:
		return float64(n), nil
	case int64:
		return float64(n), nil
	case float32:
		return float64(n), nil
	case float64:
		return n, nil
	default:
		return 0, fmt.Errorf("unsupported number type: %T", number)
	}
}
