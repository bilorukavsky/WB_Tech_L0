package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var cache = NewCache()

func main() {
	// Подключение к базе данных
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}

	initializeDB()
	defer closeDB() // Закрыть соединение с базой данных
	migrateDB()     // Создание таблиц

	cache.InitializeFromDB()

	clusterID := os.Getenv("CLUSTER_ID")
	clientID := os.Getenv("CLIENT_ID")
	channelName := os.Getenv("CHANNEL_NAME")

	nc, err := connectToNATSStreaming(clusterID, clientID)
	if err != nil {
		log.Fatalf("Ошибка подключения к NATS Streaming: %v", err)
	}
	defer nc.Close()

	subscription, err := subscribeToChannel(nc, channelName)
	if err != nil {
		log.Fatalf("Ошибка подписки на канал: %v", err)
	}
	defer subscription.Close()

	log.Printf("Подписка на канал '%s'...\n", channelName)

	waitForInterrupt()
}
