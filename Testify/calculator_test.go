package calculator

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestCalculatorBasicOperations tests basic arithmetic operations
func TestCalculatorBasicOperations(t *testing.T) {
	calc := NewCalculator()

	// Test Addition
	t.Run("Addition", func(t *testing.T) {
		result := calc.Add(5, 3)
		assert.Equal(t, 8.0, result, "Addition should be correct")
		assert.Equal(t, 8.0, calc.GetMemory(), "Memory should store the result")
	})

	// Test Subtraction
	t.Run("Subtraction", func(t *testing.T) {
		result := calc.Subtract(10, 4)
		assert.Equal(t, 6.0, result, "Subtraction should be correct")
		assert.Equal(t, 6.0, calc.GetMemory(), "Memory should be updated")
	})

	// Test Multiplication
	t.Run("Multiplication", func(t *testing.T) {
		result := calc.Multiply(7, 8)
		assert.Equal(t, 56.0, result, "Multiplication should be correct")
		assert.Equal(t, 56.0, calc.GetMemory())
	})

	// Test Division
	t.Run("Division", func(t *testing.T) {
		result, err := calc.Divide(15, 3)
		assert.NoError(t, err, "Division should not return error")
		assert.Equal(t, 5.0, result, "Division should be correct")
		assert.Equal(t, 5.0, calc.GetMemory())
	})

	// Test Division by Zero
	t.Run("DivisionByZero", func(t *testing.T) {
		result, err := calc.Divide(10, 0)
		assert.Error(t, err, "Division by zero should return error")
		assert.Equal(t, 0.0, result, "Result should be zero on error")
		assert.Contains(t, err.Error(), "division by zero")
	})
}

// TestCalculatorAdvancedOperations tests advanced mathematical operations
func TestCalculatorAdvancedOperations(t *testing.T) {
	calc := NewCalculator()

	// Test Square Root
	t.Run("SquareRoot", func(t *testing.T) {
		result, err := calc.Sqrt(16)
		assert.NoError(t, err)
		assert.Equal(t, 4.0, result)
		
		// Test negative number
		result, err = calc.Sqrt(-4)
		assert.Error(t, err)
		assert.Equal(t, 0.0, result)
		assert.Contains(t, err.Error(), "negative number")
	})

	// Test Power
	t.Run("Power", func(t *testing.T) {
		result := calc.Power(2, 3)
		assert.Equal(t, 8.0, result)
		
		result = calc.Power(5, 0)
		assert.Equal(t, 1.0, result)
		
		result = calc.Power(10, 2)
		assert.Equal(t, 100.0, result)
	})
}

// TestCalculatorMemoryOperations tests memory-related functionality
func TestCalculatorMemoryOperations(t *testing.T) {
	calc := NewCalculator()

	// Initial memory should be zero
	assert.Equal(t, 0.0, calc.GetMemory())

	// Perform operation and check memory
	calc.Add(10, 5)
	assert.Equal(t, 15.0, calc.GetMemory())

	// Clear memory
	calc.ClearMemory()
	assert.Equal(t, 0.0, calc.GetMemory())
}

// TestUtilityFunctions tests standalone utility functions
func TestUtilityFunctions(t *testing.T) {
	// Test IsEven
	t.Run("IsEven", func(t *testing.T) {
		assert.True(t, IsEven(4))
		assert.True(t, IsEven(0))
		assert.False(t, IsEven(3))
		assert.False(t, IsEven(-1))
	})

	// Test IsPositive
	t.Run("IsPositive", func(t *testing.T) {
		assert.True(t, IsPositive(1.5))
		assert.True(t, IsPositive(0.1))
		assert.False(t, IsPositive(0))
		assert.False(t, IsPositive(-1.5))
	})
}

