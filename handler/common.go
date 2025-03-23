package handler

import (
	"html/template"
	"net/http"
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
	if r.Header.Get(HTMX_HEADER) == "true" {
		// HTMX request - render just the partial template
		err := Templates.ExecuteTemplate(w, templateName, data)
		if err != nil {
			http.Error(w, "Error rendering template: "+err.Error(), http.StatusInternalServerError)
		}
	} else {
		// Regular request - render full page with index.html wrapper
		if data == nil {
			data = PageData{}
		}
		data["Content"] = templateName
		err := Templates.ExecuteTemplate(w, "index.html", data)
		if err != nil {
			http.Error(w, "Error rendering template: "+err.Error(), http.StatusInternalServerError)
		}
	}
}
