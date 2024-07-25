package app

import (
	"encoding/json"
	"github.com/Bzelijah/go-kafka-libs-comparement/internal/producer"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"io"
	"net/http"
)

type Handler struct {
	confluentKafkaProducer *producer.ConfluentKafkaProducer
	saramaProducer         *kafka.Producer
}

func NewHandler(confluentKafkaProducer *producer.ConfluentKafkaProducer,

// saramaProducer *kafka.Producer,
) *Handler {
	return &Handler{
		confluentKafkaProducer: confluentKafkaProducer,
		//saramaProducer:         saramaProducer,
	}
}

type AddMessageRequest struct {
	Message string `json:"message"`
}

func (h *Handler) HandleAddMessage(c echo.Context) error {
	buf, err := io.ReadAll(c.Request().Body)
	if err != nil {
		log.Err(err).Msg("could not read request body")
		return c.String(http.StatusInternalServerError, "could not read request body")
	}

	body := AddMessageRequest{}

	err = json.Unmarshal(buf, &body)
	if err != nil {
		log.Err(err).Msg("could not unmarshal request body")
	}

	if err := h.confluentKafkaProducer.AddMessage(body.Message); err != nil {
		log.Err(err).Msg("could not add message")
	}

	return nil
}
