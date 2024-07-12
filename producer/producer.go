package producer

import (
    "log"
    "github.com/IBM/sarama"
)

func ProduceMessage(brokerList []string, topic, message string) error {
    // Set up configuration
    config := sarama.NewConfig()
    config.Producer.Return.Successes = true
    config.Producer.Return.Errors = true

    // Initialize producer
    producer, err := sarama.NewSyncProducer(brokerList, config)
    if err != nil {
        return err
    }
    defer producer.Close()

    // Create a message object
    msg := &sarama.ProducerMessage{
        Topic: topic,
        Value: sarama.StringEncoder(message),
    }

    // Send message
    partition, offset, err := producer.SendMessage(msg)
    if err != nil {
        return err
    }

    // Print partition and offset
    log.Printf("Message sent to partition %d at offset %d\n", partition, offset)

    return nil
}
