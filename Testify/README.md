# Testify Demo

A comprehensive demonstration of the `github.com/stretchr/testify` library, the most popular testing framework for Go applications, showcasing assertions, mocks, test suites, and advanced testing patterns.

## 🚀 Features

This demo covers all major Testify features:
- **Rich Assertions** - Over 60 built-in assertion methods
- **Mocking Framework** - Powerful mocking capabilities with expectations
- **Test Suites** - Organized test execution with setup/teardown
- **Require vs Assert** - Different failure behaviors
- **Benchmarking** - Performance testing examples
- **Real-world Examples** - Practical testing scenarios

## 📦 Dependencies

```bash
go get github.com/stretchr/testify
```

## 🔧 Setup

1. **Install dependencies:**
   ```bash
   go mod tidy
   ```

2. **Run all tests:**
   ```bash
   go test
   ```

3. **Run tests with verbose output:**
   ```bash
   go test -v
   ```

4. **Run specific tests:**
   ```bash
   go test -v -run TestCalculator
   ```

5. **Run benchmarks:**
   ```bash
   go test -bench=.
   ```

6. **Run with coverage:**
   ```bash
   go test -cover
   ```

## 📋 What It Demonstrates

### 1. Rich Assertion Library

#### Basic Assertions
```go
func TestBasicAssertions(t *testing.T) {
    // Equality
    assert.Equal(t, expected, actual)
    assert.NotEqual(t, expected, actual)
    
    // Boolean
    assert.True(t, condition)
    assert.False(t, condition)
    
    // Nil checks
    assert.Nil(t, object)
    assert.NotNil(t, object)
    
    // Error handling
    assert.NoError(t, err)
    assert.Error(t, err)
}
```

#### String Assertions
```go
func TestStringAssertions(t *testing.T) {
    text := "Hello, World!"
    
    assert.Contains(t, text, "World")
    assert.NotContains(t, text, "Go")
    assert.HasPrefix(t, text, "Hello")
    assert.HasSuffix(t, text, "World!")
    assert.Regexp(t, `^Hello.*!$`, text)
}
```

#### Collection Assertions
```go
func TestCollectionAssertions(t *testing.T) {
    numbers := []int{1, 2, 3, 4, 5}
    
    assert.Len(t, numbers, 5)
    assert.NotEmpty(t, numbers)
    assert.Contains(t, numbers, 3)
    assert.ElementsMatch(t, expected, numbers)
    assert.Subset(t, numbers, subset)
}
```

### 2. Mocking Framework

#### Interface Mocking
```go
// Define interface
type EmailService interface {
    SendEmail(to, subject, body string) error
    ValidateEmail(email string) bool
}

// Create mock
type MockEmailService struct {
    mock.Mock
}

func (m *MockEmailService) SendEmail(to, subject, body string) error {
    args := m.Called(to, subject, body)
    return args.Error(0)
}

func (m *MockEmailService) ValidateEmail(email string) bool {
    args := m.Called(email)
    return args.Bool(0)
}
```

#### Using Mocks in Tests
```go
func TestWithMocks(t *testing.T) {
    mockService := new(MockEmailService)
    
    // Set expectations
    mockService.On("ValidateEmail", "test@example.com").Return(true)
    mockService.On("SendEmail", "test@example.com", "Subject", "Body").Return(nil)
    
    // Use mock in your code
    service := NewNotificationService(mockService)
    err := service.NotifyUser("test@example.com", "Hello!")
    
    // Verify expectations
    assert.NoError(t, err)
    mockService.AssertExpectations(t)
}
```

### 3. Test Suites

#### Suite Structure
```go
type CalculatorTestSuite struct {
    suite.Suite
    calculator *Calculator
}

// Setup runs before each test
func (suite *CalculatorTestSuite) SetupTest() {
    suite.calculator = NewCalculator()
}

// Teardown runs after each test
func (suite *CalculatorTestSuite) TearDownTest() {
    suite.calculator.ClearMemory()
}

func (suite *CalculatorTestSuite) TestAdd() {
    result := suite.calculator.Add(2, 3)
    suite.Equal(5.0, result)
}

// Run the suite
func TestCalculatorTestSuite(t *testing.T) {
    suite.Run(t, new(CalculatorTestSuite))
}
```

### 4. Require vs Assert

