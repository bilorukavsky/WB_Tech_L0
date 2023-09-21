package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetOrderByID(t *testing.T) {
	// Подготовка данных для теста
	cache := NewCache()
	order := Order{
		OrderUID: "test_order",
	}
	cache.Set(order.OrderUID, order)

	// Создание фейкового сервера
	ts := httptest.NewServer(http.HandlerFunc(getOrderByID))
	defer ts.Close()

	// Тестирование корректного URL
	resp, err := http.Get(ts.URL + "/order/test_order")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		t.Errorf("Ожидается статус OK, получено: %v", resp.Status)
	}

	// Тестирование некорректного URL
	resp, err = http.Get(ts.URL + "/invalid_url")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Ожидается статус BadRequest, получено: %v", resp.Status)
	}
}
