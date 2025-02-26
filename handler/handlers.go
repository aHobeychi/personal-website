package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var HTMX_HEADER = "HX-Request"

// Handlers
func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"Content": "index.html",
		"Title":   "Home",
	})
}

func ResumeHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "resume.html", gin.H{
		"Title": "Resume",
	})
}

func ProjectsHandler(c *gin.Context) {
	if c.GetHeader(HTMX_HEADER) == "true" {
		c.HTML(http.StatusOK, "projects.html", gin.H{
			"Title": "Projects",
		})
	} else {
		c.HTML(http.StatusOK, "projects.html", gin.H{
			"Title": "Page",
		})
	}
}

func ContactHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "contact.html", gin.H{
		"Title": "Contact",
	})
}

func HomeHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", gin.H{
		"Title": "Contact",
	})
}
