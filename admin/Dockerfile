# Stage 1: Build the Go application
FROM golang:1.23 AS build-stage

WORKDIR /app

# Copy go.mod and go.sum and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project and build the application
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main ./cmd/main.go
    
# Stage 2: Run tests
FROM build-stage AS run-test-stage
RUN go test -v ./...

# Stage 3: Prepare the release image
FROM debian:bullseye-slim AS build-release-stage
    
WORKDIR /app

# Install CA certificates (ensure TLS works)
COPY --from=build-stage /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy the built binary from the build stage
COPY --from=build-stage /app/main /app/main
COPY .env /app/.env

# Expose the application port
EXPOSE 50051

# Run the application directly (no shell)
ENTRYPOINT ["/app/main"]
