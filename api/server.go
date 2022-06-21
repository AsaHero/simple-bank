package api

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/AsaHero/simple-bank/api/handler"
)

type Server struct {
	server *http.Server
}

func (sr Server) Start(address string, handler http.Handler) error {
	server := &http.Server{
		Addr:           address,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    time.Second * 10,
		WriteTimeout:   time.Second * 10,
	}

	return server.ListenAndServe()
}

func (sr Server) Stop(ctx context.Context) error {
	return sr.server.Shutdown(ctx)
}

func InitRouter(h *handler.Handler) *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/signin", h.SignIn)
		auth.POST("/signup", h.SignUp)
	}

	transaction := router.Group("/")
	{
		transaction.POST("/make-transfers")
		transaction.GET("/get-all-transactions")
		transaction.GET("/get-transcation-by-id/:id")
	}

	return router
}
