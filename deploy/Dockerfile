# Stage 1: Build the Go binary
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy dependency files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the application, disabling CGO for a static binary
RUN CGO_ENABLED=0 GOOS=linux go build -o /main ./cmd/patterngarden

# Stage 2: Create the final, minimal image
FROM alpine:latest

# Copy the compiled binary from the builder stage
COPY --from=builder /main /main

# Expose the port the app will run on (optional, but good practice)
EXPOSE 8080

# Command to run the executable
CMD ["/main"]
