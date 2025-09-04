package main

import (
	"fmt"
	"log"

	"github.com/LcTheSecond/veritas"
)

func main() {
	// Create a new validator instance
	v := veritas.New()

	fmt.Println("=== Veritas Validation Library Demo ===\n")

	// Test CNPJ validation
	fmt.Println("1. CNPJ Validation:")
	cnpj := "11.222.333/0001-81"
	if err := v.ValidateCNPJ(cnpj); err != nil {
		fmt.Printf("   ❌ Invalid CNPJ '%s': %v\n", cnpj, err)
	} else {
		fmt.Printf("   ✅ Valid CNPJ: %s\n", cnpj)
	}

	// Test CPF validation
	fmt.Println("\n2. CPF Validation:")
	cpf := "123.456.789-09"
	if err := v.ValidateCPF(cpf); err != nil {
		fmt.Printf("   ❌ Invalid CPF '%s': %v\n", cpf, err)
	} else {
		fmt.Printf("   ✅ Valid CPF: %s\n", cpf)
	}

	// Test email validation
	fmt.Println("\n3. Email Validation:")
	email := "user@example.com"
	if err := v.ValidateEmail(email); err != nil {
		fmt.Printf("   ❌ Invalid email '%s': %v\n", email, err)
	} else {
		fmt.Printf("   ✅ Valid email: %s\n", email)
	}

	// Test phone validation
	fmt.Println("\n4. Phone Validation:")
	phone := "+5511999999999"
	if err := v.ValidateBrazilianPhone(phone); err != nil {
		fmt.Printf("   ❌ Invalid phone '%s': %v\n", phone, err)
	} else {
		fmt.Printf("   ✅ Valid phone: %s\n", phone)
	}

	// Test price validation
	fmt.Println("\n5. Price Validation:")
	price := "29.99"
	if err := v.ValidatePrice(price); err != nil {
		fmt.Printf("   ❌ Invalid price '%s': %v\n", price, err)
	} else {
		fmt.Printf("   ✅ Valid price: %s\n", price)
	}

	// Test SKU validation
	fmt.Println("\n6. SKU Validation:")
	sku := "PROD-123-ABC"
	if err := v.ValidateSKU(sku); err != nil {
		fmt.Printf("   ❌ Invalid SKU '%s': %v\n", sku, err)
	} else {
		fmt.Printf("   ✅ Valid SKU: %s\n", sku)
	}

	// Test multiple validations
	fmt.Println("\n7. Multiple Validations:")
	errors := v.ValidateMultiple(
		func() *veritas.ValidationError {
			if err := v.ValidateEmail("invalid-email"); err != nil {
				return veritas.NewValidationError("email", veritas.ErrorTypeInvalid, err.Error(), "invalid-email")
			}
			return nil
		},
		func() *veritas.ValidationError {
			if err := v.ValidateStringLength("short", 10, 100); err != nil {
				return veritas.NewValidationError("description", veritas.ErrorTypeTooShort, err.Error(), "short")
			}
			return nil
		},
	)

	if len(errors) > 0 {
		fmt.Printf("   ❌ Found %d validation errors:\n", len(errors))
		for _, err := range errors {
			fmt.Printf("      - %s: %s\n", err.Field, err.Message)
		}
	} else {
		fmt.Println("   ✅ All validations passed!")
	}

	// Test formatting functions
	fmt.Println("\n8. Formatting Functions:")
	if formatted, err := v.FormatCPF("12345678909"); err != nil {
		log.Printf("Error formatting CPF: %v", err)
	} else {
		fmt.Printf("   Formatted CPF: %s\n", formatted)
	}

	if formatted, err := v.FormatPrice(29.99, "USD"); err != nil {
		log.Printf("Error formatting price: %v", err)
	} else {
		fmt.Printf("   Formatted price: %s\n", formatted)
	}

	fmt.Println("\n=== Demo Complete ===")
}
