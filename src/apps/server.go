package apps

import (
	"fmt"
	"log"
	"main-service/src/configs"
	"time"

	"github.com/gofiber/fiber/v2"
)

// StartServer func for starting a simple server.
func WebServer(app *fiber.App) {

	// Run server.
	go func() {
		log.Fatal(app.Listen(fmt.Sprintf("0.0.0.0:%v", configs.GetPortWebServer())))
	}()
	time.Sleep(22)

}
