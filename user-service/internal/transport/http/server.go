package http

import (
	"log"
	"users/config"
	"users/internal/repo"
	"users/internal/service"
	"users/internal/transport/http/handler"
	"users/pkg/postgres"
	"users/pkg/rmq"
)

func Start() {
	conf, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	db, err := postgres.Connect(conf.PgDSN)
	if err != nil {
		log.Fatal(err)
	}

	rabbit, err := rmq.Connect(conf.RABBITMQ_URL)
	if err != nil {
		log.Fatal(err)
	}

	repo := repo.New(db)
	service := service.New(repo, rabbit)
	handler := handler.New(service)
	router := routes(handler)

	router.Run(conf.Port)
}
