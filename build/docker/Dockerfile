# Build stage
FROM golang:1.24-alpine AS builder

# Set working directory
WORKDIR /app

# Install dependencies for building
RUN apk add --no-cache gcc musl-dev

# Copy go mod and sum files
COPY go.mod ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-w -s" -o personalwebsite ./cmd/server/main.go

# Runtime stage
FROM alpine:3.18

# Set working directory
WORKDIR /app

# Set production environment variable
ENV APP_ENV=production

# Install runtime dependencies
RUN apk --no-cache add ca-certificates tzdata

# Copy the binary from builder stage
COPY --from=builder /app/personalwebsite /app/personalwebsite

# Copy necessary files
COPY --from=builder /app/config/production.json /app/config/production.json
COPY --from=builder /app/frontend /app/frontend

# Create app directory for outputs
RUN mkdir -p /app/app/html/blog
RUN mkdir -p /app/app/html/templates
RUN mkdir -p /app/app/html/toc

# Expose port
EXPOSE 8080

# Run the application
CMD ["/app/personalwebsite"]