# Use the official Golang image as a build stage
FROM golang:1.22.1 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Install golang-migrate
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source files from the current directory to the Working Directory inside the container
COPY . .

# Change directory to the mynance-service-api directory
WORKDIR /app/cmd/mynance-service-api

# Build the Go mynance-service-api
RUN go build -o /app/main .

# Final stage: runtime image
FROM golang:1.22.1

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main /app/main

# Install migrate binary to a directory within PATH
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Copy the migrations directory
COPY --from=builder /app/internal/database/migrations /app/migrations

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["/app/main"]
