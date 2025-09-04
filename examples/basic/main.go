package main

import (
	"fmt"

	"github.com/LcTheSecond/veritas"
)

func main() {
	fmt.Println("=== Veritas Validation Library Demo ===\n")

	// Test CNPJ validation
	fmt.Println("1. CNPJ Validation:")
	cnpj := "11.222.333/0001-81"
	if err := veritas.ValidateCNPJ(cnpj); err != nil {
		fmt.Printf("   ❌ Invalid CNPJ '%s': %v\n", cnpj, err)
	} else {
		fmt.Printf("   ✅ Valid CNPJ: %s\n", cnpj)
	}

	// Test CPF validation
	fmt.Println("\n2. CPF Validation:")
	cpf := "123.456.789-09"
	if err := veritas.ValidateCPF(cpf); err != nil {
		fmt.Printf("   ❌ Invalid CPF '%s': %v\n", cpf, err)
	} else {
		fmt.Printf("   ✅ Valid CPF: %s\n", cpf)
	}

	// Test email validation
	fmt.Println("\n3. Email Validation:")
	email := "user@example.com"
	if err := veritas.ValidateEmail(email); err != nil {
		fmt.Printf("   ❌ Invalid email '%s': %v\n", email, err)
	} else {
		fmt.Printf("   ✅ Valid email: %s\n", email)
	}

	// Test phone validation
	fmt.Println("\n4. Phone Validation:")
	phone := "+55 41 99504-8710"
	if err := veritas.ValidatePhone(phone); err != nil {
		fmt.Printf("   ❌ Invalid phone '%s': %v\n", phone, err)
	} else {
		fmt.Printf("   ✅ Valid phone: %s\n", phone)
	}

	// Test URL validation
	fmt.Println("\n5. URL Validation:")
	url := "https://www.google.com"
	if err := veritas.ValidateURL(url); err != nil {
		fmt.Printf("   ❌ Invalid URL '%s': %v\n", url, err)
	} else {
		fmt.Printf("   ✅ Valid URL: %s\n", url)
	}

	// Test string validation
	fmt.Println("\n6. String Validation:")
	str := "hello world"
	if err := veritas.ValidateString(str, 5, 20); err != nil {
		fmt.Printf("   ❌ Invalid string '%s': %v\n", str, err)
	} else {
		fmt.Printf("   ✅ Valid string: %s\n", str)
	}

	// Test number validations
	fmt.Println("\n7. Number Validations:")

	// ValidateNumber
	if err := veritas.ValidateNumber("123.45"); err != nil {
		fmt.Printf("   ❌ '123.45' is not a number: %v\n", err)
	} else {
		fmt.Printf("   ✅ '123.45' is a number\n")
	}

	// ValidatePositive
	if err := veritas.ValidatePositive(10); err != nil {
		fmt.Printf("   ❌ '10' is not positive: %v\n", err)
	} else {
		fmt.Printf("   ✅ '10' is positive\n")
	}

	// ValidateEven
	if err := veritas.ValidateEven(4); err != nil {
		fmt.Printf("   ❌ '4' is not even: %v\n", err)
	} else {
		fmt.Printf("   ✅ '4' is even\n")
	}

	// ValidateBiggerThan
	if err := veritas.ValidateBiggerThan(20, 15); err != nil {
		fmt.Printf("   ❌ '20' is not bigger than '15': %v\n", err)
	} else {
		fmt.Printf("   ✅ '20' is bigger than '15'\n")
	}

	// ValidateBetween
	if err := veritas.ValidateBetween(50, 10, 100); err != nil {
		fmt.Printf("   ❌ '50' is not between '10' and '100': %v\n", err)
	} else {
		fmt.Printf("   ✅ '50' is between '10' and '100'\n")
	}

	// ValidatePrime
	if err := veritas.ValidatePrime(7); err != nil {
		fmt.Printf("   ❌ '7' is not prime: %v\n", err)
	} else {
		fmt.Printf("   ✅ '7' is prime\n")
	}

	fmt.Println("\n=== Demo Complete ===")
}
