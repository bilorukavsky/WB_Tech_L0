package main

import (
	"log"
	"sync"
)

var (
	cache     = make(map[string]Order)
	cacheLock sync.Mutex
)

func getFromCache(key string) (Order, bool) {
	cacheLock.Lock()
	defer cacheLock.Unlock()

	order, found := cache[key]
	return order, found
}

func cacheData(order Order) {
	cacheLock.Lock()
	cache[order.OrderUID] = order
	cacheLock.Unlock()
}

func initializeCacheFromDB() {
	// Запрос к базе данных для получения данных
	var orders []Order
	if err := db.Find(&orders).Error; err != nil {
		log.Printf("Ошибка при загрузке данных из БД: %v", err)
		return
	}

	// Заполнение кэша данными из БД
	for _, order := range orders {
		cacheData(order)
	}

	log.Println("Кэш инициализирован из БД")
}
