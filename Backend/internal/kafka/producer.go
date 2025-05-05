package kafka

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

func ProduceOrderEvent(orderID string) {
	// Set up Kafka writer (connect to Kafka)
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "order-events",
		Balancer: &kafka.LeastBytes{},
	})

	// Create the message for Kafka
	message := kafka.Message{
		Key:   []byte(orderID),
		Value: []byte("New Order Created: " + orderID),
	}

	// Send the message to Kafka
	if err := writer.WriteMessages(context.Background(), message); err != nil {
		log.Fatalf("could not write message to Kafka: %v", err)
	}
	log.Println("Order event sent to Kafka: ", orderID)

	// Close the writer after use
	defer writer.Close()
}
