
# Simple Makefile for a Go project

# Build the application
all: build

build:
	@echo "Building..."
	@go build -o main cmd/api/main.go

# Run the application
run:
	@go run cmd/api/main.go -env-mode=development -config-path=environments/config.yaml

# Test the application
test:
	@echo "Testing..."
	@go test ./...

# Swag application
swag:
	@swag init --parseDependency --parseInternal --parseDepth 1 -g cmd/api/main.go

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f main

.PHONY: all build run test clean
		
