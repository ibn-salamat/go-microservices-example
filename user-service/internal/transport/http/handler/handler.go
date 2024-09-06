package handler

import (
	"users/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	Ping(c *gin.Context)
	User() UserHandler
}

type handler struct {
	service service.Service
}

func New(s service.Service) Handler {
	return handler{
		service: s,
	}
}

func (h handler) User() UserHandler {
	u := uh{
		s: h.service.User(),
	}
	return u
}

func (h handler) Ping(c *gin.Context) {
	c.JSON(200, "pong")
}
