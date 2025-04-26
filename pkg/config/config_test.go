package config_test

import (
	"bytes"
	"log/slog"
	"testing"

	"github.com/ravenxwrx/pencilpusher/pkg/config"
	"github.com/stretchr/testify/require"
)

func TestLoad(t *testing.T) {
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
			path:      "testdata/config.invalid.json",
			expectErr: true,
		},
	}

	for name, tt := range tests {
		tf := func(t *testing.T) {
			err := config.Load(tt.path)
			if tt.expectErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.Equal(t, config.LogLevelDebug, config.Get().Logging.Level)
			require.Equal(t, config.LogTypeJSON, config.Get().Logging.Format)
		}
		t.Run(name, tf)
	}
}

func TestDefaultConfig(t *testing.T) {
	output := bytes.NewBufferString("")
	slog.SetDefault(slog.New(slog.NewTextHandler(output, nil)))

	cfg := config.Get()
	require.Equal(t, config.LogLevelInfo, cfg.Logging.Level)
	require.Equal(t, config.LogTypeText, cfg.Logging.Format)

	require.Contains(t, output.String(), "Config is nil, using default config")
}
