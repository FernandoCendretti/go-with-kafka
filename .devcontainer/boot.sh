#!/bin/zsh

# Update package list
sudo apt update

# Install Go dependencies
go mod tidy

# Change ownership of Go directories to vscode user
sudo chown -R vscode:vscode

# Creating a Kafka topic
echo "Wating for Kafka to start..."
sleep 10

kafka-topics --create \
  --bootstrap-server kafka:29092 \
  --replication-factor 1 \
  --partitions 1 \
  --topic example-topic || echo "The topic already exists."

echo "Topic created."