BIN_DIR := bin
BINARY := server
PROTO_DIR := pkg/proto
PROTO_FILE := $(PROTO_DIR)/cardvalidator.proto
PROTO_OUT := $(PROTO_DIR)/cardvalidator.pb.go
WEB_DIR := web

.PHONY: all build run test test-coverage test-e2e proto lint format docker-build docker-run clean

all: build

build: proto
	go build -o $(BIN_DIR)/$(BINARY) ./cmd/server

run: build
	./$(BIN_DIR)/$(BINARY)

test:
	go test -v ./tests/unit/...

test-coverage:
	go test -coverprofile=coverage.out ./tests/unit/...
	go tool cover -html=coverage.out -o coverage.html

test-e2e:
	go test -v ./tests/e2e/...

proto: $(PROTO_OUT)

$(PROTO_OUT): $(PROTO_FILE)
	./scripts/generate_proto.sh

lint:
	golangci-lint run ./...

# Format code
.PHONY: fmt
fmt:
	@echo "🎨 Formatting code..."
	@go fmt ./...

docker-build:
	docker build -t credit-card-validator -f deployments/Dockerfile .

docker-run:
	docker run -p 8080:8080 -p 9090:9090 credit-card-validator

clean:
	rm -rf $(BIN_DIR)
	rm -f coverage.out coverage.html