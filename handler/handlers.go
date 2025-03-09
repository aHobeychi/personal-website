package handler

import (
	"net/http"

	"aHobeychi/personal-website/logger"
	"aHobeychi/personal-website/parser"

	"github.com/gin-gonic/gin"
)

var HTMX_HEADER = "HX-Request"

func ResumeHandler(c *gin.Context) {
	if c.GetHeader(HTMX_HEADER) == "true" {
		c.HTML(http.StatusOK, "resume", gin.H{
			"ActivePage": "resume",
		})
	} else {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Content":    "resume",
			"ActivePage": "resume",
		})
	}
}

func BlogHandler(c *gin.Context) {
	if c.GetHeader(HTMX_HEADER) == "true" {
		c.HTML(http.StatusOK, "blog", gin.H{
			"ActivePage": "blog",
		})
	} else {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Content":    "blog",
			"ActivePage": "blog",
		})
	}
}

func ProjectsHandler(c *gin.Context) {

	projects, err := parser.ParseProjects()

	if err != nil {
		logger.LogError("Error parsing projects: " + err.Error())
		c.HTML(http.StatusInternalServerError, "error", gin.H{
			"Message": "Error parsing projects",
		})
		return
	}

	if c.GetHeader(HTMX_HEADER) == "true" {

		c.HTML(http.StatusOK, "projects", gin.H{
			"projects":   projects,
			"ActivePage": "project",
		})
	} else {

		c.HTML(http.StatusOK, "index.html", gin.H{
			"Content":    "projects",
			"projects":   projects,
			"ActivePage": "project",
		})
	}
}

func ContactHandler(c *gin.Context) {
	if c.GetHeader(HTMX_HEADER) == "true" {
		c.HTML(http.StatusOK, "contact", gin.H{
			"ActivePage": "contact",
		})
	} else {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Content":    "contact",
			"ActivePage": "contact",
		})
	}
}

func HomeHandler(c *gin.Context) {
	// Get 3 projects for the home page
	projects, err := parser.ParseProjects(3)

	if err != nil {
		logger.LogError("Error parsing projects: " + err.Error())
		c.HTML(http.StatusInternalServerError, "error", gin.H{
			"Message": "Error parsing projects",
		})
		return
	}

	if c.GetHeader(HTMX_HEADER) == "true" {
		c.HTML(http.StatusOK, "home", gin.H{
			"projects":   projects,
			"ActivePage": "home",
		})
	} else {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Content":    "home",
			"projects":   projects,
			"ActivePage": "home",
		})
	}
}
