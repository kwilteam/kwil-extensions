# First stage: build the application
FROM golang:1.20 AS builder

# Set working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source from the current directory to the working directory inside the container
COPY . .

# Build the application
RUN go build -o main .

# Second stage: create the Docker image
FROM alpine:latest

# Set working directory
WORKDIR /root/

# Copy the binary from the first stage
COPY --from=builder /app/main .

# Expose port
EXPOSE 50051

# Command to run the application
CMD ["./main"]
