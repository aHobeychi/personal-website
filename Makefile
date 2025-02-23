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

# Build the application
build:
	@echo "Creating build directory..."
	mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(APP_NAME) $(PKG)

# Run the application
run: build
	@echo "Running $(APP_NAME)..."
	./$(BUILD_DIR)/$(APP_NAME)

generate-styles:
	npm run build

# Production build and run
prod-build:
	@echo "Generating styles"
	make generate-styles
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
	@echo "Minifying HTML & CSS files..."
	make minify
	@echo "Building application..."
	go build -o $(BUILD_DIR)/$(APP_NAME) $(PKG)
	export APP_ENV=production

# Production build and run
prod:
	make prod-build
	export APP_ENV=production && ./$(BUILD_DIR)/$(APP_NAME)

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

# Generate HTML from Markdown files
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

# Minify CSS and HTML files
minify:
	@echo "Running minification scripts..."
	./${SCRIPTS_DIR}/minify.sh
	@echo "Generating minified CSS styles..."
	minify -b frontend/assets/css/*.css -o frontend/assets/css/styles.css
	minify -b frontend/assets/css/*.css -o app/assets/css/styles.css

# Development workflow
dev: fmt lint test build run

# Run the development server
dev-server:
	make minify
	export ENV=development && build/scripts/run-server.sh

# Help message
help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  all        Build the application (default)"
	@echo "  build      Compile the application"
	@echo "  run        Run the application"
	@echo "  prod       Build and run the application in production mode"
	@echo "  clean      Remove build artifacts"
	@echo "  test       Run unit tests"
	@echo "  fmt        Format the code"
	@echo "  lint       Run linting"
	@echo "  gen        Generate HTML from Markdown files"
	@echo "  minify     Minify CSS and HTML files"
	@echo "  dev        Run all steps: format, lint, test, build, and run"
	@echo "  dev-server Run the development server"
	@echo "  help       Show this help message"
