package parser

import (
	"aHobeychi/personal-website/internal/cache"
	"aHobeychi/personal-website/internal/config"
	models "aHobeychi/personal-website/internal/domain"
	"time"
)

var certificationCache *cache.Cache[models.Certification]

func init() {
	// Initialize the certification cache
	certificationCache = cache.NewCache[models.Certification](
		config.Get().Paths.CertificationsJSON,
		time.Duration(config.Get().Features.CacheTTL*int(time.Minute)),
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
