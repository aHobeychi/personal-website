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
	PATH_TOWARDS_PROJECTS_JSON = "static/projects/projects.json"
	cache                      []models.Project
	cacheOnce                  sync.Once
	cacheErr                   error
	cacheMutex                 sync.Mutex
	cacheTicker                *time.Ticker
	disableCache               bool
	CACHE_TTL                  = 1 * time.Minute
)

// init initializes the cache refresh mechanism
// Sets up a ticker that clears the project cache at regular intervals
// defined by CACHE_TTL
func init() {
	cacheTicker = time.NewTicker(CACHE_TTL)
	go func() {
		for range cacheTicker.C {
			logger.DebugLogger.Println("Clearing project cache")
			clearCache()
		}
	}()
}

// clearCache removes all cached projects and resets the cache state
// This forces the cache to be repopulated on the next request
// Uses mutex locking to ensure thread safety
func clearCache() {
	cacheMutex.Lock()
	defer cacheMutex.Unlock()
	cache = nil
	cacheOnce = sync.Once{}
	cacheErr = nil
}

// SetDisableCache allows toggling the caching mechanism on or off
// When true, projects will always be read directly from file
// When false, projects are cached and refreshed according to CACHE_TTL
func SetDisableCache(flag bool) {
	disableCache = flag
}

// ParseProjects retrieves a list of projects, either from cache or from file
// Uses a sync.Once to ensure the cache is populated only once between cache clear operations
// Optional limit parameter controls the maximum number of projects returned
// Returns a slice of Project models and any error encountered
func ParseProjects(limit ...int) ([]models.Project, error) {

	if disableCache {
		logger.DebugLogger.Println("Cache disabled, reading projects from file")
		return parseProjectsFromFile(limit...)
	}

	cacheOnce.Do(func() {
		cacheMutex.Lock()
		defer cacheMutex.Unlock()
		logger.DebugLogger.Println("Cache enabled, reading projects from file")

		file, err := os.Open(PATH_TOWARDS_PROJECTS_JSON)
		if err != nil {
			cacheErr = err
			return
		}
		defer file.Close()

		err = json.NewDecoder(file).Decode(&cache)
		if err != nil {
			cacheErr = err
		}
	})

	logger.DebugLogger.Println("Cache populated, returning projects")

	if cacheErr != nil {
		return nil, cacheErr
	}

	// If a limit is provided, return only that number of projects
	if len(limit) > 0 && limit[0] < len(cache) {
		return cache[:limit[0]], nil
	}

	return cache, nil
}

// parseProjectsFromFile reads the projects JSON file and decodes it into Project models
// This is called directly when cache is disabled or indirectly to populate the cache
// Optional limit parameter controls the maximum number of projects returned
// Returns a slice of Project models and any error encountered
func parseProjectsFromFile(limit ...int) ([]models.Project, error) {
	file, err := os.Open(PATH_TOWARDS_PROJECTS_JSON)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var projects []models.Project
	err = json.NewDecoder(file).Decode(&projects)
	if err != nil {
		return nil, err
	}

	// If a limit is provided, return only that number of projects
	if len(limit) > 0 && limit[0] < len(projects) {
		return projects[:limit[0]], nil
	}

	return projects, nil
}
