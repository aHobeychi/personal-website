# Personal Website

A modern personal portfolio website built with Go, HTMX, and TailwindCSS. This project demonstrates a simple yet effective way to create a dynamic website with minimal JavaScript.

## Project Structure

```text
├── go.mod                 # Go module definition and dependencies
├── go.sum                 # Go module checksums
├── Makefile               # Makefile for build automation
├── package.json           # Node.js dependencies
├── tailwind.config.js     # TailwindCSS configuration
├── app/                   # Application binary and server templates
│   ├── personalwebsite    # Compiled application binary
│   ├── html/              # Compiled HTML templates
│       ├── blog/          # Compiled blog HTML files
│       ├── templates/     # Compiled template files
│       ├── toc/           # Compiled table of contents files
├── build/                 # Build-related resources
│   ├── ci/                # Continuous Integration configuration
│   ├── docker/            # Docker configuration
│   │   └── Dockerfile     # Docker build definition
│   ├── scripts/           # Build and utility scripts
│       ├── compress_image.sh  # Script for image compression
│       ├── minify.sh      # Script for asset minification
│       └── run-server.sh  # Development server script
├── cmd/                   # Application entrypoints
│   └── server/            # Server entry point
│       └── main.go        # Main application entry point
├── config/                # Configuration files
│   ├── development.json   # Development environment configuration
│   └── production.json    # Production environment configuration
├── frontend/              # Frontend assets and content
│   ├── assets/            # Frontend static assets
│   │   ├── css/           # CSS stylesheets
│   │   │   ├── styles.css # Compiled styles
│   │   │   └── tailwind.css # Tailwind configuration
│   │   ├── images/        # Image assets
│   ├── catalog/           # Content catalog in JSON format
│   │   ├── blogs.json     # Blog data
│   │   ├── certifications.json # Certification data
│   │   ├── favorites.json # Favorites data
│   │   ├── projects.json  # Project data
│   │   └── work-experience.json # Work experience data
│   ├── content/           # Content files
│   │   └── blog/          # Blog content
│   │       ├── html/      # Generated HTML blog content
│   │       │   ├── content/  # HTML content files
│   │       │   └── table-of-contents/ # Table of contents files
│   │       └── markdown/  # Source markdown files
│   ├── templates/         # HTML templates
│       ├── index.html     # Main template layout
│       ├── components/    # Reusable UI components
│       │   ├── footer.html
│       │   ├── header.html
│       │   ├── navbar.html
│       │   └── sidebar.html
│       ├── pages/         # Page-specific templates
│           ├── blog-content.html
│           ├── blog-list.html
│           ├── home.html
│           ├── projects.html
│           └── resume.html
├── internal/              # Internal application packages
│   ├── cache/             # Caching mechanisms
│   │   └── cache.go       # Generic caching functionality
│   ├── config/            # Configuration handling
│   │   └── config.go      # Configuration loading and processing
│   ├── domain/            # Domain models
│   │   ├── blog.go        # Blog domain models
│   │   ├── certification.go # Certification models
│   │   ├── project.go     # Project models
│   │   └── work_experience.go # Work experience models
│   ├── handler/           # HTTP request handlers
│   │   ├── blog.go        # Blog route handlers
│   │   ├── common.go      # Common handler utilities
│   │   ├── home.go        # Home route handlers
│   │   ├── projects.go    # Projects route handlers
│   │   └── resume.go      # Resume route handlers
│   ├── parser/            # Data parsing utilities
│   │   ├── blog_provider.go # Blog provider interface
│   │   ├── blog.go        # Blog data parsing
│   │   ├── certification.go # Certification data parsing
│   │   ├── project.go     # Project data parsing
│   │   └── work_experience.go # Work experience parsing
│   ├── preprocessor/      # Content preprocessing
│   │   └── table_of_contents.go # Table of contents generation
│   └── util/              # Utility packages
│       ├── logger/        # Logging utilities
│       │   ├── logger.go  # Logging configuration
│       │   └── logger_middleware.go # HTTP logging middleware
│       └── middleware/    # HTTP middleware
│           └── no_cache.go # Cache control middleware
```

## Features

- **Go Backend**: Is limited to the Go Standard Library
- **HTMX Integration**: For seamless, JavaScript-free dynamic content updates
- **TailwindCSS**: For responsive and modern UI design
- **Project Showcase**: Dynamically loads and displays projects from JSON
- **Blog Integration**: Dynamically loads and displays blog posts from JSON and Markdown
- **Live Reload**: Development environment with automatic rebuilding and reloading
- **HTML Minification**: Both at build time and runtime for optimal performance

