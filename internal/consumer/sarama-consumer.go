package consumer

import (
	"fmt"
	"github.com/IBM/sarama"
	"github.com/rs/zerolog/log"
)

func RunSaramaConsumer(partition int32, topic string) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.Initial = sarama.OffsetNewest

	client, err := sarama.NewClient([]string{"localhost:9092"}, config)
	if err != nil {
		log.Error().Err(err).Msg("SaramaConsumer: Failed to create Kafka client")
		return
	}
	defer client.Close()

	consumer, err := sarama.NewConsumerFromClient(client)
	if err != nil {
		log.Error().Err(err).Msg("SaramaConsumer: Failed to create Kafka consumer")
		return
	}
	defer consumer.Close()

	offsetManager, err := sarama.NewOffsetManagerFromClient("sarama-consumer", client)
	if err != nil {
		log.Error().Err(err).Msg("SaramaConsumer: Failed to create offset manager")
		return
	}
	defer offsetManager.Close()

	partitionOffsetManager, err := offsetManager.ManagePartition(topic, partition)
	if err != nil {
		log.Error().Err(err).Msg("SaramaConsumer: Failed to create partition offset manager")
		return
	}
	defer partitionOffsetManager.Close()

	nextOffset, _ := partitionOffsetManager.NextOffset()
	if nextOffset == sarama.OffsetNewest {
		nextOffset = sarama.OffsetNewest
	}

	partitionConsumer, err := consumer.ConsumePartition(topic, partition, nextOffset)
	if err != nil {
		log.Error().Err(err).Msg("SaramaConsumer: Failed to create partition consumer")
		return
	}
	defer partitionConsumer.Close()

	log.Info().Msg("SaramaConsumer is running. Consuming partition " + fmt.Sprint(partition))

	for {
		select {
		case msg := <-partitionConsumer.Messages():
			fmt.Printf("SaramaConsumer: Message received: %s\n", string(msg.Value))
			partitionOffsetManager.MarkOffset(msg.Offset+1, "") // Mark the offset as processed
		case err := <-partitionConsumer.Errors():
			fmt.Printf("SaramaConsumer error: %v\n", err)
		}
	}
}
