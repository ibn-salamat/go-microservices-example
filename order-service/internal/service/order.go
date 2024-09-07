package service

import (
	"errors"
	"orders/internal/models"
	"orders/internal/repo"
	"orders/internal/transport/http/handler/payload"
)

type OrderService interface {
	GetByID(id string) (*models.Order, error)
	Create(payload.CreateOrderPayload) (*models.Order, error)
}

type os struct {
	or repo.OrderRepo
	ur repo.UserRepo
}

func (s os) GetByID(id string) (*models.Order, error) {
	return s.or.GetByID(id)
}

func (s os) Create(payload payload.CreateOrderPayload) (*models.Order, error) {
	user, err := s.ur.GetByID(payload.UserId)
	if err != nil || user == nil {
		return nil, errors.New("user not found")
	}

	return s.or.Create(payload)
}
