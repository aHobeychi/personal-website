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

// ParseProjects retrieves a list of projects, either from cache or from file
// Uses a sync.Once to ensure the cache is populated only once between cache clear operations
// Optional limit parameter controls the maximum number of projects returned
// Returns a slice of Project models and any error encountered
func ParseBlogs(limit ...int) ([]models.Blog, error) {

	var cache []models.Blog
	file, err := os.Open(PATH_TOWARDS_BLOG_JSON)
	if err != nil {
		cacheErr = err
		return nil, err
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&cache)
	if err != nil {
		cacheErr = err
	}

	logger.DebugLogger.Println("Cache populated, returning projects")

	// print first object as json
	jsonData, err := json.Marshal(cache[0])
	if err != nil {
		logger.ErrorLogger.Println("Error marshalling JSON:", err)
		return nil, err
	}
	logger.DebugLogger.Println(string(jsonData))

	return cache, nil
}
