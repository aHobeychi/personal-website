package parser

import (
	"aHobeychi/personal-website/logger"
	"aHobeychi/personal-website/models"
	"encoding/json"
	"os"
)

var (
	PATH_TOWARDS_BLOG_JSON = "static/content-catalog/blogs.json"
)

// ParseBlogs retrieves a list of blogs from file
// Optional limit parameter controls the maximum number of blogs returned
// Returns a slice of Blog models and any error encountered
func ParseBlogs(limit ...int) ([]models.Blog, error) {
	var blogs []models.Blog
	file, err := os.Open(PATH_TOWARDS_BLOG_JSON)
	if err != nil {
		logger.ErrorLogger.Println("Error opening blog JSON file:", err)
		return nil, err
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&blogs)
	if err != nil {
		logger.ErrorLogger.Println("Error decoding blog JSON:", err)
		return nil, err
	}

	logger.DebugLogger.Println("Blogs loaded successfully")

	// Apply limit if specified
	if len(limit) > 0 && limit[0] > 0 && limit[0] < len(blogs) {
		return blogs[:limit[0]], nil
	}

	return blogs, nil
}

// GetBlogHTMLContent returns the HTML content of a blog post by its ID.
// The returned string contains raw HTML that should be served with
// content type "text/html" to be properly rendered in a browser.
func GetBlogHTMLContent(blogId string) (string, error) {
	// create formatted path from blogId
	blogPath := "static/blog-posts/markdown/" + blogId + ".html"

	// Read the file content directly
	content, err := os.ReadFile(blogPath)
	if err != nil {
		logger.ErrorLogger.Println("Error reading blog content file:", err)
		return "", err
	}

	// convert content to string
	contentString := string(content)
	logger.DebugLogger.Printf("HTML content retrieved for blog ID: %s", blogId)

	return contentString, nil
}
