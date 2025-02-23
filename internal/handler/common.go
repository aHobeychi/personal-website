package handler

import (
	"html/template"
	"net/http"

	"aHobeychi/personal-website/internal/config"
)

// HTMX_HEADER is the header name that HTMX sends to indicate an HTMX request
const HTMX_HEADER = "HX-Request"

// PageData represents the common data structure for template rendering
type PageData map[string]any

// Templates stores all parsed templates
var Templates *template.Template

// InitializeTemplates parses all HTML templates and stores them for later use
func InitializeTemplates(templateFiles []string) {
	var err error
	Templates, err = template.ParseFiles(templateFiles...)
	if err != nil {
		panic("Error parsing templates: " + err.Error())
	}
}

// RenderTemplate renders the appropriate template based on whether it's an HTMX request
func RenderTemplate(w http.ResponseWriter, r *http.Request, templateName string, data PageData) {
	// Check if this is an HTMX request
	if data == nil {
		data = PageData{}
	}

	// Add displayBlogs configuration to all templates
	cfg := config.Get()
	data["DisplayBlogs"] = cfg.Features.DisplayBlogs

	if r.Header.Get(HTMX_HEADER) == "true" {
		// HTMX request - render just the partial template
		err := Templates.ExecuteTemplate(w, templateName, data)
		if err != nil {
			http.Error(w, "Error rendering template: "+err.Error(), http.StatusInternalServerError)
		}
	} else {
		// Regular request - render full page with index.html wrapper
		data["Content"] = templateName
		err := Templates.ExecuteTemplate(w, "index.html", data)
		if err != nil {
			http.Error(w, "Error rendering template: "+err.Error(), http.StatusInternalServerError)
		}
	}
}

// GetRefererPage determines the source page based on the Referer header
// If the Referer header is empty or does not match known pages, it defaults to nil"
func GetRefererPage(r *http.Request) string {
	// Determine the source page by checking the Referer header
	return r.Header.Get("Referer")
}
