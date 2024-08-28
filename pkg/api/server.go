package api

import (
	"context"
	"log"
	"net/http"
	"time"
)

type Server struct {
	Inner            *http.Server
	ShutdownDeadline int
}

func NewServer(address string, router *Router) *Server {
	return &Server{
		Inner: &http.Server{
			Addr:         address,
			Handler:      router.Inner,
			WriteTimeout: time.Second * 15,
			ReadTimeout:  time.Second * 15,
			IdleTimeout:  time.Second * 60,
		},
		ShutdownDeadline: 3,
	}
}

func (server *Server) Run() {
	go func() {
		if err := server.Inner.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()
}

func (server *Server) Shutdown(ctx context.Context) error {
	return server.Inner.Shutdown(ctx)
}
