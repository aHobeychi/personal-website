#!/bin/bash

# Define directories
SOURCE_DIR="templates"
OUTPUT_DIR="bin/html"

# Ensure output directory exists
mkdir -p "$OUTPUT_DIR"

# Find and minify HTML files in templates directory
find "$SOURCE_DIR" -type f -name '*.html' | while read -r file; do
  # Extract filename only (without path)
  filename=$(basename "$file")

  # Minify and move to bin/html
  minify -o "$OUTPUT_DIR/$filename" "$file"

  echo "Minified: $file -> $OUTPUT_DIR/$filename"
done

echo "✅ Minification complete! Minified files are in $OUTPUT_DIR"
