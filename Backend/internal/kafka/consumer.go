package kafka

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

func ConsumeOrderEvents() {
	// Create Kafka reader (connect to Kafka)
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "order-events",
		GroupID: "order-consumer-group", // Consumers in a group will share the load
	})

	for {
		// Read messages from Kafka
		message, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatalf("could not read message: %v", err)
		}

		// Process the message (you can do anything here, e.g., logging, notifying, etc.)
		log.Printf("Received order event: %s", string(message.Value))
	}
}
