package http

import (
	"log/slog"
	"net/http"
)

type Server struct {
	*http.Server
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

func (s *Server) Shutdown() error {
	return s.Close()
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
	}
}

func mux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("Hello, World!"))
	})

	return mux
}
