package parser

import (
	"aHobeychi/personal-website/cache"
	"aHobeychi/personal-website/logger"
	"aHobeychi/personal-website/models"
	"os"
	"time"
)

const (
	PATH_TOWARDS_BLOG_JSON = "static/content-catalog/blogs.json"
	PATH_TOWARDS_BLOG_HTML = "static/blog-posts/html/"
	BLOG_CACHE_TTL         = 10 * time.Minute
)

// Global blogCache instance
var blogCache *cache.Cache[models.Blog]

func init() {
	// Initialize the blog cache
	blogCache = cache.NewCache[models.Blog](
		PATH_TOWARDS_BLOG_JSON,
		BLOG_CACHE_TTL,
		"blog",
	)
}

// SetDisableBlogCache allows toggling the blog caching mechanism on or off
func SetDisableBlogCache(flag bool) {
	blogCache.SetDisabled(flag)
}

// ParseBlogs retrieves a list of blogs, either from cache or from file
// Optional limit parameter controls the maximum number of blogs returned
// Returns a slice of Blog models and any error encountered
func ParseBlogs(limit ...int) ([]models.Blog, error) {
	return blogCache.Get(limit...)
}

// GetBlogHTMLContent returns the HTML content of a blog post by its ID.
func GetBlogHTMLContent(blogId string) (string, error) {
	blogPath := PATH_TOWARDS_BLOG_HTML + blogId + ".html"
	content, err := os.ReadFile(blogPath)
	if err != nil {
		logger.ErrorLogger.Println("Error reading blog content file:", err)
		return "", err
	}
	contentString := string(content)
	logger.DebugLogger.Printf("HTML content retrieved for blog ID: %s", blogId)
	return contentString, nil
}

func GetBlogByID(id string) (models.Blog, error) {
	blogs, err := ParseBlogs()
	if err != nil {
		return models.Blog{}, err
	}

	for _, blog := range blogs {
		if blog.Id == id {
			return blog, nil
		}
	}

	return models.Blog{}, os.ErrNotExist
}
