package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

// Config struct to hold our configuration
type Config struct {
	// Database
	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string

	// Server
	ServerPort int
	ServerHost string
	DebugMode  bool

	// Security
	JWTSecret string
	APIKey    string

	// Email
	SMTPHost  string
	SMTPPort  int
	EmailFrom string

	// Redis
	RedisHost     string
	RedisPort     int
	RedisPassword string

	// App
	AppName    string
	AppVersion string
	LogLevel   string
}

func main() {
	fmt.Println("🚀 GoDotEnv Demo Application")
	fmt.Println("============================")

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Printf("Warning: Error loading .env file: %v", err)
		fmt.Println("⚠️  No .env file found, using system environment variables")
	} else {
		fmt.Println("✅ Successfully loaded .env file")
	}

	// Load configuration
	config := loadConfig()

	// Display configuration
	displayConfig(config)

	// Demonstrate different ways to access env vars
	demonstrateUsage()

	// Create database connection string example
	createDatabaseURL(config)

	// Prevent terminal window from closing on Windows
	if runtime.GOOS == "windows" {
		fmt.Println("\nPress Enter to exit...")
		bufio.NewScanner(os.Stdin).Scan()
	}
}

func loadConfig() Config {
	return Config{
		// Database configuration with defaults
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnvAsInt("DB_PORT", 5432),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBName:     getEnv("DB_NAME", "testdb"),
		DBSSLMode:  getEnv("DB_SSL_MODE", "disable"),

		// Server configuration
		ServerPort: getEnvAsInt("SERVER_PORT", 8080),
		ServerHost: getEnv("SERVER_HOST", "localhost"),
		DebugMode:  getEnvAsBool("DEBUG_MODE", false),

		// Security
		JWTSecret: getEnv("JWT_SECRET", ""),
		APIKey:    getEnv("API_KEY", ""),

		// Email
		SMTPHost:  getEnv("SMTP_HOST", ""),
		SMTPPort:  getEnvAsInt("SMTP_PORT", 587),
		EmailFrom: getEnv("EMAIL_FROM", ""),

		// Redis
		RedisHost:     getEnv("REDIS_HOST", "localhost"),
		RedisPort:     getEnvAsInt("REDIS_PORT", 6379),
		RedisPassword: getEnv("REDIS_PASSWORD", ""),

		// Application
		AppName:    getEnv("APP_NAME", "GoDotEnv Demo"),
		AppVersion: getEnv("APP_VERSION", "1.0.0"),
		LogLevel:   getEnv("LOG_LEVEL", "info"),
	}
}

func displayConfig(config Config) {
	fmt.Println("\n📋 Loaded Configuration:")
	fmt.Println("------------------------")

	// Database
	fmt.Printf("🗄️  Database:\n")
	fmt.Printf("   Host: %s:%d\n", config.DBHost, config.DBPort)
	fmt.Printf("   User: %s\n", config.DBUser)
	fmt.Printf("   Password: %s\n", maskPassword(config.DBPassword))
	fmt.Printf("   Database: %s\n", config.DBName)
	fmt.Printf("   SSL Mode: %s\n", config.DBSSLMode)

	// Server
	fmt.Printf("\n🌐 Server:\n")
	fmt.Printf("   Address: %s:%d\n", config.ServerHost, config.ServerPort)
	fmt.Printf("   Debug Mode: %t\n", config.DebugMode)

	// Security
	fmt.Printf("\n🔐 Security:\n")
	fmt.Printf("   JWT Secret: %s\n", maskSecret(config.JWTSecret))
	fmt.Printf("   API Key: %s\n", maskSecret(config.APIKey))

	// Email
	if config.SMTPHost != "" {
		fmt.Printf("\n📧 Email:\n")
		fmt.Printf("   SMTP: %s:%d\n", config.SMTPHost, config.SMTPPort)
		fmt.Printf("   From: %s\n", config.EmailFrom)
	}

	// Redis
	fmt.Printf("\n🔴 Redis:\n")
	fmt.Printf("   Address: %s:%d\n", config.RedisHost, config.RedisPort)
	fmt.Printf("   Password: %s\n", maskPassword(config.RedisPassword))

	// App
	fmt.Printf("\n📱 Application:\n")
	fmt.Printf("   Name: %s\n", config.AppName)
	fmt.Printf("   Version: %s\n", config.AppVersion)
	fmt.Printf("   Log Level: %s\n", config.LogLevel)
}

