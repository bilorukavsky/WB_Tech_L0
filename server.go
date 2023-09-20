package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

func getOrderByID(w http.ResponseWriter, r *http.Request) {
	// Извлечь ID заказа из URL-пути
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 3 {
		http.Error(w, "Неверный URL", http.StatusBadRequest)
		return
	}

	orderID := parts[2]

	cachedOrder, found := cache.Get(orderID)
	if !found {
		http.Error(w, "Заказ не найден", http.StatusNotFound)
		return
	}

	// Преобразование данных заказа в JSON
	jsonOrder, err := json.Marshal(cachedOrder)
	if err != nil {
		http.Error(w, "Ошибка при кодировании JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonOrder)
}