## How Blogs.json Works

The `blogs.json` file in the `frontend/catalog/` directory stores information about your blog posts. Each blog entry contains:

- `id`: Unique identifier for the blog post
- `title`: Blog post title
- `summary`: Brief summary of the blog post
- `tags`: Array of tags associated with the blog post
- `link`: URL to the blog post

Example:

```json
{
    "id": 1,
    "title": "Blog Post Title",
    "summary": "Summary of the blog post",
    "tags": ["Go", "HTMX", "TailwindCSS"],
    "link": "https://yourwebsite.com/blog/post"
}
```

The `internal/parser/blog.go` module handles loading and caching this data, with the following features:

- Implements caching to reduce file I/O operations
- Supports configurable cache TTL (Time To Live)
- Allows limiting the number of blog posts returned (useful for homepage previews)

## HTMX Integration

This project uses [HTMX](https://htmx.org/) to create dynamic content without writing JavaScript. HTMX allows for:

- Partial page updates without full page reloads
- Clean separation of concerns (HTML for structure, CSS for presentation)
- Progressive enhancement approach to web development

The handlers check for the `HX-Request` header to determine if a request is coming from HTMX, then return either a full page or just the content fragment as appropriate.

## TailwindCSS Integration

[TailwindCSS](https://tailwindcss.com/) is used for styling through a CDN for simplicity. It provides:

- Utility-first CSS framework for rapid UI development
- Consistent design system
- Responsive design out of the box

The project also uses Flowbite components to enhance the UI.

## Configuration System

The application now uses a robust configuration system with environment-specific JSON files located in the `config/` directory:

- `development.json`: Configuration for development environments
- `production.json`: Configuration for production environments

### Configuration Structure

The configuration files follow this structure:

```json
{
  "server": {
    "port": 8080,
    "host": "localhost"
  },
  "logging": {
    "level": "debug"
  },
  "caching": {
    "ttl": 60
  },
  "content": {
    "catalogPath": "./frontend/catalog",
    "templatesPath": "./frontend/templates"
  }
}
```

The application automatically loads the appropriate configuration file based on the `APP_ENV` environment variable (defaults to "development" if not specified).

### Accessing Configuration

The configuration is loaded and managed by the `internal/config/config.go` module, which provides a clean API for accessing configuration values throughout the application.

## Content Preprocessing

The application now includes a preprocessing system located in `internal/preprocessor/` that transforms content before serving:

### Table of Contents Generator

The `table_of_contents.go` preprocessor automatically generates table of contents for blog posts by:

1. Parsing Markdown files from `frontend/content/blog/markdown/`
2. Extracting headings and creating a hierarchical structure
3. Generating HTML table of contents files in `frontend/content/blog/html/table-of-contents/`
4. These table of contents files are then injected into blog content when viewed

This enhances navigation within blog posts without requiring JavaScript for generation.

## Environment Variables

The application uses the following environment variables:

| Variable     | Default       | Description                               |
|-------------|---------------|-------------------------------------------|
| APP_ENV     | "development" | Environment to run the application in     |
| SERVER_PORT | "8080"        | The port the web server will listen on    |
| CACHE_TTL   | "60"          | Cache time-to-live in seconds             |
| LOG_LEVEL   | "debug"       | Logging level (debug, info, warn, error)  |

## Development

### Live Reload with Nodemon

The `build/scripts/run-server.sh` script sets up a development environment with automatic rebuilding and reloading:

```bash
npx nodemon \
  --watch "**" \
  --ext "go,html,js,json,css" \
  --signal SIGTERM \
  --exec "go run ${PWD}/cmd/server/main.go"
```

This script uses Nodemon to watch for changes in any file with the specified extensions and restarts the Go application when changes are detected.

### Running the Application

1. Make sure you have Go installed
2. Install dependencies:

   ```bash
   go mod download
   ```

3. For development with live reload:

   ```bash
   sh build/scripts/run-server.sh
   ```

4. For production:

   ```bash
   go run cmd/server/main.go
   ```

## Handlers

The application has the following route handlers defined in `internal/handler/`:

- `HomeHandler`: Serves the homepage with up to 3 featured projects
- `ResumeHandler`: Serves the resume page
- `ProjectsHandler`: Serves the projects page with all projects
- `ContactHandler`: Serves the contact page
- `BlogHandler`: Serves the blog list and individual blog posts

Each handler uses a smart rendering approach that checks if the request is coming from HTMX (partial content) or a direct browser request (full page).