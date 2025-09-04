// Package veritas provides Brazilian CNPJ validation functions.
package veritas

import (
	"fmt"
	"regexp"
	"strconv"
)

// ValidateCNPJ validates a Brazilian CNPJ (Cadastro Nacional da Pessoa Jurídica).
func ValidateCNPJ(cnpj interface{}) error {
	cnpjStr, ok := cnpj.(string)
	if !ok {
		return fmt.Errorf("CNPJ must be a string")
	}

	// Clean the CNPJ string (remove non-numeric characters)
	re := regexp.MustCompile(`\D`)
	cnpjStr = re.ReplaceAllString(cnpjStr, "")

	// Check if CNPJ has exactly 14 digits
	if len(cnpjStr) != 14 {
		return fmt.Errorf("CNPJ must have exactly 14 digits")
	}

	// Check for invalid sequences (all same digits)
	firstDigit := cnpjStr[0]
	allSame := true
	for _, digit := range cnpjStr {
		if byte(digit) != firstDigit {
			allSame = false
			break
		}
	}
	if allSame {
		return fmt.Errorf("CNPJ cannot be a sequence of identical digits")
	}

	// Validate CNPJ check digits
	baseDigits := cnpjStr[:12]
	checkDigits := cnpjStr[12:]

	// Calculate first check digit
	weights1 := []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
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
	weights2 := []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
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
		return fmt.Errorf("invalid CNPJ check digits")
	}

	return nil
}
