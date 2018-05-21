package http

import (
	"context"
	"net/http"
	"time"
	"fmt"
)

type Server struct {
	server *http.Server
}

func New(port string, handler http.Handler) *Server {
	return &Server{
		server: &http.Server{
			Addr:         ":" + port,
			Handler:      handler,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 55 * time.Second,
		},
	}
}

func (s *Server) ListenAndServe() {
	go func() {
		fmt.Println("server up!")
		if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("erro: %s\n", err)
		}
	}()
}

func (s *Server) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	s.server.Shutdown(ctx)
	fmt.Println("server down!")
}
