package main

import (
	"flag"
	"log/slog"
	"os"

	"github.com/ravenxwrx/pencilpusher/pkg/config"
)

var (
	configPath string
)

func main() {
	flag.StringVar(&configPath, "config", "./config.yaml", "Path to the config file")
	flag.Parse()

	if err := config.Load(configPath); err != nil {
		slog.Error("Failed to load config", "error", err)
		os.Exit(1)
	}
}
