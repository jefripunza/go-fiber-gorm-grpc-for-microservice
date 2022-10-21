package rabbitmq

import (
	"main-service/src/apps"

	"github.com/streadway/amqp"
)

func Send(exchange string, msg string) error {
	channel, err_conn := apps.RabbitConnect()
	if err_conn != nil {
		return err_conn
	}
	err := channel.Publish(
		exchange, // exchange
		"",       // routing key
		false,    // mandatory
		false,    // immediate
		amqp.Publishing{
			//Headers:     headers,
			ContentType: "text/plain",
			Body:        []byte(msg),
		})
	if err != nil {
		return err
	}
	return nil
}