// TestFibonacci tests the Fibonacci function with various edge cases
func TestFibonacci(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
		hasError bool
	}{
		{"Fibonacci of 0", 0, 0, false},
		{"Fibonacci of 1", 1, 1, false},
		{"Fibonacci of 2", 2, 1, false},
		{"Fibonacci of 3", 3, 2, false},
		{"Fibonacci of 5", 5, 5, false},
		{"Fibonacci of 10", 10, 55, false},
		{"Negative input", -1, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Fibonacci(tt.input)
			
			if tt.hasError {
				assert.Error(t, err)
				assert.Equal(t, 0, result)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

// TestStringProcessor tests string processing functionality
func TestStringProcessor(t *testing.T) {
	sp := NewStringProcessor()

	t.Run("Reverse", func(t *testing.T) {
		assert.Equal(t, "olleh", sp.Reverse("hello"))
		assert.Equal(t, "a", sp.Reverse("a"))
		assert.Equal(t, "", sp.Reverse(""))
		assert.Equal(t, "tset", sp.Reverse("test"))
	})

	t.Run("IsPalindrome", func(t *testing.T) {
		assert.True(t, sp.IsPalindrome("racecar"))
		assert.True(t, sp.IsPalindrome("A man a plan a canal Panama"))
		assert.True(t, sp.IsPalindrome("race car"))
		assert.False(t, sp.IsPalindrome("hello"))
		assert.True(t, sp.IsPalindrome(""))
		assert.True(t, sp.IsPalindrome("a"))
	})

	t.Run("WordCount", func(t *testing.T) {
		assert.Equal(t, 0, sp.WordCount(""))
		assert.Equal(t, 1, sp.WordCount("hello"))
		assert.Equal(t, 2, sp.WordCount("hello world"))
		assert.Equal(t, 3, sp.WordCount("  hello   world  test  "))
		assert.Equal(t, 1, sp.WordCount("test"))
	})
}

// TestUserService tests the UserService functionality
func TestUserService(t *testing.T) {
	userService := NewUserService()

	t.Run("GetUser", func(t *testing.T) {
		// Test existing user
		user, err := userService.GetUser(1)
		assert.NoError(t, err)
		require.NotNil(t, user) // Use require for critical assertions
		assert.Equal(t, 1, user.ID)
		assert.Equal(t, "John Doe", user.Name)
		assert.Equal(t, "john@example.com", user.Email)

		// Test non-existing user
		user, err = userService.GetUser(999)
		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Contains(t, err.Error(), "user not found")
	})

	t.Run("AddUser", func(t *testing.T) {
		newUser := User{
			ID:       3,
			Name:     "Alice Johnson",
			Email:    "alice@example.com",
			Username: "alice_j",
			Age:      28,
		}

		err := userService.AddUser(newUser)
		assert.NoError(t, err)

		// Verify user was added
		user, err := userService.GetUser(3)
		assert.NoError(t, err)
		assert.Equal(t, newUser.Name, user.Name)

		// Test adding user with empty name
		invalidUser := User{ID: 4, Name: "", Email: "test@example.com"}
		err = userService.AddUser(invalidUser)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "name cannot be empty")
	})

	t.Run("UpdateUser", func(t *testing.T) {
		updatedUser := User{
			ID:       1,
			Name:     "John Updated",
			Email:    "john.updated@example.com",
			Username: "john_updated",
			Age:      31,
		}

		err := userService.UpdateUser(1, updatedUser)
		assert.NoError(t, err)

		// Verify user was updated
		user, err := userService.GetUser(1)
		assert.NoError(t, err)
		assert.Equal(t, "John Updated", user.Name)
		assert.Equal(t, "john.updated@example.com", user.Email)

		// Test updating non-existing user
		err = userService.UpdateUser(999, updatedUser)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "user not found")
	})

	t.Run("DeleteUser", func(t *testing.T) {
		// First, get the initial count
		initialCount := userService.GetUserCount()

		err := userService.DeleteUser(2)
		assert.NoError(t, err)

		// Verify user was deleted
		user, err := userService.GetUser(2)
		assert.Error(t, err)
		assert.Nil(t, user)

		// Verify count decreased
		assert.Equal(t, initialCount-1, userService.GetUserCount())

		// Test deleting non-existing user
		err = userService.DeleteUser(999)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "user not found")
	})

	t.Run("GetAllUsers", func(t *testing.T) {
		users := userService.GetAllUsers()
		assert.NotEmpty(t, users)
		assert.IsType(t, []User{}, users)
	})

	t.Run("ValidateUser", func(t *testing.T) {
		// Valid user
		validUser := User{
			ID:       1,
			Name:     "Valid User",
			Email:    "valid@example.com",
			Username: "valid_user",
			Age:      25,
		}
		errors := userService.ValidateUser(validUser)
		assert.Empty(t, errors)

		// Invalid user - multiple errors
		invalidUser := User{
			ID:       2,
			Name:     "",     // Empty name
			Email:    "",     // Empty email
			Username: "",     // Empty username
			Age:      -5,     // Negative age
		}
		errors = userService.ValidateUser(invalidUser)
		assert.NotEmpty(t, errors)
		assert.Contains(t, errors, "name cannot be empty")
		assert.Contains(t, errors, "email cannot be empty")
		assert.Contains(t, errors, "username cannot be empty")
		assert.Contains(t, errors, "age cannot be negative")

		// Age too high
		oldUser := User{
			ID:       3,
			Name:     "Old User",
			Email:    "old@example.com",
			Username: "old_user",
			Age:      200, // Too old
		}
		errors = userService.ValidateUser(oldUser)
		assert.Contains(t, errors, "age cannot be greater than 150")
	})
}

