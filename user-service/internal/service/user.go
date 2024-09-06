package service

import (
	"users/internal/repo"
)

type UserService interface {
	Get() error
}

type user struct {
	r repo.UserRepo
}

func (u user) Get() error {
	u.r.Get()
	return nil
}
