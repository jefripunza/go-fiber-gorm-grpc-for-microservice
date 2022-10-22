package routers

import (
	"main-service/src/configs"
	"main-service/src/routers/endpoints"

	"github.com/gofiber/fiber/v2"
)

func Register(app *fiber.App) {

	// Register All API Routers
	api := app.Group(configs.GetBaseEndpoint())
	endpoints.Basic(api) //> Connect to gRPC Service
	endpoints.RabbitMQ(api)
	endpoints.Products(api)

}
