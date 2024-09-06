package service

import "orders/internal/repo"

type Service interface {
	Order() OrderService
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
	u := us{
		r: s.r.Order(),
	}
	return u
}
