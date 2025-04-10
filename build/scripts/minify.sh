#!/bin/bash

# Define directories
SOURCE_DIR="frontend/templates"
BLOG_DIR="frontend/content/blog/html/content"
BLOG_TOC_DIR="frontend/content/blog/html/table-of-contents"

OUTPUT_DIR="app/html"
SRC_DIR=../

# Ensure output directory exists
mkdir -p "$OUTPUT_DIR"

# Find and minify HTML files in templates directory
find "$SOURCE_DIR" -type f -name '*.html' | while read -r file; do
  # Extract filename only (without path)
  filename=$(basename "$file")

  # Minify and move to app/html
  minify -o "$OUTPUT_DIR/templates/$filename" "$file"

  echo "Minified: $file -> $OUTPUT_DIR/$filename"
done

# Find and minify HTML files in blog directory
find "$BLOG_DIR" -type f -name '*.html' | while read -r file; do
  # Extract filename only (without path)
  filename=$(basename "$file")

  # Minify and move to app/html
  minify -o "$OUTPUT_DIR/blog/$filename" "$file"

  echo "Minified: $file -> $OUTPUT_DIR/$filename"
done

# Find and minify HTML files in blog table-of-contents directory
find "$BLOG_TOC_DIR" -type f -name '*.html' | while read -r file; do
  # Extract filename only (without path)
  filename=$(basename "$file")

  # Minify and move to app/html
  minify -o "$OUTPUT_DIR/toc/$filename" "$file"

  echo "Minified: $file -> $OUTPUT_DIR/$filename"
done

echo "✅ Minification of html complete! Minified files are in $OUTPUT_DIR"
