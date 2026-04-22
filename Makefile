.PHONY: build run clean fmt test help

BINARY_NAME=concurrency-patterns
MAIN_PATH=main.go

help:
	@echo "Available targets:"
	@echo "  make build   - Compile the Go program"
	@echo "  make run     - Run the program"
	@echo "  make clean   - Remove build artifacts"
	@echo "  make fmt     - Format Go code"
	@echo "  make test    - Run tests"

build:
	@echo "Building..."
	go build -o $(BINARY_NAME).exe $(MAIN_PATH)

run:
	@echo "Running..."
	go run $(MAIN_PATH)

clean:
	@echo "Cleaning..."
	rm -f $(BINARY_NAME).exe

fmt:
	@echo "Formatting code..."
	go fmt ./...

test:
	@echo "Running tests..."
	go test -v ./...
