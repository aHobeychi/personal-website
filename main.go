package main

import (
	"aHobeychi/personal-website/handler"
	"aHobeychi/personal-website/logger"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

type Config struct {
	SERVER_PORT string
	CACHE_TTL   string
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
func noCacheMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Set headers to prevent caching for CSS files
		if filepath.Ext(c.Request.URL.Path) == ".css" {
			c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
			c.Header("Pragma", "no-cache")
			c.Header("Expires", "0")
		}
		c.Next()
	}
}

func main() {

	// Load configuration
	config := getConfig()
	// Set the log level based on the configuration
	logger.SetLogLevel(config.LOG_LEVEL)

	r := gin.New()

	// Set Gin to release mode if not in debug
	if !strings.EqualFold(config.LOG_LEVEL, "debug") {
		gin.SetMode(gin.ReleaseMode)
	}

	// Apply no-cache middleware
	r.Use(noCacheMiddleware())
	r.Use(logger.CustomLoggerMiddleware())

	// Load the html files form the dynamically generated html array
	htmlFiles := getHtmlFiles()
	r.LoadHTMLFiles(htmlFiles...)

	// Explicitly set static file serving with proper path
	r.Static("/static", "./static")

	// Routes
	r.GET("/", handler.HomeHandler)
	r.GET("/home", handler.HomeHandler)
	r.GET("/resume", handler.ResumeHandler)
	r.GET("/project", handler.ProjectsHandler)
	r.GET("/blog", handler.BlogHandler)
	r.GET("/contact", handler.ContactHandler)

	r.Run(":" + config.SERVER_PORT)
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
		CACHE_TTL:   getEnv("CACHE_TTL", "60"),
		LOG_LEVEL:   getEnv("LOG_LEVEL", "debug"),
	}
}
