package handler

import (
	"aHobeychi/personal-website/logger"
	"aHobeychi/personal-website/parser"
	"net/http"
)

// ServeHomepage handles the home page
func ServeHomepage(w http.ResponseWriter, r *http.Request) {

	// Get 3 projects for the home page

	projects, projects_err := parser.ParseProjects(3)
	blogs, blogs_err := parser.ParseBlogs(3)

	if projects_err != nil {
		logger.LogError("Error parsing projects: " + projects_err.Error())
		return
	}

	if blogs_err != nil {
		logger.LogError("Error parsing blogs: " + blogs_err.Error())
		return
	}

	data := PageData{
		"projects": projects,
		"blogs":    blogs,
	}
	RenderTemplate(w, r, "home", data)
}
