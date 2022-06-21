package api

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	server *http.Server
}

func (sr Server) Start(address string, handler http.Handler) error {
	server := &http.Server{
		Addr: address,
		Handler: handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout: time.Second * 10,
		WriteTimeout: time.Second * 10,
	}

	return server.ListenAndServe()
}

func (sr Server) Stop(ctx context.Context) error {
	return sr.server.Shutdown(ctx)
}