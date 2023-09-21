package main

import (
	"testing"
)

func TestCacheSetAndGet(t *testing.T) {
	cache := NewCache()

	// Тест на установку и получение данных из кэша
	order := Order{OrderUID: "test_order", TrackNumber: "123456"}
	cache.Set(order.OrderUID, order)

	cachedOrder, found := cache.Get(order.OrderUID)
	if !found {
		t.Error("Данные не найдены в кэше")
	}

	if cachedOrder.OrderUID != order.OrderUID && cachedOrder.TrackNumber != order.TrackNumber {
		t.Errorf("Ожидается: %+v, Получено: %+v", order, cachedOrder)
	}
}
