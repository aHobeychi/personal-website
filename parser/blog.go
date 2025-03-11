package parser

import (
	"aHobeychi/personal-website/logger"
	"aHobeychi/personal-website/models"
	"encoding/json"
	"os"
	"sync"
	"time"
)

var (
	PATH_TOWARDS_BLOG_JSON = "static/content-catalog/blogs.json"
	PATH_TOWARDS_BLOG_HTML = "static/blog-posts/html/"
	blogCache              []models.Blog
	blogCacheOnce          sync.Once
	blogCacheErr           error
	blogCacheMutex         sync.Mutex
	blogCacheTicker        *time.Ticker
	blogDisableCache       bool
	BLOG_CACHE_TTL         = 10 * time.Minute
)

// init initializes the blog cache refresh mechanism
func init() {
	blogCacheTicker = time.NewTicker(BLOG_CACHE_TTL)
	go func() {
		for range blogCacheTicker.C {
			logger.DebugLogger.Println("Clearing blog cache")
			clearBlogCache()
		}
	}()
}

// clearBlogCache removes all cached blogs and resets the cache state
func clearBlogCache() {
	blogCacheMutex.Lock()
	defer blogCacheMutex.Unlock()
	blogCache = nil
	blogCacheOnce = sync.Once{}
	blogCacheErr = nil
}

// SetDisableBlogCache allows toggling the blog caching mechanism on or off
func SetDisableBlogCache(flag bool) {
	blogDisableCache = flag
}

// ParseBlogs retrieves a list of blogs, either from cache or from file
func ParseBlogs(limit ...int) ([]models.Blog, error) {
	if blogDisableCache {
		logger.DebugLogger.Println("Blog cache disabled, reading blogs from file")
		return parseBlogsFromFile(limit...)
	}
	blogCacheOnce.Do(func() {
		blogCacheMutex.Lock()
		defer blogCacheMutex.Unlock()
		logger.DebugLogger.Println("Blog cache enabled, reading blogs from file")
		file, err := os.Open(PATH_TOWARDS_BLOG_JSON)
		if err != nil {
			blogCacheErr = err
			return
		}
		defer file.Close()
		err = json.NewDecoder(file).Decode(&blogCache)
		if err != nil {
			blogCacheErr = err
		}
	})
	logger.DebugLogger.Println("Blog cache populated, returning blogs")
	if blogCacheErr != nil {
		return nil, blogCacheErr
	}
	if len(limit) > 0 && limit[0] < len(blogCache) {
		return blogCache[:limit[0]], nil
	}
	return blogCache, nil
}

// parseBlogsFromFile reads the blogs JSON file and decodes it into Blog models
func parseBlogsFromFile(limit ...int) ([]models.Blog, error) {
	file, err := os.Open(PATH_TOWARDS_BLOG_JSON)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var blogs []models.Blog
	err = json.NewDecoder(file).Decode(&blogs)
	if err != nil {
		return nil, err
	}
	if len(limit) > 0 && limit[0] < len(blogs) {
		return blogs[:limit[0]], nil
	}
	return blogs, nil
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
