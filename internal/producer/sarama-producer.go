package producer

import (
	"fmt"
	"github.com/IBM/sarama"
	"github.com/rs/zerolog/log"
)

type SaramaProducer struct {
	producer  sarama.SyncProducer
	partition int32
	topic     string
}

func RunSaramaProducer(partition int32, topic string) (*SaramaProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Err(err).Msg("Failed to start Sarama producer")
		return nil, err
	}

	return &SaramaProducer{producer: producer, partition: partition, topic: topic}, nil
}

func (p *SaramaProducer) AddMessage(message string) error {
	producerMessage := &sarama.ProducerMessage{
		Topic:     p.topic,
		Partition: p.partition,
		Value:     sarama.StringEncoder(message),
	}

	_, _, err := p.producer.SendMessage(producerMessage)
	if err != nil {
		fmt.Printf("SaramaProducer: Failed to send message: %v\n", err)
		return err
	}

	fmt.Printf("SaramaProducer: Message sent to partition %d\n", p.partition)

	return nil
}
