package logger

import (
	"net/http"
	"path/filepath"
	"time"
)

// CustomLoggerMiddleware creates a middleware that logs HTTP requests
// It filters out CSS file requests to reduce log noise
func CustomLoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Skip logging for CSS files
		if filepath.Ext(r.URL.Path) == ".css" {
			next.ServeHTTP(w, r)
			return
		}

		// Log other requests
		startTime := time.Now()
		path := r.URL.Path
		raw := r.URL.RawQuery

		// Create a response writer wrapper to capture the status code
		rw := &responseWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK, // Default to 200 OK
		}

		// Process request
		next.ServeHTTP(rw, r)

		// Log request details
		endTime := time.Now()
		latency := endTime.Sub(startTime)
		clientIP := r.RemoteAddr
		statusCode := rw.statusCode
		method := r.Method

		if raw != "" {
			path = path + "?" + raw
		}

		// Format the log line
		logLine := "[HTTP] " + time.Now().Format("2006/01/02 - 15:04:05") + " | " +
			statusCodeToString(statusCode) + " | " +
			latency.String() + " | " +
			clientIP + " | " +
			method + " | " +
			path

		LogDebug(logLine)
	})
}

// responseWriter is a wrapper to capture the status code
type responseWriter struct {
	http.ResponseWriter
	statusCode int
	written    bool
}

// WriteHeader captures the status code before writing it
func (rw *responseWriter) WriteHeader(code int) {
	if !rw.written {
		rw.statusCode = code
		rw.ResponseWriter.WriteHeader(code)
		rw.written = true
	}
}

// Write captures that the body has been written to
func (rw *responseWriter) Write(b []byte) (int, error) {
	if !rw.written {
		// If WriteHeader was not explicitly called, it's called with StatusOK
		// before the first Write, as per the http.ResponseWriter interface
		rw.WriteHeader(http.StatusOK)
	}
	return rw.ResponseWriter.Write(b)
}

// Flush implements the http.Flusher interface
func (rw *responseWriter) Flush() {
	if f, ok := rw.ResponseWriter.(http.Flusher); ok {
		f.Flush()
	}
}

// Helper function to convert status code to string
func statusCodeToString(code int) string {
	return http.StatusText(code)
}
