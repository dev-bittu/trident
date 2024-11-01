# Phony targets that do not correspond to actual files
.PHONY: all build build-cross runc runs rc rs clean help

# Default target - builds and cross-compiles the project
all: build build-cross

# Build for the current OS and architecture
build:
	@echo "Building client and server for the current OS..."
	go build -o build/client.exe cmd/client/main.go
	go build -o build/server.exe cmd/server/main.go

# Cross-compile for multiple OS and architectures
build-cross:
	@echo "Cross-compiling for Windows, Linux, and macOS on multiple architectures..."

	# Windows cross-compilation for amd64, arm64, and 386
	GOOS=windows GOARCH=amd64 go build -o build/client-windows-amd64.exe cmd/client/main.go
	GOOS=windows GOARCH=arm64 go build -o build/client-windows-arm64.exe cmd/client/main.go
	GOOS=windows GOARCH=386 go build -o build/client-windows-386.exe cmd/client/main.go
	GOOS=windows GOARCH=amd64 go build -o build/server-windows-amd64.exe cmd/server/main.go
	GOOS=windows GOARCH=arm64 go build -o build/server-windows-arm64.exe cmd/server/main.go
	GOOS=windows GOARCH=386 go build -o build/server-windows-386.exe cmd/server/main.go

	# Linux cross-compilation for amd64, arm64, and 386
	GOOS=linux GOARCH=amd64 go build -o build/client-linux-amd64 cmd/client/main.go
	GOOS=linux GOARCH=arm64 go build -o build/client-linux-arm64 cmd/client/main.go
	GOOS=linux GOARCH=386 go build -o build/client-linux-386 cmd/client/main.go
	GOOS=linux GOARCH=amd64 go build -o build/server-linux-amd64 cmd/server/main.go
	GOOS=linux GOARCH=arm64 go build -o build/server-linux-arm64 cmd/server/main.go
	GOOS=linux GOARCH=386 go build -o build/server-linux-386 cmd/server/main.go

	# macOS (Darwin) cross-compilation for amd64 and arm64
	GOOS=darwin GOARCH=amd64 go build -o build/client-darwin-amd64 cmd/client/main.go
	GOOS=darwin GOARCH=arm64 go build -o build/client-darwin-arm64 cmd/client/main.go
	GOOS=darwin GOARCH=amd64 go build -o build/server-darwin-amd64 cmd/server/main.go
	GOOS=darwin GOARCH=arm64 go build -o build/server-darwin-arm64 cmd/server/main.go

# Run client after building for the current OS
runc:
	@echo "Building and running client..."
	@make build
	./build/client.exe

# Run server after building for the current OS
runs:
	@echo "Building and running server..."
	@make build
	./build/server.exe

# Run client directly from the source
rc:
	@echo "Running client directly from source..."
	go run cmd/client/main.go

# Run server directly from the source
rs:
	@echo "Running server directly from source..."
	go run cmd/server/main.go

# Clean up build artifacts
clean:
	@echo "Cleaning up build artifacts..."
	rm -rf build

# Display help information
help:
	@echo "Available commands:"
	@echo "  make build        - Build client and server for the current OS"
	@echo "  make build-cross  - Cross-compile for Windows, Linux, and macOS (amd64, arm64, 386)"
	@echo "  make runc         - Build and run the client"
	@echo "  make runs         - Build and run the server"
	@echo "  make rc           - Run the client from source"
	@echo "  make rs           - Run the server from source"
	@echo "  make clean        - Remove build artifacts"
	@echo "  make help         - Display this help information"
