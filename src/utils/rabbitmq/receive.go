package rabbitmq

import (
	"fmt"
	"log"
	"main-service/src/apps"

	"github.com/streadway/amqp"
)

func CreateListener(queue_name string) <-chan amqp.Delivery {
	channelRabbitMQ, err_conn := apps.RabbitConnect()
	if err_conn != nil {
		panic(err_conn)
	}
	// Subscribing to QueueService1 for getting messages.
	messages, err_consume := channelRabbitMQ.Consume(
		queue_name, // queue name
		"",         // consumer
		false,      // auto-ack
		false,      // exclusive
		false,      // no local
		false,      // no wait
		nil,        // arguments
	)
	if err_consume != nil {
		log.Println(err_consume)
	} else {
		fmt.Printf("Waiting messages from %s...\n", queue_name)
	}
	return messages
}
