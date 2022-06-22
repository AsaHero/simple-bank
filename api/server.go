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

	transaction := router.Group("/v1")
	{	
		transaction.POST("/transfer", h.MakeTransfer)
		transaction.GET("/all-transfers", h.GetAllTransfers)
		transaction.GET("/transfer/:id", h.GetTransferById)
		transaction.GET("/all-entries", h.GetAllEntries)
		transaction.GET("/entry/:id", h.GetEntryById)
	}

	return router
}
