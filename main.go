package main

import (
    "log"
    "kafka-example/producer"
    "kafka-example/consumer"
)

func main() {
    // Kafka broker list
    brokerList := []string{"localhost:9092"}

    // Example topic and message
    topic := "test-topic"
    message := "Hello, Kafka!"

    // Produce message
    err := producer.ProduceMessage(brokerList, topic, message)
    if err != nil {
        log.Fatalf("Error producing message: %v", err)
    }

    // Consume messages
    err = consumer.ConsumeMessages(brokerList, topic)
    if err != nil {
        log.Fatalf("Error consuming messages: %v", err)
    }
}
