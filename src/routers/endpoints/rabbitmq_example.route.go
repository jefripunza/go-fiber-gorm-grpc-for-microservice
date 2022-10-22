package endpoints

import (
	"main-service/src/controllers"

	"github.com/gofiber/fiber/v2"
)

func RabbitMQ(route fiber.Router) {
	route_v1 := route.Group("/v1")
	route_v1.Get("/send-queue", controllers.RabbitExample)
}
