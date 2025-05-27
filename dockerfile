# Use the official Go image
FROM golang:1.22-alpine

# Install git for module support
RUN apk add --no-cache git

# Set working directory inside container
WORKDIR /app

# Copy go.mod and go.sum
COPY ./app/go.mod ./app/go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of your code
COPY app .

# Build the application
RUN go build -o main .

# Expose the port your app listens on
EXPOSE 3000

# Run the compiled binary
CMD ["./main"]
