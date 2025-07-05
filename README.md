# 💳 Credit Card Validator Microservice

A high-performance microservice for validating credit card numbers using the Luhn algorithm with payment network identification.

## 🚀 Features

- **REST API** with Echo framework
- **gRPC API** for high-performance communication
- **Web Interface** for testing
- **Luhn Algorithm** validation
- **Payment Network Detection** (Visa, Mastercard, American Express, Discover)
- **Prometheus Metrics** for monitoring
- **Structured Logging** with Logrus
- **Configuration Management** with Viper
- **Comprehensive Testing**
- **Docker Support** for containerization

## 🏗️ Architecture

```
┌─────────────────┐    ┌─────────────────┐
│   Web Client    │    │   gRPC Client   │
└─────────┬───────┘    └─────────┬───────┘
          │                      │
          │ HTTP/REST            │ gRPC
          │                      │
┌─────────▼──────────────────────▼───────┐
│           Echo HTTP Server             │
│    ┌─────────────┐  ┌─────────────┐   │
│    │ REST Handler│  │ gRPC Handler│   │
│    └─────────────┘  └─────────────┘   │
│              │                        │
│    ┌─────────▼─────────────┐          │
│    │  Validator Service    │          │
│    │  (Luhn Algorithm)     │          │
│    └───────────────────────┘          │
└────────────────────────────────────────┘
```

## 🛠️ Quick Start

### Prerequisites

- Go 1.21+
- Docker (optional)
- protoc compiler (for gRPC)

### Installation

```bash
# Clone the repository
git clone <repo-url>
cd credit-card-validator

# Install dependencies
go mod download

# Generate protobuf files
make proto

# Build the application
make build

# Run the server
make run
```

### Using Docker

```bash
# Build Docker image
make docker-build

# Run with Docker
make docker-run
```

## 📡 API Endpoints

### REST API

**Base URL:** `http://localhost:8080`

#### Validate Card Number

```bash
POST /api/v1/validate
Content-Type: application/json

{
  "card_number": "4532015112830366"
}
```

**Response:**
```json
{
  "valid": true,
  "card_type": "visa",
  "card_number": "4532015112830366"
}
```

#### Health Check

```bash
GET /health
```

#### Metrics

```bash
GET /metrics
```

### gRPC API

**Address:** `localhost:9090`

```protobuf
service CardValidator {
  rpc ValidateCard(ValidateCardRequest) returns (ValidateCardResponse);
}
```

### Web Interface

Visit `http://localhost:8080` to access the web interface for testing.

## 🧪 Testing

```bash
# Run unit tests
make test


```

## 📊 Monitoring

The service exposes Prometheus metrics at `/metrics`:

- `card_validation_requests_total` - Total number of validation requests
- `card_validation_duration_seconds` - Request duration histogram
- `card_validation_errors_total` - Total number of validation errors

## ⚙️ Configuration

Configuration can be set via environment variables or config file:

```env
PORT=8080
GRPC_PORT=9090
LOG_LEVEL=info
METRICS_ENABLED=true
```

## 🔧 Development

### Available Make Commands

```bash
make help          # Show available commands
make build         # Build the application
make run           # Run the application
make test          # Run unit tests
make proto         # Generate protobuf files
make lint          # Run linter
make format        # Format code
make docker-build  # Build Docker image
make docker-run    # Run Docker container
make clean         # Clean build artifacts
```

### Project Structure

```
credit-card-validator/
├── cmd/server/          # Application entrypoint
├── internal/
│   ├── api/            # API handlers (REST & gRPC)
│   ├── service/        # Business logic
│   ├── config/         # Configuration
│   └── middleware/     # HTTP middleware
├── pkg/proto/          # Protocol buffer definitions
├── web/                # Web interface
├── test/              # Test files
└── deployments/        # Deployment files
```

## 📋 Supported Card Types

- **Visa**: 4xxx-xxxx-xxxx-xxxx
- **Mastercard**: 5xxx-xxxx-xxxx-xxxx
- **American Express**: 34xx-xxxxxx-xxxxx, 37xx-xxxxxx-xxxxx
- **Discover**: 6xxx-xxxx-xxxx-xxxx

## 🐛 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests
5. Run `make test`
6. Submit a pull request

## 📝 License

MIT License