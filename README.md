# Personal Website
A modern personal portfolio website built with Go, HTMX, and TailwindCSS. This project demonstrates a simple yet effective way to create a dynamic website with minimal JavaScript.

## Project Structure
```
├── go.mod                 # Go module definition and dependencies
├── go.sum                 # Go module checksums
├── main.go                # Application entry point and setup
├── Makefile               # Makefile for build automation
├── handler/
│   ├── blog.go            # Blog route handlers
│   ├── common.go          # Common handlers and utilities
│   ├── home.go            # Home route handlers
│   ├── projects.go        # Projects route handlers
│   └── resume.go          # Resume route handlers
├── logger/
│   ├── logger.go          # Custom logging functionality
│   └── logger_middleware.go # Logging middleware for HTTP requests
├── models/
│   ├── blog.go            # Data structures for blog posts
│   └── project.go         # Data structures for projects
├── parser/
│   ├── blog.go            # JSON parsing for blog data
│   └── project.go         # JSON parsing for project data
├── static/
│   ├── blog-posts/
│   │   ├── html/          # HTML blog posts
│   │   └── markdown/      # Markdown blog posts
│   ├── content-catalog/
│   │   ├── blogs.json     # Blog data in JSON format
│   │   └── projects.json  # Project data in JSON format
│   ├── css/
│   │   ├── blogs.css      # Blog-specific styles
│   │   ├── colors.css     # Color variables and theming
│   │   └── styles.css     # Custom styles
├── templates/
│   ├── index.html         # Main template layout
│   ├── components/        # Reusable UI components
│   │   ├── footer.html
│   │   ├── header.html
│   │   ├── navbar.html
│   │   └── sidebar.html
│   ├── pages/             # Page-specific templates
│   │   ├── blog-content.html
│   │   ├── blog-list.html
│   │   ├── blog.html
│   │   ├── contact.html
│   │   ├── home.html
│   │   ├── projects.html
│   │   └── resume.html
└── bin/
    └── myapp              # Compiled binary
```

## Features
- **Go Backend**: Is limited to the Go Standard Library
- **HTMX Integration**: For seamless, JavaScript-free dynamic content updates
- **TailwindCSS**: For responsive and modern UI design
- **Project Showcase**: Dynamically loads and displays projects from JSON
- **Blog Integration**: Dynamically loads and displays blog posts from JSON and Markdown
- **Live Reload**: Development environment with automatic rebuilding and reloading

## How Blogs.json Works
The `blogs.json` file in the `static/content-catalog/` directory stores information about your blog posts. Each blog entry contains:
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

The `parser/blog.go` module handles loading and caching this data, with the following features:
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

## Environment Variables
The application uses the following environment variables:
| Variable     | Default | Description                               |
|-------------|---------|-------------------------------------------|
| SERVER_PORT | "8080"  | The port the web server will listen on    |
| CACHE_TTL   | "60"    | Cache time-to-live in seconds             |
| LOG_LEVEL   | "debug" | Logging level (debug, info, warn, error)  |

## Development
### Live Reload with Nodemon
The `.run-server.sh` script sets up a development environment with automatic rebuilding and reloading:
```bash
npx nodemon \
  --watch "**" \
  --ext "go,html,js,json,css" \
  --signal SIGTERM \
  --exec "go run ${PWD}/main.go"
```

This script uses Nodemon to watch for changes in any file with the specified extensions and restarts the Go application when changes are detected.

### Running the Application
1. Make sure you have Go installed
2. Install dependencies:
   ```
   go mod download
   ```
3. For development with live reload:
   ```
   sh .run-server.sh
   ```
4. For production:
   ```
   go run main.go
   ```

## Handlers

The application has the following route handlers defined in `handler/`:

- `HomeHandler`: Serves the homepage with up to 3 featured projects
- `ResumeHandler`: Serves the resume page
- `ProjectsHandler`: Serves the projects page with all projects
- `ContactHandler`: Serves the contact page
- `BlogHandler`: Serves the blog list and individual blog posts

Each handler uses a smart rendering approach that checks if the request is coming from HTMX (partial content) or a direct browser request (full page).

## License
[Your license information here]
