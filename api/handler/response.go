package handler

import (
	"log"

	"github.com/gin-gonic/gin"
)

type Error struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, status_code int, message string) {
	log.Println(message)
	c.AbortWithStatusJSON(status_code, Error{
		Message: message,
	})
}
