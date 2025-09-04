# Veritas

A simple Go validation library

[![Go Version](https://img.shields.io/badge/go-%3E%3D1.19-blue.svg)](https://golang.org/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Veritas is a simple Go validation library that provides essential validation functions for common data types.

## Features

- **Document Validation**: CNPJ, CPF validation with proper algorithms
- **Contact Validation**: Email, phone number, URL validation
- **String Validation**: Length validation
- **Number Validation**: IsNumber, IsPositive, IsNegative, IsEven, BiggerThan, SmallerThan, Between, IsPrime
- **Simple Error Handling**: Standard Go error pattern (nil = valid, error = invalid)

## Installation

```bash
go get github.com/LcTheSecond/veritas
```

## Quick Start

```go
package main

import (
    "fmt"
    "github.com/LcTheSecond/veritas"
)

func main() {
    // Validate CNPJ
    if err := veritas.ValidateCNPJ("11.222.333/0001-81"); err != nil {
        fmt.Printf("Invalid CNPJ: %v\n", err)
    } else {
        fmt.Println("Valid CNPJ!")
    }
    
    // Validate CPF
    if err := veritas.ValidateCPF("123.456.789-09"); err != nil {
        fmt.Printf("Invalid CPF: %v\n", err)
    } else {
        fmt.Println("Valid CPF!")
    }
    
    // Validate Email
    if err := veritas.ValidateEmail("user@example.com"); err != nil {
        fmt.Printf("Invalid email: %v\n", err)
    } else {
        fmt.Println("Valid email!")
    }
}
```

## Usage

### Document Validation

```go
// CNPJ validation (Brazilian corporate tax ID)
err := veritas.ValidateCNPJ("11.222.333/0001-81")
if err != nil {
    log.Printf("CNPJ validation failed: %v", err)
}

// CPF validation (Brazilian individual tax ID)
err = veritas.ValidateCPF("123.456.789-09")
if err != nil {
    log.Printf("CPF validation failed: %v", err)
}
```

### Contact Information

```go
// Email validation
err := veritas.ValidateEmail("user@example.com")

// Brazilian phone validation (mobile and landline)
err = veritas.ValidatePhone("+55 41 9.9504-8710")  // Mobile
err = veritas.ValidatePhone("+55 41 3346-4468")    // Landline
err = veritas.ValidatePhone("41 9.9504-8710")      // Mobile without +55
err = veritas.ValidatePhone("41 3346-4468")        // Landline without +55

// URL validation (checks format + HTTP 200 status)
err = veritas.ValidateURL("https://example.com")
```

### String Validation

```go
// String length validation
err := veritas.ValidateString("hello", 3, 10)  // min 3, max 10 characters
```

### Number Validation

```go
// Basic number validation
err := veritas.ValidateNumber(42)           // Check if it's a valid number
err := veritas.ValidatePositive(42)         // Check if positive
err := veritas.ValidateNegative(-5)         // Check if negative
err := veritas.ValidateEven(8)              // Check if even

// Comparison validation
err := veritas.ValidateBiggerThan(10, 5)      // 10 > 5
err := veritas.ValidateSmallerThan(3, 10)     // 3 < 10
err := veritas.ValidateBetween(15, 10, 20)    // 10 <= 15 <= 20

// Prime number validation
err := veritas.ValidatePrime(17)            // Check if prime
```

## API Reference

### Core Functions

| Function | Description | Example |
|----------|-------------|---------|
| `ValidateCNPJ(cnpj interface{}) error` | Validates Brazilian CNPJ | `"11.222.333/0001-81"` |
| `ValidateCPF(cpf interface{}) error` | Validates Brazilian CPF | `"123.456.789-09"` |
| `ValidateEmail(email interface{}) error` | Validates email format | `"user@example.com"` |
| `ValidatePhone(phone interface{}) error` | Validates Brazilian phone | `"+55 41 9.9504-8710"` |
| `ValidateURL(url interface{}) error` | Validates URL + HTTP 200 | `"https://example.com"` |
| `ValidateString(str interface{}, min, max int) error` | Validates string length | `"hello", 3, 10` |
| `ValidateNumber(num interface{}) error` | Validates if number | `42` |
| `ValidatePositive(num interface{}) error` | Validates if positive | `42` |
| `ValidateNegative(num interface{}) error` | Validates if negative | `-5` |
| `ValidateEven(num interface{}) error` | Validates if even | `8` |
| `ValidateBiggerThan(num interface{}, than float64) error` | Validates if bigger | `10, 5` |
| `ValidateSmallerThan(num interface{}, than float64) error` | Validates if smaller | `3, 10` |
| `ValidateBetween(num interface{}, min, max float64) error` | Validates if between | `15, 10, 20` |
| `ValidatePrime(num interface{}) error` | Validates if prime | `17` |

### Error Handling

All validation functions follow the standard Go error pattern:

```go
v := veritas.New()

err := v.CNPJ("invalid-cnpj")
if err != nil {
    fmt.Printf("Validation error: %v\n", err)
} else {
    fmt.Println("Valid!")
}
```

## Brazilian Phone Number Format

The phone validation supports Brazilian phone numbers:

**Mobile numbers:**
- `+55 41 9.9504-8710` (with country code)
- `41 9.9504-8710` (without country code)

**Landline numbers:**
- `+55 41 3346-4468` (with country code)
- `41 3346-4468` (without country code)

The validation checks:
- Valid Brazilian area codes (DDD)
- Mobile numbers must start with 9 after area code
- Proper digit validation
- Handles formatted input (spaces, dots, hyphens, parentheses)

## URL Validation

The URL validation not only checks the format but also verifies the URL is accessible:

- Validates URL format (scheme, host)
- Makes HTTP HEAD request
- Verifies HTTP 200 status code
- 10-second timeout

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

### Development Setup

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git checkout -b feature/amazing-feature`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## Testing

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run specific test
go test -run TestCNPJ
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Brazilian tax ID validation algorithms based on official specifications
- Email validation follows RFC 5322 standards
- Phone number validation supports Brazilian formats

## Changelog

### v0.1.0
- Initial release
- CNPJ, CPF validation with proper algorithms
- Email, phone, URL validation
- String length validation
- Number validation (IsNumber, IsPositive, IsNegative, IsEven, BiggerThan, SmallerThan, Between, IsPrime)
- Simple error handling with standard Go pattern

---

**Veritas** - *Where data meets truth*