package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// Config holds all configuration for the application
type Config struct {
	Server struct {
		Port        int    `json:"port"`
		Host        string `json:"host"`
		Domain      string `json:"domain"`
		Environment string `json:"environment"`
	} `json:"server"`
	Paths struct {
		Templates          string `json:"templates"`
		AssetFiles         string `json:"assetFiles"`
		BlogHTML           string `json:"blogHTML"`
		TocHTML            string `json:"tocHTML"`
		ProjectsJSON       string `json:"projectsJSON"`
		BlogsJSON          string `json:"blogsJSON"`
		WorkExperienceJSON string `json:"workExperienceJSON"`
		CertificationsJSON string `json:"certificationsJSON"`
	} `json:"paths"`
	Features struct {
		CacheEnabled bool `json:"cacheEnabled"`
		CacheTTL     int  `json:"cacheTTL"`
		DebugMode    bool `json:"debugMode"`
		DisplayBlogs bool `json:"displayBlogs"`
	} `json:"features"`
	Logging struct {
		Level string `json:"level"`
	}
}

// global config instance
var cfg *Config

// Load loads the configuration from files and environment variables
func Load() (*Config, error) {
	if cfg != nil {
		return cfg, nil
	}

	// Determine environment
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	// Initialize default config
	cfg = &Config{}

	// Load environment-specific config file
	configPath := fmt.Sprintf("config/%s.json", env)

	err := loadConfigFile(configPath, cfg)

	if err != nil && !os.IsNotExist(err) {
		return nil, fmt.Errorf("failed to load config file: %w", err)
	}

	// Normalize paths
	cfg.normalizePaths()

	return cfg, nil
}

// loadConfigFile loads config from a JSON file
func loadConfigFile(path string, c *Config) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, c)
	if err != nil {
		return fmt.Errorf("invalid JSON in config file: %w", err)
	}

	return nil
}

// normalizePaths converts relative paths to absolute paths
func (c *Config) normalizePaths() {
	projectRoot := getProjectRoot()

	c.Paths.Templates = makeAbsolute(c.Paths.Templates, projectRoot)
	c.Paths.AssetFiles = makeAbsolute(c.Paths.AssetFiles, projectRoot)
	c.Paths.BlogHTML = makeAbsolute(c.Paths.BlogHTML, projectRoot)
	c.Paths.TocHTML = makeAbsolute(c.Paths.TocHTML, projectRoot)
	c.Paths.BlogsJSON = makeAbsolute(c.Paths.BlogsJSON, projectRoot)
	c.Paths.WorkExperienceJSON = makeAbsolute(c.Paths.WorkExperienceJSON, projectRoot)
	c.Paths.CertificationsJSON = makeAbsolute(c.Paths.CertificationsJSON, projectRoot)
	c.Paths.ProjectsJSON = makeAbsolute(c.Paths.ProjectsJSON, projectRoot)
}

// makeAbsolute converts a path to absolute if it's not already
func makeAbsolute(path string, basePath string) string {
	if filepath.IsAbs(path) {
		return path
	}
	return filepath.Join(basePath, path)
}

// getProjectRoot attempts to determine the project root directory
func getProjectRoot() string {
	// Use working directory
	workDir, err := os.Getwd()
	if err != nil {
		// Fallback to executable directory if working dir fails
		if execPath, err := os.Executable(); err == nil {
			return filepath.Dir(execPath)
		}
		return "."
	}
	return workDir
}

// Get returns the global configuration
func Get() *Config {
	if cfg == nil {
		cfg, _ = Load()
	}
	return cfg
}
