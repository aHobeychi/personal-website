package handler

import (
	"net/http"

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
	if c.GetHeader(HTMX_HEADER) == "true" {
		c.HTML(http.StatusOK, "projects", gin.H{})
	} else {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Content": "projects",
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
	if c.GetHeader(HTMX_HEADER) == "true" {
		c.HTML(http.StatusOK, "home", gin.H{})
	} else {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Content": "home",
		})
	}
}
