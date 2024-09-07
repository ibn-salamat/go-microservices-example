package service

import (
	"errors"
	"orders/internal/models"
	"orders/internal/repo"
	"orders/internal/transport/grpc/response"
	"orders/internal/transport/http/handler/payload"
)

type OrderService interface {
	GetByID(id string) (*models.Order, *response.GetUserByIdResponse, error)
	Create(payload.CreateOrderPayload) (*models.Order, error)
}

type os struct {
	or      repo.OrderRepo
	ur      repo.UserRepo
	service Service
}

func (s os) GetByID(id string) (*models.Order, *response.GetUserByIdResponse, error) {
	order, err := s.or.GetByID(id)
	if err != nil {
		return nil, nil, err
	}

	user, err := s.service.GetByIDFromGRPC(order.UserId)
	if err != nil {
		return nil, nil, err
	}

	return order, user, err
}

func (s os) Create(payload payload.CreateOrderPayload) (*models.Order, error) {
	user, err := s.ur.GetByID(payload.UserId)
	if err != nil || user == nil {
		return nil, errors.New("user not found")
	}

	return s.or.Create(payload)
}