func demonstrateUsage() {
	fmt.Println("\n🔍 Environment Variable Access Methods:")
	fmt.Println("---------------------------------------")

	// Method 1: Direct os.Getenv
	dbHost := os.Getenv("DB_HOST")
	fmt.Printf("1. os.Getenv(\"DB_HOST\"): '%s'\n", dbHost)

	// Method 2: With default value
	port := getEnv("SERVER_PORT", "3000")
	fmt.Printf("2. With default: SERVER_PORT = '%s'\n", port)

	// Method 3: As integer
	dbPort := getEnvAsInt("DB_PORT", 5432)
	fmt.Printf("3. As integer: DB_PORT = %d\n", dbPort)

	// Method 4: As boolean
	debug := getEnvAsBool("DEBUG_MODE", false)
	fmt.Printf("4. As boolean: DEBUG_MODE = %t\n", debug)

	// Method 5: Check if variable exists
	if value, exists := os.LookupEnv("API_KEY"); exists {
		fmt.Printf("5. API_KEY exists: '%s'\n", maskSecret(value))
	} else {
		fmt.Printf("5. API_KEY not set\n")
	}
}

func createDatabaseURL(config Config) {
	fmt.Println("\n🔗 Database Connection Examples:")
	fmt.Println("---------------------------------")

	// PostgreSQL connection string
	postgresURL := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		config.DBUser,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBName,
		config.DBSSLMode,
	)

	// GORM DSN
	gormDSN := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		config.DBHost,
		config.DBUser,
		config.DBPassword,
		config.DBName,
		config.DBPort,
		config.DBSSLMode,
	)

	fmt.Printf("PostgreSQL URL: %s\n", maskConnectionString(postgresURL))
	fmt.Printf("GORM DSN: %s\n", maskConnectionString(gormDSN))
}

// Helper functions
func getEnv(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultVal
}

func getEnvAsBool(name string, defaultVal bool) bool {
	valStr := getEnv(name, "")
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}
	return defaultVal
}

func maskPassword(password string) string {
	if password == "" {
		return "<empty>"
	}
	if len(password) <= 4 {
		return strings.Repeat("*", len(password))
	}
	return password[:2] + strings.Repeat("*", len(password)-4) + password[len(password)-2:]
}

func maskSecret(secret string) string {
	if secret == "" {
		return "<not set>"
	}
	if len(secret) <= 8 {
		return strings.Repeat("*", len(secret))
	}
	return secret[:4] + strings.Repeat("*", len(secret)-8) + secret[len(secret)-4:]
}

func maskConnectionString(connStr string) string {
	// Simple password masking in connection strings
	if strings.Contains(connStr, "password=") {
		parts := strings.Split(connStr, " ")
		for i, part := range parts {
			if strings.HasPrefix(part, "password=") {
				passwordPart := strings.Split(part, "=")
				if len(passwordPart) == 2 {
					parts[i] = "password=" + maskPassword(passwordPart[1])
				}
			}
		}
		return strings.Join(parts, " ")
	}

	// For postgres:// URLs
	if strings.Contains(connStr, "://") && strings.Contains(connStr, ":") {
		// This is a basic implementation - in production use proper URL parsing
		return strings.ReplaceAll(connStr, strings.Split(strings.Split(connStr, "://")[1], "@")[0], "***:***")
	}

	return connStr
}
