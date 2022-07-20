# Start from golang:1.18-alpine as base image.
FROM golang:1.18-alpine as builder

ENV GO111MODULE=on

# Install git.
RUN apk update && apk add --no-cache git

# Set the current working directory inside the container.
WORKDIR /app

# Copy go mod and sum files.
COPY go.mod go.sum ./

# Download all dependencies.
# Dependencies will be cached if the go.mod and the go.sum files are not changed in the meantime.
RUN go mod download

# Copy everything to the current working directory.
COPY . .

# Build the Go application.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o main .


# Start a dev stage as a new stage from the base image.
FROM alpine as dev
RUN apk add --no-cache ca-certificates bash

# Set the working directory for the application.
WORKDIR /app

# Copy the pre-built binary files from the previous stage.
COPY . .
COPY --from=builder /app/main .

# Expose port 8080 to the outside world.
EXPOSE 8080

# Command to run the executable.
CMD ["./main"]