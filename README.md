# Personal Website

A modern personal portfolio website built with Go, HTMX, and TailwindCSS. This project demonstrates a simple yet effective way to create a dynamic website with minimal JavaScript.

## Project Structure

```text
├── app/                   # Application binary and server templates
│   ├── assets/            # Compiled assets (css, js, images)
│   ├── html/              # Compiled HTML templates
├── build/                 # Build-related resources
│   ├── ci/                # Continuous Integration configuration
│   ├── docker/            # Docker configuration
│   ├── scripts/           # Build and utility scripts
├── cmd/                   # Application entrypoints
│   └── server/            # Server entry point
├── config/                # Configuration files
├── frontend/              # Frontend assets and content
│   ├── assets/            # Frontend static assets (css, js, images)
│   ├── catalog/           # Content catalog in JSON format
│   ├── content/           # Content files (blog posts, etc.)
│   ├── templates/         # HTML templates
│       ├── components/    # Reusable UI components
│       ├── pages/         # Page-specific templates
├── internal/              # Internal application packages
    ├── cache/             # Caching mechanisms
    ├── config/            # Configuration handling
    ├── domain/            # Domain models
    ├── handler/           # HTTP request handlers
    ├── parser/            # Data parsing utilities
    ├── preprocessor/      # Content preprocessing
    └── util/              # Utility packages
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

[TailwindCSS](https://tailwindcss.com/) is used for styling. It provides:

- Utility-first CSS framework for rapid UI development
- Consistent design system
- Responsive design out of the box
- Support for translucent UI elements with opacity and backdrop blur utilities

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

## Deployment with Fly.io

This project supports seamless deployment to [Fly.io](https://fly.io) using the `fly.toml` configuration file. The file contains the following key configurations:

- **App Name**: Defines the unique application name on Fly.io platform
- **Primary Region**: Sets the primary deployment region (e.g., 'yul' for Montreal)
- **Build Configuration**: Points to the Docker build definition
- **Environment Variables**: Configures production environment settings
- **HTTP Service**: Defines server configuration including:
  - Internal port mapping
  - HTTPS enforcement
  - Machine scaling configuration
  - Resource allocation
- **Static Assets**: Configures serving of static assets from multiple directories:
  - `/app/app/assets` available at `/assets`
  - `/app/frontend/assets` available at `/frontend/assets`

### Deploying to Fly.io

To deploy your application:

```bash
fly deploy
```

The deployment uses the Docker configuration from `build/docker/Dockerfile` and sets up the application with production environment variables.

## Enhanced UI Features

### Sidebar Navigation

The project includes a responsive sidebar navigation system implemented in `frontend/assets/js/sidebar.js`. Key features:

- **Responsive Design**: Automatically shows/hides based on screen size
- **Touch-Friendly Mobile Navigation**: Slide-out menu on mobile devices
- **HTMX Integration**: Preserves sidebar state during partial page loads
- **Smart Event Handling**: Closes sidebar when clicking navigation links or outside the sidebar
- **Resize Handling**: Ensures proper sidebar state when resizing the browser window

### Scroll Spy

The `frontend/assets/js/scroll-spy.js` provides an intelligent table of contents navigation for blog posts:

- **Dynamic Highlighting**: Automatically highlights the current section in the table of contents based on scroll position
- **Performance Optimized**: Uses Intersection Observer API with fallback scroll listeners
- **HTMX Compatible**: Reinitializes when content changes via HTMX
- **Responsive**: Recalculates heading positions on window resize and when images finish loading
- **Error Handling**: Gracefully handles edge cases and prevents script errors

This enhances the reading experience by providing visual feedback about the reader's current position in longer articles.
