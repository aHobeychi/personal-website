# Building My Personal Website with Go, HTMX, and TailwindCSS

A personal website is more than a digital résumé—it’s a space to showcase work and experiment with technology. I built mine to be fast, minimal, and developer-friendly, using **Go (standard library only)** for the backend, **HTMX** for interactivity, and **TailwindCSS** for styling.

## Stack Overview

- **Go**: Using only the `net/http` package, the backend remains lightweight, performant, and dependency-free.
- **HTMX**: Enables dynamic content loading via HTML attributes—no need for JavaScript frameworks.
- **TailwindCSS**: Provides utility-first styling with minimal custom CSS.

## Project Structure

Key directories:

- `cmd/`: Entry point (`main.go`)
- `internal/`: Core logic—handlers, caching, config, parsing
- `frontend/`: Assets, templates, JSON catalogs, Markdown content
- `app/`: Compiled output
- `build/`: Scripts and Docker/CI setup

## Backend Design

Handlers serve full pages or partial fragments depending on the `HX-Request` header. Configuration is loaded from JSON files based on `APP_ENV`. A TTL-based caching layer minimizes redundant file reads.

## Frontend Architecture

Templates use a layout-based structure with shared components (`navbar`, `footer`, etc.) rendered using Go’s `html/template`. TailwindCSS is compiled from source for optimized output. HTMX powers partial updates—for example, loading blog content without reloading the entire page.

## Content System

Content is stored in Markdown or JSON under `frontend/catalog/` and `frontend/content/blog/markdown/`. On startup, Go parsers load and cache content. A preprocessor generates blog TOCs from Markdown headings for enhanced in-page navigation.

## Development Workflow

The script `run-server.sh` uses `nodemon` to watch Go, HTML, CSS, JS, and JSON files for changes, automatically restarting the server. Assets are minified and compressed via utility scripts orchestrated by the `Makefile`.

## Deployment with Fly.io

Deployment is managed with Fly.io using `fly.toml` and a Docker-based workflow. Fly.io handles TLS, scaling, health checks, and static asset serving. Deployment is as simple as:

```bash
fly deploy
```

## Enhanced UI Features

Custom JavaScript improves UX:

- **Sidebar Navigation**: Responsive and mobile-friendly, preserving state during HTMX interactions.
- **Scroll Spy**: Highlights blog TOC sections based on scroll position using the Intersection Observer API.

## Conclusion

This site demonstrates how a performant, modern web experience can be built using simple, powerful tools. Go provides a robust backend, HTMX enables interactivity without bloat, and TailwindCSS ensures a clean, responsive UI. Future enhancements may include CMS integration, search, or localization.
