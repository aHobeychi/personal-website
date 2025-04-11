package handler

import (
	"aHobeychi/personal-website/internal/parser"
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
		"BlogID":      blog.Id,
		"ContentData": template.HTML(contentData), // Convert to template.HTML to prevent escaping
		"SourcePage":  sourcePage,                 // Add source page information
	}

	// RenderTemplate already checks for HTMX headers and renders appropriately
	RenderTemplate(w, r, "blog-content", data)
}

// ServeBlogTableOfContents serves the pre-generated table of contents for a blog post
func ServeBlogTableOfContents(w http.ResponseWriter, r *http.Request) {
	// Extract blog ID from URL path
	path := strings.TrimPrefix(r.URL.Path, "/blog/")

	parts := strings.Split(path, "/")
	if len(parts) < 2 {
		http.Error(w, "Invalid blog ID", http.StatusBadRequest)
		return
	}

	blogID := parts[0] // The first part should be the blog ID

	// Get the blog post
	blog, err := parser.GetBlogByID(blogID)
	if err != nil {
		http.Error(w, "Blog not found", http.StatusNotFound)
		return
	}

	// Get the pre-generated table of contents - now using the simplified parser method
	tocContent, err := parser.GetBlogTableOfContents(blog.Id)
	if err != nil {
		http.Error(w, "Failed to load table of contents", http.StatusInternalServerError)
		return
	}

	// Set content type
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// Write the pre-generated content directly to the response
	w.Write([]byte(tocContent))
}

// ExtractBlogIDFromPath extracts the blog ID from the URL path
func extractBlogIDFromPath(path string) string {
	// Extracts the last segment from a URL path like "/blog/my-blog-post"
	return filepath.Base(path)
}
