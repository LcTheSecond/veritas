// Package veritas provides numeric validation functions.
package veritas

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// IsNumber validates that a value is a valid number.
func (v *Validator) IsNumber(num interface{}) error {
	_, err := v.parseNumber(num)
	if err != nil {
		return err
	}
	return nil
}

// IsPositive validates that a number is positive (> 0).
func (v *Validator) IsPositive(num interface{}) error {
	numValue, err := v.parseNumber(num)
	if err != nil {
		return err
	}
	if numValue <= 0 {
		return fmt.Errorf("number must be positive")
	}
	return nil
}

// IsNegative validates that a number is negative (< 0).
func (v *Validator) IsNegative(num interface{}) error {
	numValue, err := v.parseNumber(num)
	if err != nil {
		return err
	}
	if numValue >= 0 {
		return fmt.Errorf("number must be negative")
	}
	return nil
}

// IsEven validates that a number is even.
func (v *Validator) IsEven(num interface{}) error {
	numValue, err := v.parseNumber(num)
	if err != nil {
		return err
	}
	if int(numValue)%2 != 0 {
		return fmt.Errorf("number must be even")
	}
	return nil
}

// BiggerThan validates that a number is bigger than the given value.
func (v *Validator) BiggerThan(num interface{}, than float64) error {
	numValue, err := v.parseNumber(num)
	if err != nil {
		return err
	}
	if numValue <= than {
		return fmt.Errorf("number must be bigger than %v", than)
	}
	return nil
}

// SmallerThan validates that a number is smaller than the given value.
func (v *Validator) SmallerThan(num interface{}, than float64) error {
	numValue, err := v.parseNumber(num)
	if err != nil {
		return err
	}
	if numValue >= than {
		return fmt.Errorf("number must be smaller than %v", than)
	}
	return nil
}

// Between validates that a number is between min and max (inclusive).
func (v *Validator) Between(num interface{}, min, max float64) error {
	numValue, err := v.parseNumber(num)
	if err != nil {
		return err
	}
	if numValue < min || numValue > max {
		return fmt.Errorf("number must be between %v and %v", min, max)
	}
	return nil
}

// IsPrime validates that a number is a prime number.
func (v *Validator) IsPrime(num interface{}) error {
	numValue, err := v.parseNumber(num)
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
func (v *Validator) parseNumber(number interface{}) (float64, error) {
	switch n := number.(type) {
	case string:
		n = strings.TrimSpace(n)
		if v.IsEmpty(n) {
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
