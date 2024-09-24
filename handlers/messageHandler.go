package handlers

import (
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
	"simple-micro-rabbitmq/config"
	"simple-micro-rabbitmq/controllers"
)

func ConsumeMessage() error {
	conn, err := amqp.Dial(config.GetRabbitDSN())
	if err != nil {
		log.Fatal("Failed to connect to RabbitMQ:", err)
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("Failed to open a channel:", err)
		return err
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"template_queue", // queue name
		true,             // durable
		false,            // delete when unused
		false,            // exclusive
		false,            // no-wait
		nil,              // arguments
	)
	if err != nil {
		log.Fatal("Failed to declare a queue:", err)
		return err
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Fatal("Failed to register a consumer:", err)
		return err
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)

			var msg controllers.Message
			err := json.Unmarshal(d.Body, &msg)
			if err != nil {
				log.Println("Failed to parse message:", err)
				continue
			}

			// Передаем сообщение в контроллер и получаем результат
			response := controllers.HandleMessage(msg)

			// Если запрос требует ответа, отправляем его в очередь, указанную в ReplyTo
			if d.ReplyTo != "" {
				responseBytes, _ := json.Marshal(response)

				err = ch.Publish(
					"",        // exchange
					d.ReplyTo, // routing key (reply queue)
					false,     // mandatory
					false,     // immediate
					amqp.Publishing{
						ContentType:   "application/json",
						CorrelationId: d.CorrelationId,
						Body:          responseBytes,
					},
				)

				if err != nil {
					log.Println("Failed to publish response:", err)
				} else {
					log.Println("Response sent to:", d.ReplyTo)
				}
			}
		}
	}()

	log.Println("Waiting for messages. To exit press CTRL+C")
	<-forever

	return nil
}
