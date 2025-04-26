package http

import (
	"context"
	"log/slog"
	"net/http"
	"time"
)

type Server struct {
	*http.Server

	closed chan struct{}
}

func (s *Server) Start() error {
	slog.Info("Starting HTTP service", "address", s.Addr)

	if err := s.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			slog.Debug("Server closed")

			return nil
		}

		return err
	}

	return nil
}

func New() *Server {
	s := &http.Server{
		Addr:         BindAddr(),
		ReadTimeout:  ReadTimeout(),
		WriteTimeout: WriteTimeout(),
		Handler:      mux(),
	}

	return &Server{
		Server: s,

		closed: make(chan struct{}),
	}
}

func (s *Server) Shutdown(ctx context.Context) error {
	timeout := 5 * time.Second

	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	err := s.Server.Shutdown(ctx)
	close(s.closed)

	return err
}

func (s *Server) Closed() <-chan struct{} {
	return s.closed
}

func mux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("Hello, World!"))
	})

	return mux
}
