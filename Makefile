# Makefile

# Output binary name and location
BINARY_NAME=bin/server

# Docker image name
IMAGE_NAME=credit-card-validator

# Proto definitions
PROTO_DIR := pkg/proto
PROTO_FILE := $(PROTO_DIR)/cardvalidator.proto
PROTO_OUT := $(PROTO_DIR)/cardvalidator.pb.go

# Default target: clean -> fmt -> proto -> test -> build
.PHONY: all
all: clean fmt proto test build

# Build the binary
.PHONY: build
build:
	@echo "🔨 Building $(BINARY_NAME)..."
	@mkdir -p bin
	@go build -o $(BINARY_NAME) ./cmd/server

# Run the app
.PHONY: run
run: build
	@echo "🚀 Running $(BINARY_NAME)..."
	@./$(BINARY_NAME)

# Run unit tests
.PHONY: test
test:
	@echo "✅ Running unit tests..."
	@go test -v ./test/...

# Generate Go code from proto file using script
.PHONY: proto
proto: $(PROTO_OUT)

$(PROTO_OUT): $(PROTO_FILE)
	@echo "📦 Generating protobuf files using generate_proto.sh..."
	@./scripts/generate_proto.sh

# Format code
.PHONY: fmt
fmt:
	@echo "🎨 Formatting code..."
	@go fmt ./...

# Lint code
.PHONY: lint
lint:
	@echo "🔍 Linting code..."
	@golangci-lint run ./...

# Build Docker image
.PHONY: docker-build
docker-build:
	@echo "🐳 Building Docker image..."
	@docker build -t $(IMAGE_NAME) -f deployments/Dockerfile .

# Run Docker container
.PHONY: docker-run
docker-run:
	@echo "🚀 Running Docker container on ports 8080 (HTTP) and 9090 (gRPC)..."
	@docker run -p 8080:8080 -p 9090:9090 $(IMAGE_NAME)

# Clean up binaries and generated files
.PHONY: clean
clean:
	@echo "🧹 Cleaning up..."
	@rm -rf bin
