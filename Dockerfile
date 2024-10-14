# Step 1: Build the Go application
FROM golang:1.19 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o myapp -buildvcs=false

# Step 2: Build a small image with just the Go binary
FROM alpine:latest

# Set a working directory
WORKDIR /root/

# Copy the built Go binary from the builder image
COPY --from=builder /app/myapp .

# Expose the port on which the app will run (change this if needed)
EXPOSE 8080

# Command to run the binary
CMD ["./myapp"]
 