Kafka Example
This project demonstrates basic interaction with Apache Kafka using the Sarama library in Golang. It includes functionalities for producing and consuming messages from Kafka topics.

Requirements
Go 1.11+
Apache Kafka
Sarama library (github.com/Shopify/sarama or github.com/IBM/sarama)
Installation
Apache Kafka
Ensure Apache Kafka is running on your system. If not already installed, you can use Docker to set it up:


docker pull apache/kafka:3.7.1
docker run -p 9092:9092 --network host --name kafka \
  -e KAFKA_ADVERTISED_LISTENERS=PLAINTEXT://localhost:9092 \
  -e KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR=1 \
  apache/kafka:3.7.1
Sarama Library
Install the Sarama library using go get:

go get -u github.com/Shopify/sarama
# or
go get -u github.com/IBM/sarama
Adjust the import in your Go code accordingly based on which Sarama repository you choose to use (github.com/Shopify/sarama or github.com/IBM/sarama).

Usage
Follow the examples in your code to produce and consume messages from Kafka topics. Ensure your Kafka broker configuration (brokerList) matches the configuration used in the Kafka Docker container (typically localhost:9092).
