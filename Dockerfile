# Use the official Golang image to build the binary
FROM golang:latest AS builder

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum first to cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go binary
RUN go build -o main .

CMD ["./main"]