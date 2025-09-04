// Package veritas provides comprehensive unit tests for numeric validation functions.
package veritas

import (
	"testing"
)

// TestValidateNumber_ValidCases tests valid number inputs
func TestValidateNumber_ValidCases(t *testing.T) {
	tests := []struct {
		name     string
		number   interface{}
		expected error
	}{
		{
			name:     "Valid integer",
			number:   123,
			expected: nil,
		},
		{
			name:     "Valid float64",
			number:   123.45,
			expected: nil,
		},
		{
			name:     "Valid float32",
			number:   float32(123.45),
			expected: nil,
		},
		{
			name:     "Valid int64",
			number:   int64(123456789),
			expected: nil,
		},
		{
			name:     "Valid string number",
			number:   "123.45",
			expected: nil,
		},
		{
			name:     "Valid string integer",
			number:   "123",
			expected: nil,
		},
		{
			name:     "Valid negative number",
			number:   -123.45,
			expected: nil,
		},
		{
			name:     "Valid zero",
			number:   0,
			expected: nil,
		},
		{
			name:     "Valid scientific notation string",
			number:   "1.23e+02",
			expected: nil,
		},
		{
			name:     "Valid string with spaces",
			number:   " 123.45 ",
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateNumber(tt.number)
			if err != tt.expected {
				t.Errorf("ValidateNumber() error = %v, expected %v", err, tt.expected)
			}
		})
	}
}

// TestValidateNumber_InvalidCases tests invalid number inputs
func TestValidateNumber_InvalidCases(t *testing.T) {
	tests := []struct {
		name     string
		number   interface{}
		expected string
	}{
		{
			name:     "Invalid string",
			number:   "not a number",
			expected: "strconv.ParseFloat: parsing \"not a number\": invalid syntax",
		},
		{
			name:     "Empty string",
			number:   "",
			expected: "number cannot be empty",
		},
		{
			name:     "String with only spaces",
			number:   "   ",
			expected: "number cannot be empty",
		},
		{
			name:     "Boolean input",
			number:   true,
			expected: "unsupported number type: bool",
		},
		{
			name:     "Nil input",
			number:   nil,
			expected: "unsupported number type: <nil>",
		},
		{
			name:     "Slice input",
			number:   []int{1, 2, 3},
			expected: "unsupported number type: []int",
		},
		{
			name:     "Map input",
			number:   map[string]int{"key": 1},
			expected: "unsupported number type: map[string]int",
		},
		{
			name:     "String with letters and numbers",
			number:   "123abc",
			expected: "strconv.ParseFloat: parsing \"123abc\": invalid syntax",
		},
		{
			name:     "String with special characters",
			number:   "123.45.67",
			expected: "strconv.ParseFloat: parsing \"123.45.67\": invalid syntax",
		},
		{
			name:     "String with currency symbol",
			number:   "$123.45",
			expected: "strconv.ParseFloat: parsing \"$123.45\": invalid syntax",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateNumber(tt.number)
			if err == nil {
				t.Errorf("ValidateNumber() expected error, got nil")
			} else if err.Error() != tt.expected {
				t.Errorf("ValidateNumber() error = %v, expected %v", err.Error(), tt.expected)
			}
		})
	}
}

// TestValidatePositive_ValidCases tests valid positive numbers
func TestValidatePositive_ValidCases(t *testing.T) {
	tests := []struct {
		name     string
		number   interface{}
		expected error
	}{
		{
			name:     "Positive integer",
			number:   123,
			expected: nil,
		},
		{
			name:     "Positive float",
			number:   123.45,
			expected: nil,
		},
		{
			name:     "Positive string number",
			number:   "123.45",
			expected: nil,
		},
		{
			name:     "Very small positive number",
			number:   0.000001,
			expected: nil,
		},
		{
			name:     "Large positive number",
			number:   999999999.99,
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidatePositive(tt.number)
			if err != tt.expected {
				t.Errorf("ValidatePositive() error = %v, expected %v", err, tt.expected)
			}
		})
	}
}

