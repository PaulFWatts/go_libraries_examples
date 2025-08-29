package calculator

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

// EmailService interface for dependency injection
type EmailService interface {
	SendEmail(to, subject, body string) error
	ValidateEmail(email string) bool
}

// NotificationService depends on EmailService
type NotificationService struct {
	emailService EmailService
}

func NewNotificationService(emailService EmailService) *NotificationService {
	return &NotificationService{emailService: emailService}
}

func (ns *NotificationService) NotifyUser(email, message string) error {
	if !ns.emailService.ValidateEmail(email) {
		return errors.New("invalid email address")
	}

	return ns.emailService.SendEmail(email, "Notification", message)
}

func (ns *NotificationService) SendWelcomeEmail(user User) error {
	if !ns.emailService.ValidateEmail(user.Email) {
		return errors.New("invalid email address")
	}

	subject := "Welcome!"
	body := "Welcome to our service, " + user.Name + "!"

	return ns.emailService.SendEmail(user.Email, subject, body)
}

// MockEmailService is a mock implementation of EmailService
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

// Test using mocks
func TestNotificationService(t *testing.T) {
	t.Run("NotifyUser_Success", func(t *testing.T) {
		// Create mock
		mockEmailService := new(MockEmailService)

		// Set expectations
		mockEmailService.On("ValidateEmail", "test@example.com").Return(true)
		mockEmailService.On("SendEmail", "test@example.com", "Notification", "Hello!").Return(nil)

		// Create service with mock
		notificationService := NewNotificationService(mockEmailService)

		// Test the method
		err := notificationService.NotifyUser("test@example.com", "Hello!")

		// Assertions
		assert.NoError(t, err)

		// Verify all expectations were met
		mockEmailService.AssertExpectations(t)
	})

	t.Run("NotifyUser_InvalidEmail", func(t *testing.T) {
		mockEmailService := new(MockEmailService)

		// Set expectations - invalid email
		mockEmailService.On("ValidateEmail", "invalid-email").Return(false)

		notificationService := NewNotificationService(mockEmailService)

		err := notificationService.NotifyUser("invalid-email", "Hello!")

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid email address")

		// Verify ValidateEmail was called, but SendEmail was not
		mockEmailService.AssertExpectations(t)
	})

	t.Run("NotifyUser_SendEmailFails", func(t *testing.T) {
		mockEmailService := new(MockEmailService)

		// Set expectations - email sending fails
		mockEmailService.On("ValidateEmail", "test@example.com").Return(true)
		mockEmailService.On("SendEmail", "test@example.com", "Notification", "Hello!").
			Return(errors.New("SMTP error"))

		notificationService := NewNotificationService(mockEmailService)

		err := notificationService.NotifyUser("test@example.com", "Hello!")

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "SMTP error")

		mockEmailService.AssertExpectations(t)
	})

	t.Run("SendWelcomeEmail_Success", func(t *testing.T) {
		mockEmailService := new(MockEmailService)

		user := User{
			ID:    1,
			Name:  "John Doe",
			Email: "john@example.com",
		}

		// Set expectations
		mockEmailService.On("ValidateEmail", user.Email).Return(true)
		mockEmailService.On("SendEmail", user.Email, "Welcome!",
			"Welcome to our service, John Doe!").Return(nil)

		notificationService := NewNotificationService(mockEmailService)

		err := notificationService.SendWelcomeEmail(user)

		assert.NoError(t, err)
		mockEmailService.AssertExpectations(t)
	})

	t.Run("MockWithAnyArgs", func(t *testing.T) {
		mockEmailService := new(MockEmailService)

		// Using mock.Anything for flexible matching
		mockEmailService.On("ValidateEmail", mock.Anything).Return(true)
		mockEmailService.On("SendEmail", mock.AnythingOfType("string"),
			mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(nil)

		notificationService := NewNotificationService(mockEmailService)

		err := notificationService.NotifyUser("any@example.com", "Any message")

		assert.NoError(t, err)
		mockEmailService.AssertExpectations(t)
	})

	t.Run("MockWithReturnValues", func(t *testing.T) {
		mockEmailService := new(MockEmailService)

		// Different return values for different inputs
		mockEmailService.On("ValidateEmail", "valid@example.com").Return(true)
		mockEmailService.On("ValidateEmail", "invalid@example.com").Return(false)

		// Test valid email
		assert.True(t, mockEmailService.ValidateEmail("valid@example.com"))

		// Test invalid email
		assert.False(t, mockEmailService.ValidateEmail("invalid@example.com"))

		mockEmailService.AssertExpectations(t)
	})
}

