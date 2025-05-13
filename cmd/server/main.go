package main

import (
	"aHobeychi/personal-website/internal/config"
	"aHobeychi/personal-website/internal/handler"
	"aHobeychi/personal-website/internal/parser"
	"aHobeychi/personal-website/internal/preprocessor"
	"aHobeychi/personal-website/internal/util/logger"
	"aHobeychi/personal-website/internal/util/middleware"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// returns an array of all html files under the templates folder
func getHtmlFiles(path string) []string {
	var htmlFiles []string
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
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

func GenerateTableOfContents() {
	provider := parser.GetBlogProvider()
	if err := preprocessor.GenerateAllTableOfContents(provider); err != nil {
		logger.LogError("Failed to initialize table of contents: " + err.Error())
	}
}

func main() {
	config, err := config.Load()
	if err != nil {
		logger.LogError("Failed to load configuration: " + err.Error())
		return
	}

	// Set the log level based on the configuration
	logger.SetLogLevel(config.Logging.Level)
	logger.LogDebug("Environment set to: " + config.Server.Environment)

	if config.Server.Environment == "production" {
		logger.LogDebug("Production mode enabled")
		GenerateTableOfContents()
	} else {
		logger.LogDebug("Development mode enabled")
	}

	htmlFiles := getHtmlFiles(config.Paths.Templates)
	handler.InitializeTemplates(htmlFiles)

	// Create a new router using the standard library
	mux := http.NewServeMux()

	// Setup static file server
	fileServer := http.FileServer(http.Dir(config.Paths.AssetFiles))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	// Routes
	mux.HandleFunc("/", handler.ServeHomepage)
	mux.HandleFunc("/home", handler.ServeHomepage)
	mux.HandleFunc("/resume", handler.ServeResume)
	mux.HandleFunc("/project", handler.ServeProjectsList)
	mux.HandleFunc("/blog", handler.ServeBlogList)
	mux.HandleFunc("/blog/", func(w http.ResponseWriter, r *http.Request) {
		// Check if the request is for the table of contents
		if strings.Contains(r.URL.Path, "/table-of-contents") {
			handler.ServeBlogTableOfContents(w, r)
			return
		}
		// Otherwise, serve the regular blog content
		handler.ServeBlogContent(w, r)
	})

	// Apply middleware chain
	var handler http.Handler = mux

	if !config.Features.CacheEnabled {
		handler = middleware.NoCacheMiddleware(handler)
	}

	if config.Server.Environment == "production" {
		handler = middleware.DomainRedirectMiddleware(handler)
	}

	handler = logger.CustomLoggerMiddleware(handler)

	// Start the server
	logger.LogDebug("Server starting on port " + strconv.Itoa(config.Server.Port))
	err = http.ListenAndServe(":"+strconv.Itoa(config.Server.Port), handler)
	if err != nil {
		logger.LogError("Server failed to start: " + err.Error())
	}
}