// TestValidatePositive_InvalidCases tests invalid positive numbers
func TestValidatePositive_InvalidCases(t *testing.T) {
	tests := []struct {
		name     string
		number   interface{}
		expected string
	}{
		{
			name:     "Zero",
			number:   0,
			expected: "number must be positive",
		},
		{
			name:     "Negative number",
			number:   -123.45,
			expected: "number must be positive",
		},
		{
			name:     "Negative string number",
			number:   "-123.45",
			expected: "number must be positive",
		},
		{
			name:     "Very small negative number",
			number:   -0.000001,
			expected: "number must be positive",
		},
		{
			name:     "Large negative number",
			number:   -999999999.99,
			expected: "number must be positive",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidatePositive(tt.number)
			if err == nil {
				t.Errorf("ValidatePositive() expected error, got nil")
			} else if err.Error() != tt.expected {
				t.Errorf("ValidatePositive() error = %v, expected %v", err.Error(), tt.expected)
			}
		})
	}
}

// TestValidateNegative_ValidCases tests valid negative numbers
func TestValidateNegative_ValidCases(t *testing.T) {
	tests := []struct {
		name     string
		number   interface{}
		expected error
	}{
		{
			name:     "Negative integer",
			number:   -123,
			expected: nil,
		},
		{
			name:     "Negative float",
			number:   -123.45,
			expected: nil,
		},
		{
			name:     "Negative string number",
			number:   "-123.45",
			expected: nil,
		},
		{
			name:     "Very small negative number",
			number:   -0.000001,
			expected: nil,
		},
		{
			name:     "Large negative number",
			number:   -999999999.99,
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateNegative(tt.number)
			if err != tt.expected {
				t.Errorf("ValidateNegative() error = %v, expected %v", err, tt.expected)
			}
		})
	}
}

// TestValidateNegative_InvalidCases tests invalid negative numbers
func TestValidateNegative_InvalidCases(t *testing.T) {
	tests := []struct {
		name     string
		number   interface{}
		expected string
	}{
		{
			name:     "Zero",
			number:   0,
			expected: "number must be negative",
		},
		{
			name:     "Positive number",
			number:   123.45,
			expected: "number must be negative",
		},
		{
			name:     "Positive string number",
			number:   "123.45",
			expected: "number must be negative",
		},
		{
			name:     "Very small positive number",
			number:   0.000001,
			expected: "number must be negative",
		},
		{
			name:     "Large positive number",
			number:   999999999.99,
			expected: "number must be negative",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateNegative(tt.number)
			if err == nil {
				t.Errorf("ValidateNegative() expected error, got nil")
			} else if err.Error() != tt.expected {
				t.Errorf("ValidateNegative() error = %v, expected %v", err.Error(), tt.expected)
			}
		})
	}
}

// TestValidateEven_ValidCases tests valid even numbers
func TestValidateEven_ValidCases(t *testing.T) {
	tests := []struct {
		name     string
		number   interface{}
		expected error
	}{
		{
			name:     "Even integer",
			number:   2,
			expected: nil,
		},
		{
			name:     "Even negative integer",
			number:   -4,
			expected: nil,
		},
		{
			name:     "Even string number",
			number:   "6",
			expected: nil,
		},
		{
			name:     "Large even number",
			number:   1000000,
			expected: nil,
		},
		{
			name:     "Zero",
			number:   0,
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateEven(tt.number)
			if err != tt.expected {
				t.Errorf("ValidateEven() error = %v, expected %v", err, tt.expected)
			}
		})
	}
}

// TestValidateEven_InvalidCases tests invalid even numbers
func TestValidateEven_InvalidCases(t *testing.T) {
	tests := []struct {
		name     string
		number   interface{}
		expected string
	}{
		{
			name:     "Odd integer",
			number:   3,
			expected: "number must be even",
		},
		{
			name:     "Odd negative integer",
			number:   -5,
			expected: "number must be even",
		},
		{
			name:     "Odd string number",
			number:   "7",
			expected: "number must be even",
		},
		{
			name:     "Large odd number",
			number:   1000001,
			expected: "number must be even",
		},
		{
			name:     "Float number",
			number:   2.5,
			expected: "number must be even",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateEven(tt.number)
			if err == nil {
				t.Errorf("ValidateEven() expected error, got nil")
			} else if err.Error() != tt.expected {
				t.Errorf("ValidateEven() error = %v, expected %v", err.Error(), tt.expected)
			}
		})
	}
}

