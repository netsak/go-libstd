package logger

import (
	"fmt"
	"log/slog"
	"os"
)

// Config for the logger
type Config struct {
	// Level must be one of [debug, info, warn, error]
	Level string `env:"LOG_LEVEL"`
	// Format must be one of [text, json]
	Format string `env:"LOG_FORMAT"`
}

// New create a new logger based on the config and sets it as default.
func New(conf Config) (*slog.Logger, error) {
	var level slog.Level
	if err := level.UnmarshalText([]byte(conf.Level)); err != nil {
		return nil, fmt.Errorf("unknown log level '%s'", conf.Level)
	}
	options := slog.HandlerOptions{
		Level: level,
	}
	var handler slog.Handler
	if conf.Format == "json" {
		handler = slog.NewJSONHandler(os.Stderr, &options)
	} else {
		handler = slog.NewTextHandler(os.Stderr, &options)
	}
	logger := slog.New(handler)
	slog.SetDefault(logger)
	return logger, nil
}

// MustNewFromEnv creates a new logger configured with environment variables.
// Panics when an error occurs.
func MustNewFromEnv() *slog.Logger {
	conf := Config{
		Level:  os.Getenv("LOG_LEVEL"),
		Format: os.Getenv("LOG_FORMAT"),
	}
	logger, err := New(conf)
	if err != nil {
		panic(err)
	}
	return logger
}
