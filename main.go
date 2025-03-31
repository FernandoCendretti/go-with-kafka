package main

import (
	"context"
	"log"
	"os"

	"github.com/segmentio/kafka-go"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Uso: %s [pub|sub]", os.Args[0])
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
		log.Fatalf("Modo inválido: %s. Use 'pub' para publicar ou 'sub' para consumir.", mode)
	}
}

func publish(broker string, topic string) {
	// Configuração do writer (produtor)
	writer := kafka.Writer{
		Addr:     kafka.TCP(broker),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}

	defer writer.Close()

	// Mensagem a ser enviada
	message := kafka.Message{
		Key:   []byte("Key"),
		Value: []byte("Olá, Kafka com kafka-go!"),
	}

	// Envia a mensagem
	err := writer.WriteMessages(context.Background(), message)
	if err != nil {
		log.Fatalf("Erro ao enviar mensagem: %v", err)
	}

	log.Println("Mensagem enviada com sucesso!")
}

func subscribe(broker string, topic string) {
	// Configuração do reader (consumidor)
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{broker},
		Topic:     topic,
		Partition: 0,
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
	})

	defer reader.Close()

	log.Println("Aguardando mensagens...")

	// Loop para consumir mensagens
	for {
		message, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatalf("Erro ao ler mensagem: %v", err)
		}

		log.Printf("Mensagem recebida: chave=%s valor=%s\n", string(message.Key), string(message.Value))
	}
}
