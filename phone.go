// Package veritas provides phone number validation functions.
package veritas

import (
	"fmt"
	"regexp"
	"strings"
)

// ValidatePhone validates a Brazilian phone number format.
func ValidatePhone(phone interface{}) error {
	phoneStr, ok := phone.(string)
	if !ok {
		return fmt.Errorf("phone must be a string")
	}

	// Clean the phone string (remove spaces, dots, hyphens)
	phoneStr = cleanPhone(phoneStr)

	// Check if phone is empty after cleaning
	if isEmpty(phoneStr) {
		return fmt.Errorf("phone cannot be empty")
	}

	// Check if it's a mobile number (11 digits total: +55 + DDD + 9 + 8 digits)
	if len(phoneStr) == 15 && strings.HasPrefix(phoneStr, "+55") {
		return validateMobile(phoneStr)
	}

	// Check if it's a landline number (12 digits total: +55 + DDD + 8 digits)
	if len(phoneStr) == 14 && strings.HasPrefix(phoneStr, "+55") {
		return validateLandline(phoneStr)
	}

	// Check if it's without country code (mobile: 11 digits, landline: 10 digits)
	if len(phoneStr) == 11 {
		return validateMobile("+55" + phoneStr)
	}
	if len(phoneStr) == 10 {
		return validateLandline("+55" + phoneStr)
	}

	return fmt.Errorf("invalid Brazilian phone number format")
}

// validateMobile validates a Brazilian mobile phone number.
func validateMobile(phone string) error {
	// Mobile format: +55 + DDD + 9 + 8 digits
	// Example: +5541995048710

	// Check DDD (area code) - must be 2 digits, first digit 1-9, second digit 1-9
	ddd := phone[3:5]
	if !isValidDDD(ddd) {
		return fmt.Errorf("invalid area code (DDD)")
	}

	// Check if 5th digit is 9 (mobile indicator)
	if phone[5] != '9' {
		return fmt.Errorf("mobile number must start with 9 after area code")
	}

	// Check remaining 8 digits
	number := phone[6:]
	if !isValidPhoneDigits(number) {
		return fmt.Errorf("invalid phone number digits")
	}

	return nil
}

// validateLandline validates a Brazilian landline phone number.
func validateLandline(phone string) error {
	// Landline format: +55 + DDD + 8 digits
	// Example: +554133464468

	// Check DDD (area code) - must be 2 digits, first digit 1-9, second digit 1-9
	ddd := phone[3:5]
	if !isValidDDD(ddd) {
		return fmt.Errorf("invalid area code (DDD)")
	}

	// Check remaining 8 digits
	number := phone[5:]
	if !isValidPhoneDigits(number) {
		return fmt.Errorf("invalid phone number digits")
	}

	return nil
}

// isValidDDD validates Brazilian area codes (DDD).
func isValidDDD(ddd string) bool {
	// Valid DDDs in Brazil (2 digits, first 1-9, second 1-9)
	validDDDs := []string{
		"11", "12", "13", "14", "15", "16", "17", "18", "19", // São Paulo
		"21", "22", "24", // Rio de Janeiro
		"27", "28", // Espírito Santo
		"31", "32", "33", "34", "35", "37", "38", // Minas Gerais
		"41", "42", "43", "44", "45", "46", // Paraná
		"47", "48", "49", // Santa Catarina
		"51", "53", "54", "55", // Rio Grande do Sul
		"61",       // Distrito Federal
		"62", "64", // Goiás
		"63",       // Tocantins
		"65", "66", // Mato Grosso
		"67",                         // Mato Grosso do Sul
		"68",                         // Acre
		"69",                         // Rondônia
		"71", "73", "74", "75", "77", // Bahia
		"79",       // Sergipe
		"81", "87", // Pernambuco
		"82",       // Alagoas
		"83",       // Paraíba
		"84",       // Rio Grande do Norte
		"85", "88", // Ceará
		"86", "89", // Piauí
		"91", "93", "94", // Pará
		"92", "97", // Amazonas
		"95",       // Roraima
		"96",       // Amapá
		"98", "99", // Maranhão
	}

	for _, validDDD := range validDDDs {
		if ddd == validDDD {
			return true
		}
	}
	return false
}

// isValidPhoneDigits validates phone number digits.
func isValidPhoneDigits(digits string) bool {
	// Check if all characters are digits
	matched, _ := regexp.MatchString(`^\d+$`, digits)
	return matched
}

// cleanPhone removes spaces, dots, hyphens from phone number.
func cleanPhone(phone string) string {
	// Remove spaces, dots, hyphens, parentheses
	re := regexp.MustCompile(`[\s\.\-\(\)]`)
	return re.ReplaceAllString(phone, "")
}
