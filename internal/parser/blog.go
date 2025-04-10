package parser

import (
	"aHobeychi/personal-website/internal/cache"
	"aHobeychi/personal-website/internal/config"
	models "aHobeychi/personal-website/internal/domain"
	"aHobeychi/personal-website/internal/util/logger"
	"os"
	"time"
)

const (
	BLOG_CACHE_TTL = 60 * time.Minute
)

var blogCache *cache.Cache[models.Blog]

func init() {
	// Initialize the blog cache
	blogCache = cache.NewCache[models.Blog](
		config.Get().Paths.BlogsJSON,
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
	blogPath := config.Get().Paths.BlogHTML + "/" + blogId + ".html"
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

// GetBlogTableOfContents returns the pre-generated table of contents HTML for a blog post
func GetBlogTableOfContents(blogId string) (string, error) {
	tocPath := config.Get().Paths.TocHTML + blogId + "-toc.html"
	content, err := os.ReadFile(tocPath)
	if err != nil {
		logger.ErrorLogger.Println("Error reading blog table of contents file:", err)
		return "", err
	}
	contentString := string(content)
	logger.DebugLogger.Printf("Table of contents retrieved for blog ID: %s", blogId)
	return contentString, nil
}
