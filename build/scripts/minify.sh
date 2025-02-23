#!/bin/bash

# Define directories
SOURCE_DIR="frontend/templates"
BLOG_DIR="frontend/content/blog/html/content"
BLOG_TOC_DIR="frontend/content/blog/html/table-of-contents"
JS_SCRIPT_DIR="frontend/assets/js/"

OUTPUT_DIR="app/html"
OUTPUT_DIR_ASSETS="app/assets/"
SRC_DIR=../

# Ensure output directory exists
mkdir -p "$OUTPUT_DIR"
mkdir -p "$OUTPUT_DIR_ASSETS"

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

find "$JS_SCRIPT_DIR" -type f -name '*.js' | while read -r file; do

  filename=$(basename "$file")
  minify -o "$OUTPUT_DIR_ASSETS/js/$filename" "$file"

  echo "Minified: $file -> $OUTPUT_DIR_ASSETS/$filename"
done

echo "✅ Minification of html and js complete! Minified files are in $OUTPUT_DIR"
