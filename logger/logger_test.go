package logger_test

import (
	"fmt"
	"testing"

	"github.com/netsak/go-libstd/logger"
)

func TestLogLevel(t *testing.T) {
	testCases := []struct {
		level string
		err   error
	}{
		{"debug", nil},
		{"info", nil},
		{"warn", nil},
		{"error", nil},
		{"krosch", fmt.Errorf("unknown log level 'krosch'")},
	}
	for _, tc := range testCases {
		t.Run(tc.level, func(t *testing.T) {
			conf := logger.Config{
				Level:  tc.level,
				Format: "text",
			}
			_, err := logger.New(conf)
			if fmt.Sprintf("%v", err) != fmt.Sprintf("%v", tc.err) {
				t.Errorf("want: %q, got: %q", tc.err, err)
			}
		})
	}
}

func TestLogFormat(t *testing.T) {
	testCases := []struct {
		format   string
		expected string
	}{
		{"text", "*slog.TextHandler"},
		{"json", "*slog.JSONHandler"},
		{"krosch", "*slog.TextHandler"}, // fallback to text
	}
	for _, tc := range testCases {
		t.Run(tc.format, func(t *testing.T) {
			conf := logger.Config{
				Level:  "debug",
				Format: tc.format,
			}
			l, err := logger.New(conf)
			if err != nil {
				t.Fatalf("expected no error, but got: %v", err)
			}
			got := fmt.Sprintf("%T", l.Handler())
			if tc.expected != got {
				t.Errorf("want: %q, got: %q", tc.expected, got)
			}
		})
	}
}

func TestMustNewFromEnv(t *testing.T) {
	t.Setenv("LOG_LEVEL", "error")
	t.Setenv("LOG_FORMAT", "json")
	logger.MustNewFromEnv()
}

func TestMustNewFromEnvPanic(t *testing.T) {
	defer func() { _ = recover() }()
	logger.MustNewFromEnv()
	t.Error("should have caused a panic")
}
