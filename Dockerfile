# Use the official Golang image to build the binary
FROM golang:1.21 AS builder

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum first to cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go binary
RUN go build -o main .

# Use a lightweight image for the final container
FROM alpine:latest

# Set the working directory
WORKDIR /root/

# Copy the built binary from the builder stage
COPY --from=builder /app/main .


# Run the binary
CMD ["./main"]
