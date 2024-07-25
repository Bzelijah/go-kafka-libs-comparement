package consumer

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/pkg/errors"
	"log"
)

func RunConfluentKafkaConsumer(partition int32, topic string) {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "group1",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		log.Fatal("Failed to start Confluent Kafka consumer:", errors.Wrap(err, "NewConfluentKafkaConsumer"))
	}

	tp := kafka.TopicPartition{Topic: &topic, Partition: partition}

	err = consumer.Assign([]kafka.TopicPartition{tp})
	if err != nil {
		log.Fatal("Failed to assign topic partition:", err)
	}

	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			fmt.Printf("ConfluentKafkaConsumer: Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
			topics, err := consumer.CommitMessage(msg)
			if err != nil {
				fmt.Printf("Failed to commit message: %v\n", err)
			}

			fmt.Println("topics", topics)
		} else {
			fmt.Printf("ConfluentKafkaConsumer error: %v\n", err)
		}
	}
}
