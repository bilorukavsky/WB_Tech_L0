package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
	//_ "github.com/lib/pq"
	//"gorm.io/gorm"
)

type Delivery struct {
	gorm.Model
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Zip     string `json:"zip"`
	City    string `json:"city"`
	Address string `json:"address"`
	Region  string `json:"region"`
	Email   string `json:"email"`
	OrderID int    `json:"order_id"`
}

type Payment struct {
	gorm.Model
	Transaction  string `json:"transaction"`
	RequestID    string `json:"request_id"`
	Currency     string `json:"currency"`
	Provider     string `json:"provider"`
	Amount       int    `json:"amount"`
	PaymentDt    int64  `json:"payment_dt"`
	Bank         string `json:"bank"`
	DeliveryCost int    `json:"delivery_cost"`
	GoodsTotal   int    `json:"goods_total"`
	CustomFee    int    `json:"custom_fee"`
	OrderID      int    `json:"order_id"`
}

type Item struct {
	gorm.Model
	ChrtID      int    `json:"chrt_id"`
	TrackNumber string `json:"track_number"`
	Price       int    `json:"price"`
	Rid         string `json:"rid"`
	Name        string `json:"name"`
	Sale        int    `json:"sale"`
	Size        string `json:"size"`
	TotalPrice  int    `json:"total_price"`
	NmID        int    `json:"nm_id"`
	Brand       string `json:"brand"`
	Status      int    `json:"status"`
	OrderID     int    `json:"order_id"`
}

type Order struct {
	gorm.Model
	OrderUID        *string  `json:"order_uid"`
	TrackNumber     *string  `json:"track_number"`
	Entry           *string  `json:"entry"`
	Delivery        Delivery `json:"delivery"`
	Payment         Payment  `json:"payment"`
	Items           []Item   `json:"items"`
	Locale          *string  `json:"locale"`
	InternalSig     *string  `json:"internal_signature"`
	CustomerID      *string  `json:"customer_id"`
	DeliveryService *string  `json:"delivery_service"`
	ShardKey        *string  `json:"shardkey"`
	SMID            int      `json:"sm_id"`
	DateCreated     *string  `json:"date_created"`
	OOFShard        *string  `json:"oof_shard"`
}

func main() {
	// Подключение к базе данных
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}

	initializeDB()
	defer closeDB() // Закрыть соединение с базой данных
	migrateDB()     // Создание таблиц

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
