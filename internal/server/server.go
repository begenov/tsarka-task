package server

import (
	"context"
	"net/http"

	"github.com/begenov/tsarka-task/internal/config"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(cfg config.HTTPServerConfig, handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:           cfg.Host + ":" + cfg.Port,
			ReadTimeout:    cfg.ReadTimeout,
			WriteTimeout:   cfg.WriteTimeout,
			MaxHeaderBytes: cfg.MaxHeaderBytes << 20,
			Handler:        handler,
		},
	}
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
