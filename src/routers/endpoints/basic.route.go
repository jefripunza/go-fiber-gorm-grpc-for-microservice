package endpoints

import (
	"microservice/src/controllers"

	"github.com/gofiber/fiber/v2"
)

func Basic(route fiber.Router) {
	route_v1 := route.Group("/v1")
	route_v1.Get("/add/:a/:b", controllers.BasicAdd)
	route_v1.Get("/multiply/:a/:b", controllers.BasicMultiply)
}
