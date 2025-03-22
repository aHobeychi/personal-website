package parser

import (
	"aHobeychi/personal-website/cache"
	"aHobeychi/personal-website/models"
	"time"
)

const (
	PATH_TOWARDS_CERTIFICATIONS = "static/content-catalog/certifications.json"
	CERTIFICATION_CACHE_TTL     = 10 * time.Minute
)

// Global certificationCache instance
var certificationCache *cache.Cache[models.Certification]

func init() {
	// Initialize the certification cache
	certificationCache = cache.NewCache[models.Certification](
		PATH_TOWARDS_CERTIFICATIONS,
		CERTIFICATION_CACHE_TTL,
		"certification",
	)
}

// SetCertificationDisableCache allows toggling the caching mechanism on or off
func SetCertificationDisableCache(flag bool) {
	certificationCache.SetDisabled(flag)
}

// ParseCertifications retrieves a list of certifications, either from cache or from file
// Optional limit parameter controls the maximum number of certifications returned
// Returns a slice of Certification models and any error encountered
func ParseCertifications(limit ...int) ([]models.Certification, error) {
	return certificationCache.Get(limit...)
}
