# Build stage
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY . .
RUN apk add --no-cache make protoc
RUN make proto
RUN make build

# Final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /app/bin/server .
COPY --from=builder /app/web ./web

EXPOSE 8080 9090
CMD ["./server"]