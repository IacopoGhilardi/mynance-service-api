# Use the official Golang image as a build stage
FROM golang:1.22.1 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Install a compatible version of air
RUN go install github.com/cosmtrek/air@v1.41.0

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source files from the current directory to the Working Directory inside the container
COPY . .

# Change directory to the app directory
WORKDIR /app/cmd/app

# Build the Go app
RUN go build -o /cmd/app/main .

# Start a new stage from scratch
FROM golang:1.22.1

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Pre-built binary files from the previous stage
COPY --from=builder /cmd/app/main .
COPY --from=builder /go/bin/air /bin/air

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["air"]
