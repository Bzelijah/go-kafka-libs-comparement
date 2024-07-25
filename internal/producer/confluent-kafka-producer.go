package producer

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type ConfluentKafkaProducer struct {
	producer  *kafka.Producer
	partition int32
}

func NewConfluentKafkaProducer(partition int32) (*ConfluentKafkaProducer, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092"})
	if err != nil {
		return nil, err
	}

	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	return &ConfluentKafkaProducer{
		producer: p,
	}, nil
}

var topic = "quickstart-topic"

func (p *ConfluentKafkaProducer) AddMessage(message string) error {
	if err := p.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: p.partition},
		Value:          []byte(message),
	}, nil); err != nil {
		return err
	}

	return nil
}