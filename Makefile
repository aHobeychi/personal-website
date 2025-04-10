# Go parameters
APP_NAME=personalwebsite
BUILD_DIR=app
SRC_DIR=cmd/server/
PKG=$(SRC_DIR)/main.go
MARKDOWN_DIR=frontend/content/blog/markdown
HTML_DIR=frontend/content/blog/html
SCRIPTS_DIR=build/scripts

# Default target: Build the application
all: build

# Run the application
run: build
	@echo "Running $(APP_NAME)..."
	./$(BUILD_DIR)/$(APP_NAME)

prod:
	@echo "Minifying html & css files"
	make minify
	@echo "Building application"
	go build -o $(BUILD_DIR)/$(APP_NAME) $(PKG)
	export APP_ENV=production && $(BUILD_DIR)/$(APP_NAME)
	
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
			pandoc $$file -o $(HTML_DIR)/content/$$filename.html; \
		done; \
		echo "Conversion complete."; \
	fi

dev-server:
	export ENV=development
	build/scripts/run-server.sh
	
minify:
	./${SCRIPTS_DIR}/minify.sh
	@echo "Generating minified css styles"
	minify -b frontend/assets/css/*.css -o frontend/assets/css/styles.css


build:
	@echo "Creating build directory..."
	mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(APP_NAME) $(PKG)

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
