# Use the official Go image with the desired version as the base image
FROM golang:1.20.1-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download the Go module dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Copy the .env file to the working directory
COPY .env .

# Build the Go application
RUN go build -o main ./cmd

# Set the entry point for the container
CMD ["./main"]
