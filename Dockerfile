# base image
FROM golang:1.24-alpine

# Goes to the app directory
WORKDIR /src/app

RUN apk add --no-cache git

#  Copy the application code into the container
COPY app ./

# Install dependencies
RUN go mod tidy && go mod verify

# Expose the port your app listens on
EXPOSE 3000

# Run the compiled binary
ENTRYPOINT [ "go", "run", "." ]
