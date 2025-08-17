package logger

import (
	"log/slog"
	"os"
	"strings"
)

func New() *slog.Logger {
	opts := &slog.HandlerOptions{
		Level:     readLogLevel(),
		AddSource: false,
	}
	h := newHandler(slog.NewTextHandler(os.Stdout, opts))
	return slog.New(h)
}

func readLogLevel() slog.Level {
	level := os.Getenv("LOG_LEVEL")
	switch strings.ToUpper(level) {
	case "INFO":
		return slog.LevelInfo
	case "WARN":
		return slog.LevelWarn
	case "ERROR":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
