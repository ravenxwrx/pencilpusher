package logger_test

import (
	"log/slog"
	"testing"

	"github.com/ravenxwrx/pencilpusher/pkg/logger"
	"github.com/stretchr/testify/require"
)

func TestInitLogger(t *testing.T) {
	tests := map[string]struct {
		level       string
		format      string
		expectError bool
	}{
		logger.LogLevelInfo: {
			level:  logger.LogLevelInfo,
			format: logger.LogTypeText,
		},
		logger.LogLevelDebug: {
			level:  logger.LogLevelDebug,
			format: logger.LogTypeJSON,
		},
		logger.LogLevelWarn: {
			level:  logger.LogLevelWarn,
			format: logger.LogTypeText,
		},
		logger.LogLevelError: {
			level:  logger.LogLevelError,
			format: logger.LogTypeJSON,
		},
		"invalid level": {
			level:       "invalid",
			format:      logger.LogTypeText,
			expectError: true,
		},
		"invalid format": {
			level:       logger.LogLevelInfo,
			format:      "invalid",
			expectError: true,
		},
	}

	d := slog.Default()

	for name, tt := range tests {
		tf := func(t *testing.T) {
			slog.SetDefault(d)
			logger.SetLogLevel(tt.level)
			logger.SetLogFormat(tt.format)

			err := logger.InitLogger()
			if tt.expectError {
				require.Error(t, err)

				return
			}

			require.NoError(t, err)
		}

		t.Run(name, tf)
	}
}
