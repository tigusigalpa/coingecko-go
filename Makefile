.PHONY: help build test fmt vet lint clean examples install

help:
	@echo "Available targets:"
	@echo "  build      - Build the package"
	@echo "  test       - Run tests"
	@echo "  fmt        - Format code"
	@echo "  vet        - Run go vet"
	@echo "  lint       - Run linters"
	@echo "  clean      - Clean build artifacts"
	@echo "  examples   - Build all examples"
	@echo "  install    - Install dependencies"

build:
	go build -v ./...

test:
	go test -v -race -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

fmt:
	go fmt ./...

vet:
	go vet ./...

lint: fmt vet
	@echo "Linting complete"

clean:
	rm -f coverage.out coverage.html
	go clean

examples:
	@echo "Building examples..."
	cd examples/basic && go build
	cd examples/portfolio && go build
	cd examples/market_analysis && go build
	cd examples/defi_tracker && go build
	@echo "Examples built successfully"

install:
	go mod download
	go mod tidy
