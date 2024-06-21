package logger

import (
	"log/slog"
	"os"
)

var logger *slog.Logger

func init() {
	logLevel := os.Getenv("LOG_LEVEL")
	slogLevel := slog.LevelInfo
	if logLevel == "DEBUG" {
		slogLevel = slog.LevelDebug
	} else if logLevel == "INFO" {
		slogLevel = slog.LevelInfo
	} else if logLevel == "WARN" {
		slogLevel = slog.LevelWarn
	} else if logLevel == "ERROR" {
		slogLevel = slog.LevelError
	}
	logger = slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		Level: slogLevel,
	}))
	logger.Debug("Initializing logger finished.")
}

func GetLogger() *slog.Logger {
	return logger
}
