# Start with a Go base image
FROM golang:1.23-alpine as builder

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum files first for better caching
COPY go.mod ./
# COPY go.sum ./ (uncomment if you have a go.sum file)

# Download dependencies
RUN go mod download

# Copy the rest of the application
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/personal-website .

# Use a minimal alpine image for the final image
FROM alpine:3.18

# Install ca-certificates for HTTPS requests
RUN apk --no-cache add ca-certificates

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/bin/personal-website /app/personal-website

# Copy templates and static files which are required at runtime
COPY --from=builder /app/templates /app/templates
COPY --from=builder /app/static /app/static

# Expose the port the app runs on
EXPOSE 8080

# Set environment variables
ENV SERVER_PORT=8080
ENV LOG_LEVEL=info

# Command to run when the container starts
CMD ["/app/personal-website"]