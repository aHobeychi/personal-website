package handler

import (
	"net/http"
)

// ServeResume handles the resume page
func ServeResume(w http.ResponseWriter, r *http.Request) {
	// The ServeMux ensures this handler is only called for the exact path "/resume"
	// so we don't need to check r.URL.Path here

	data := PageData{
		"ActivePage": "resume",
	}
	RenderTemplate(w, r, "resume", data)
}
