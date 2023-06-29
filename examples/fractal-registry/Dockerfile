# First stage: build the application
FROM golang:alpine AS stage

# Set working directory
WORKDIR /app

# Copy the source from the current directory to the working directory inside the container
COPY . .

# Download all dependencies
RUN go mod download

# Build the application
RUN go build -o main

# Second stage: create the Docker image
FROM alpine:latest

# Set working directory
WORKDIR /root/

# Copy the binary from the first stage
COPY --from=stage /app/main .

# Expose port
EXPOSE 50051

# Command to run the application
CMD ["./main"]