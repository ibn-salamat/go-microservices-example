package service

import (
	"orders/internal/models"
	"orders/internal/repo"
)

type UserService interface {
	GetByID(id int) (*models.User, error)
	Create(id int) error
}

type us struct {
	ur repo.UserRepo
}

func (s us) GetByID(id int) (*models.User, error) {
	return s.ur.GetByID(id)
}

func (s us) Create(id int) error {
	return s.ur.Create(id)
}
