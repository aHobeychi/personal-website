#!/bin/bash
npx nodemon \
  --watch "cmd/" \
  --watch "internal/" \
  --watch "frontend/templates/" \
  --watch "frontend/assets/" \
  --watch "frontend/content/blog/html/content" \
  --watch "frontend/content/markdown" \
  --watch "config/" \
  --ext "go,html,js,json,css" \
  --ignore "app/" \
  --ignore "frontend/assets/css/styles.css" \
  --ignore "*.tmp" \
  --ignore "*.log" \
  --signal SIGTERM \
  --exec "make build run"
