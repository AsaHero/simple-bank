package handler

import (
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
