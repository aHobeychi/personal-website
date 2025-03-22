package main

import (
	"aHobeychi/personal-website/handler"
	"aHobeychi/personal-website/logger"
	"net/http"
	"os"
	"path/filepath"
)

type Config struct {
	SERVER_PORT string
	LOG_LEVEL   string
}

// returns an array of all html files under the templates folder
func getHtmlFiles() []string {
	var htmlFiles []string
	err := filepath.Walk("templates", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".html" {
			htmlFiles = append(htmlFiles, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	return htmlFiles
}

// noCacheMiddleware sets headers to prevent caching for CSS files.
// This middleware is useful during development to ensure the latest changes are always loaded.
// Note: This middleware should be removed before deploying to production.
func noCacheMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set headers to prevent caching for CSS files
		if filepath.Ext(r.URL.Path) == ".css" {
			w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
			w.Header().Set("Pragma", "no-cache")
			w.Header().Set("Expires", "0")
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	// Load configuration
	config := getConfig()

	// Set the log level based on the configuration
	logger.SetLogLevel(config.LOG_LEVEL)

	// Parse all templates
	htmlFiles := getHtmlFiles()
	handler.InitializeTemplates(htmlFiles)

	// Create a new router using the standard library
	mux := http.NewServeMux()

	// Setup static file server
	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	// Routes
	mux.HandleFunc("/", handler.ServeHomepage)
	mux.HandleFunc("/home", handler.ServeHomepage)
	mux.HandleFunc("/resume", handler.ServeResume)
	mux.HandleFunc("/project", handler.ServeProjectsList)
	mux.HandleFunc("/blog", handler.ServeBlogList)
	mux.HandleFunc("/blog/", handler.ServeBlogPost) // Will handle /blog/{blogId} paths

	// Apply middleware chain
	var handler http.Handler = mux
	handler = noCacheMiddleware(handler)
	handler = logger.CustomLoggerMiddleware(handler)

	// Start the server
	logger.LogDebug("Server starting on port " + config.SERVER_PORT)
	err := http.ListenAndServe(":"+config.SERVER_PORT, handler)
	if err != nil {
		logger.LogError("Server failed to start: " + err.Error())
	}
}

// getConfig retrieves configuration values from environment variables or defaults
// to predefined values. This function is useful for setting up application
// configurations without hardcoding them in the source code.
func getConfig() Config {
	getEnv := func(key, defaultValue string) string {
		if value, exists := os.LookupEnv(key); exists {
			return value
		}
		return defaultValue
	}

	return Config{
		SERVER_PORT: getEnv("SERVER_PORT", "8080"),
		LOG_LEVEL:   getEnv("LOG_LEVEL", "debug"),
	}
}
