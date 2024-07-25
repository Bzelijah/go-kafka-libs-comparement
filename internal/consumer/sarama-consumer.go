package consumer

import (
	"fmt"
	"github.com/IBM/sarama"
	"github.com/pkg/errors"
	"log"
)

func RunSaramaConsumer(partition int32) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatal("Failed to start Sarama consumer:", errors.Wrap(err, "NewSaramaConsumer"))
	}
	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition("quickstart-topic", partition, sarama.OffsetOldest)
	if err != nil {
		log.Fatal("Failed to start partition consumer:", errors.Wrap(err, "NewSaramaConsumer"))
	}
	defer partitionConsumer.Close()

	for {
		select {
		case msg := <-partitionConsumer.Messages():
			fmt.Printf("SaramaConsumer: Message received: %s\n", string(msg.Value))
		case err := <-partitionConsumer.Errors():
			fmt.Printf("SaramaConsumer error: %v\n", err)
		}
	}

}
