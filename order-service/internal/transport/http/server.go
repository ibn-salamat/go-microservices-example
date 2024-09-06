package http

import (
	"log"
	"orders/config"
	"orders/internal/repo"
	"orders/internal/service"
	"orders/internal/transport/http/handler"
	"orders/pkg/mongo"
)

func Start() {
	conf, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	db, err := mongo.Connect(conf.MongoURL, conf.MongoUsername, conf.MongoPassword)
	if err != nil {
		log.Fatal(err)
	}

	repo := repo.New(db)
	service := service.New(repo)
	handler := handler.New(service)
	router := routes(handler)

	router.Run(conf.Port)
}
