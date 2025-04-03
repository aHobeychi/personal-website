#!/bin/bash

# Ensure pngquant is installed
if ! command -v pngquant &>/dev/null; then
  echo "Error: pngquant is not installed. Install it using: brew install pngquant (macOS) or sudo apt install pngquant (Linux)"
  exit 1
fi

# Check if correct arguments are passed
if [ "$#" -ne 2 ]; then
  echo "Usage: $0 <input_file> <output_file>"
  exit 1
fi

input_file="$1"
output_file="$2"

# Check if input file exists
if [ ! -f "$input_file" ]; then
  echo "Error: Input file '$input_file' does not exist."
  exit 1
fi

# Compress PNG using pngquant
pngquant --quality=65-80 --speed 1 --output "$output_file" --force "$input_file"

# Verify output file creation
if [ -f "$output_file" ]; then
  echo "Compression successful! Output saved to: $output_file"
else
  echo "Error: Compression failed."
  exit 1
fi
