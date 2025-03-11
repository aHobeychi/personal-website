package handler

import (
	"aHobeychi/personal-website/logger"
	"aHobeychi/personal-website/parser"
	"net/http"
)

// ServeProjectsList handles the projects page
func ServeProjectsList(w http.ResponseWriter, r *http.Request) {
	// The ServeMux ensures this handler is only called for the exact path "/project"
	// so we don't need to check r.URL.Path here

	projects, err := parser.ParseProjects()
	if err != nil {
		logger.LogError("Error parsing projects: " + err.Error())
		return
	}

	data := PageData{
		"projects":   projects,
		"ActivePage": "project",
	}
	RenderTemplate(w, r, "projects", data)
}
