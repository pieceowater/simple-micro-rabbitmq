package controllers

import (
	"log"
	"simple-micro-rabbitmq/services"
)

type Message struct {
	Action string      `json:"action"`
	Data   interface{} `json:"data"`
}

func HandleMessage(msg Message) interface{} {
	switch msg.Action {
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

	default:
		log.Println("Unknown action:", msg.Action)
		return nil
	}
}
