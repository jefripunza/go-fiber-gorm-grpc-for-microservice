package endpoints

import (
	"github.com/gofiber/fiber/v2"

	"main-service/src/controllers"
)

func Products(route fiber.Router) {
	route_v1 := route.Group("/v1")
	route_v1.Post("/add", controllers.ProductCreateOne)
	route_v1.Get("/", controllers.ProductReadAllData)
	route_v1.Get("/id/:id", controllers.ProductById)
	route_v1.Get("/code/:code", controllers.ProductByCode)
	route_v1.Put("/", controllers.ProductUpdateById)
	route_v1.Delete("/:id", controllers.ProductDeleteById)
}
