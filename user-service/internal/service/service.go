package service

import "users/internal/repo"

type Service interface {
	User() UserService
}

type service struct {
	r repo.Repo
}

func New(r repo.Repo) Service {
	return service{
		r: r,
	}
}

func (s service) User() UserService {
	u := us{
		r: s.r.User(),
	}
	return u
}
