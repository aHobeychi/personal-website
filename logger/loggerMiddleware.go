package logger

import (
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// Custom logger middleware that filters out CSS file requests
func CustomLoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Skip logging for CSS files
		if filepath.Ext(c.Request.URL.Path) == ".css" {
			c.Next()
			return
		}

		// Log other requests using Gin's default logger format
		startTime := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		// Log only if the request is not for a CSS file
		endTime := time.Now()
		latency := endTime.Sub(startTime)
		clientIP := c.ClientIP()
		statusCode := strconv.Itoa(c.Writer.Status())
		method := c.Request.Method

		if raw != "" {
			path = path + "?" + raw
		}

		// Format similar to Gin's default logger
		logLine := "[GIN] " + time.Now().Format("2006/01/02 - 15:04:05") + " | " +
			statusCode + " | " +
			latency.String() + " | " +
			clientIP + " | " +
			method + " | " +
			path

		LogDebug(logLine)
	}
}
