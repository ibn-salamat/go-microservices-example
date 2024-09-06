package service

import (
	"users/internal/models"
	"users/internal/repo"
	"users/internal/transport/http/handler/payload"
)

type UserService interface {
	GetByID(id int) (*models.User, error)
	Create(payload.CreateUserPayload) (*models.User, error)
}

type us struct {
	r repo.UserRepo
}

func (s us) GetByID(id int) (*models.User, error) {
	return s.r.GetByID(id)
}

func (s us) Create(payload payload.CreateUserPayload) (*models.User, error) {
	return s.r.Create(payload)
}
