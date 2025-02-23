package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handlers
func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "base.html", gin.H{
		"Content": "index.html",
		"Title":   "Home",
	})
}

func AboutHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "base.html", gin.H{
		"Content": "about.html",
		"Title":   "About Me",
	})
}

func ResumeHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "base.html", gin.H{
		"Content": "resume.html",
		"Title":   "Resume",
	})
}

func ProjectsHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "base.html", gin.H{
		"Content": "projects.html",
		"Title":   "Projects",
	})
}

func ContactHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "base.html", gin.H{
		"Content": "contact.html",
		"Title":   "Contact",
	})
}
