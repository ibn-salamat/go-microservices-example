package http

import (
	"orders/config"
	"orders/internal/service"
	"orders/internal/transport/http/handler"
)

func Start(conf config.Config, service service.Service) {
	handler := handler.New(service)
	router := routes(handler)

	router.Run(conf.Port)
}
