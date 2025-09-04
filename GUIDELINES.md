
# PROJECT RULES: Google Style Guide Compliant Standards

## 1. OVERVIEW & OBJECTIVE
Clean, simple Go backend standards aligned with [Google's Go Style Guide](https://google.github.io/styleguide/go/best-practices), focused on domain-driven design with flat structure. Emphasizes readability, maintainability, and single sources of truth using Ent ORM and following OpenAPI 3.0 API contracts.

## 2. NAMING CONVENTIONS

### 2.1 Function and Method Names

#### Avoid Repetition
When choosing function or method names, consider the context in which the name will be read:

* **Omit redundant information from function names:**
  * Package names (when clear from context)
  * Receiver types
  * Parameter types (when no collision)
  * Pointer indicators

```go
// Bad: Repetitive naming
package users
func CreateUser(user *User) error
func (u *User) GetUserName() string

// Good: Context-aware naming
package users
func Create(user *User) error
func (u *User) Name() string
```

* **For methods, do not repeat the receiver name:**
```go
// Bad:
func (c *Config) WriteConfigTo(w io.Writer) (int64, error)

// Good:
func (c *Config) WriteTo(w io.Writer) (int64, error)
```

* **Do not repeat parameter names:**
```go
// Bad:
func OverrideFirstWithSecond(dest, source *Config) error

// Good:
func Override(dest, source *Config) error
```

#### Naming Conventions
* **Functions that return values use noun-like names:**
```go
// Good:
func (c *Config) JobName(key string) (value string, ok bool)
```

* **Functions that perform actions use verb-like names:**
```go
// Good:
func (c *Config) WriteDetail(w io.Writer) (int64, error)
```

* **Avoid "Get" prefix for accessors:**
```go
// Bad:
func (c *Config) GetJobName(key string) (value string, ok bool)

// Good:
func (c *Config) JobName(key string) (value string, ok bool)
```

* **Include type names for disambiguation:**
```go
// Good:
func ParseInt(input string) (int, error)
func ParseInt64(input string) (int64, error)
```

### 2.2 General Naming Rules
* **Variables:** `camelCase` with short, descriptive names (e.g., `user`, `orderID`, `isValid`)
* **Constants:** `UPPER_SNAKE_CASE` (e.g., `MAX_RETRY`, `DEFAULT_LIMIT`)
* **Types/Structs:** `PascalCase` (e.g., `UserHandler`, `CreateUserInput`)
* **Packages:** `lowercase` (e.g., `users`, `orders`)
* **Files:** `snake_case.go` (e.g., `handlers.go`, `service.go`, `routes.go`)

## 3. TEST DOUBLE AND HELPER PACKAGES

### 3.1 Test Package Naming
Use ENT's built-in test client for testing:

```go
// Use enttest for testing ENT operations
import "github.com/LcTheSecond/back25/internal/ent/enttest"

func TestService_Create(t *testing.T) {
    client := enttest.NewClient(t)
    service := NewService(client)
    // ... test logic
}
```

### 3.2 Test Double Naming Patterns

#### Using ENT Test Client
```go
// Good: Use ENT's test client for testing
func TestService_Create(t *testing.T) {
    client := enttest.NewClient(t)
    service := NewService(client)
    // ... test logic
}
```

#### Testing with Real Database
```go
// Good: Test with actual database operations
func TestService_Create_Integration(t *testing.T) {
    client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
    defer client.Close()
    
    service := NewService(client)
    // ... integration test logic
}
```

#### Mocking External Dependencies
```go
// Good: Mock external services, not ENT operations
type MockEmailService struct{}

func (m *MockEmailService) SendWelcomeEmail(email string) error {
    return nil // Mock implementation
}
```

### 3.3 Test Variable Naming
Choose names that clearly differentiate test clients from production types:

```go
// Good: Clear differentiation
func TestProcessor_Process(t *testing.T) {
    testClient := enttest.NewClient(t)
    service := NewService(testClient)
    // ... test logic
}
```

## 4. GLOBAL STATE AVOIDANCE

### 4.1 Litmus Tests for Global State
Global state is unsafe when:
* Multiple functions interact via global state when otherwise independent
* Independent test cases interact through global state
* Users are tempted to replace global state with test doubles
* Users must consider special ordering requirements

### 4.2 Safe Global State Conditions
Global state is acceptable only when:
* The state is logically constant
* The package's observable behavior is stateless
* The state doesn't bleed into external systems
* There's no expectation of predictable behavior

### 4.3 Dependency Injection Pattern
Prefer explicit dependency passing over global state:

```go
// Bad: Global state
package users
var client *Client

func Create(user *User) error {
    return client.Create(user)
}

// Good: Dependency injection
type Service struct {
    client *Client
}

func NewService(client *Client) *Service {
    return &Service{client: client}
}

func (s *Service) Create(user *User) error {
    return s.client.Create(user)
}
```

## 5. ARCHITECTURAL & STRUCTURAL PATTERNS

### 5.1 Domain-Driven Design
* **Organization:** Business domains in `internal/domains/{domain}/`
* **Flat Package Structure:** Each domain contains: `handlers.go`, `service.go`, `routes.go`, `types.go`
* **Single Package Per Domain:** All files in same package for easy access
* **Direct ENT Integration:** Service layer uses ENT client directly for data access
* **Grouped Routes:** Careful when setting up routes, since there is a grouped route prefix in server package (e.g., `api/v1`)

### 5.2 Layer Responsibilities
* **Handler Pattern:** Thin handlers that parse request, call service, return response
* **Service Layer:** Business logic and data access using ENT directly
* **No Repository Layer:** Use ENT client directly in service layer
* **Dependency Injection:** Simple constructor functions for dependencies

## 6. FUNCTION DESIGN PRINCIPLES

### 6.1 Core Principles
* **Single Responsibility:** Each function does ONE thing well
* **DRY Principle:** Eliminate code duplication through shared utilities
* **Function Size:** Keep functions small (under 30 lines, ideally under 15)
* **Parameter Count:** Limit parameters (max 3-4, use structs for more)
* **Return Values:** Return meaningful values, avoid side effects
* **Abstraction Levels:** Functions operate at single level of abstraction
* **Grouped Routes:** Careful when setting up routes, since there is a grouped route prefix in server package (e.g., `api/v1`)

### 6.2 Function Composition
Break complex functions into smaller, focused functions:

```go
// Bad: Monolithic function
func ProcessOrder(order *Order) error {
    // 50+ lines of mixed concerns
}

// Good: Composed functions
func ProcessOrder(order *Order) error {
    if err := validateOrder(order); err != nil {
        return fmt.Errorf("validating order: %w", err)
    }
    
    if err := calculateTotals(order); err != nil {
        return fmt.Errorf("calculating totals: %w", err)
    }
    
    return saveOrder(order)
}
```

## 7. ERROR HANDLING STANDARDS

### 7.1 Error Handling Patterns
* **Simple Error Handling:** Use `return err` with context when needed
* **Error Context:** Use `fmt.Errorf("context: %w", err)` for additional context
* **Error Types:** Use standard Go errors, avoid complex error hierarchies
* **Validation Errors:** Return clear, actionable validation error messages
* **HTTP Status Codes:** Map errors to appropriate HTTP status codes

### 7.2 Error Wrapping
```go
// Good: Proper error wrapping
func (s *Service) CreateUser(input *CreateUserInput) (*ent.User, error) {
    if err := s.validateUser(input); err != nil {
        return nil, fmt.Errorf("validating user: %w", err)
    }
    
    user, err := s.client.User.Create().
        SetName(input.Name).
        SetEmail(input.Email).
        Save(ctx)
    if err != nil {
        return nil, fmt.Errorf("creating user: %w", err)
    }
    
    return user, nil
}
```

## 8. DATABASE & ENT ORM PATTERNS

### 8.1 Ent ORM Usage
* **Direct Usage:** Use ENT client directly in service layer
* **No Repository Wrapper:** ENT is already a complete DAL
* **Service Layer Integration:** Embed ENT operations in business logic
* **Connection Pooling:** Implement efficient connection pooling
* **Transactions:** Wrap related operations in database transactions
* **Query Optimization:** Use ENT's query builders directly
* **Migrations:** Use ENT's migration system for schema changes

### 8.2 Service Layer with ENT
```go
// Good: Direct ENT usage in service layer
type Service struct {
    client *ent.Client
}

func (s *Service) Create(ctx context.Context, input *CreateUserInput) (*ent.User, error) {
    user, err := s.client.User.Create().
        SetName(input.Name).
        SetEmail(input.Email).
        Save(ctx)
    if err != nil {
        return nil, fmt.Errorf("creating user: %w", err)
    }
    return user, nil
}
```

## 9. API DESIGN & IMPLEMENTATION

### 9.1 OpenAPI 3.0 Compliance
* **Mandatory:** Follow API-CONTRACT.md exactly
* **RESTful Endpoints:** Use standard HTTP methods and status codes
* **Response Format:** Return data directly in `data` field, no custom envelopes
* **Error Format:** Use OpenAPI standard error response format
* **Validation:** Use struct tags for input validation
* **Pagination:** Implement standard pagination with `page` and `limit` parameters
* **Grouped Routes:** Careful when setting up routes, since there is a grouped route prefix in server package (e.g., `api/v1`)

### 9.2 Response Handling
* **Use Response Package:** Always use `internal/shared/response` package helpers
* **Success Responses:** Use `response.Success()` for all successful responses
* **Error Responses:** Use standardized error helpers like `response.BadRequest()`, `response.NotFound()`
* **OpenAPI 3.0 Format:** Maintain exact error structure with `status`, `title`, and `detail`

### 9.3 Handler Patterns
```go
// Good: Clean handler structure using response package
import "github.com/LcTheSecond/gomx/backend/internal/shared/response"

func (h *Handler) Create(c *fiber.Ctx) error {
    var input CreateUserInput
    if err := c.BodyParser(&input); err != nil {
        return response.BadRequest(c, "Invalid JSON payload")
    }
    
    if err := h.validateCreateInput(&input); err != nil {
        return response.ValidationError(c, err.Error())
    }
    
    user, err := h.service.Create(c.Context(), &input)
    if err != nil {
        return response.InternalError(c, "Failed to create user")
    }
    
    return response.Success(c, fiber.StatusCreated, user)
}

func (h *Handler) Get(c *fiber.Ctx) error {
    id := c.Params("id")
    user, err := h.service.Get(c.Context(), id)
    if err != nil {
        return response.NotFound(c, "User not found")
    }
    
    return response.Success(c, fiber.StatusOK, user)
}
```

## 10. TESTING & QUALITY ASSURANCE

### 10.1 Testing Framework
* **Testing Framework:** Use Go's standard `testing` package with `testify` for assertions
* **Unit Tests:** Write unit tests for all service methods and utility functions
* **Test Coverage:** Aim for minimum 80% test coverage for business logic
* **Test Naming:** Use `Test{FunctionName}_{Scenario}` format (e.g., `TestCreateUser_Success`)
* **Test Files:** Place tests in same package with `_test.go` suffix

### 10.2 Test Organization
```go
// Good: Well-organized test structure
func TestService_CreateUser(t *testing.T) {
    tests := []struct {
        name    string
        input   *CreateUserInput
        mock    *enttest.Client
        wantErr bool
    }{
        {
            name: "successful_creation",
            input: &CreateUserInput{Name: "John", Email: "john@example.com"},
            mock:  enttest.NewClient(t),
            wantErr: false,
        },
        {
            name: "validation_failure",
            input: &CreateUserInput{Name: "", Email: "invalid-email"},
            mock:  enttest.NewClient(t),
            wantErr: true,
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            service := NewService(tt.mock)
            _, err := service.Create(context.Background(), tt.input)
            
            if tt.wantErr && err == nil {
                t.Error("expected error, got nil")
            }
            if !tt.wantErr && err != nil {
                t.Errorf("unexpected error: %v", err)
            }
        })
    }
}
```

### 10.3 Test Isolation
Ensure tests don't interfere with each other:

```go
// Good: Isolated test setup
func TestMain(m *testing.M) {
    // Setup test database
    setupTestDB()
    
    // Run tests
    code := m.Run()
    
    // Cleanup
    cleanupTestDB()
    
    os.Exit(code)
}
```

## 11. CODE ORGANIZATION & FILE STRUCTURE

```
internal/
├── domains/
│   ├── users/
│   │   ├── handlers.go       # HTTP handlers
│   │   ├── service.go        # business logic + data access
│   │   ├── routes.go         # route definitions
│   │   ├── types.go          # request/response types
│   │   └── users_test.go     # tests
│   └── orders/
│       ├── handlers.go
│       ├── service.go
│       ├── routes.go
│       ├── types.go
│       └── orders_test.go
├── shared/
│   ├── response/             # response utilities (OpenAPI 3.0 compliant)
│   ├── middleware/           # HTTP middleware
│   ├── utils/               # shared utilities
│   └── errors/              # error management
├── server/
└── ent/                     # Ent ORM generated code
```

## 12. RESPONSE PACKAGE STRUCTURE

### 12.1 Response Helpers
```go
// internal/shared/response/response.go
package response

// Success sends OpenAPI 3.0 compliant success response
func Success(c *fiber.Ctx, status int, data interface{}) error

// Error sends OpenAPI 3.0 compliant error response
func Error(c *fiber.Ctx, status int, title string, detail string) error

// Common helpers
func BadRequest(c *fiber.Ctx, detail string) error
func NotFound(c *fiber.Ctx, detail string) error
func InternalError(c *fiber.Ctx, detail string) error
func ValidationError(c *fiber.Ctx, detail string) error
func Unauthorized(c *fiber.Ctx, detail string) error
func Forbidden(c *fiber.Ctx, detail string) error
```

### 12.2 Response Usage Examples
```go
// Success responses
return response.Success(c, fiber.StatusOK, data)
return response.Success(c, fiber.StatusCreated, createdResource)

// Error responses
return response.BadRequest(c, "Invalid input format")
return response.NotFound(c, "Resource not found")
return response.ValidationError(c, "Field is required")
return response.InternalError(c, "Service temporarily unavailable")
```

## 13. SECURITY & PERFORMANCE

### 13.1 Security Standards
* **Input Validation:** Validate all user inputs at handler level
* **SQL Injection:** Use Ent ORM's parameterized queries (already safe)
* **Authentication:** Implement JWT-based authentication with proper validation
* **Authorization:** Check user permissions at service layer

### 13.2 Performance Considerations
* **Connection Pooling:** Use efficient database connection pooling
* **Caching:** Implement appropriate caching strategies
* **Memory Management:** Avoid memory leaks by properly closing resources
* **Query Optimization:** Use Ent's query builders efficiently

## 14. DEVELOPMENT WORKFLOW & TOOLS

### 14.1 Code Quality Tools
* **Formatting:** Use `gofmt` with default settings. Run `go fmt ./...` before commits
* **Linting:** Use `golangci-lint` for code quality checks
* **Hot Reloading:** Use Air for development with `make runb` command
* **Build Process:** Use Makefile for common development tasks

### 14.2 Code Generation
* **Ent Generation:** Run `go generate ./internal/ent` for Ent code generation
* **Version Control:** Use conventional commit messages

## 15. CODE REVIEW CHECKLIST

### 15.1 Function Quality
* [ ] Functions are small and focused (under 30 lines)
* [ ] Function names avoid repetition and are context-aware
* [ ] No code duplication (DRY principle followed)
* [ ] Single responsibility principle applied
* [ ] Proper error handling with context wrapping

### 15.2 Naming and Structure
* [ ] Function names clearly describe their purpose
* [ ] Test doubles use appropriate naming patterns
* [ ] No global state without passing litmus tests
* [ ] Dependencies are explicitly injected

### 15.3 Testing and Compliance
* [ ] Tests cover the implemented functionality
* [ ] Test doubles use behavior-based naming
* [ ] Tests are properly isolated
* [ ] Code follows Go idioms and best practices
* [ ] API responses follow OpenAPI 3.0 contract

### 15.4 Response Handling
* [ ] Uses `internal/shared/response` package helpers
* [ ] Follows OpenAPI 3.0 response format exactly
* [ ] Proper HTTP status codes for all responses
* [ ] Consistent error message formatting

### 15.5 Documentation
* [ ] Comments explain WHY, not WHAT
* [ ] All public functions have godoc comments
* [ ] Complex logic is documented
* [ ] API endpoints are documented

## 16. GOOGLE STYLE GUIDE INTEGRATION

This document incorporates key principles from [Google's Go Style Guide](https://google.github.io/styleguide/go/best-practices):

* **Function Naming:** Context-aware naming that avoids repetition
* **Test Double Patterns:** Behavior-based naming for stubs and mocks
* **Global State Avoidance:** Litmus tests for acceptable global state
* **Test Isolation:** Ensuring tests don't interfere with each other
* **Error Handling:** Proper error wrapping and context
* **Dependency Injection:** Explicit dependency passing over global state

*This refactored guideline ensures your Go backend maintains high quality, consistency, and follows both Google's Go Style Guide and your project's specific requirements for domain-driven architecture and OpenAPI 3.0 compliance with standardized response handling.*