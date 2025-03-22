// Package cache provides a generic caching mechanism for JSON data
package cache

import (
	"aHobeychi/personal-website/logger"
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"
)

// Cache provides a generic caching mechanism for any type of data
type Cache[T any] struct {
	path        string
	data        []T
	once        sync.Once
	err         error
	mutex       sync.Mutex
	ticker      *time.Ticker
	disableFlag bool
	ttl         time.Duration
	name        string
}

// NewCache creates a new cache with the specified parameters
func NewCache[T any](path string, ttl time.Duration, name string) *Cache[T] {
	c := &Cache[T]{
		path:        path,
		ttl:         ttl,
		name:        name,
		disableFlag: false,
	}

	// Initialize the ticker for cache invalidation
	c.ticker = time.NewTicker(ttl)
	go func() {
		for range c.ticker.C {
			logger.DebugLogger.Printf("Clearing %s cache", c.name)
			c.Clear()
		}
	}()

	return c
}

// Clear removes all cached data and resets the cache state
func (c *Cache[T]) Clear() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.data = nil
	c.once = sync.Once{}
	c.err = nil
}

// SetDisabled allows toggling the caching mechanism on or off
func (c *Cache[T]) SetDisabled(flag bool) {
	c.disableFlag = flag
}

// IsDisabled returns whether caching is disabled
func (c *Cache[T]) IsDisabled() bool {
	return c.disableFlag
}

// Get retrieves data either from cache or directly from file
// Uses a loader function to load data from file when cache is disabled or needs population
func (c *Cache[T]) Get(limit ...int) ([]T, error) {
	// If caching is disabled, read directly from file
	if c.disableFlag {
		logger.DebugLogger.Printf("%s cache disabled, reading from file", c.name)
		return c.loadFromFile(limit...)
	}

	// Initialize cache if not already done
	c.once.Do(func() {
		c.mutex.Lock()
		defer c.mutex.Unlock()
		logger.DebugLogger.Printf("%s cache enabled, reading from file", c.name)

		file, err := os.Open(c.path)
		if err != nil {
			c.err = err
			return
		}
		defer file.Close()

		err = json.NewDecoder(file).Decode(&c.data)
		if err != nil {
			c.err = err
		}
	})

	logger.DebugLogger.Printf("%s cache populated, returning data", c.name)

	if c.err != nil {
		return nil, c.err
	}

	// If a limit is provided, return only that number of items
	if len(limit) > 0 && limit[0] < len(c.data) {
		return c.data[:limit[0]], nil
	}

	return c.data, nil
}

// loadFromFile reads the JSON file and decodes it into the specified type
func (c *Cache[T]) loadFromFile(limit ...int) ([]T, error) {
	file, err := os.Open(c.path)
	if err != nil {
		return nil, fmt.Errorf("failed to open %s: %w", c.name, err)
	}
	defer file.Close()

	var data []T
	err = json.NewDecoder(file).Decode(&data)
	if err != nil {
		return nil, fmt.Errorf("failed to decode %s JSON: %w", c.name, err)
	}

	// If a limit is provided, return only that number of items
	if len(limit) > 0 && limit[0] < len(data) {
		return data[:limit[0]], nil
	}

	return data, nil
}
