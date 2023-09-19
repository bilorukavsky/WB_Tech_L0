package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/nats-io/stan.go"
)

func connectToNATSStreaming(clusterID, clientID string) (stan.Conn, error) {
	return stan.Connect(clusterID, clientID)
}

func subscribeToChannel(nc stan.Conn, channelName string) (stan.Subscription, error) {
	subscription, err := nc.Subscribe(channelName, func(msg *stan.Msg) {
		log.Printf("Получено сообщение: %s\n", string(msg.Data))

		order := Order{}

		if err := json.Unmarshal(msg.Data, &order); err != nil {
			log.Printf("Ошибка при распаковке JSON: %v", err)
			return
		}

		// Кэшировать полученные данные
		cachedOrder, found := getFromCache(order.OrderUID)
		if found {
			log.Printf("Данные уже есть в кэше: %+v", cachedOrder)
		}
		if err := db.Create(&order).Error; err != nil {
			log.Printf("Ошибка при записи в БД: %v", err)
		}

		log.Printf("Данные успешно сохранены в БД")
		cacheData(order)

	})

	if err != nil {
		log.Printf("Ошибка при подписке на канал: %v", err)
	}

	return subscription, err
}

func waitForInterrupt() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig
	fmt.Println("Завершение работы...")
}
