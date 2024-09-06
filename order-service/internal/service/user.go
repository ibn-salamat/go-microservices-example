package service

import (
	"orders/internal/models"
	"orders/internal/repo"
	"orders/internal/transport/http/handler/payload"
)

type OrderService interface {
	GetByID(id string) (*models.Order, error)
	Create(payload.CreateOrderPayload) (*models.Order, error)
}

type us struct {
	r repo.OrderRepo
}

func (s us) GetByID(id string) (*models.Order, error) {
	return s.r.GetByID(id)
}

func (s us) Create(payload payload.CreateOrderPayload) (*models.Order, error) {
	return s.r.Create(payload)
}
