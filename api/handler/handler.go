package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/AsaHero/simple-bank/service"
)

type Handler struct {
	service service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: *service,
	}
}

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/signin", h.SignIn)
		auth.POST("/signup", h.SignUp)
	}

	// transaction := router.Group("/")
	// {
	// 	transaction.GET("/get-all-transfers")
	// }

	return router
}
