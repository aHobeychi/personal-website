package handler

import (
	"aHobeychi/personal-website/parser"
	"html/template"
	"net/http"
	"path/filepath"
	"strings"
)

// ServeBlogList handles the blog list page
func ServeBlogList(w http.ResponseWriter, r *http.Request) {
	// The ServeMux ensures this handler is only called for the exact path "/blog"
	// so we don't need to check r.URL.Path here

	blogs, err := parser.ParseBlogs()
	if err != nil {
		http.Error(w, "Error loading blog data", http.StatusInternalServerError)
		return
	}

	data := PageData{
		"blogs": blogs,
	}

	// RenderTemplate already checks for HTMX headers and renders appropriately
	RenderTemplate(w, r, "blog-list", data)
}

// ServeBlogContent handles rendering a specific blog post
func ServeBlogContent(w http.ResponseWriter, r *http.Request) {

	// Extract blog ID from URL path using the existing helper function
	id := extractBlogIDFromPath(r.URL.Path)

	blog, err := parser.GetBlogByID(id)
	if err != nil {
		http.Error(w, "Blog not found", http.StatusNotFound)
		return
	}

	contentData, err := parser.GetBlogHTMLContent(blog.Id)
	if err != nil {
		http.Error(w, "Failed to load blog content", http.StatusInternalServerError)
		return
	}

	// Determine the source page by checking the Referer header
	sourcePage := "notes" // Default to "notes"
	referer := GetRefererPage(r)
	if referer != "" {
		if strings.Contains(referer, "/home") {
			sourcePage = "home"
		}
	}

	data := PageData{
		"BlogTitle":   blog.Title,
		"ContentData": template.HTML(contentData), // Convert to template.HTML to prevent escaping
		"SourcePage":  sourcePage,                 // Add source page information
	}

	// RenderTemplate already checks for HTMX headers and renders appropriately
	RenderTemplate(w, r, "blog-content", data)
}

// ExtractBlogIDFromPath extracts the blog ID from the URL path
func extractBlogIDFromPath(path string) string {
	// Extracts the last segment from a URL path like "/blog/my-blog-post"
	return filepath.Base(path)
}
