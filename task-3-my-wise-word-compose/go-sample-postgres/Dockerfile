# Start with the official Go image
FROM golang:1.23.2

# Set the working directory inside the container
WORKDIR /myapp


# Check Go version
RUN go version

# Copy go.mod and go.sum files to the working directory and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code files
COPY *.go ./

# Copy AUTHORS.md to the image before the build
COPY AUTHORS.md ./
# COPY .env ./

# Build the Go application, outputting a binary named "my-go-app"
RUN CGO_ENABLED=0 GOOS=linux go build -o my-go-app

# Copy LINKS.md to the image after the build
COPY LINKS.md ./

# Expose the application port
EXPOSE 80

# Command to run the built binary
CMD ["./my-go-app"]
