# Use the official Go image
FROM golang:1.24-alpine

WORKDIR /src/app

COPY app ./

RUN go mod tidy && go mod verify

# Expose the port your app listens on
EXPOSE 3000

# Run the compiled binary
ENTRYPOINT [ "go", "run", "." ]
