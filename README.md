# Veritas

A collection of input validations

[![Go Version](https://img.shields.io/badge/go-%3E%3D1.19-blue.svg)](https://golang.org/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Report Card](https://goreportcard.com/badge/github.com/yourusername/veritas)](https://goreportcard.com/report/github.com/yourusername/veritas)
[![GoDoc](https://godoc.org/github.com/yourusername/veritas?status.svg)](https://godoc.org/github.com/yourusername/veritas)

Veritas is a comprehensive Go validation library that provides a collection of input validation functions for various data formats and business requirements.

## Features

- **Document Validation**: CNPJ, CPF validation with proper algorithms
- **Contact Validation**: Email, phone number, URL validation
- **Business Validation**: SKU, price, stock quantity, barcode validation
- **String Validation**: Length, pattern, username, password, slug, UUID validation
- **Numeric Validation**: Price, stock, percentage, number range validation
- **Custom Validators**: Easy to extend with custom validation rules
- **Error Handling**: Detailed validation error messages with types
- **Performance**: Optimized for high-performance applications

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
    v := veritas.New()
    
    // Validate CNPJ
    if err := v.ValidateCNPJ("11.222.333/0001-81"); err != nil {
        fmt.Printf("Invalid CNPJ: %v\n", err)
    } else {
        fmt.Println("Valid CNPJ!")
    }
    
    // Validate CPF
    if err := v.ValidateCPF("123.456.789-09"); err != nil {
        fmt.Printf("Invalid CPF: %v\n", err)
    } else {
        fmt.Println("Valid CPF!")
    }
    
    // Validate Email
    if err := v.ValidateEmail("user@example.com"); err != nil {
        fmt.Printf("Invalid email: %v\n", err)
    } else {
        fmt.Println("Valid email!")
    }
}
```

## Usage

### Document Validation

```go
v := veritas.New()

// CNPJ validation (Brazilian corporate tax ID)
err := v.ValidateCNPJ("11.222.333/0001-81")
if err != nil {
    log.Printf("CNPJ validation failed: %v", err)
}

// CPF validation (Brazilian individual tax ID)
err = v.ValidateCPF("123.456.789-09")
if err != nil {
    log.Printf("CPF validation failed: %v", err)
}

// CPF or CNPJ validation (auto-detects based on length)
err = v.ValidateCPFOrCNPJ("123.456.789-09")
if err != nil {
    log.Printf("Document validation failed: %v", err)
}
```

### Contact Information

```go
// Email validation
err := v.ValidateEmail("user@example.com")

// Phone number validation (international)
err = v.ValidatePhone("+5511999999999")

// Brazilian phone validation
err = v.ValidateBrazilianPhone("11999999999")

// US phone validation
err = v.ValidateUSPhone("1234567890")

// URL validation
err = v.ValidateURL("https://example.com")
```

### Business Data

```go
// SKU validation
err := v.ValidateSKU("ABC-123-XYZ")

// Price validation (positive number with up to 2 decimal places)
err = v.ValidatePrice("29.99")

// Stock quantity validation
err = v.ValidateStockQuantity(100)
```

### String Validation

```go
// String length validation
err := v.ValidateStringLength("hello", 3, 10)

// Username validation
err = v.ValidateUsername("user123")

// Password validation
err = v.ValidatePassword("MyPass123")

// Strong password validation
err = v.ValidateStrongPassword("MyStr0ng!Pass")

// URL slug validation
err = v.ValidateSlug("my-awesome-slug")

// UUID validation
err = v.ValidateUUID("550e8400-e29b-41d4-a716-446655440000")

// Hex color validation
err = v.ValidateHexColor("#FF5733")
```

### Numeric Validation

```go
// Positive number validation
err := v.ValidatePositiveNumber(42)

// Number range validation
err = v.ValidateNumberRange(50, 10, 100)

// Percentage validation
err = v.ValidatePercentage(85.5)

// Barcode validation (EAN-13, UPC-A)
err = v.ValidateBarcode("1234567890123")
```

## API Reference

### Core Functions

| Function | Description | Example |
|----------|-------------|---------|
| `ValidateCNPJ(cnpj interface{}) error` | Validates Brazilian CNPJ | `"11.222.333/0001-81"` |
| `ValidateCPF(cpf interface{}) error` | Validates Brazilian CPF | `"123.456.789-09"` |
| `ValidateEmail(email interface{}) error` | Validates email format | `"user@example.com"` |
| `ValidatePhone(phone interface{}) error` | Validates international phone | `"+5511999999999"` |
| `ValidateBrazilianPhone(phone interface{}) error` | Validates Brazilian phone | `"11999999999"` |
| `ValidateURL(url interface{}) error` | Validates URL format | `"https://example.com"` |
| `ValidateSKU(sku interface{}) error` | Validates SKU format | `"ABC-123-XYZ"` |
| `ValidatePrice(price interface{}) error` | Validates price format | `"29.99"` |
| `ValidateStockQuantity(qty interface{}) error` | Validates stock quantity | `100` |
| `ValidateStringLength(str interface{}, min, max int) error` | Validates string length | `"hello", 3, 10` |
| `ValidateUsername(username interface{}) error` | Validates username format | `"user123"` |
| `ValidatePassword(password interface{}) error` | Validates password format | `"MyPass123"` |
| `ValidateUUID(uuid interface{}) error` | Validates UUID format | `"550e8400-e29b-41d4-a716-446655440000"` |

### Error Handling

```go
v := veritas.New()

err := v.ValidateCNPJ("invalid-cnpj")
if err != nil {
    fmt.Printf("Validation error: %v\n", err)
}

// Using ValidationError type for detailed error handling
validationErr := v.Validate("cnpj", "invalid-cnpj", v.ValidateCNPJ)
if validationErr != nil {
    fmt.Printf("Field: %s, Type: %s, Message: %s\n", 
        validationErr.Field, validationErr.Type, validationErr.Message)
}
```

## Custom Validators

You can easily extend Veritas with custom validators:

```go
// Custom validator function
func ValidateCustomField(value interface{}) error {
    str, ok := value.(string)
    if !ok {
        return fmt.Errorf("value must be a string")
    }
    if len(str) < 5 {
        return fmt.Errorf("field must be at least 5 characters long")
    }
    return nil
}

// Usage with Veritas
v := veritas.New()
err := v.Validate("custom_field", "test", ValidateCustomField)
```

## Performance

Veritas is optimized for performance:

- **Minimal allocations** - reuses regex patterns
- **Fast algorithms** - optimized validation logic
- **Concurrent safe** - can be used in goroutines
- **Memory efficient** - no global state

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

### Development Setup

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

### Adding New Validators

1. Add the validator function to the appropriate file
2. Add comprehensive tests
3. Update documentation
4. Ensure all tests pass

## Testing

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run specific test
go test -run TestValidateCNPJ
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Brazilian tax ID validation algorithms based on official specifications
- Email validation follows RFC 5322 standards
- Phone number validation supports international formats

## Changelog

### v1.0.0
- Initial release
- CNPJ, CPF validation with proper algorithms
- Email, phone, URL validation
- String validation (length, pattern, username, password, slug, UUID)
- Numeric validation (price, stock, percentage, barcode)
- Business data validation (SKU, price, stock quantity)
- Comprehensive error handling with types

---

**Veritas** - *Where data meets truth*
