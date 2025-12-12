# Build stage
FROM golang:1.25.5 AS builder

# Install Node.js and npm
RUN curl -fsSL https://deb.nodesource.com/setup_23.x | bash - && \
  apt-get update && apt-get install -y nodejs

# Install pandoc for Markdown to HTML conversion
RUN apt-get install -y pandoc

# Install minify tool
RUN go install github.com/tdewolff/minify/v2/cmd/minify@latest

# Set working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod ./

# Copy the source code
COPY . .

# Install Node.js dependencies
RUN npm install

# Build the application with CGO_ENABLED=0 for static linking
ENV CGO_ENABLED=0
RUN make prod-build

# Runtime stage - using Alpine Linux
FROM alpine:latest

# Install CA certificates for HTTPS requests
RUN apk add --no-cache ca-certificates

# Set working directory
WORKDIR /app

# Copy the built application and all necessary directories
COPY --from=builder /app/app/personalwebsite /app/personalwebsite
COPY --from=builder /app/app /app/app
COPY --from=builder /app/config /app/config
COPY --from=builder /app/frontend /app/frontend

# Set environment variable for production
ENV APP_ENV=production

# Expose the port the app runs on
EXPOSE 8080

# Command to run the executable
CMD ["./personalwebsite"]
