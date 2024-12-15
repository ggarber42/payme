# Stage 1: Build the Go application
FROM golang:1.22.1-alpine as builder

# Set environment variables
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Set the working directory
WORKDIR /app
COPY . .

# Copy Go modules manifests and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire application code

# Build the application binary
RUN go build -o /go/bin/server cmd/api/api.go

# Stage 2: Create a minimal runtime image
FROM alpine:latest

# Install certificates for HTTPS support
RUN apk --no-cache add ca-certificates

# Set the working directory
WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /go/bin/server /go/bin/server

# Expose the application's port (adjust as necessary)
EXPOSE 8080

# Command to run the application
ENTRYPOINT ["/go/bin/server"]
