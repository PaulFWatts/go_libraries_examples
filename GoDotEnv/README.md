# GoDotEnv Demo

This is a comprehensive demonstration of using the `godotenv` library in Go to manage environment variables from `.env` files.

## 🚀 Features

- **Complete `.env` file handling**
- **Type conversion helpers** (string, int, bool)
- **Default value support**
- **Security-conscious display** (masks passwords and secrets)
- **Multiple configuration examples**
- **Database connection string generation**
- **Cross-platform terminal handling**

## 📁 Files

- `main.go` - Main application with comprehensive examples
- `.env` - Environment variables (modify for your setup)
- `.env.example` - Template file for sharing
- `README.md` - This documentation
- `go.mod` - Go module definition

## 🔧 Setup

1. **Install dependencies:**
   ```bash
   go mod tidy
   ```

2. **Modify the `.env` file** with your actual values

3. **Run the demo:**
   ```bash
   go run main.go
   ```

   Or build and run:
   ```bash
   go build -o godotenv-demo
   ./godotenv-demo
   ```

## 📋 Environment Variables

### Database Configuration
- `DB_HOST` - Database host (default: localhost)
- `DB_PORT` - Database port (default: 5432)
- `DB_USER` - Database username
- `DB_PASSWORD` - Database password
- `DB_NAME` - Database name
- `DB_SSL_MODE` - SSL mode (default: disable)

### Server Configuration
- `SERVER_PORT` - Server port (default: 8080)
- `SERVER_HOST` - Server host (default: localhost)
- `DEBUG_MODE` - Enable debug mode (default: false)

### Security
- `JWT_SECRET` - JWT signing secret
- `API_KEY` - API key for external services

### Email Configuration
- `SMTP_HOST` - SMTP server host
- `SMTP_PORT` - SMTP server port (default: 587)
- `EMAIL_FROM` - Default sender email

### Redis Configuration
- `REDIS_HOST` - Redis host (default: localhost)
- `REDIS_PORT` - Redis port (default: 6379)
- `REDIS_PASSWORD` - Redis password

### Application Settings
- `APP_NAME` - Application name
- `APP_VERSION` - Application version
- `LOG_LEVEL` - Logging level

## 🛠️ Code Examples

### Basic Usage
```go
// Load .env file
err := godotenv.Load()
if err != nil {
    log.Fatal("Error loading .env file")
}

// Get environment variable
dbHost := os.Getenv("DB_HOST")
```

### With Default Values
```go
func getEnv(key, defaultVal string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return defaultVal
}

dbHost := getEnv("DB_HOST", "localhost")
```

### Type Conversion
```go
func getEnvAsInt(name string, defaultVal int) int {
    valueStr := getEnv(name, "")
    if value, err := strconv.Atoi(valueStr); err == nil {
        return value
    }
    return defaultVal
}

port := getEnvAsInt("SERVER_PORT", 8080)
```

### Configuration Struct
```go
type Config struct {
    DBHost     string
    DBPort     int
    ServerPort int
    DebugMode  bool
}

config := Config{
    DBHost:     getEnv("DB_HOST", "localhost"),
    DBPort:     getEnvAsInt("DB_PORT", 5432),
    ServerPort: getEnvAsInt("SERVER_PORT", 8080),
    DebugMode:  getEnvAsBool("DEBUG_MODE", false),
}
```

## 🔒 Security Best Practices

1. **Never commit `.env` files** to version control
2. **Use `.env.example`** as a template
3. **Mask sensitive data** when displaying
4. **Use strong, unique values** for secrets
5. **Rotate secrets regularly**

## 📝 Sample Output

```
🚀 GoDotEnv Demo Application
============================
✅ Successfully loaded .env file

📋 Loaded Configuration:
------------------------
🗄️  Database:
   Host: localhost:5432
   User: postgres
   Password: yo**********re
   Database: testdb
   SSL Mode: disable

🌐 Server:
   Address: localhost:8080
   Debug Mode: true

🔐 Security:
   JWT Secret: your****key
   API Key: abc1****f456
```

## 🚀 Integration Examples

### With Gin Framework
```go
router := gin.Default()
if getEnvAsBool("DEBUG_MODE", false) {
    gin.SetMode(gin.DebugMode)
} else {
    gin.SetMode(gin.ReleaseMode)
}
```

### With GORM
```go
dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
    os.Getenv("DB_HOST"),
    os.Getenv("DB_USER"),
    os.Getenv("DB_PASSWORD"),
    os.Getenv("DB_NAME"),
    os.Getenv("DB_PORT"),
    os.Getenv("DB_SSL_MODE"))

db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
```

## 🌟 Why godotenv?

- **Simple and reliable**
- **Zero dependencies**
- **Works with standard `os.Getenv()`**
- **Most popular choice** (10k+ GitHub stars)
- **Great documentation and community**

---
*This demo showcases best practices for environment variable management in Go applications.*
