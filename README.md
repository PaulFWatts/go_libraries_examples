# Go Libraries Examples 🚀

A comprehensive collection of working examples demonstrating the most popular and essential Go libraries. This repository serves as a practical reference for developers looking to implement common functiona1. **[hello](./hello/)** - Basic Go syntax and concepts
2. **[greetings](./greetings/)** - Package creation and imports
3. **[TimeDemo](./TimeDemo/)** - Standard library usage
4. **[Testify](./Testify/)** - Testing and quality assurance
5. **[GoDotEnv](./GoDotEnv/)** - Environment configuration basics
6. **[Viper](./Viper/)** - Advanced configuration management
7. **[JWT](./JWT/)** - Authentication and security
8. **[MapStructure](./MapStructure/)** - Data handling
9. **[HTTPRouter](./HTTPRouter/)** - High-performance HTTP routing
10. **[Templ](./Templ/)** - Type-safe HTML template development
11. **[Echo](./Echo/)** - Minimalist web framework development
12. **[Gorm](./Gorm/)** - Database operations
13. **[Gin](./Gin/)** - Full-featured web development
14. **[GoQuery](./GoQuery/)** - Web scraping
15. **[Cobra](./Cobra/)** - CLI applicationsll-established Go packages.

## 📋 Table of Contents

