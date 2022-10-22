package queues

import (
	"log"
	"microservice/src/configs"
	"microservice/src/utils/rabbitmq"
	"time"
)

func Example() {
	messages := rabbitmq.CreateListener(configs.ExampleQueue())
	go func() {
		for message := range messages {
			// For example, show received message in a console.
			log.Printf(" > Received queue message: %s\n", message.Body)
			time.Sleep(1 * time.Second)
			message.Ack(true)
			time.Sleep(1 * time.Second)
		}
	}()
}
