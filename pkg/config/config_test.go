package config

import (
	"bytes"
	"log/slog"
	"testing"

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
			err := Load(tt.path)
			if tt.expectErr {
				require.Error(t, err)

				return
			}

			require.NoError(t, err)
			require.Equal(t, LogLevelDebug, Get().Logging.Level)
			require.Equal(t, LogTypeJSON, Get().Logging.Format)
		}
		t.Run(name, tf)
	}
}

func TestDefaultConfig(t *testing.T) {
	output := bytes.NewBufferString("")
	slog.SetDefault(slog.New(slog.NewTextHandler(output, nil)))

	cfg := Get()
	require.Equal(t, LogLevelInfo, cfg.Logging.Level)
	require.Equal(t, LogTypeText, cfg.Logging.Format)

	require.Contains(t, output.String(), "Config is nil, using default config")
}
