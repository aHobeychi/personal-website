package handler

import (
	"html/template"
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

func BlogContentHandler(c *gin.Context) {
	blogId := c.Param("blogId")

	// Get blog metadata to display title in breadcrumbs
	blogs, _ := parser.ParseBlogs()
	var blogTitle string = "Blog Post" // Default title if not found

	// Find the blog with matching ID to get its title
	for _, blog := range blogs {
		if blog.Id == blogId {
			blogTitle = blog.Title
			break
		}
	}

	// Parse the blog content from the file
	content, err := parser.GetBlogHTMLContent(blogId)
	if err != nil {
		logger.LogError("Error parsing blog content: " + err.Error())
		c.HTML(http.StatusInternalServerError, "error", gin.H{
			"Message": "Error parsing blog content",
		})
		return
	}

	// Convert the HTML content string to template.HTML to prevent escaping
	htmlContent := template.HTML(content)

	if c.GetHeader(HTMX_HEADER) == "true" {
		c.HTML(http.StatusOK, "blog-content", gin.H{
			"ContentData": htmlContent,
			"BlogTitle":   blogTitle,
			"ActivePage":  "blog",
		})
	} else {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Content":     "blog-content",
			"ContentData": htmlContent,
			"BlogTitle":   blogTitle,
			"ActivePage":  "blog",
		})
	}
}

func BlogHandler(c *gin.Context) {
	blogs, _ := parser.ParseBlogs() // TODO: Add error handling

	if c.GetHeader(HTMX_HEADER) == "true" {
		c.HTML(http.StatusOK, "blog-list", gin.H{
			"blogs":      blogs,
			"ActivePage": "blog",
		})
	} else {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"blogs":      blogs,
			"Content":    "blog-list",
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
