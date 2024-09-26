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
	case "templateGetItem":
		item := services.TemplateGetItem(msg.Data)
		log.Printf("templateGetItem: %v", item)
		return item

	case "templateGetItems":
		items := services.TemplateGetItems()
		log.Printf("templateGetItems: %v", items)
		return items

	case "templateCreateItem":
		created := services.TemplateCreateItem(msg.Data)
		log.Printf("templateCreateItem: %v", created)
		return created

	case "templateUpdateItem":
		updated := services.TemplateUpdateItem(msg.Data)
		log.Printf("templateUpdateItem: %v", updated)
		return updated

	case "templateRemoveItem":
		removed := services.TemplateRemoveItem(msg.Data)
		log.Printf("templateRemoveItem: %v", removed)
		return removed

	case "ping":
		log.Println("Received PING request")
		return "PONG"

	default:
		log.Println("Unknown action:", msg.Pattern)
		return "Unknown action"
	}
}
