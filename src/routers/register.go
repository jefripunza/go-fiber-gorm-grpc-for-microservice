package routers

import (
	"microservice/src/configs"
	"microservice/src/routers/endpoints"

	"github.com/gofiber/fiber/v2"
)

func Register(app *fiber.App) {

	// Register All API Routers
	api := app.Group(configs.GetBaseEndpoint())
	endpoints.Basic(api) //> Connect to gRPC Service
	endpoints.RabbitMQ(api)
	endpoints.Products(api)

}
