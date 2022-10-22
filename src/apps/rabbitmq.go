package apps

import (
	"microservice/src/configs"

	"github.com/streadway/amqp"
)

func RabbitConnect() (*amqp.Channel, error) {
	conn, err_conn := amqp.Dial(configs.RabbitURL())
	if err_conn != nil {
		return nil, err_conn
	}

	channel, err_channel := conn.Channel()
	if err_channel != nil {
		return nil, err_conn
	}

	return channel, nil
}
