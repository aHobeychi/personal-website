package parser

import (
	"aHobeychi/personal-website/cache"
	"aHobeychi/personal-website/models"
	"time"
)

const (
	PATH_TOWARDS_PROJECTS_JSON = "static/content-catalog/projects.json"
	PROJECT_CACHE_TTL          = 60 * time.Minute
)

// Global projectCache instance
var projectCache *cache.Cache[models.Project]

func init() {
	// Initialize the project cache
	projectCache = cache.NewCache[models.Project](
		PATH_TOWARDS_PROJECTS_JSON,
		PROJECT_CACHE_TTL,
		"project",
	)
}

// SetDisableCache allows toggling the caching mechanism on or off
func SetDisableCache(flag bool) {
	projectCache.SetDisabled(flag)
}

// ParseProjects retrieves a list of projects, either from cache or from file
// Optional limit parameter controls the maximum number of projects returned
// Returns a slice of Project models and any error encountered
func ParseProjects(limit ...int) ([]models.Project, error) {
	return projectCache.Get(limit...)
}
