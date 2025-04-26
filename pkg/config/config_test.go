package config

import (
	"bytes"
	"log/slog"
	"testing"

	"github.com/ravenxwrx/pencilpusher/pkg/http"
	"github.com/ravenxwrx/pencilpusher/pkg/logger"
	"github.com/stretchr/testify/require"
)

func TestLoad(t *testing.T) {
	t.Cleanup(func() {
		cfg = nil
	})

	tests := map[string]struct {
		path      string
		expectErr bool
	}{
		"valid config": {
			path:      "testdata/config.yaml",
			expectErr: false,
		},
		"missing file": {
			path:      "testdata/config.bogus.yaml",
			expectErr: true,
		},
		"invalid file": {
			path:      "testdata/config.invalid.yaml",
			expectErr: true,
		},
	}

	for name, tt := range tests {
		tf := func(t *testing.T) {
			cfg = nil

			err := Load(tt.path)
			if tt.expectErr {
				require.Error(t, err)

				return
			}

			require.NoError(t, err)
			require.Equal(t, logger.LogLevelDebug, Get().Logging.Level)
			require.Equal(t, logger.LogTypeJSON, Get().Logging.Format)
		}
		t.Run(name, tf)
	}
}

func TestDefaultConfig(t *testing.T) {
	output := bytes.NewBufferString("")
	slog.SetDefault(slog.New(slog.NewTextHandler(output, nil)))

	cfg := Get()
	require.Equal(t, logger.LogLevelInfo, cfg.Logging.Level)
	require.Equal(t, logger.LogTypeText, cfg.Logging.Format)

	require.Equal(t, ":8080", cfg.Http.Address)

	require.Contains(t, output.String(), "Config is nil, using default config")
}

func TestPropagate(t *testing.T) {
	err := Load("testdata/config.yaml")
	require.NoError(t, err)

	require.Equal(t, logger.LogLevelDebug, logger.LogLevel())
	require.Equal(t, logger.LogTypeJSON, logger.LogFormat())

	require.Equal(t, ":9090", http.BindAddr())
}
