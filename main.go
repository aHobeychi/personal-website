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

func main() {
	r := gin.Default()

	// Load the html files form the dynamically generated html array
	htmlFiles := getHtmlFiles()
	r.LoadHTMLFiles(htmlFiles...)

	// Serve static files
	r.Static("/static", "./static")

	// Routes
	r.GET("/", handler.IndexHandler)
	r.GET("/home", handler.HomeHandler)
	r.GET("/resume", handler.ResumeHandler)
	r.GET("/project", handler.ProjectsHandler)
	r.GET("/contact", handler.ContactHandler)

	r.Run(":8080")
}