// Database interface for testing
type Database interface {
	GetUser(id int) (*User, error)
	SaveUser(user *User) error
	DeleteUser(id int) error
}

// UserRepository depends on Database
type UserRepository struct {
	db Database
}

func NewUserRepository(db Database) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) FindUser(id int) (*User, error) {
	return ur.db.GetUser(id)
}

func (ur *UserRepository) CreateUser(user *User) error {
	if user.Name == "" {
		return errors.New("name is required")
	}
	return ur.db.SaveUser(user)
}

func (ur *UserRepository) RemoveUser(id int) error {
	// Check if user exists first
	_, err := ur.db.GetUser(id)
	if err != nil {
		return err
	}
	return ur.db.DeleteUser(id)
}

// MockDatabase is a mock implementation of Database
type MockDatabase struct {
	mock.Mock
}

func (m *MockDatabase) GetUser(id int) (*User, error) {
	args := m.Called(id)
	return args.Get(0).(*User), args.Error(1)
}

func (m *MockDatabase) SaveUser(user *User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockDatabase) DeleteUser(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestUserRepository(t *testing.T) {
	t.Run("FindUser_Success", func(t *testing.T) {
		mockDB := new(MockDatabase)
		expectedUser := &User{ID: 1, Name: "John Doe", Email: "john@example.com"}

		mockDB.On("GetUser", 1).Return(expectedUser, nil)

		repo := NewUserRepository(mockDB)
		user, err := repo.FindUser(1)

		assert.NoError(t, err)
		assert.Equal(t, expectedUser, user)
		mockDB.AssertExpectations(t)
	})

	t.Run("FindUser_NotFound", func(t *testing.T) {
		mockDB := new(MockDatabase)

		mockDB.On("GetUser", 999).Return((*User)(nil), errors.New("user not found"))

		repo := NewUserRepository(mockDB)
		user, err := repo.FindUser(999)

		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Contains(t, err.Error(), "user not found")
		mockDB.AssertExpectations(t)
	})

	t.Run("CreateUser_Success", func(t *testing.T) {
		mockDB := new(MockDatabase)
		newUser := &User{Name: "Jane Doe", Email: "jane@example.com"}

		mockDB.On("SaveUser", newUser).Return(nil)

		repo := NewUserRepository(mockDB)
		err := repo.CreateUser(newUser)

		assert.NoError(t, err)
		mockDB.AssertExpectations(t)
	})

	t.Run("CreateUser_InvalidData", func(t *testing.T) {
		mockDB := new(MockDatabase)
		invalidUser := &User{Name: "", Email: "test@example.com"}

		// Note: SaveUser should not be called for invalid data
		repo := NewUserRepository(mockDB)
		err := repo.CreateUser(invalidUser)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "name is required")

		// Verify SaveUser was never called
		mockDB.AssertNotCalled(t, "SaveUser")
	})

	t.Run("RemoveUser_Success", func(t *testing.T) {
		mockDB := new(MockDatabase)
		existingUser := &User{ID: 1, Name: "John Doe"}

		mockDB.On("GetUser", 1).Return(existingUser, nil)
		mockDB.On("DeleteUser", 1).Return(nil)

		repo := NewUserRepository(mockDB)
		err := repo.RemoveUser(1)

		assert.NoError(t, err)
		mockDB.AssertExpectations(t)
	})

	t.Run("RemoveUser_NotFound", func(t *testing.T) {
		mockDB := new(MockDatabase)

		mockDB.On("GetUser", 999).Return((*User)(nil), errors.New("user not found"))

		repo := NewUserRepository(mockDB)
		err := repo.RemoveUser(999)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "user not found")

		// Verify DeleteUser was never called
		mockDB.AssertNotCalled(t, "DeleteUser")
		mockDB.AssertExpectations(t)
	})
}

