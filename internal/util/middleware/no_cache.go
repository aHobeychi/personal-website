package middleware

import (
	"net/http"
	"path/filepath"
)

// noCacheMiddleware sets headers to prevent caching for CSS files.
// This middleware is useful during development to ensure the latest changes are always loaded.
// Note: This middleware should be removed before deploying to production.
func NoCacheMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set headers to prevent caching for CSS files
		if filepath.Ext(r.URL.Path) == ".css" {
			w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
			w.Header().Set("Pragma", "no-cache")
			w.Header().Set("Expires", "0")
		}
		next.ServeHTTP(w, r)
	})
}
