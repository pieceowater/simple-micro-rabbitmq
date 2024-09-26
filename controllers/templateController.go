package controllers

import (
	gossiper "github.com/pieceowater-dev/lotof.lib.gossiper"
	"log"
	"simple-micro-rabbitmq/services"
)

type Message struct {
	Pattern string      `json:"pattern"`
	Data    interface{} `json:"data"`
}

func HandleMessage(msg gossiper.AMQMessage) interface{} {
	switch msg.Pattern { // Используем Pattern вместо Action
	case "findOneItem":
		item := services.TemplateGetItem(msg.Data)
		log.Printf("findOneItem: %v", item)
		return item

	case "findAllItem":
		items := services.TemplateGetItems()
		log.Printf("findAllItem: %v", items)
		return items

	case "createItem":
		created := services.TemplateCreateItem(msg.Data)
		log.Printf("createItem: %v", created)
		return created

	case "updateItem":
		updated := services.TemplateUpdateItem(msg.Data)
		log.Printf("updateItem: %v", updated)
		return updated

	case "ping":
		log.Println("Received PING request")
		return "PONG"

	default:
		log.Println("Unknown action:", msg.Pattern)
		return "Unknown action"
	}
}