// Test Suite Example using testify/suite
type CalculatorTestSuite struct {
	suite.Suite
	calculator *Calculator
}

// SetupTest runs before each test method
func (suite *CalculatorTestSuite) SetupTest() {
	suite.calculator = NewCalculator()
}

// TearDownTest runs after each test method
func (suite *CalculatorTestSuite) TearDownTest() {
	suite.calculator.ClearMemory()
}

// SetupSuite runs once before all tests in the suite
func (suite *CalculatorTestSuite) SetupSuite() {
	// Setup code that runs once for the entire suite
}

// TearDownSuite runs once after all tests in the suite
func (suite *CalculatorTestSuite) TearDownSuite() {
	// Cleanup code that runs once for the entire suite
}

func (suite *CalculatorTestSuite) TestAdd() {
	result := suite.calculator.Add(2, 3)
	suite.Equal(5.0, result)
	suite.Equal(5.0, suite.calculator.GetMemory())
}

func (suite *CalculatorTestSuite) TestSubtract() {
	result := suite.calculator.Subtract(10, 3)
	suite.Equal(7.0, result)
	suite.Equal(7.0, suite.calculator.GetMemory())
}

func (suite *CalculatorTestSuite) TestDivideByZero() {
	result, err := suite.calculator.Divide(10, 0)
	suite.Error(err)
	suite.Equal(0.0, result)
	suite.Contains(err.Error(), "division by zero")
}

func (suite *CalculatorTestSuite) TestMemoryOperations() {
	// Test initial memory
	suite.Equal(0.0, suite.calculator.GetMemory())

	// Perform operation
	suite.calculator.Multiply(4, 5)
	suite.Equal(20.0, suite.calculator.GetMemory())

	// Clear memory
	suite.calculator.ClearMemory()
	suite.Equal(0.0, suite.calculator.GetMemory())
}

// Run the test suite
func TestCalculatorTestSuite(t *testing.T) {
	suite.Run(t, new(CalculatorTestSuite))
}

// String Processing Test Suite
type StringProcessorTestSuite struct {
	suite.Suite
	processor *StringProcessor
}

func (suite *StringProcessorTestSuite) SetupTest() {
	suite.processor = NewStringProcessor()
}

func (suite *StringProcessorTestSuite) TestReverse() {
	testCases := map[string]string{
		"hello":   "olleh",
		"world":   "dlrow",
		"testify": "yfitset",
		"":        "",
		"a":       "a",
		"racecar": "racecar",
	}

	for input, expected := range testCases {
		result := suite.processor.Reverse(input)
		suite.Equal(expected, result, "Failed for input: %s", input)
	}
}

func (suite *StringProcessorTestSuite) TestIsPalindrome() {
	palindromes := []string{
		"racecar",
		"A man a plan a canal Panama",
		"race car",
		"",
		"a",
		"Was it a car or a cat I saw",
	}

	notPalindromes := []string{
		"hello",
		"world",
		"testify",
		"almost a palindrome",
	}

	for _, p := range palindromes {
		suite.True(suite.processor.IsPalindrome(p), "Should be palindrome: %s", p)
	}

	for _, np := range notPalindromes {
		suite.False(suite.processor.IsPalindrome(np), "Should not be palindrome: %s", np)
	}
}

func (suite *StringProcessorTestSuite) TestWordCount() {
	testCases := map[string]int{
		"":                   0,
		"hello":              1,
		"hello world":        2,
		"  hello   world  ":  2,
		"one two three four": 4,
		"\thello\nworld\t":   2,
	}

	for input, expected := range testCases {
		result := suite.processor.WordCount(input)
		suite.Equal(expected, result, "Failed for input: '%s'", input)
	}
}

func TestStringProcessorTestSuite(t *testing.T) {
	suite.Run(t, new(StringProcessorTestSuite))
}
