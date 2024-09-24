package main

import (
	"log"
	"simple-micro-rabbitmq/config"
	"simple-micro-rabbitmq/handlers"
)

func main() {
	// Загружаем конфигурацию
	config.LoadEnv()

	// Начинаем чтение сообщений из RabbitMQ
	log.Println("Starting to consume messages...")
	err := handlers.ConsumeMessage()
	if err != nil {
		log.Fatal("Failed to consume messages:", err)
	}
}
