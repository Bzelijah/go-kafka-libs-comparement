package main

import (
	"github.com/Bzelijah/go-kafka-libs-comparement/internal/app"
	"github.com/Bzelijah/go-kafka-libs-comparement/internal/consumer"
	"github.com/Bzelijah/go-kafka-libs-comparement/internal/producer"
	"github.com/labstack/echo/v4"
	"log"
)

func main() {
	topic := "quickstart-topic"

	go consumer.RunSaramaConsumer(1, topic)
	go consumer.RunConfluentKafkaConsumer(0, topic)

	confluentKafkaProducer, err := producer.NewConfluentKafkaProducer(1, topic)
	if err != nil {
		log.Fatal(err)
	}

	saramaProducer, err := producer.RunSaramaProducer(0, topic)
	if err != nil {
		log.Fatal(err)
	}

	appHandler := app.NewHandler(confluentKafkaProducer, saramaProducer)

	e := echo.New()
	e.POST("/add-message", appHandler.HandleAddMessage)

	if err := e.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}
