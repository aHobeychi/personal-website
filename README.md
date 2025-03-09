# Personal Website

A modern personal portfolio website built with Go (Gin framework), HTMX, and TailwindCSS. This project demonstrates a simple yet effective way to create a dynamic website with minimal JavaScript.

## Project Structure

```
├── go.mod                 # Go module definition and dependencies
├── go.sum                 # Go module checksums
├── main.go                # Application entry point and setup
├── .run-server.sh         # Script for development with live-reload
├── handler/
│   └── handlers.go        # HTTP route handlers
├── logger/
│   ├── logger.go          # Custom logging functionality
│   └── loggerMiddleware.go # Logging middleware for HTTP requests
├── models/
│   └── project.go         # Data structures for projects
├── parser/
│   └── project.go         # JSON parsing for project data
├── static/
│   ├── css/
│   │   ├── colors.css     # Color variables and theming
│   │   └── styles.css     # Custom styles
│   └── projects/
│       └── projects.json  # Project data in JSON format
├── templates/
│   ├── index.html         # Main template layout
│   ├── components/        # Reusable UI components
│   │   ├── footer.html
│   │   ├── header.html
│   │   └── sidebar.html
│   └── pages/             # Page-specific templates
│       ├── contact.html
│       ├── home.html
│       ├── projects.html
│       └── resume.html
└── tmp/                   # Temporary files for development
    ├── build-errors.log
    └── main
```

## Features

- **Go Backend**: Uses the Gin framework for routing and handling HTTP requests
- **HTMX Integration**: For seamless, JavaScript-free dynamic content updates
- **TailwindCSS**: For responsive and modern UI design
- **Project Showcase**: Dynamically loads and displays projects from JSON
- **Live Reload**: Development environment with automatic rebuilding and reloading

## How Projects.json Works

The `projects.json` file in the `static/projects/` directory stores information about your portfolio projects. Each project entry contains:

- `id`: Unique identifier for the project
- `name`: Project title
- `description`: Brief description of the project
- `tags`: Array of technologies used
- `link`: URL to the project or repository

Example:
```json
{
    "id": 1,
    "name": "Project Name",
    "description": "Description of the project",
    "tags": ["Go", "HTMX", "TailwindCSS"],
    "link": "https://github.com/username/project"
}
```

The `parser/project.go` module handles loading and caching this data, with the following features:
- Implements caching to reduce file I/O operations
- Supports configurable cache TTL (Time To Live)
- Allows limiting the number of projects returned (useful for homepage previews)

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

The application has the following route handlers defined in `handler/handlers.go`:

- `HomeHandler`: Serves the homepage with up to 3 featured projects
- `ResumeHandler`: Serves the resume page
- `ProjectsHandler`: Serves the projects page with all projects
- `ContactHandler`: Serves the contact page

Each handler uses a smart rendering approach that checks if the request is coming from HTMX (partial content) or a direct browser request (full page).

## License

[Your license information here]
