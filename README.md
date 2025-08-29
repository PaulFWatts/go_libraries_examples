# Go Libraries Examples 🚀

A comprehensive collection of working examples demonstrating the most popular and essential Go libraries. This repository serves as a practical reference for developers looking to implement common functionality using well-established Go packages.

## 📋 Table of Contents

- [Libraries Included](#libraries-included)
- [Getting Started](#getting-started)
- [Repository Structure](#repository-structure)
- [Usage](#usage)
- [Contributing](#contributing)

## 🛠️ Libraries Included

### Web Development
- **[Gin](./Gin/)** - High-performance HTTP web framework
  - REST API endpoints
  - JSON responses
  - Middleware support

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
├── Gin/                        # Web framework example
├── GoDotEnv/                   # Environment variable management
├── GoQuery/                    # Web scraping example
├── Gorm/                       # ORM database example
├── JWT/                        # JSON Web Token example
├── MapStructure/               # Data mapping example
├── TimeDemo/                   # Time package demonstration
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

# CLI tool example
cd Cobra && go run main.go

# Database operations
cd Gorm && go run main.go

# JSON Web Token authentication
cd JWT && go run main.go

# Environment variables
cd GoDotEnv && go run main.go

# Web scraping
cd GoQuery && go run main.go

# Data mapping
cd MapStructure && go run main.go

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
4. **[GoDotEnv](./GoDotEnv/)** - Configuration management
5. **[JWT](./JWT/)** - Authentication and security
6. **[MapStructure](./MapStructure/)** - Data handling
7. **[Gorm](./Gorm/)** - Database operations
8. **[Gin](./Gin/)** - Web development
9. **[GoQuery](./GoQuery/)** - Web scraping
10. **[Cobra](./Cobra/)** - CLI applications

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
- **Echo** - Minimalist web framework
- **Testify** - Testing toolkit
- **Viper** - Configuration management
- **Logrus** - Structured logging
- **Redis-Go** - Redis client
- **Zap** - High-performance logging
- **Gorilla/Mux** - HTTP router and URL matcher

---

**Happy coding with Go! 🐹**

> This repository is maintained as a learning resource. Each example is designed to be simple, well-documented, and immediately runnable.
