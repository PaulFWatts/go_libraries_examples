# AI Agent Instructions for Go Libraries Examples

## Project Architecture

This is a **multi-module Go workspace** showcasing comprehensive library demonstrations. Each subdirectory is an independent Go module with working examples, integrated via `go.work` for unified development.

### Key Structural Patterns

- **Workspace Management**: Use `go work use ./NewLibrary` to add modules, not manual `go.mod` editing
- **Module Independence**: Each library demo has isolated dependencies in its own `go.mod`
- **Comprehensive Demos**: Examples are production-ready with full web servers, CLIs, or test suites (not minimal snippets)
- **Template Generation**: Templ uses `.templ` files that compile to `_templ.go` - always run `templ generate` after editing

## Development Workflows

### Adding New Library Demonstrations

1. Create directory: `mkdir NewLibrary && cd NewLibrary`
2. Initialize module: `go mod init newlibrary-demo`
3. Add to workspace: `go work use ./NewLibrary` (from root)
4. Follow established patterns from flagship demos (Echo, Templ, Viper, Testify)

### Critical Build Commands

```bash
# Workspace-wide dependency management
go work sync

# For Templ specifically (required before building)
cd Templ && templ generate

# Testing pattern for comprehensive suites
cd Testify && go test -v ./... -cover

# Running web servers with proper context
cd Templ && go run main.go  # Starts on :8080
cd Echo && go run main.go   # Interactive API on :8080
```

## Library-Specific Implementation Patterns

### Web Servers (Echo, Templ, Gin)
- Always include middleware chain for logging/CORS
- Use structured data types shared between templates and handlers
- Implement health check endpoints at `/health`
- Example: `type Todo = templates.Todo` for type aliasing in Templ

### Configuration Management (Viper)
- Multi-format support (YAML/JSON/TOML) with sample generation
- Environment variable overrides with `VIPERAPP_` prefix pattern
- Live configuration watching with `viper.WatchConfig()`
- Complex nested structs with `mapstructure` tags

### Template Systems (Templ)
- Type definitions must exist in template files for cross-package usage
- Generated `_templ.go` files are build artifacts, never edit directly
- Component patterns: `templ ComponentName(props Type) { ... }`
- Always run `templ generate` after editing `.templ` files

### Testing Frameworks (Testify)
- Comprehensive test categories: unit, integration, benchmarks, mocks
- Mock generation pattern: interfaces in separate files for easier mocking
- Test suites with setup/teardown using `suite.Suite`
- Coverage analysis with detailed assertions (60+ assertion types)

## Cross-Component Integration

### Type Sharing Strategy
Libraries redefine common types locally rather than shared packages:
```go
// In templates/simple_todo.templ
type Todo struct { ... }

// In main.go  
type Todo = templates.Todo  // Type alias
```

### Data Flow Patterns
- In-memory stores for demos (arrays with global state)
- RESTful API patterns with proper HTTP status codes
- Middleware chains for logging: `fmt.Printf("📝 %s %s - %v\n", method, path, duration)`

## Key Files for Understanding Patterns

- `go.work` - Workspace configuration and module registry
- `GO_COMMANDS_CHEAT_SHEET.md` - Project-specific tooling guide
- `Templ/main.go` - Complex web server with type aliasing patterns
- `Viper/main.go` - CLI application architecture with Cobra integration
- `Echo/main.go` - Middleware chain and REST API patterns
- `Testify/*_test.go` - Comprehensive testing strategies

## Critical Dependencies

- Templ requires CLI tool: `go install github.com/a-h/templ/cmd/templ@latest`
- Workspace requires Go 1.21+ for proper multi-module support
- Most demos use Bootstrap CSS via CDN (no local assets to manage)

## Documentation Standards

Each demo includes:
- Comprehensive `README.md` with feature breakdown and usage examples
- Code examples showing both basic and advanced patterns
- Integration with main repository README in "Featured Comprehensive Demonstrations"

When adding new libraries, prioritize comprehensive, production-ready examples over simple Hello World implementations.
