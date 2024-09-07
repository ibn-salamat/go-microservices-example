package cmd

import (
	"log"
	"orders/config"
	"orders/internal/repo"
	"orders/internal/service"
	"orders/internal/transport/http"
	"orders/pkg/mongo"

	"github.com/rabbitmq/amqp091-go"
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

	conn, err := amqp091.Dial(conf.RABBITMQ_URL)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	repo := repo.New(db)
	service := service.New(repo, conn)

	go func() {
		service.ListenRMQ()
	}()

	http.Start(conf, service)
}
