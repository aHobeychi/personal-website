package parser

import (
	"aHobeychi/personal-website/internal/cache"
	"aHobeychi/personal-website/internal/config"
	models "aHobeychi/personal-website/internal/domain"
	"time"
)

const (
	WORK_EXPERIENCE_CACHE_TTL = 60 * time.Minute
)

var workExperienceCache *cache.Cache[models.WorkExperience]

func init() {
	// Initialize the work experience cache
	workExperienceCache = cache.NewCache[models.WorkExperience](
		config.Get().Paths.WorkExperienceJSON,
		WORK_EXPERIENCE_CACHE_TTL,
		"work experience",
	)
}

// SetWorkExperienceDisableCache allows toggling the caching mechanism on or off
func SetWorkExperienceDisableCache(flag bool) {
	workExperienceCache.SetDisabled(flag)
}

// ParseWorkExperiences retrieves a list of work experiences, either from cache or from file
// Optional limit parameter controls the maximum number of work experiences returned
// Returns a slice of WorkExperience models and any error encountered
func ParseWorkExperiences(limit ...int) ([]models.WorkExperience, error) {
	return workExperienceCache.Get(limit...)
}
