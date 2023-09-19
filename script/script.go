package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/nats-io/stan.go"
)

type Message struct {
	Text string `json:"text"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}

	clusterID := os.Getenv("CLUSTER_ID")
	clientID := os.Getenv("CLIENT_ID")
	channelName := os.Getenv("CHANNEL_NAME")

	// Подключение к NATS Streaming
	nc, err := stan.Connect(clusterID, clientID)
	if err != nil {
		log.Fatalf("Ошибка подключения к NATS Streaming: %v", err)
	}
	defer nc.Close()

	// Бесконечный цикл для публикации сообщений
	for {
		message := Message{Text: "Сообщение"}

		messageJSON, err := json.Marshal(message)
		if err != nil {
			log.Printf("Ошибка при маршалинге JSON: %v", err)
			continue
		}

		err = nc.Publish(channelName, messageJSON)
		if err != nil {
			log.Printf("Ошибка при публикации сообщения: %v", err)
		} else {
			fmt.Println("Сообщение успешно опубликовано.")
		}

		// Подождать некоторое время перед отправкой следующего сообщения
		time.Sleep(5 * time.Second)
	}
}
