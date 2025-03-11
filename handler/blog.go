package handler

import (
	"aHobeychi/personal-website/logger"
	"aHobeychi/personal-website/parser"
	"html/template"
	"net/http"
	"strings"
)

// ServeBlogList handles the blog listing page
func ServeBlogList(w http.ResponseWriter, r *http.Request) {
	// The ServeMux ensures this handler is only called for the exact path "/blog"
	// so we don't need to check r.URL.Path here

	blogs, err := parser.ParseBlogs()
	if err != nil {
		logger.LogError("Error parsing blogs: " + err.Error())
		return
	}

	data := PageData{
		"blogs":      blogs,
		"ActivePage": "blog",
	}

	RenderTemplate(w, r, "blog-list", data)
}

// ServeBlogPost handles rendering the content of a specific blog post
func ServeBlogPost(w http.ResponseWriter, r *http.Request) {

	// Make sure the path is in the format /blog/{blogId}
	if r.URL.Path == "/blog/" || r.URL.Path == "/blog" {
		http.Redirect(w, r, "/blog", http.StatusSeeOther)
		return
	}

	// Extract blogId from URL path
	blogId := strings.TrimPrefix(r.URL.Path, "/blog/")

	// Get blog metadata to display title in breadcrumbs
	blogs, err := parser.ParseBlogs()
	var blogTitle string = "Blog Post" // Default title if not found

	if err != nil {
		logger.LogError("Error parsing blogs for metadata: " + err.Error())
		// Continue with default title if we can't get blog metadata
	} else {
		// Find the blog with matching ID to get its title
		for _, blog := range blogs {
			if blog.Id == blogId {
				blogTitle = blog.Title
				break
			}
		}
	}

	// Parse the blog content from the file
	content, err := parser.GetBlogHTMLContent(blogId)
	if err != nil {
		logger.LogError("Error parsing blog content: " + err.Error())
		return
	}

	// Convert the HTML content string to template.HTML to prevent escaping
	htmlContent := template.HTML(content)

	data := PageData{
		"ContentData": htmlContent,
		"BlogTitle":   blogTitle,
		"ActivePage":  "blog",
	}
	RenderTemplate(w, r, "blog-content", data)
}
