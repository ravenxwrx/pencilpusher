package config

import (
	"log/slog"
	"os"

	yaml "sigs.k8s.io/yaml/goyaml.v3"
)

var cfg *Config

var defaultConfig = &Config{
	Logging: Logging{
		Level:  LogLevelInfo,
		Format: LogTypeText,
	},
}

func Load(path string) error {
	fp, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		return err
	}
	defer fp.Close()

	decoder := yaml.NewDecoder(fp)
	if err := decoder.Decode(&cfg); err != nil {
		return err
	}

	return nil
}

func Get() *Config {
	if cfg == nil {
		slog.Warn("Config is nil, using default config")
		return defaultConfig
	}

	return cfg
}