// Benchmark tests for performance measurement
func BenchmarkCalculatorAdd(b *testing.B) {
	calc := NewCalculator()
	for i := 0; i < b.N; i++ {
		calc.Add(float64(i), float64(i+1))
	}
}

func BenchmarkFibonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = Fibonacci(20)
	}
}

func BenchmarkStringReverse(b *testing.B) {
	sp := NewStringProcessor()
	testString := "This is a test string for benchmarking"
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sp.Reverse(testString)
	}
}

// Example test showing different assertion styles
func TestAssertionStyles(t *testing.T) {
	t.Run("BasicAssertions", func(t *testing.T) {
		// Basic equality
		assert.Equal(t, 42, 42)
		assert.NotEqual(t, 42, 43)

		// Approximate equality for floats
		assert.InDelta(t, 1.0, 1.1, 0.2)

		// Boolean assertions
		assert.True(t, true)
		assert.False(t, false)

		// Nil checks
		var ptr *int
		assert.Nil(t, ptr)
		
		value := 42
		ptr = &value
		assert.NotNil(t, ptr)
	})

	t.Run("StringAssertions", func(t *testing.T) {
		text := "Hello, World!"
		
		assert.Contains(t, text, "World")
		assert.NotContains(t, text, "Go")
		
		assert.True(t, strings.HasPrefix(text, "Hello"))
		assert.True(t, strings.HasSuffix(text, "World!"))
		
		assert.Regexp(t, `^Hello.*!$`, text)
	})

	t.Run("CollectionAssertions", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		
		assert.Len(t, numbers, 5)
		assert.NotEmpty(t, numbers)
		assert.Contains(t, numbers, 3)
		assert.NotContains(t, numbers, 10)
		
		// Element-wise comparison
		expected := []int{1, 2, 3, 4, 5}
		assert.ElementsMatch(t, expected, numbers)
		
		// Subset checking
		subset := []int{2, 4}
		assert.Subset(t, numbers, subset)
	})

	t.Run("TypeAssertions", func(t *testing.T) {
		var value interface{} = "hello"
		
		assert.IsType(t, "", value)
		assert.Implements(t, (*error)(nil), &CustomError{})
	})

	t.Run("PanicAssertions", func(t *testing.T) {
		assert.Panics(t, func() {
			panic("test panic")
		})
		
		assert.NotPanics(t, func() {
			// This should not panic
			_ = 1 + 1
		})
	})
}

// CustomError is a custom error type for testing
type CustomError struct {
	message string
}

func (e *CustomError) Error() string {
	return e.message
}

// TestRequireVsAssert demonstrates the difference between require and assert
func TestRequireVsAssert(t *testing.T) {
	t.Run("RequireStopsOnFailure", func(t *testing.T) {
		data := []int{1, 2, 3}
		
		// If this fails, the test stops here
		require.NotEmpty(t, data)
		require.Len(t, data, 3)
		
		// These will only run if the requires above pass
		assert.Equal(t, 1, data[0])
		assert.Equal(t, 2, data[1])
		assert.Equal(t, 3, data[2])
	})

	t.Run("AssertContinuesOnFailure", func(t *testing.T) {
		// All these assertions will run even if earlier ones fail
		assert.Equal(t, 1, 1, "First assertion")
		assert.Equal(t, 2, 2, "Second assertion")
		assert.Equal(t, 3, 3, "Third assertion")
	})
}
