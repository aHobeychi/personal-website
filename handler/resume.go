package handler

import (
	"aHobeychi/personal-website/parser"
	"net/http"
)

// ServeResume handles the resume page
func ServeResume(w http.ResponseWriter, r *http.Request) {
	// The ServeMux ensures this handler is only called for the exact path "/resume"
	// so we don't need to check r.URL.Path here

	// Get the work experience data
	workExperience, err := parser.ParseWorkExperiences()
	if err != nil {
		http.Error(w, "Error loading work experience data", http.StatusInternalServerError)
		return
	}

	// Get the certification data
	certifications, err := parser.ParseCertifications()
	if err != nil {
		http.Error(w, "Error loading certification data", http.StatusInternalServerError)
		return
	}

	data := PageData{
		"ActivePage":     "resume",
		"WorkExperience": workExperience,
		"Certifications": certifications,
	}
	RenderTemplate(w, r, "resume", data)
}
