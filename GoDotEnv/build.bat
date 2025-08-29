@echo off
REM Build script for GoDotEnv demo (Windows)

echo Building GoDotEnv Demo...

REM Build for Windows
go build -o godotenv-demo.exe main.go

REM Build for Linux
echo Building for Linux...
set GOOS=linux
set GOARCH=amd64
go build -o godotenv-demo-linux main.go

REM Build for macOS  
echo Building for macOS...
set GOOS=darwin
set GOARCH=amd64
go build -o godotenv-demo-macos main.go

REM Reset environment
set GOOS=
set GOARCH=

echo Build complete!
echo Run with: godotenv-demo.exe
