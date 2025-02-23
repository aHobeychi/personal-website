// Package preprocessor handles pre-generation of content that is used by the web server
package preprocessor

import (
	"aHobeychi/personal-website/internal/config"
	"aHobeychi/personal-website/internal/util/logger"
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// Blog represents the basic structure needed for table of contents generation
type Blog struct {
	Id    string
	Title string
}

// BlogProvider is an interface for retrieving blog information
type BlogProvider interface {
	GetAllBlogs() ([]Blog, error)
	GetBlogContent(blogId string) (string, error)
}

// GenerateTableOfContents parses the HTML content and extracts headers to create a table of contents
func GenerateTableOfContents(htmlContent string) (string, error) {
	// Regular expressions to find header tags and their content
	headerRegex := regexp.MustCompile(`<h([1-6])[^>]*>(.*?)</h[1-6]>`)
	idRegex := regexp.MustCompile(`id=["']([^"']*)["']`)

	// Find all headers in the HTML content
	headers := headerRegex.FindAllStringSubmatch(htmlContent, -1)

	var buffer bytes.Buffer
	buffer.WriteString("<ul class=\"toc-list\">")

	// Track the current heading level to create proper nesting
	currentLevel := 0

	// Process each heading
	for i, header := range headers {
		if len(header) < 3 {
			continue // Skip if we don't have proper matches
		}

		// Get level and content from regex match
		level := 0
		fmt.Sscanf(header[1], "%d", &level) // Convert the level string to int
		headingText := extractTextFromHTML(header[2])

		// Look for an ID in the header tag
		idMatch := idRegex.FindStringSubmatch(header[0])
		var id string
		if len(idMatch) >= 2 {
			id = idMatch[1]
		} else {
			// Generate an ID based on the text if none exists
			id = strings.ToLower(strings.ReplaceAll(headingText, " ", "-"))
			// Ensure the ID is URL-friendly
			id = cleanIDString(id)
		}

		// Adjust nesting based on heading level
		if i == 0 {
			// First heading
		} else if level > currentLevel {
			// Going deeper in nesting
			for j := 0; j < level-currentLevel; j++ {
				buffer.WriteString("<ul>")
			}
		} else if level < currentLevel {
			// Coming out of nesting
			for j := 0; j < currentLevel-level; j++ {
				buffer.WriteString("</ul></li>")
			}
		} else {
			// Same level, close previous item
			buffer.WriteString("</li>")
		}

		// Write the list item
		buffer.WriteString(fmt.Sprintf("<li><a href=\"#%s\" class=\"sidebar-close\">%s</a>", id, headingText))

		currentLevel = level
	}

	// Close any open lists
	for j := 0; j < currentLevel; j++ {
		buffer.WriteString("</li></ul>")
	}

	return buffer.String(), nil
}

// GetBlogTableOfContentsPath returns the path to the table of contents file for a blog post
func GetBlogTableOfContentsPath(blogId string) string {
	return config.Get().Paths.TocHTML + "/" + blogId + "-toc.html"
}

// GenerateAndSaveTableOfContents generates the table of contents for a blog post and saves it to a file
func GenerateAndSaveTableOfContents(blog Blog, content string) error {
	// Generate the table of contents
	toc, err := GenerateTableOfContents(content)
	if err != nil {
		return err
	}

	// Create the table of contents HTML wrapper
	tocHTML := fmt.Sprintf(`<div class="blog-toc"><h2>Table of Contents</h2>%s</div>`, toc)

	// Ensure the directory exists
	err = os.MkdirAll(config.Get().Paths.TocHTML, 0755)
	if err != nil {
		return err
	}

	// Write the table of contents to file
	tocPath := GetBlogTableOfContentsPath(blog.Id)
	err = os.WriteFile(tocPath, []byte(tocHTML), 0644)
	if err != nil {
		logger.ErrorLogger.Printf("Error writing table of contents file for blog ID %s: %v", blog.Id, err)
		return err
	}

	logger.DebugLogger.Printf("Generated and saved table of contents for blog ID %s", blog.Id)
	return nil
}

// GenerateAllTableOfContents generates table of contents files for all blog posts
func GenerateAllTableOfContents(provider BlogProvider) error {
	blogs, err := provider.GetAllBlogs()
	if err != nil {
		return err
	}

	for _, blog := range blogs {
		content, err := provider.GetBlogContent(blog.Id)
		if err != nil {
			logger.ErrorLogger.Printf("Error getting content for blog ID %s: %v", blog.Id, err)
			// Continue with other blogs even if one fails
			continue
		}

		err = GenerateAndSaveTableOfContents(blog, content)
		if err != nil {
			logger.ErrorLogger.Printf("Error generating table of contents for blog ID %s: %v", blog.Id, err)
			// Continue with other blogs even if one fails
			continue
		}
	}

	logger.DebugLogger.Println("Generated table of contents for all blogs")
	return nil
}

// GetBlogTableOfContents returns the pre-generated table of contents HTML for a blog post
func GetBlogTableOfContents(blogId string, provider BlogProvider) (string, error) {
	tocPath := GetBlogTableOfContentsPath(blogId)

	// Check if the file exists
	if _, err := os.Stat(tocPath); os.IsNotExist(err) {
		// If the ToC file doesn't exist, generate it
		logger.DebugLogger.Printf("Table of contents file for blog ID %s does not exist, generating it", blogId)

		// Get the blog
		blogs, err := provider.GetAllBlogs()
		if err != nil {
			return "", err
		}

		var targetBlog Blog
		found := false
		for _, blog := range blogs {
			if blog.Id == blogId {
				targetBlog = blog
				found = true
				break
			}
		}

		if !found {
			return "", fmt.Errorf("blog with ID %s not found", blogId)
		}

		// Get the content
		content, err := provider.GetBlogContent(blogId)
		if err != nil {
			return "", err
		}

		// Generate and save the table of contents
		if err := GenerateAndSaveTableOfContents(targetBlog, content); err != nil {
			return "", err
		}
	}

	content, err := os.ReadFile(tocPath)
	if err != nil {
		logger.ErrorLogger.Printf("Error reading table of contents file for blog ID %s: %v", blogId, err)
		return "", err
	}

	return string(content), nil
}

// extractTextFromHTML removes HTML tags from a string
func extractTextFromHTML(html string) string {
	// Simple regex to remove HTML tags
	tagRegex := regexp.MustCompile(`<[^>]*>`)
	text := tagRegex.ReplaceAllString(html, "")
	return strings.TrimSpace(text)
}

// cleanIDString makes a string safe for use as an HTML ID
func cleanIDString(s string) string {
	// Keep only alphanumeric characters and hyphens
	return regexp.MustCompile(`[^a-z0-9-]`).ReplaceAllString(s, "")
}
