package consumer

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/rs/zerolog/log"
)

func RunConfluentKafkaConsumer(partition int32, topic string) {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":  "localhost:9092",
		"group.id":           "confluent-kafka-consumer",
		"enable.auto.commit": false,
		"auto.offset.reset":  "latest",
	})
	if err != nil {
		log.Error().Err(err).Msg("ConfluentKafkaConsumer: Failed to start Confluent Kafka consumer")
		return
	}
	defer consumer.Close()

	tp := kafka.TopicPartition{Topic: &topic, Partition: partition, Offset: kafka.OffsetStored}

	err = consumer.Assign([]kafka.TopicPartition{tp})
	if err != nil {
		log.Error().Err(err).Msg("ConfluentKafkaConsumer: Failed to assign topic partition")
		return
	}

	log.Info().Msg("ConfluentKafkaConsumer is running. Consuming partition " + fmt.Sprint(partition))

	for {
		msg, err := consumer.ReadMessage(-1)
		if err != nil {
			fmt.Printf("ConfluentKafkaConsumer error: %v\n", err)
			continue
		}

		fmt.Printf("ConfluentKafkaConsumer: Message received: %s\n", string(msg.Value))

		_, err = consumer.CommitMessage(msg)
		if err != nil {
			fmt.Printf("ConfluentKafkaConsumer: Failed to commit message: %v\n", err)
		}
	}
}
