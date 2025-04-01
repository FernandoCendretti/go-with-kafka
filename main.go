package main

import (
	"context"
	"log"
	"os"

	"github.com/segmentio/kafka-go"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s [pub|sub]", os.Args[0])
	}

	mode := os.Args[1]
	brokers := []string{"kafka:29092"}
	topic := "example-topic"

	switch mode {
	case "pub":
		publish(brokers[0], topic)
	case "sub":
		subscribe(brokers[0], topic)
	default:
		log.Fatalf("Invalid mode: %s. Use 'pub' to publish or 'sub' to consume.", mode)
	}
}

func publish(broker string, topic string) {
	// Writer (producer) configuration
	writer := kafka.Writer{
		Addr:     kafka.TCP(broker),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}

	defer writer.Close()

	// Message to be sent
	message := kafka.Message{
		Key:   []byte("Key"),
		Value: []byte("Hello, Kafka with kafka-go!"),
	}

	// Send the message
	err := writer.WriteMessages(context.Background(), message)
	if err != nil {
		log.Fatalf("Error sending message: %v", err)
	}

	log.Println("Message sent successfully!")
}

func subscribe(broker string, topic string) {
	// Reader (consumer) configuration
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{broker},
		Topic:     topic,
		Partition: 0,
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
	})

	defer reader.Close()

	log.Println("Waiting for messages...")

	// Loop to consume messages
	for {
		message, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatalf("Error reading message: %v", err)
		}

		log.Printf("Message received: key=%s value=%s\n", string(message.Key), string(message.Value))
	}
}
