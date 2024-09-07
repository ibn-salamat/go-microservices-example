package service

import "orders/internal/repo"

type Service interface {
	Order() OrderService
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

func (s service) Order() OrderService {
	o := os{
		or: s.r.Order(),
		ur: s.r.User(),
	}
	return o
}

func (s service) User() UserService {
	u := us{
		ur: s.r.User(),
	}
	return u
}
