package consumer

import (
    "log"
    "github.com/IBM/sarama"
)

func ConsumeMessages(brokerList []string, topic string) error {
    // Set up configuration
    config := sarama.NewConfig()
    config.Consumer.Return.Errors = true
    config.Consumer.Offsets.Initial = sarama.OffsetOldest // Start consuming from the oldest message

    // Initialize consumer
    consumer, err := sarama.NewConsumer(brokerList, config)
    if err != nil {
        return err
    }
    defer consumer.Close()

    // Obtain a PartitionConsumer
    partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetOldest)
    if err != nil {
        return err
    }
    defer partitionConsumer.Close()

    // Consume messages
    for {
        select {
        case err := <-partitionConsumer.Errors():
            log.Fatalf("Error consuming message: %v", err)
        case msg := <-partitionConsumer.Messages():
            log.Printf("Received message: %s\n", msg.Value)
            // Optionally: Process the received message
        }
    }
}
