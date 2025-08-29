package calculator

import (
	"errors"
	"math"
)

// Calculator represents a simple calculator
type Calculator struct {
	memory float64
}

// NewCalculator creates a new calculator instance
func NewCalculator() *Calculator {
	return &Calculator{memory: 0}
}

// Add performs addition
func (c *Calculator) Add(a, b float64) float64 {
	result := a + b
	c.memory = result
	return result
}

// Subtract performs subtraction
func (c *Calculator) Subtract(a, b float64) float64 {
	result := a - b
	c.memory = result
	return result
}

// Multiply performs multiplication
func (c *Calculator) Multiply(a, b float64) float64 {
	result := a * b
	c.memory = result
	return result
}

// Divide performs division
func (c *Calculator) Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	result := a / b
	c.memory = result
	return result, nil
}

// Sqrt calculates square root
func (c *Calculator) Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, errors.New("negative number")
	}
	result := math.Sqrt(a)
	c.memory = result
	return result, nil
}

// Power calculates a^b
func (c *Calculator) Power(base, exponent float64) float64 {
	result := math.Pow(base, exponent)
	c.memory = result
	return result
}

// GetMemory returns the last calculated value
func (c *Calculator) GetMemory() float64 {
	return c.memory
}

// ClearMemory resets the memory to zero
func (c *Calculator) ClearMemory() {
	c.memory = 0
}

// IsEven checks if a number is even
func IsEven(n int) bool {
	return n%2 == 0
}

// IsPositive checks if a number is positive
func IsPositive(n float64) bool {
	return n > 0
}

// Fibonacci generates the nth Fibonacci number
func Fibonacci(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("negative input not allowed")
	}
	if n <= 1 {
		return n, nil
	}

	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b, nil
}

// StringProcessor represents a string processing utility
type StringProcessor struct{}

// NewStringProcessor creates a new string processor
func NewStringProcessor() *StringProcessor {
	return &StringProcessor{}
}

// Reverse reverses a string
func (sp *StringProcessor) Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// IsPalindrome checks if a string is a palindrome
func (sp *StringProcessor) IsPalindrome(s string) bool {
	cleaned := ""
	for _, r := range s {
		if r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' || r >= '0' && r <= '9' {
			if r >= 'A' && r <= 'Z' {
				r = r + 32 // Convert to lowercase
			}
			cleaned += string(r)
		}
	}
	return cleaned == sp.Reverse(cleaned)
}

// WordCount counts words in a string
func (sp *StringProcessor) WordCount(s string) int {
	if len(s) == 0 {
		return 0
	}

	count := 0
	inWord := false

	for _, r := range s {
		if r == ' ' || r == '\t' || r == '\n' {
			inWord = false
		} else if !inWord {
			inWord = true
			count++
		}
	}
	return count
}

// User represents a user in our system
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Age      int    `json:"age"`
}

// UserService represents a user service
type UserService struct {
	users []User
}

// NewUserService creates a new user service
func NewUserService() *UserService {
	return &UserService{
		users: []User{
			{ID: 1, Name: "John Doe", Email: "john@example.com", Username: "john_doe", Age: 30},
			{ID: 2, Name: "Jane Smith", Email: "jane@example.com", Username: "jane_smith", Age: 25},
		},
	}
}

// GetUser retrieves a user by ID
func (us *UserService) GetUser(id int) (*User, error) {
	for _, user := range us.users {
		if user.ID == id {
			return &user, nil
		}
	}
	return nil, errors.New("user not found")
}

// AddUser adds a new user
func (us *UserService) AddUser(user User) error {
	if user.Name == "" {
		return errors.New("name cannot be empty")
	}
	if user.Email == "" {
		return errors.New("email cannot be empty")
	}

	// Check for duplicate ID
	for _, existingUser := range us.users {
		if existingUser.ID == user.ID {
			return errors.New("user with this ID already exists")
		}
	}

	us.users = append(us.users, user)
	return nil
}

// UpdateUser updates an existing user
func (us *UserService) UpdateUser(id int, updatedUser User) error {
	for i, user := range us.users {
		if user.ID == id {
			updatedUser.ID = id // Preserve the original ID
			us.users[i] = updatedUser
			return nil
		}
	}
	return errors.New("user not found")
}

// DeleteUser deletes a user by ID
func (us *UserService) DeleteUser(id int) error {
	for i, user := range us.users {
		if user.ID == id {
			us.users = append(us.users[:i], us.users[i+1:]...)
			return nil
		}
	}
	return errors.New("user not found")
}

// GetAllUsers returns all users
func (us *UserService) GetAllUsers() []User {
	return us.users
}

// GetUserCount returns the number of users
func (us *UserService) GetUserCount() int {
	return len(us.users)
}

// ValidateUser validates user data
func (us *UserService) ValidateUser(user User) []string {
	var errors []string

	if user.Name == "" {
		errors = append(errors, "name cannot be empty")
	}
	if user.Email == "" {
		errors = append(errors, "email cannot be empty")
	}
	if user.Username == "" {
		errors = append(errors, "username cannot be empty")
	}
	if user.Age < 0 {
		errors = append(errors, "age cannot be negative")
	}
	if user.Age > 150 {
		errors = append(errors, "age cannot be greater than 150")
	}

	return errors
}
