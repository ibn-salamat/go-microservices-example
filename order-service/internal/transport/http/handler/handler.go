package handler

import (
	"orders/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	Ping(c *gin.Context)
	Order() OrderHandler
}

type handler struct {
	service service.Service
}

func New(s service.Service) Handler {
	return handler{
		service: s,
	}
}

func (h handler) Order() OrderHandler {
	o := oh{
		s: h.service.Order(),
	}
	return o
}

func (h handler) Ping(c *gin.Context) {
	c.JSON(200, "pong")
}
