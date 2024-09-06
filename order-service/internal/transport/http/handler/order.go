package handler

import (
	"errors"
	"net/http"
	"orders/internal/service"
	"orders/internal/transport/http/handler/payload"
	"orders/internal/transport/http/handler/response"

	"github.com/gin-gonic/gin"
)

type OrderHandler interface {
	GetByID(c *gin.Context)
	Create(c *gin.Context)
}

type oh struct {
	s service.OrderService
}

func (h oh) GetByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("id is required")})
		return
	}

	order, err := h.s.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	content := (*response.CreateOrderResponse)(order)
	c.JSON(http.StatusBadRequest, gin.H{"success": true, "content": content})
}

func (h oh) Create(c *gin.Context) {
	var payload payload.CreateOrderPayload

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order, err := h.s.Create(payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	content := (*response.CreateOrderResponse)(order)
	c.JSON(http.StatusBadRequest, gin.H{"success": true, "content": content})
}