```go
func TestRequireVsAssert(t *testing.T) {
    data := []int{1, 2, 3}
    
    // Require stops execution on failure
    require.NotEmpty(t, data)    // If this fails, test stops
    require.Len(t, data, 3)      // This won't run if above fails
    
    // Assert continues execution on failure
    assert.Equal(t, 1, data[0])  // All assertions run
    assert.Equal(t, 2, data[1])  // even if earlier ones fail
    assert.Equal(t, 3, data[2])
}
```

## 🧪 Testing Scenarios Covered

### 1. Calculator Operations
- Basic arithmetic (add, subtract, multiply, divide)
- Advanced operations (sqrt, power)
- Error handling (division by zero, negative sqrt)
- Memory operations
- Edge cases and boundary conditions

### 2. String Processing
- String reversal algorithms
- Palindrome detection with normalization
- Word counting with whitespace handling
- Unicode and special character support

### 3. User Management System
- CRUD operations testing
- Input validation
- Error condition testing
- Business logic validation
- Data integrity checks

### 4. Service Layer Testing
- Dependency injection with mocks
- External service integration
- Error propagation
- Complex business workflows

### 5. Repository Pattern Testing
- Database interaction mocking
- Query result testing
- Transaction handling
- Error condition simulation

## 🎯 Advanced Testing Patterns

### 1. Table-Driven Tests
```go
func TestFibonacci(t *testing.T) {
    tests := []struct {
        name     string
        input    int
        expected int
        hasError bool
    }{
        {"Fibonacci of 0", 0, 0, false},
        {"Fibonacci of 1", 1, 1, false},
        {"Fibonacci of 5", 5, 5, false},
        {"Negative input", -1, 0, true},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result, err := Fibonacci(tt.input)
            
            if tt.hasError {
                assert.Error(t, err)
            } else {
                assert.NoError(t, err)
                assert.Equal(t, tt.expected, result)
            }
        })
    }
}
```

### 2. Subtests Organization
```go
func TestCalculatorOperations(t *testing.T) {
    calc := NewCalculator()

    t.Run("Addition", func(t *testing.T) {
        result := calc.Add(5, 3)
        assert.Equal(t, 8.0, result)
    })

    t.Run("Division", func(t *testing.T) {
        t.Run("ValidDivision", func(t *testing.T) {
            result, err := calc.Divide(10, 2)
            assert.NoError(t, err)
            assert.Equal(t, 5.0, result)
        })

        t.Run("DivisionByZero", func(t *testing.T) {
            result, err := calc.Divide(10, 0)
            assert.Error(t, err)
            assert.Equal(t, 0.0, result)
        })
    })
}
```

### 3. Mock Expectations
```go
func TestMockExpectations(t *testing.T) {
    mockService := new(MockEmailService)
    
    // Exact parameter matching
    mockService.On("SendEmail", "user@example.com", "Subject", "Body").Return(nil)
    
    // Flexible parameter matching
    mockService.On("ValidateEmail", mock.Anything).Return(true)
    mockService.On("SendEmail", mock.AnythingOfType("string"), 
        mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(nil)
    
    // Conditional returns
    mockService.On("ValidateEmail", "valid@example.com").Return(true)
    mockService.On("ValidateEmail", "invalid@example.com").Return(false)
    
    // Verify specific methods were called
    mockService.AssertCalled(t, "ValidateEmail", "user@example.com")
    mockService.AssertNotCalled(t, "SendEmail")
    mockService.AssertExpectations(t)
}
```

## 🏃‍♂️ Running Tests

### Basic Test Execution
```bash
# Run all tests
go test

# Verbose output
go test -v

# Run specific test
go test -v -run TestCalculator

# Run specific subtest
go test -v -run TestCalculator/Addition
```

### Coverage Analysis
```bash
# Basic coverage
go test -cover

# Detailed coverage report
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
```

### Benchmarking
```bash
# Run all benchmarks
go test -bench=.

# Run specific benchmark
go test -bench=BenchmarkCalculator

# Benchmark with memory stats
go test -bench=. -benchmem
```

### Parallel Testing
```bash
# Run tests in parallel
go test -parallel 4

# Control parallelism
go test -cpu=1,2,4
```

## 🔍 Test Output Examples

### Successful Test Run
```
=== RUN   TestCalculatorBasicOperations
=== RUN   TestCalculatorBasicOperations/Addition
=== RUN   TestCalculatorBasicOperations/Subtraction
=== RUN   TestCalculatorBasicOperations/Multiplication
=== RUN   TestCalculatorBasicOperations/Division
=== RUN   TestCalculatorBasicOperations/DivisionByZero
--- PASS: TestCalculatorBasicOperations (0.00s)
    --- PASS: TestCalculatorBasicOperations/Addition (0.00s)
    --- PASS: TestCalculatorBasicOperations/Subtraction (0.00s)
    --- PASS: TestCalculatorBasicOperations/Multiplication (0.00s)
    --- PASS: TestCalculatorBasicOperations/Division (0.00s)
    --- PASS: TestCalculatorBasicOperations/DivisionByZero (0.00s)
```

