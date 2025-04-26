package main

import (
	"flag"
	"log/slog"
	"os"
	"os/signal"

	"github.com/ravenxwrx/pencilpusher/pkg/config"
	"github.com/ravenxwrx/pencilpusher/pkg/http"
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

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	httpServer := http.New()

	go httpServer.Start()

	<-signalChan
	slog.Info("Received interrupt signal, shutting down")

	if err := httpServer.Shutdown(); err != nil {
		slog.Error("Failed to shut down server", "error", err)
		os.Exit(1)
	}
}
