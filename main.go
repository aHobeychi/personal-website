package main

import (
	"aHobeychi/personal-website/handler"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

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
	r := gin.Default()

	// Apply no-cache middleware
	r.Use(noCacheMiddleware())

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
	r.GET("/contact", handler.ContactHandler)

	r.Run(":8080")
}
