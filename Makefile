# Go parameters
APP_NAME=myapp
BUILD_DIR=bin
SRC_DIR=.
PKG=$(SRC_DIR)/main.go
MARKDOWN_DIR=static/blog-posts/markdown
HTML_DIR=static/blog-posts/html

# Default target: Build the application
all: build

# Build the Go application
build:
	@echo "Building $(APP_NAME)..."
	mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(APP_NAME) $(PKG)

# Run the application
run: build
	@echo "Running $(APP_NAME)..."
	./$(BUILD_DIR)/$(APP_NAME)

prod:
	export ENV=prod
	
# Clean build artifacts
clean:
	@echo "Cleaning build directory..."
	rm -rf $(BUILD_DIR)

# Run tests
test:
	@echo "Running tests..."
	go test ./...

# Format the code
fmt:
	@echo "Formatting code..."
	go fmt ./...

# Lint the code (requires `golangci-lint` installed)
lint:
	@echo "Linting code..."
	golangci-lint run

# Build and run
dev: fmt lint test build run

gen:
	@echo "Checking for Markdown files..."
	@if [ -z "$$(ls -A $(MARKDOWN_DIR)/*.md 2>/dev/null)" ]; then \
		echo "No Markdown files found. Skipping conversion."; \
	else \
		echo "Generating HTML from Markdown..."; \
		mkdir -p $(HTML_DIR); \
		for file in $(MARKDOWN_DIR)/*.md; do \
			filename=$$(basename $$file .md); \
			pandoc $$file -o $(HTML_DIR)/$$filename.html; \
		done; \
		echo "Conversion complete."; \
	fi

dev-server:
	export ENV=development
	${SRC_DIR}/scripts/run-server.sh
	
gen-styles:
	@echo "Generating minified css styles"
	minify -b ${SRC_DIR}/static/css/*.css -o ${SRC_DIR}/static/css/styles.css

# Help message
help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  all        Build the application (default)"
	@echo "  build      Compile the application"
	@echo "  run        Run the application"
	@echo "  clean      Remove build artifacts"
	@echo "  test       Run unit tests"
	@echo "  fmt        Format the code"
	@echo "  lint       Run linting"
	@echo "  dev        Run all steps: format, lint, test, build, and run"
	@echo "  help       Show this help message"
	@echo "  gen        Generate HTML from Markdown files"
	@echo "  dev-server Run the development server"
