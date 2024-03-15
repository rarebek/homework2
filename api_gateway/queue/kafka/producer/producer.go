package producer

import (
	"fmt"
	"log"

	"github.com/IBM/sarama"
)

var producer sarama.SyncProducer

func init() {
	var err error
	producer, err = NewProducer()
	if err != nil {
		log.Fatalf("Error creating Kafka producer: %v", err)
	}
}

func NewProducer() (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	// Optionally, set configuration options such as retries, compression, etc.
	// config.Producer.RequiredAcks = sarama.WaitForAll
	// config.Producer.Retry.Max = 3
	// config.Producer.Compression = sarama.CompressionSnappy

	// Create a new sync producer
	producer, err := sarama.NewSyncProducer([]string{"kafka:9092"}, config)
	if err != nil {
		return nil, fmt.Errorf("error creating sync producer: %v", err)
	}
	return producer, nil
}

func ProduceMessage(topic string, message string) error {
	// Create a new producer if it hasn't been initialized yet (unlikely due to the init() function, but good practice)
	if producer == nil {
		return fmt.Errorf("Kafka producer is not initialized")
	}

	// Construct the message to be sent
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}

	// Send the message and handle any errors
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Printf("Failed to send message: %v\n", err)
		return err
	}

	// Print the partition and offset of the produced message
	log.Printf("Message sent successfully! Partition: %d, Offset: %d\n", partition, offset)
	return nil
}
