package environment

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func RabbitMQ() {
	if err_exchanges := godotenv.Load(".rabbitmq.exchanges"); err_exchanges != nil {
		fmt.Println(".rabbitmq.exchanges is not loaded properly")
		fmt.Println(err_exchanges)
		os.Exit(2)
	}
	if err_queues := godotenv.Load(".rabbitmq.queues"); err_queues != nil {
		fmt.Println(".rabbitmq.queues is not loaded properly")
		fmt.Println(err_queues)
		os.Exit(2)
	}
}
