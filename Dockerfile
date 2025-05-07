# Stage 1: Build the Go binary
FROM golang:1.22-alpine3.19 as builder

# Set the current working directory inside the container
WORKDIR /app

# Copy the go mod and sum files
COPY go.mod .

# Download the dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod tidy

# Copy the rest of the application
COPY . .

# Build the Go app
RUN go build -o main .

# Stage 2: Create a smaller image with the Go binary
FROM alpine:latest  

# Install required dependencies for Go app to run (e.g., for static files)
RUN apk --no-cache add ca-certificates

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Copy static files (HTML, CSS, JS)
COPY --from=builder /app/static /root/static

# Expose port 8080 to the outside world
EXPOSE 8080

# Run the Go app
CMD ["./main"]
