package configs

import "os"

func RabbitURL() string {
	return os.Getenv("RABBIT_URL")
}
