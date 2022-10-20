package configs

import "os"

var (
	RabbitURL = os.Getenv("RABBIT_URL")
)
