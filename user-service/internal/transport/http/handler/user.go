package handler

import (
	"users/internal/service"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	Get(c *gin.Context)
}

type userHandler struct {
	s service.UserService
}

func (h userHandler) Get(c *gin.Context) {
	h.s.Get()

	c.JSON(200, "get user")
}
