package handler

import (
	"errors"
	"net/http"
	"strconv"
	"users/internal/service"
	"users/internal/transport/http/handler/payload"
	"users/internal/transport/http/handler/response"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	GetByID(c *gin.Context)
	Create(c *gin.Context)
}

type uh struct {
	s service.UserService
}

func (h uh) GetByID(c *gin.Context) {
	idParam := c.Param("id")
	if idParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("id is required")})
		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.s.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	content := (*response.CreateUserResponse)(user)
	c.JSON(http.StatusBadRequest, gin.H{"success": true, "content": content})
}

func (h uh) Create(c *gin.Context) {
	var payload payload.CreateUserPayload

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.s.Create(payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	content := (*response.CreateUserResponse)(user)
	c.JSON(http.StatusBadRequest, gin.H{"success": true, "content": content})
}
