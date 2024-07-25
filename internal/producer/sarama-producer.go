package producer

import (
	"fmt"
	"github.com/IBM/sarama"
	"log"
)

func RunSaramaProducer(partition int32) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatal("Failed to start Sarama producer:", err)
	}
	defer producer.Close()

	topic := "quickstart-topic"
	messages := []string{"Welcome", "to", "the", "Sarama", "client"}

	for _, msg := range messages {
		message := &sarama.ProducerMessage{
			Topic:     topic,
			Partition: partition,
			Value:     sarama.StringEncoder(msg),
		}
		partition, offset, err := producer.SendMessage(message)
		if err != nil {
			fmt.Printf("Failed to send message: %v\n", err)
		} else {
			fmt.Printf("Message sent to partition %d with offset %d\n", partition, offset)
		}
	}
}
