package parser

import (
	"aHobeychi/personal-website/internal/cache"
	"aHobeychi/personal-website/internal/config"
	models "aHobeychi/personal-website/internal/domain"
	"time"
)

var projectCache *cache.Cache[models.Project]

func init() {
	// Initialize the project cache
	projectCache = cache.NewCache[models.Project](
		config.Get().Paths.ProjectsJSON,
		time.Duration(config.Get().Features.CacheTTL*int(time.Minute)),
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