// TestValidatePrime_ValidCases tests valid prime numbers
func TestValidatePrime_ValidCases(t *testing.T) {
	tests := []struct {
		name     string
		number   interface{}
		expected error
	}{
		{
			name:     "Smallest prime",
			number:   2,
			expected: nil,
		},
		{
			name:     "Small prime",
			number:   3,
			expected: nil,
		},
		{
			name:     "Medium prime",
			number:   17,
			expected: nil,
		},
		{
			name:     "Large prime",
			number:   97,
			expected: nil,
		},
		{
			name:     "Prime as string",
			number:   "13",
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidatePrime(tt.number)
			if err != tt.expected {
				t.Errorf("ValidatePrime() error = %v, expected %v", err, tt.expected)
			}
		})
	}
}

// TestValidatePrime_InvalidCases tests invalid prime numbers
func TestValidatePrime_InvalidCases(t *testing.T) {
	tests := []struct {
		name     string
		number   interface{}
		expected string
	}{
		{
			name:     "One",
			number:   1,
			expected: "number must be at least 2 to be prime",
		},
		{
			name:     "Zero",
			number:   0,
			expected: "number must be at least 2 to be prime",
		},
		{
			name:     "Negative number",
			number:   -5,
			expected: "number must be at least 2 to be prime",
		},
		{
			name:     "Composite number",
			number:   4,
			expected: "number is not prime",
		},
		{
			name:     "Large composite number",
			number:   100,
			expected: "number is not prime",
		},
		{
			name:     "Float number",
			number:   2.5,
			expected: "prime number must be an integer",
		},
		{
			name:     "String float",
			number:   "3.14",
			expected: "prime number must be an integer",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidatePrime(tt.number)
			if err == nil {
				t.Errorf("ValidatePrime() expected error, got nil")
			} else if err.Error() != tt.expected {
				t.Errorf("ValidatePrime() error = %v, expected %v", err.Error(), tt.expected)
			}
		})
	}
}

// TestValidateBiggerThan_ValidCases tests valid bigger than comparisons
func TestValidateBiggerThan_ValidCases(t *testing.T) {
	tests := []struct {
		name     string
		number   interface{}
		than     float64
		expected error
	}{
		{
			name:     "Integer bigger than threshold",
			number:   10,
			than:     5,
			expected: nil,
		},
		{
			name:     "Float bigger than threshold",
			number:   10.5,
			than:     10.0,
			expected: nil,
		},
		{
			name:     "String number bigger than threshold",
			number:   "15",
			than:     10,
			expected: nil,
		},
		{
			name:     "Negative number bigger than more negative threshold",
			number:   -5,
			than:     -10,
			expected: nil,
		},
		{
			name:     "Zero bigger than negative threshold",
			number:   0,
			than:     -1,
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateBiggerThan(tt.number, tt.than)
			if err != tt.expected {
				t.Errorf("ValidateBiggerThan() error = %v, expected %v", err, tt.expected)
			}
		})
	}
}

// TestValidateBiggerThan_InvalidCases tests invalid bigger than comparisons
func TestValidateBiggerThan_InvalidCases(t *testing.T) {
	tests := []struct {
		name     string
		number   interface{}
		than     float64
		expected string
	}{
		{
			name:     "Number equal to threshold",
			number:   5,
			than:     5,
			expected: "number must be bigger than 5",
		},
		{
			name:     "Number smaller than threshold",
			number:   3,
			than:     5,
			expected: "number must be bigger than 5",
		},
		{
			name:     "String number smaller than threshold",
			number:   "2",
			than:     5,
			expected: "number must be bigger than 5",
		},
		{
			name:     "Negative number smaller than threshold",
			number:   -10,
			than:     -5,
			expected: "number must be bigger than -5",
		},
		{
			name:     "Zero smaller than positive threshold",
			number:   0,
			than:     1,
			expected: "number must be bigger than 1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateBiggerThan(tt.number, tt.than)
			if err == nil {
				t.Errorf("ValidateBiggerThan() expected error, got nil")
			} else if err.Error() != tt.expected {
				t.Errorf("ValidateBiggerThan() error = %v, expected %v", err.Error(), tt.expected)
			}
		})
	}
}

// TestValidateSmallerThan_ValidCases tests valid smaller than comparisons
func TestValidateSmallerThan_ValidCases(t *testing.T) {
	tests := []struct {
		name     string
		number   interface{}
		than     float64
		expected error
	}{
		{
			name:     "Integer smaller than threshold",
			number:   3,
			than:     5,
			expected: nil,
		},
		{
			name:     "Float smaller than threshold",
			number:   9.5,
			than:     10.0,
			expected: nil,
		},
		{
			name:     "String number smaller than threshold",
			number:   "5",
			than:     10,
			expected: nil,
		},
		{
			name:     "Negative number smaller than less negative threshold",
			number:   -10,
			than:     -5,
			expected: nil,
		},
		{
			name:     "Negative number smaller than positive threshold",
			number:   -1,
			than:     0,
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateSmallerThan(tt.number, tt.than)
			if err != tt.expected {
				t.Errorf("ValidateSmallerThan() error = %v, expected %v", err, tt.expected)
			}
		})
	}
}

