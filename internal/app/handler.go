package app

import (
	"encoding/json"
	"github.com/Bzelijah/go-kafka-libs-comparement/internal/producer"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"io"
	"net/http"
)

type Handler struct {
	confluentKafkaProducer *producer.ConfluentKafkaProducer
	saramaProducer         *producer.SaramaProducer
}

func NewHandler(confluentKafkaProducer *producer.ConfluentKafkaProducer, saramaProducer *producer.SaramaProducer) *Handler {
	return &Handler{
		confluentKafkaProducer: confluentKafkaProducer,
		saramaProducer:         saramaProducer,
	}
}

type AddMessageRequest struct {
	Message string `json:"message"`
	LibName string `json:"lib-name"`
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

	if body.LibName == "confluent-kafka-go" {
		if err = h.confluentKafkaProducer.AddMessage(body.Message); err != nil {
			log.Err(err).Msg("could not add message")
			return c.String(http.StatusInternalServerError, "confluent-kafka-go: could not add message")
		}

		return c.String(http.StatusOK, "confluent-kafka-go: added message")
	} else if body.LibName == "sarama" {
		if err = h.saramaProducer.AddMessage(body.Message); err != nil {
			log.Err(err).Msg("could not add message")
			return c.String(http.StatusInternalServerError, "sarama: could not add message")
		}

		return c.String(http.StatusOK, "sarama: added message")
	}

	return c.String(http.StatusNotFound, "unknown request")
}
