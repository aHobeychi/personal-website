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
	PATH_TOWARDS_CERTIFICATIONS = "static/content-catalog/certifications.json"
	certificationCache          []models.Certification
	certificationCacheOnce      sync.Once
	certificationCacheErr       error
	certificationCacheMutex     sync.Mutex
	certificationCacheTicker    *time.Ticker
	certificationDisableCache   bool
	CERTIFICATION_CACHE_TTL     = 10 * time.Minute
)

// init initializes the cache refresh mechanism
// Sets up a ticker that clears the certification cache at regular intervals
// defined by CERTIFICATION_CACHE_TTL
func init() {
	certificationCacheTicker = time.NewTicker(CERTIFICATION_CACHE_TTL)
	go func() {
		for range certificationCacheTicker.C {
			logger.DebugLogger.Println("Clearing certification cache")
			clearCertificationCache()
		}
	}()
}

// clearCertificationCache removes all cached certifications and resets the cache state
// This forces the cache to be repopulated on the next request
// Uses mutex locking to ensure thread safety
func clearCertificationCache() {
	certificationCacheMutex.Lock()
	defer certificationCacheMutex.Unlock()
	certificationCache = nil
	certificationCacheOnce = sync.Once{}
	certificationCacheErr = nil
}

// SetCertificationDisableCache allows toggling the caching mechanism on or off
// When true, certifications will always be read directly from file
// When false, certifications are cached and refreshed according to CERTIFICATION_CACHE_TTL
func SetCertificationDisableCache(flag bool) {
	certificationDisableCache = flag
}

// ParseCertifications retrieves a list of certifications, either from cache or from file
// Uses a sync.Once to ensure the cache is populated only once between cache clear operations
// Optional limit parameter controls the maximum number of certifications returned
// Returns a slice of Certification models and any error encountered
func ParseCertifications(limit ...int) ([]models.Certification, error) {
	if certificationDisableCache {
		logger.DebugLogger.Println("Cache disabled, reading certifications from file")
		return parseCertificationsFromFile(limit...)
	}

	certificationCacheOnce.Do(func() {
		certificationCacheMutex.Lock()
		defer certificationCacheMutex.Unlock()
		logger.DebugLogger.Println("Cache enabled, reading certifications from file")

		file, err := os.Open(PATH_TOWARDS_CERTIFICATIONS)
		if err != nil {
			certificationCacheErr = err
			return
		}
		defer file.Close()

		err = json.NewDecoder(file).Decode(&certificationCache)
		if err != nil {
			certificationCacheErr = err
		}
	})

	logger.DebugLogger.Println("Cache populated, returning certifications")

	if certificationCacheErr != nil {
		return nil, certificationCacheErr
	}

	// If a limit is provided, return only that number of certifications
	if len(limit) > 0 && limit[0] < len(certificationCache) {
		return certificationCache[:limit[0]], nil
	}

	return certificationCache, nil
}

// parseCertificationsFromFile parses the certification data from a JSON file.
// This is called directly when cache is disabled or indirectly to populate the cache
// Optional limit parameter controls the maximum number of certifications returned
// Returns a slice of Certification models and any error encountered
func parseCertificationsFromFile(limit ...int) ([]models.Certification, error) {
	file, err := os.Open(PATH_TOWARDS_CERTIFICATIONS)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var certifications []models.Certification
	err = json.NewDecoder(file).Decode(&certifications)
	if err != nil {
		return nil, err
	}

	if len(limit) > 0 && limit[0] < len(certifications) {
		return certifications[:limit[0]], nil
	}

	return certifications, nil
}