### Failed Test Output
```
=== RUN   TestUserService/ValidateUser
--- FAIL: TestUserService/ValidateUser (0.00s)
    calculator_test.go:425: 
        Error:          Not equal: 
                        expected: []string{"name cannot be empty"}
                        actual  : []string{"name cannot be empty", "email cannot be empty"}
                        
                        Diff:
                        --- Expected
                        +++ Actual
                        @@ -1,3 +1,4 @@
                         []string{
                           "name cannot be empty",
                        +  "email cannot be empty",
                         }
```

### Mock Verification Output
```
=== RUN   TestNotificationService/NotifyUser_Success
--- PASS: TestNotificationService/NotifyUser_Success (0.00s)

=== RUN   TestNotificationService/NotifyUser_InvalidEmail
--- PASS: TestNotificationService/NotifyUser_InvalidEmail (0.00s)
    mock_test.go:95: PASS: ValidateEmail(string)
```

### Benchmark Results
```
BenchmarkCalculatorAdd-8        100000000    10.2 ns/op       0 B/op    0 allocs/op
BenchmarkFibonacci-8              300000    4521 ns/op         0 B/op    0 allocs/op
BenchmarkStringReverse-8         5000000     298 ns/op        32 B/op    1 allocs/op
```

## 🛠️ Best Practices Demonstrated

### 1. Test Organization
- **Descriptive test names** that explain what is being tested
- **Logical grouping** using subtests and suites
- **Table-driven tests** for multiple similar scenarios
- **Proper setup and teardown** for test isolation

### 2. Assertion Selection
- **Use `require`** when failure should stop the test immediately
- **Use `assert`** when you want to continue checking other conditions
- **Specific assertions** over generic ones (`assert.Empty` vs `assert.Len(x, 0)`)

### 3. Mock Usage
- **Mock external dependencies** not internal logic
- **Set specific expectations** rather than using `mock.Anything` everywhere
- **Verify all expectations** are met
- **Use meaningful mock names** that reflect their purpose

### 4. Error Testing
- **Test both success and failure paths**
- **Verify error messages** contain expected content
- **Test edge cases** and boundary conditions
- **Ensure proper error propagation**

### 5. Test Data Management
- **Use test fixtures** for complex data setup
- **Isolate tests** to prevent data contamination
- **Create realistic test data** that represents actual usage

## 🚀 Integration with CI/CD

### GitHub Actions Example
```yaml
name: Tests
on: [push, pull_request]
jobs:
  test:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v2
    - uses: actions/setup-go@v2
      with:
        go-version: 1.21
    - run: go test -v -cover ./...
```

### Coverage Reporting
```bash
# Generate coverage report
go test -coverprofile=coverage.out ./...

# View coverage in terminal
go tool cover -func=coverage.out

# Generate HTML report
go tool cover -html=coverage.out -o coverage.html
```

## 🆚 Testify vs Standard Testing

| Feature | Standard `testing` | Testify |
|---------|-------------------|---------|
| Basic assertions | Manual `if` statements | Rich assertion library |
| Error messages | Custom formatting | Automatic descriptive messages |
| Mocking | Manual or third-party | Built-in mock framework |
| Test organization | Basic | Suites with setup/teardown |
| Readability | Verbose | Concise and expressive |
| Learning curve | Minimal | Moderate |

## 📚 Real-World Applications

### 1. Web API Testing
- HTTP handler testing with mocked dependencies
- Request/response validation
- Authentication and authorization testing
- Error handling and status codes

### 2. Database Layer Testing
- Repository pattern testing with mock databases
- Query validation and result processing
- Transaction handling and rollback scenarios
- Connection pooling and error recovery

### 3. Business Logic Testing
- Complex calculation validation
- Workflow and state machine testing
- Business rule enforcement
- Integration between multiple services

### 4. Microservices Testing
- Service interaction mocking
- Message queue and event testing
- Circuit breaker and retry logic
- Service discovery and health checks

---

This comprehensive Testify demo provides a solid foundation for implementing robust testing strategies in Go applications, covering everything from basic assertions to advanced mocking patterns and test organization techniques.
