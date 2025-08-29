# Go Tooling Commands Cheat Sheet

## Module Management

| Command | Description |
|---------|-------------|
| `go mod init <module-name>` | Initialize a new Go module |
| `go mod tidy` | Add missing and remove unused dependencies |
| `go mod download` | Download dependencies to local cache |
| `go mod verify` | Verify dependencies have expected content |
| `go mod edit -go=<version>` | Update Go version in go.mod |
| `go mod edit -require=<module>@<version>` | Add or update a dependency |
| `go mod edit -droprequire=<module>` | Remove a dependency |
| `go mod graph` | Print module requirement graph |
| `go mod why <module>` | Explain why a module is needed |
| `go mod vendor` | Make vendored copy of dependencies |

## Workspace Management

| Command | Description |
|---------|-------------|
| `go work init <modules>` | Initialize workspace with modules |
| `go work use <modules>` | Add modules to workspace |
| `go work edit -go=<version>` | Update Go version in go.work |
| `go work sync` | Sync workspace dependencies |

## Build & Run

| Command | Description |
|---------|-------------|
| `go run .` | Run the current package |
| `go run <file.go>` | Run a specific Go file |
| `go build` | Build executable for current package |
| `go build -o <name>` | Build with custom executable name |
| `go build <package>` | Build a specific package |
| `go install` | Build and install package to $GOPATH/bin |
| `go clean` | Remove build artifacts |

## Testing

| Command | Description |
|---------|-------------|
| `go test` | Run tests in current package |
| `go test ./...` | Run all tests in module |
| `go test -v` | Run tests with verbose output |
| `go test -cover` | Run tests with coverage |
| `go test -race` | Run tests with race detection |
| `go test -bench=.` | Run benchmarks |
| `go test -run <pattern>` | Run specific tests matching pattern |

## Dependencies

| Command | Description |
|---------|-------------|
| `go get <module>` | Add dependency to current module |
| `go get <module>@<version>` | Get specific version of module |
| `go get <module>@latest` | Get latest version of module |
| `go get -u` | Update all dependencies |
| `go get -u <module>` | Update specific dependency |
| `go list -m all` | List all dependencies |
| `go list -u -m all` | List dependencies with available updates |

## Code Quality

| Command | Description |
|---------|-------------|
| `go fmt` | Format current package |
| `go fmt ./...` | Format all packages in module |
| `gofmt -w .` | Format and write changes to files |
| `go vet` | Report likely mistakes in current package |
| `go vet ./...` | Vet all packages in module |
| `golint` | Run Go linter (requires installation) |
| `go doc <package>` | Show documentation for package |
| `go doc <package>.<symbol>` | Show documentation for symbol |

## Environment & Info

| Command | Description |
|---------|-------------|
| `go version` | Show Go version |
| `go env` | Print Go environment information |
| `go env <VAR>` | Print specific environment variable |
| `go list` | List packages in current module |
| `go list ./...` | List all packages in module |
| `go list -json` | List packages in JSON format |

## Cross Compilation

| Command | Description |
|---------|-------------|
| `GOOS=linux go build` | Build for Linux |
| `GOOS=windows go build` | Build for Windows |
| `GOOS=darwin go build` | Build for macOS |
| `GOARCH=amd64 go build` | Build for 64-bit architecture |
| `GOOS=linux GOARCH=arm go build` | Build for Linux ARM |

## Debug & Profile

| Command | Description |
|---------|-------------|
| `go build -race` | Build with race detection |
| `go build -ldflags="-s -w"` | Build stripped binary (smaller size) |
| `go tool pprof <binary> <profile>` | Analyze performance profile |
| `go tool trace <trace-file>` | Analyze execution trace |

## Generate & Tools

| Command | Description |
|---------|-------------|
| `go generate` | Run generate commands in current package |
| `go generate ./...` | Run generate commands in all packages |
| `go tool` | Run Go tool |
| `go tool compile` | Run Go compiler |
| `go tool link` | Run Go linker |

## Common Flags

| Flag | Description |
|------|-------------|
| `-v` | Verbose output |
| `-a` | Force rebuild of packages |
| `-n` | Print commands but do not run them |
| `-x` | Print commands as they are executed |
| `-work` | Print temporary work directory and don't delete it |
| `-tags <list>` | Build with specific build tags |

## Examples

### Quick Start
```bash
# Create new module
go mod init example.com/myproject

# Add dependency
go get github.com/gin-gonic/gin

# Run with live reload (requires air)
go run main.go

# Build optimized binary
go build -ldflags="-s -w" -o myapp
```

### Testing Workflow
```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run specific test
go test -run TestMyFunction

# Run benchmarks
go test -bench=. -benchmem
```

### Module Maintenance
```bash
# Clean up dependencies
go mod tidy

# Update all dependencies
go get -u ./...

# Check for vulnerabilities
go list -json -deps | nancy sleuth
```

---
*Generated on: $(date)*
