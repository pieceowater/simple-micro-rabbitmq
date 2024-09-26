package main

import (
	"encoding/json"
	gossiper "github.com/pieceowater-dev/lotof.lib.gossiper"
	"log"
	"simple-micro-rabbitmq/config"
	"simple-micro-rabbitmq/controllers"
)

func main() {
	gossiper.Setup(config.GossiperConf, func(msg []byte) interface{} {
		var message gossiper.AMQMessage
		err := json.Unmarshal(msg, &message)
		if err != nil {
			log.Println("Failed to unmarshal custom message:", err)
			return nil
		}
		return controllers.HandleMessage(message)
	})
}
