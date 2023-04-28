# Use the official Golang image as the base image
FROM golang:1.17-alpine

RUN apk add --no-cache gcc musl-dev

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum files into the container
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the application
RUN go build -o user-service .

# Expose the port for the user-service
EXPOSE 8080

# Run the user-service
CMD ["./user-service"]
