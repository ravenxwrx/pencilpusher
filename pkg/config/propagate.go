package config

import (
	"github.com/ravenxwrx/pencilpusher/pkg/http"
	"github.com/ravenxwrx/pencilpusher/pkg/logger"
)

func setConfigValues(c *Config) {
	cfg = c

	logger.SetLogLevel(c.Logging.Level)
	logger.SetLogFormat(c.Logging.Format)

	http.SetBindAddr(c.Http.Address)
}
