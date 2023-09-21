package main

import (
	"log"
	"net/http"
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

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Обработчик для отображения данных заказа по ID
	http.HandleFunc("/order/", getOrderByID)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/index.html")
	})

	log.Fatal(http.ListenAndServe(":8080", nil)) // Запуск сервера на порту 8080

	waitForInterrupt()
}
