package apps

import (
	"fmt"
	"log"
	"microservice/src/configs"

	"github.com/gofiber/fiber/v2"
)

// StartServer func for starting a simple server.
func WebServer(app *fiber.App) {

	// Run server.
	log.Fatal(app.Listen(fmt.Sprintf("0.0.0.0:%v", configs.GetPortWebServer())))

}
