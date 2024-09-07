package handler

import (
	"errors"
	"fmt"
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
	s service.Service
}

func (h oh) GetByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("id is required")})
		return
	}

	order, user, err := h.s.Order().GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(order)
	fmt.Println(user)

	content := response.GetOrderByIdResponse{
		OrderId:   order.OrderId,
		Quantity:  order.Quantity,
		ProductId: order.ProductId,
		UserId:    user.ID,
		Email:     user.Email,
		Username:  user.Username,
	}

	c.JSON(http.StatusBadRequest, gin.H{"success": true, "content": content})
}

func (h oh) Create(c *gin.Context) {
	var payload payload.CreateOrderPayload

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order, err := h.s.Order().Create(payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	content := (*response.CreateOrderResponse)(order)
	c.JSON(http.StatusBadRequest, gin.H{"success": true, "content": content})
}
