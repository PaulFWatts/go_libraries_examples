#!/bin/bash
# Build script for GoDotEnv demo

echo "Building GoDotEnv Demo..."

# Build for current platform
go build -o godotenv-demo main.go

# Build for Windows (if not on Windows)
if [[ "$OSTYPE" != "msys" && "$OSTYPE" != "win32" ]]; then
    echo "Building for Windows..."
    GOOS=windows GOARCH=amd64 go build -o godotenv-demo.exe main.go
fi

# Build for Linux
echo "Building for Linux..."
GOOS=linux GOARCH=amd64 go build -o godotenv-demo-linux main.go

# Build for macOS
echo "Building for macOS..."
GOOS=darwin GOARCH=amd64 go build -o godotenv-demo-macos main.go

echo "Build complete!"
echo "Run with: ./godotenv-demo"
