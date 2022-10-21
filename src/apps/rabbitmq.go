package apps

import (
	"main-service/src/configs"

	"github.com/streadway/amqp"
)

func RabbitConnect() (*amqp.Channel, error) {
	conn, err_conn := amqp.Dial(configs.RabbitURL())
	if err_conn != nil {
		return nil, err_conn
	}
	// defer conn.Close()

	channel, err_channel := conn.Channel()
	if err_channel != nil {
		return nil, err_conn
	}
	// defer channel.Close()

	return channel, nil
}