- [Libraries Included](#libraries-included)
- [Getting Started](#getting-started)
- [Repository Structure](#repository-structure)
- [Usage](#usage)
- [Contributing](#contributing)

## 🛠️ Libraries Included

### Web Development
- **[Templ](./Templ/)** - Type-safe HTML templates
  - Compile-time HTML validation and type checking
  - Zero-overhead template compilation to Go functions
  - Component-based architecture with props and composition
  - IDE support with syntax highlighting and refactoring
  - Modern web development patterns with Bootstrap integration
  - Working Todo application with CRUD operations

- **[Echo](./Echo/)** - High-performance, minimalist web framework
  - Zero memory allocation HTTP router
  - Rich middleware ecosystem
  - JSON binding and validation
  - WebSocket and Server-Sent Events support
  - Template rendering and static file serving

- **[Gin](./Gin/)** - High-performance HTTP web framework
  - REST API endpoints
  - JSON responses
  - Middleware support

- **[HTTPRouter](./HTTPRouter/)** - High-performance HTTP request router
  - Zero-allocation routing
  - Path parameters and wildcards
  - RESTful API patterns
  - Custom error handling

- **[Cobra](./Cobra/)** - Powerful CLI application framework
  - Command-line interface creation
  - Command hierarchy and flags
  - Used by popular tools like Hugo, Docker, and Kubernetes

### Database & ORM
- **[GORM](./Gorm/)** - Feature-rich ORM library
  - Database migrations
  - CRUD operations
  - SQLite integration example

### Configuration & Environment
- **[Viper](./Viper/)** - Complete configuration solution
  - Multiple config file formats (YAML, JSON, TOML, HCL)
  - Environment variable integration with custom prefixes
  - Command-line flag binding via Cobra
  - Live configuration watching and reloading
  - Configuration precedence and validation
  - Real-world enterprise configuration patterns

- **[GoDotEnv](./GoDotEnv/)** - Environment variable management
  - `.env` file loading
  - Type conversion helpers
  - Security-conscious configuration handling
  - Cross-platform support

### Authentication & Security
- **[JWT](./JWT/)** - JSON Web Token implementation
  - Token creation and validation
  - HMAC and RSA signing methods
  - Custom claims and expiration handling
  - Refresh token patterns
  - Security best practices

### Testing & Quality Assurance
- **[Testify](./Testify/)** - Comprehensive testing framework
  - Rich assertion library with over 60+ assertion methods
  - Powerful mocking capabilities with expectations
  - Test suites with setup and teardown methods
  - Benchmarking and performance testing
  - Real-world testing scenarios and best practices

### Data Processing
- **[MapStructure](./MapStructure/)** - Map to struct conversion
  - JSON to struct mapping
  - Custom field tags
  - Type conversion and validation
  - Nested structure handling

### Web Scraping
- **[GoQuery](./GoQuery/)** - jQuery-like HTML parsing
  - Web scraping capabilities
  - CSS selector support
  - HTTP client best practices

### Standard Library Examples
- **[TimeDemo](./TimeDemo/)** - Go's built-in time package
  - Time formatting and parsing
  - Duration operations
  - Timezone handling
  - Timers and tickers

### Basic Go Concepts
- **[hello](./hello/)** - Simple "Hello World" example
- **[greetings](./greetings/)** - Custom package creation and usage

## 🌟 Featured Comprehensive Demonstrations

The following examples showcase production-ready implementations with extensive features:

### 🎯 **[Templ](./Templ/)** - Type-Safe HTML Templates
- **Live Demo**: Full web server with Todo application
- **Features**: Component architecture, Bootstrap UI, CRUD operations
- **Innovation**: Compile-time HTML validation, zero-overhead templates
- **Best For**: Modern web applications with type safety

### ⚙️ **[Viper](./Viper/)** - Configuration Management
- **Live Demo**: Multi-format config CLI with hot-reloading
- **Features**: YAML/JSON/TOML support, environment overrides, validation
- **Innovation**: Configuration precedence, live watching, enterprise patterns
- **Best For**: Complex applications requiring flexible configuration

### 🚀 **[Echo](./Echo/)** - Web Framework
- **Live Demo**: REST API with interactive documentation
- **Features**: Middleware chain, route grouping, JSON handling
- **Innovation**: Zero-allocation router, WebSocket support
- **Best For**: High-performance web services

### 🧪 **[Testify](./Testify/)** - Testing Framework
- **Live Demo**: Comprehensive test suite with mocks
- **Features**: 60+ assertions, mock expectations, benchmarks
- **Innovation**: Real-world testing patterns, coverage analysis
- **Best For**: Professional testing and quality assurance

## 🚀 Getting Started

### Prerequisites
- Go 1.19 or higher
- Git

### Quick Setup
1. **Clone the repository:**
   ```bash
   git clone <repository-url>
   cd go_libraries_examples
   ```

2. **Initialize the workspace:**
   ```bash
   go work sync
   ```

3. **Run any example:**
   ```bash
   cd <library-name>
   go mod tidy
   go run main.go
   ```

## 📁 Repository Structure

```
go_libraries_examples/
├── README.md                    # This file
├── GO_COMMANDS_CHEAT_SHEET.md  # Go tooling reference
├── go.work                     # Go workspace configuration
├── go.work.sum                 # Workspace checksums
├── test.db                     # SQLite database for examples
│
├── Cobra/                      # CLI framework example
├── Echo/                       # Minimalist web framework example
├── Gin/                        # Web framework example
├── GoDotEnv/                   # Environment variable management
├── GoQuery/                    # Web scraping example
├── Gorm/                       # ORM database example
├── HTTPRouter/                 # High-performance HTTP router
├── JWT/                        # JSON Web Token example
├── MapStructure/               # Data mapping example
├── Templ/                      # Type-safe HTML templates
├── Testify/                    # Testing framework example
├── TimeDemo/                   # Time package demonstration
├── Viper/                      # Configuration management
├── greetings/                  # Custom package example
└── hello/                      # Basic Go example
```

Each directory contains:
- `main.go` - Working example code
- `go.mod` - Module dependencies
- `README.md` - Detailed documentation (where applicable)

## 🎯 Usage

### Running Individual Examples

Each library example is self-contained and can be run independently:

```bash
# Web server example
cd Gin && go run main.go

# Minimalist web framework
cd Echo && go run main.go

# Type-safe HTML templates with Todo app
cd Templ && go run main.go

# High-performance HTTP router
cd HTTPRouter && go run main.go

# CLI tool example
cd Cobra && go run main.go

# Database operations
cd Gorm && go run main.go

# JSON Web Token authentication
cd JWT && go run main.go

# Environment variables
cd GoDotEnv && go run main.go

# Configuration management
cd Viper && go run main.go --help

# Web scraping
cd GoQuery && go run main.go

# Data mapping
cd MapStructure && go run main.go

# Testing framework (run tests)
cd Testify && go test -v

# Time operations
cd TimeDemo && go run main.go
```

### Building Executables

```bash
cd <library-directory>
go build -o example-name main.go
./example-name  # Linux/macOS
.\example-name.exe  # Windows
```

### Understanding Dependencies

View the dependency graph:
```bash
go mod graph
```

Update all dependencies:
```bash
go work sync
```

## 📚 Learning Path

**Recommended order for beginners:**

1. **[hello](./hello/)** - Basic Go syntax
2. **[greetings](./greetings/)** - Package creation
3. **[TimeDemo](./TimeDemo/)** - Standard library usage
4. **[Testify](./Testify/)** - Testing and quality assurance
5. **[GoDotEnv](./GoDotEnv/)** - Configuration management
6. **[JWT](./JWT/)** - Authentication and security
7. **[MapStructure](./MapStructure/)** - Data handling
8. **[HTTPRouter](./HTTPRouter/)** - High-performance HTTP routing
9. **[Echo](./Echo/)** - Minimalist web framework development
10. **[Gorm](./Gorm/)** - Database operations
11. **[Gin](./Gin/)** - Full-featured web development
12. **[GoQuery](./GoQuery/)** - Web scraping
13. **[Cobra](./Cobra/)** - CLI applications

## 🔧 Tools & Commands

This repository includes a comprehensive [Go Commands Cheat Sheet](./GO_COMMANDS_CHEAT_SHEET.md) covering:
- Module management
- Workspace operations
- Build and run commands
- Testing and benchmarking
- Code formatting and analysis

## 🤝 Contributing

Contributions are welcome! Please feel free to:
- Add new library examples
- Improve existing documentation
- Fix bugs or enhance code quality
- Suggest popular libraries to include

### Adding a New Library Example

1. Create a new directory with the library name
2. Include `main.go` with working examples
3. Add `go.mod` with proper dependencies
4. Create `README.md` with usage instructions
5. Update this main README.md

## 📄 License

This project is intended for educational purposes. Please refer to individual library licenses for their respective terms and conditions.

## ⭐ Popular Go Libraries Not Yet Included

Future additions may include:
- **Fiber** - Express-inspired web framework
- **Logrus** - Structured logging
- **Redis-Go** - Redis client
- **Zap** - High-performance logging
- **Gorilla/Mux** - HTTP router and URL matcher

---

**Happy coding with Go! 🐹**

> This repository is maintained as a learning resource. Each example is designed to be simple, well-documented, and immediately runnable.
