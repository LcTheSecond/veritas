# Veritas

A collection of input validations

[![Go Version](https://img.shields.io/badge/go-%3E%3D1.19-blue.svg)](https://golang.org/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Report Card](https://goreportcard.com/badge/github.com/yourusername/veritas)](https://goreportcard.com/report/github.com/yourusername/veritas)
[![GoDoc](https://godoc.org/github.com/yourusername/veritas?status.svg)](https://godoc.org/github.com/yourusername/veritas)

Veritas is a comprehensive Go validation library that provides a collection of input validation functions for various data formats and business requirements.

## Features

- **Document Validation**: CNPJ, CPF, RG validation with proper algorithms
- **Contact Validation**: Email, phone number, URL validation
- **Business Validation**: SKU, price, stock quantity validation
- **Data Format Validation**: Date, time, postal code validation
- **Custom Validators**: Easy to extend with custom validation rules
- **Error Handling**: Detailed validation error messages
- **Performance**: Optimized for high-performance applications

## Installation

```bash
go get github.com/yourusername/veritas
```

## Quick Start

```go
package main

import (
    "fmt"
    "github.com/yourusername/veritas"
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

// RG validation (Brazilian identity card)
err = v.ValidateRG("12.345.678-9")
if err != nil {
    log.Printf("RG validation failed: %v", err)
}
```

### Contact Information

```go
// Email validation
err := v.ValidateEmail("user@example.com")

// Phone number validation
err = v.ValidatePhone("+55 11 99999-9999")

// URL validation
err = v.ValidateURL("https://example.com")
```

### Business Data

```go
// SKU validation
err := v.ValidateSKU("ABC-123-XYZ")

// Price validation (non-negative, reasonable range)
err = v.ValidatePrice(1000) // price in cents

// Stock quantity validation
err = v.ValidateStockQuantity(100)
```

### Postal and Location

```go
// Brazilian postal code (CEP)
err := v.ValidateCEP("01234-567")

// Brazilian state (UF)
err = v.ValidateUF("SP")
```

## API Reference

### Core Functions

| Function | Description | Example |
|----------|-------------|---------|
| `ValidateCNPJ(cnpj string) error` | Validates Brazilian CNPJ | `11.222.333/0001-81` |
| `ValidateCPF(cpf string) error` | Validates Brazilian CPF | `123.456.789-09` |
| `ValidateEmail(email string) error` | Validates email format | `user@example.com` |
| `ValidatePhone(phone string) error` | Validates phone number | `+55 11 99999-9999` |
| `ValidateURL(url string) error` | Validates URL format | `https://example.com` |
| `ValidateSKU(sku string) error` | Validates SKU format | `ABC-123-XYZ` |
| `ValidatePrice(price int64) error` | Validates price (cents) | `1000` |
| `ValidateStockQuantity(qty int) error` | Validates stock quantity | `100` |

### Error Handling

```go
v := veritas.New()

err := v.ValidateCNPJ("invalid-cnpj")
if err != nil {
    switch e := err.(type) {
    case veritas.ValidationError:
        fmt.Printf("Field: %s, Message: %s\n", e.Field, e.Message)
    default:
        fmt.Printf("Validation error: %v\n", err)
    }
}
```

## Custom Validators

You can easily extend Veritas with custom validators:

```go
// Custom validator function
func ValidateCustomField(value string) error {
    if len(value) < 5 {
        return veritas.ValidationError{
            Field:   "custom_field",
            Message: "Field must be at least 5 characters long",
        }
    }
    return nil
}

// Usage
err := ValidateCustomField("test")
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
- CNPJ, CPF, RG validation
- Email, phone, URL validation
- Business data validation
- Comprehensive test coverage

---

**Veritas** - *Where data meets truth*
