APP_NAME=personalwebsite
BUILD_DIR=app
SRC_DIR=cmd/server/
VERSION := $(shell git describe --tags --always --dirty 2>/dev/null || echo "unknown")
BUILD_TIME := $(shell date -u '+%Y-%m-%d_%H:%M:%S_UTC')

SRC_DIR=cmd/server/
PKG=$(SRC_DIR)/main.go
MARKDOWN_DIR=frontend/content/blog/markdown
HTML_BASE_DIR=frontend/content/blog/html
HTML_CONTENT_DIR=$(HTML_BASE_DIR)/content
SCRIPTS_DIR=build/scripts

FRONTEND_CSS_SRC_DIR=frontend/assets/css
FRONTEND_CSS_MIN_FILE=$(FRONTEND_CSS_SRC_DIR)/styles.css
APP_ASSETS_DIR=$(BUILD_DIR)/assets
APP_CSS_DIR=$(APP_ASSETS_DIR)/css
APP_CSS_MIN_FILE=$(APP_CSS_DIR)/styles.css

BINARY=$(BUILD_DIR)/$(APP_NAME)

GO=go
PANDOC=pandoc
MINIFY=minify
NPM=npm
GOLANGCI_LINT=golangci-lint

.PHONY: all build run generate-styles prod-build prod clean test fmt lint \
        generate-html minify dev dev-server help \
        check-deps check-pandoc check-minify check-golint \
        create-dirs watch version

create-dirs:
	@echo "Ensuring required directories exist..."
	@mkdir -p $(BUILD_DIR) $(HTML_CONTENT_DIR) $(APP_CSS_DIR)

check-deps: check-pandoc check-minify
	@echo "All checked dependencies are present."

check-pandoc:
	@command -v $(PANDOC) >/dev/null 2>&1 || { echo >&2 "Error: $(PANDOC) is required but not installed. Please install it (e.g., 'brew install pandoc')."; exit 1; }

check-minify:
	@command -v $(MINIFY) >/dev/null 2>&1 || { echo >&2 "Error: $(MINIFY) is required but not installed. Please install it (e.g., 'go install github.com/tdewolff/minify/cmd/minify@latest')."; exit 1; }

check-golint:
	@command -v $(GOLANGCI_LINT) >/dev/null 2>&1 || { echo >&2 "Error: $(GOLANGCI_LINT) is required but not installed. See https://golangci-lint.run/usage/install/"; exit 1; }

build: create-dirs
	@echo "Building Go application (Version: $(VERSION), Build Time: $(BUILD_TIME))..."
	$(GO) build $(LDFLAGS) -o $(BINARY) $(PKG)

run: build
	@echo "Running $(APP_NAME)..."
	./$(BINARY)

generate-styles:
	@echo "Generating CSS styles via npm..."
	$(NPM) run build

generate-html: check-pandoc create-dirs
	@echo "Checking for Markdown files in $(MARKDOWN_DIR)..."
	@if [ -z "$$(ls -A $(MARKDOWN_DIR)/*.md 2>/dev/null)" ]; then \
		echo "No Markdown files found. Skipping conversion."; \
	else \
		echo "Generating HTML from Markdown into $(HTML_CONTENT_DIR)..."; \
		for file in $(MARKDOWN_DIR)/*.md; do \
			filename=$$(basename "$$file" .md); \
			echo "  Converting $$file to $(HTML_CONTENT_DIR)/$${filename}.html..."; \
			$(PANDOC) "$$file" -o "$(HTML_CONTENT_DIR)/$${filename}.html" || echo "  Warning: Failed to convert $${filename}.md"; \
		done; \
		echo "HTML generation complete."; \
	fi

minify: check-minify create-dirs
	@echo "Running minification script for HTML/JS (./$(SCRIPTS_DIR)/minify.sh)..."
	@if [ -f "./$(SCRIPTS_DIR)/minify.sh" ]; then \
		./$(SCRIPTS_DIR)/minify.sh; \
	else \
		echo "Warning: Minify script ./$(SCRIPTS_DIR)/minify.sh not found. Skipping."; \
	fi
	@echo "Minifying CSS styles..."
	# Minify CSS for frontend assets (served directly or for reference)
	$(MINIFY) -b $(FRONTEND_CSS_SRC_DIR)/*.css -o $(FRONTEND_CSS_MIN_FILE)
	# Minify CSS for application bundle (copied into build/app/assets/css)
	$(MINIFY) -b $(FRONTEND_CSS_SRC_DIR)/*.css -o $(APP_CSS_MIN_FILE)
	@echo "CSS minification complete."

prod-build: check-deps generate-styles generate-html minify
	@echo "Building production Go application (Version: $(VERSION), Build Time: $(BUILD_TIME))..."
	$(GO) build $(LDFLAGS) -o $(BINARY) $(PKG)
	@echo "Production build complete. Artifacts are in $(BUILD_DIR)"

prod: prod-build
	@echo "Running $(APP_NAME) in production mode..."
	APP_ENV=production ./$(BINARY)

clean:
	@echo "Cleaning build directory $(BUILD_DIR)..."
	rm -rf $(BUILD_DIR)
	@echo "Cleaning generated HTML from $(HTML_CONTENT_DIR)..."
	rm -rf $(HTML_CONTENT_DIR) # Corrected path for rm
	@echo "Cleaning minified frontend CSS $(FRONTEND_CSS_MIN_FILE)..."
	rm -f $(FRONTEND_CSS_MIN_FILE)
	@echo "Cleaning complete."

fmt:
	@echo "Formatting Go code..."
	$(GO) fmt ./...

lint: check-golint
	@echo "Linting Go code..."
	$(GOLANGCI_LINT) run

dev-server: minify
	@echo "Starting development server..."
	ENV=development $(SCRIPTS_DIR)/run-server.sh

version:
	@echo "$(APP_NAME) Makefile"
	@echo "  Version (Git):    $(VERSION)"
	@echo "  Build Time (UTC): $(BUILD_TIME)"
	@echo "  Binary:           $(BINARY)"

help:
	@echo "$(APP_NAME) Makefile - Version: $(VERSION)"
	@echo "Build Time (UTC): $(BUILD_TIME)"
	@echo ""
	@echo "Usage: make [target]"
	@echo ""
	@echo "Main Targets:"
	@echo "  build            Compile the Go application with version info"
	@echo "  run              Build and run the Go application"
	@echo "  prod-build       Create a full production build (assets + Go binary)"
	@echo "  prod             Build and run the application in production mode"
	@echo "  dev-server       Run the development server (minifies assets first)"
	@echo ""
	@echo "Asset Generation:"
	@echo "  generate-styles  Generate CSS styles (e.g., Tailwind via npm)"
	@echo "  generate-html    Generate HTML from Markdown files"
	@echo "  minify           Minify CSS and project-specific HTML/JS files"
	@echo ""
	@echo "Code Quality & Maintenance:"
	@echo "  fmt              Format Go code"
	@echo "  lint             Run Go linting (requires golangci-lint)"
	@echo "  clean            Remove build artifacts and generated files"
	@echo ""
	@echo "Development Utilities:"
	@echo "  watch            Watch for Go & frontend file changes and rebuild automatically (requires fswatch)"
	@echo "  version          Show application version and build information"
	@echo "  check-deps       Check for all external tool dependencies"
	@echo "  help             Show this help message"
