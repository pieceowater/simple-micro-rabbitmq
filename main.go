package main

import (
	gossiper "github.com/pieceowater-dev/lotof.lib.gossiper"
	"log"
	"simple-micro-rabbitmq/config"
	"simple-micro-rabbitmq/handlers"
)

func main() {
	gossiper.Setup(config.GossiperConf)

	// Начинаем чтение сообщений из RabbitMQ
	log.Println("Starting to consume messages...")
	err := handlers.ConsumeMessage()
	if err != nil {
		log.Fatal("Failed to consume messages:", err)
	}
}
