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
	PATH_TOWARDS_WORK_EXPERIENCE = "static/content-catalog/work-experience.json"
	workExperienceCache          []models.WorkExperience
	workExperienceCacheOnce      sync.Once
	workExperienceCacheErr       error
	workExperienceCacheMutex     sync.Mutex
	workExperienceCacheTicker    *time.Ticker
	workExperienceDisableCache   bool
	WORK_EXPERIENCE_CACHE_TTL    = 10 * time.Minute
)

// init initializes the cache refresh mechanism
// Sets up a ticker that clears the work experience cache at regular intervals
// defined by WORK_EXPERIENCE_CACHE_TTL
func init() {
	workExperienceCacheTicker = time.NewTicker(WORK_EXPERIENCE_CACHE_TTL)
	go func() {
		for range workExperienceCacheTicker.C {
			logger.DebugLogger.Println("Clearing work experience cache")
			clearWorkExperienceCache()
		}
	}()
}

// clearWorkExperienceCache removes all cached work experiences and resets the cache state
// This forces the cache to be repopulated on the next request
// Uses mutex locking to ensure thread safety
func clearWorkExperienceCache() {
	workExperienceCacheMutex.Lock()
	defer workExperienceCacheMutex.Unlock()
	workExperienceCache = nil
	workExperienceCacheOnce = sync.Once{}
	workExperienceCacheErr = nil
}

// SetWorkExperienceDisableCache allows toggling the caching mechanism on or off
// When true, work experiences will always be read directly from file
// When false, work experiences are cached and refreshed according to WORK_EXPERIENCE_CACHE_TTL
func SetWorkExperienceDisableCache(flag bool) {
	workExperienceDisableCache = flag
}

// ParseWorkExperiences retrieves a list of work experiences, either from cache or from file
// Uses a sync.Once to ensure the cache is populated only once between cache clear operations
// Optional limit parameter controls the maximum number of work experiences returned
// Returns a slice of WorkExperience models and any error encountered
func ParseWorkExperiences(limit ...int) ([]models.WorkExperience, error) {
	if workExperienceDisableCache {
		logger.DebugLogger.Println("Cache disabled, reading work experiences from file")
		return parseWorkExperiencesFromFile(limit...)
	}

	workExperienceCacheOnce.Do(func() {
		workExperienceCacheMutex.Lock()
		defer workExperienceCacheMutex.Unlock()
		logger.DebugLogger.Println("Cache enabled, reading work experiences from file")

		file, err := os.Open(PATH_TOWARDS_WORK_EXPERIENCE)
		if err != nil {
			workExperienceCacheErr = err
			return
		}
		defer file.Close()

		err = json.NewDecoder(file).Decode(&workExperienceCache)
		if err != nil {
			workExperienceCacheErr = err
		}
	})

	logger.DebugLogger.Println("Cache populated, returning work experiences")

	if workExperienceCacheErr != nil {
		return nil, workExperienceCacheErr
	}

	// If a limit is provided, return only that number of work experiences
	if len(limit) > 0 && limit[0] < len(workExperienceCache) {
		return workExperienceCache[:limit[0]], nil
	}

	return workExperienceCache, nil
}

// parseWorkExperiencesFromFile parses the work experience data from a JSON file.
// This is called directly when cache is disabled or indirectly to populate the cache
// Optional limit parameter controls the maximum number of work experiences returned
// Returns a slice of WorkExperience models and any error encountered
func parseWorkExperiencesFromFile(limit ...int) ([]models.WorkExperience, error) {
	file, err := os.Open(PATH_TOWARDS_WORK_EXPERIENCE)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var workExperience []models.WorkExperience
	err = json.NewDecoder(file).Decode(&workExperience)
	if err != nil {
		return nil, err
	}

	if len(limit) > 0 && limit[0] < len(workExperience) {
		return workExperience[:limit[0]], nil
	}

	return workExperience, nil
}
