package http

import (
	"users/internal/repo"
	"users/internal/service"
	"users/internal/transport/http/handler"
)

func Start(listen string) {
	repo := repo.New()
	service := service.New(repo)
	handler := handler.New(service)
	router := routes(handler)

	router.Run(listen)
}
