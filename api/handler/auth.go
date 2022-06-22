package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/AsaHero/simple-bank/api/models"
)

func (h *Handler) SignIn(c *gin.Context) {

}

func (h *Handler) SignUp(c *gin.Context) {
	var req models.CreateAccountRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "cannot parse json from body")
		return
	}

	resp, err := h.service.Authorization.CreateAccount(req)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, resp)
}
