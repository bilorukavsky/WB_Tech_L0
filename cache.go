package main

import (
	"log"
	"sync"
)

type Cache struct {
	cache     map[string]Order
	cacheLock sync.Mutex
}

func NewCache() *Cache {
	return &Cache{
		cache: make(map[string]Order),
	}
}

func (c *Cache) Get(key string) (Order, bool) {
	c.cacheLock.Lock()
	defer c.cacheLock.Unlock()

	order, found := c.cache[key]
	return order, found
}

func (c *Cache) Set(key string, order Order) {
	c.cacheLock.Lock()
	c.cache[key] = order
	c.cacheLock.Unlock()
}

func (c *Cache) InitializeFromDB() {
	// Запрос к базе данных для получения данных
	var orders []Order
	if err := db.Find(&orders).Error; err != nil {
		log.Printf("Ошибка при загрузке данных из БД: %v", err)
		return
	}

	// Заполнение кэша данными из БД
	for _, order := range orders {
		c.Set(order.OrderUID, order)
	}

	log.Println("Кэш инициализирован из БД")
}
