package apps

import (
	"fmt"
	"microservice/src/configs"
	"time"

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

func RabbitTryConnection() {
	go func() {
		time.Sleep(22)
		_, err_conn := amqp.Dial(configs.RabbitURL())
		if err_conn != nil {
			panic(fmt.Sprintf("Rabbit is not connect !!, %v", err_conn))
		}
	}()
}
