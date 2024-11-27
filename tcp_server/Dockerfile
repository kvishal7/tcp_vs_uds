# Use the official Golang image as the base
FROM golang:1.19-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Copy the rest of the application code
COPY . .

# Download the dependencies
RUN go mod tidy

# Build the Go application
RUN go build -o main .

# Expose the port your server listens on
EXPOSE 8080

# Start the application
CMD ["./main"]