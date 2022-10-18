package configs

import (
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

// FiberConfig func for configuration Fiber app.
// See: https://docs.gofiber.io/api/fiber#config
func FiberConfig() fiber.Config {
	// Define server settings.
	readTimeoutSecondsCount, _ := strconv.Atoi(os.Getenv("WEBSERVER_TIMEOUT"))

	// Return Fiber configuration.
	return fiber.Config{
		ReadTimeout: time.Second * time.Duration(readTimeoutSecondsCount),
	}
}

func GetPortWebServer() string {
	return os.Getenv("WEBSERVER_PORT")
}

func GetBaseEndpoint() string {
	return os.Getenv("ENDPOINT")
}

func GetEnvironmentSelect() string {
	return os.Getenv("ENVIRONMENT")
}