// TestValidateSmallerThan_InvalidCases tests invalid smaller than comparisons
func TestValidateSmallerThan_InvalidCases(t *testing.T) {
	tests := []struct {
		name     string
		number   interface{}
		than     float64
		expected string
	}{
		{
			name:     "Number equal to threshold",
			number:   5,
			than:     5,
			expected: "number must be smaller than 5",
		},
		{
			name:     "Number bigger than threshold",
			number:   7,
			than:     5,
			expected: "number must be smaller than 5",
		},
		{
			name:     "String number bigger than threshold",
			number:   "8",
			than:     5,
			expected: "number must be smaller than 5",
		},
		{
			name:     "Negative number bigger than threshold",
			number:   -3,
			than:     -5,
			expected: "number must be smaller than -5",
		},
		{
			name:     "Positive number bigger than negative threshold",
			number:   1,
			than:     -1,
			expected: "number must be smaller than -1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateSmallerThan(tt.number, tt.than)
			if err == nil {
				t.Errorf("ValidateSmallerThan() expected error, got nil")
			} else if err.Error() != tt.expected {
				t.Errorf("ValidateSmallerThan() error = %v, expected %v", err.Error(), tt.expected)
			}
		})
	}
}

// TestValidateBetween_ValidCases tests valid between comparisons
func TestValidateBetween_ValidCases(t *testing.T) {
	tests := []struct {
		name     string
		number   interface{}
		min      float64
		max      float64
		expected error
	}{
		{
			name:     "Number between bounds",
			number:   5,
			min:      1,
			max:      10,
			expected: nil,
		},
		{
			name:     "Number at minimum bound",
			number:   1,
			min:      1,
			max:      10,
			expected: nil,
		},
		{
			name:     "Number at maximum bound",
			number:   10,
			min:      1,
			max:      10,
			expected: nil,
		},
		{
			name:     "Float between bounds",
			number:   5.5,
			min:      1.0,
			max:      10.0,
			expected: nil,
		},
		{
			name:     "String number between bounds",
			number:   "5",
			min:      1,
			max:      10,
			expected: nil,
		},
		{
			name:     "Negative number between negative bounds",
			number:   -5,
			min:      -10,
			max:      -1,
			expected: nil,
		},
		{
			name:     "Zero between bounds",
			number:   0,
			min:      -1,
			max:      1,
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateBetween(tt.number, tt.min, tt.max)
			if err != tt.expected {
				t.Errorf("ValidateBetween() error = %v, expected %v", err, tt.expected)
			}
		})
	}
}

// TestValidateBetween_InvalidCases tests invalid between comparisons
func TestValidateBetween_InvalidCases(t *testing.T) {
	tests := []struct {
		name     string
		number   interface{}
		min      float64
		max      float64
		expected string
	}{
		{
			name:     "Number below minimum",
			number:   0,
			min:      1,
			max:      10,
			expected: "number must be between 1 and 10",
		},
		{
			name:     "Number above maximum",
			number:   11,
			min:      1,
			max:      10,
			expected: "number must be between 1 and 10",
		},
		{
			name:     "String number below minimum",
			number:   "0",
			min:      1,
			max:      10,
			expected: "number must be between 1 and 10",
		},
		{
			name:     "String number above maximum",
			number:   "11",
			min:      1,
			max:      10,
			expected: "number must be between 1 and 10",
		},
		{
			name:     "Negative number below negative minimum",
			number:   -11,
			min:      -10,
			max:      -1,
			expected: "number must be between -10 and -1",
		},
		{
			name:     "Negative number above negative maximum",
			number:   0,
			min:      -10,
			max:      -1,
			expected: "number must be between -10 and -1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateBetween(tt.number, tt.min, tt.max)
			if err == nil {
				t.Errorf("ValidateBetween() expected error, got nil")
			} else if err.Error() != tt.expected {
				t.Errorf("ValidateBetween() error = %v, expected %v", err.Error(), tt.expected)
			}
		})
	}
}
