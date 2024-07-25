package main

import (
	"github.com/Bzelijah/go-kafka-libs-comparement/internal/app"
	"github.com/Bzelijah/go-kafka-libs-comparement/internal/consumer"
	"github.com/Bzelijah/go-kafka-libs-comparement/internal/producer"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {
	go consumer.RunSaramaConsumer(0)
	//go producer.RunSaramaProducer(0)
	//go consumer.RunConfluentKafkaConsumer()
	confluentKafkaProducer, err := producer.NewConfluentKafkaProducer(0)
	if err != nil {
		log.Fatal(err)
	}

	appHandler := app.NewHandler(confluentKafkaProducer)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/add-message", appHandler.HandleAddMessage)

	if err := e.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}
