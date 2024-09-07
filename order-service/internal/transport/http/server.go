package http

import (
	"orders/config"
	"orders/internal/repo"
	"orders/internal/service"
	"orders/internal/transport/http/handler"
)

func Start(conf config.Config, repo repo.Repo) {
	service := service.New(repo)
	handler := handler.New(service)
	router := routes(handler)

	router.Run(conf.Port)
}
