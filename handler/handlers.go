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
		c.HTML(http.StatusOK, "resume", gin.H{})
	} else {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Content": "resume",
		})
	}
}

func ProjectsHandler(c *gin.Context) {

	projects, err := parser.ParseProjects(3)

	if err != nil {
		logger.LogError("Error parsing projects: " + err.Error())
		c.HTML(http.StatusInternalServerError, "error", gin.H{
			"Message": "Error parsing projects",
		})
		return
	}

	if c.GetHeader(HTMX_HEADER) == "true" {

		c.HTML(http.StatusOK, "projects", gin.H{
			"projects": projects,
		})
	} else {

		c.HTML(http.StatusOK, "index.html", gin.H{
			"Content":  "projects",
			"projects": projects,
		})
	}
}

func ContactHandler(c *gin.Context) {
	if c.GetHeader(HTMX_HEADER) == "true" {
		c.HTML(http.StatusOK, "contact", gin.H{})
	} else {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Content": "contact",
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
			"projects": projects,
		})
	} else {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Content":  "home",
			"projects": projects,
		})
	}
}
