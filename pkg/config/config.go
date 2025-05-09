package config

import (
	"log/slog"
	"os"

	"github.com/goccy/go-yaml"
	"github.com/ravenxwrx/pencilpusher/pkg/http"
	"github.com/ravenxwrx/pencilpusher/pkg/logger"
)

var cfg *Config

var defaultConfig = &Config{
	Logging: Logging{
		Level:  logger.LogLevelInfo,
		Format: logger.LogTypeText,
	},
	Http: Http{
		Address: http.BindAddr(),
	},
}

func Load(path string) error {
	fp, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		return err
	}
	defer fp.Close()

	decoder := yaml.NewDecoder(fp)

	var c Config

	if err := decoder.Decode(&c); err != nil {
		return err
	}

	setConfigValues(&c)

	return nil
}

func Get() *Config {
	if cfg == nil {
		slog.Warn("Config is nil, using default config")

		return defaultConfig
	}

	return cfg
}
