package logger

import (
	"fmt"
	"log/slog"
	"os"
)

func InitLogger() error {
	var level slog.Level
	switch LogLevel() {
	case LogLevelDebug:
		level = slog.LevelDebug
	case LogLevelInfo:
		level = slog.LevelInfo
	case LogLevelWarn:
		level = slog.LevelWarn
	case LogLevelError:
		level = slog.LevelError
	default:
		return fmt.Errorf("invalid log level: %s", logLevel)
	}

	handlerOpts := &slog.HandlerOptions{
		Level: level,
	}

	var handler slog.Handler

	switch LogFormat() {
	case LogTypeText:
		handler = slog.NewTextHandler(os.Stdout, handlerOpts)
	case LogTypeJSON:
		handler = slog.NewJSONHandler(os.Stdout, handlerOpts)
	default:
		return fmt.Errorf("invalid log format: %s", logFormat)
	}

	slog.SetDefault(slog.New(handler))

	return nil
}
