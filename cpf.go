// Package veritas provides Brazilian CPF validation functions.
package veritas

import (
	"fmt"
	"regexp"
	"strconv"
)

// CPF validates a Brazilian CPF (Cadastro de Pessoas Físicas).
func (v *Validator) CPF(cpf interface{}) error {
	cpfStr, ok := cpf.(string)
	if !ok {
		return fmt.Errorf("CPF must be a string")
	}

	// Clean the CPF string (remove non-numeric characters)
	re := regexp.MustCompile(`\D`)
	cpfStr = re.ReplaceAllString(cpfStr, "")

	// Check if CPF has exactly 11 digits
	if len(cpfStr) != 11 {
		return fmt.Errorf("CPF must have exactly 11 digits")
	}

	// Check for invalid sequences (all same digits)
	firstDigit := cpfStr[0]
	for _, digit := range cpfStr {
		if byte(digit) != firstDigit {
			break
		}
		if digit == rune(cpfStr[len(cpfStr)-1]) {
			return fmt.Errorf("CPF cannot be a sequence of identical digits")
		}
	}

	// Validate CPF check digits
	baseDigits := cpfStr[:9]
	checkDigits := cpfStr[9:]

	// Calculate first check digit
	weights1 := []int{10, 9, 8, 7, 6, 5, 4, 3, 2}
	sum1 := 0
	for i, digit := range baseDigits {
		digitValue, _ := strconv.Atoi(string(digit))
		sum1 += digitValue * weights1[i]
	}
	remainder1 := sum1 % 11
	firstCheckDigit := 0
	if remainder1 >= 2 {
		firstCheckDigit = 11 - remainder1
	}

	// Calculate second check digit
	weights2 := []int{11, 10, 9, 8, 7, 6, 5, 4, 3, 2}
	sum2 := 0
	for i, digit := range baseDigits + string(rune(firstCheckDigit+'0')) {
		digitValue, _ := strconv.Atoi(string(digit))
		sum2 += digitValue * weights2[i]
	}
	remainder2 := sum2 % 11
	secondCheckDigit := 0
	if remainder2 >= 2 {
		secondCheckDigit = 11 - remainder2
	}

	// Compare with provided check digits
	expectedCheckDigits := fmt.Sprintf("%d%d", firstCheckDigit, secondCheckDigit)
	if checkDigits != expectedCheckDigits {
		return fmt.Errorf("invalid CPF check digits")
	}

	return nil
}
