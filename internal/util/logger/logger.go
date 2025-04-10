package logger

import (
	"log"
	"strings"
)

var (
	WarningLogger *log.Logger
	ErrorLogger   *log.Logger
	DebugLogger   *log.Logger
	logLevel      string
)

func init() {
	// Initialize loggers with appropriate prefixes
	WarningLogger = log.New(log.Writer(), "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(log.Writer(), "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	DebugLogger = log.New(log.Writer(), "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// LogWarning logs a warning message
func LogWarning(message string) {
	WarningLogger.Println(message)
}

// LogError logs an error message
func LogError(message string) {
	ErrorLogger.Println(message)
}

// LogDebug logs a debug message
func LogDebug(message string) {
	if !strings.EqualFold(logLevel, "DEBUG") {
		return
	}
	DebugLogger.Println(message)
}

func SetLogLevel(level string) {
	logLevel = level
}
