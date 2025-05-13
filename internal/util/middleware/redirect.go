package middleware

import (
	"net/http"
	"strings"
)

// DomainRedirectMiddleware redirects requests from fly.io domain to custom domain
func DomainRedirectMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the request is coming from fly.io domain
		if strings.Contains(r.Host, "fly.dev") {
			targetURL := "https://alexhobeychi.com" + r.URL.Path
			if r.URL.RawQuery != "" {
				targetURL += "?" + r.URL.RawQuery
			}
			http.Redirect(w, r, targetURL, http.StatusMovedPermanently) // 301 redirect
			return
		}
		next.ServeHTTP(w, r)
	})
}
