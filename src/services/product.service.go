package services

import (
	"fmt"
	"main-service/src/dto/request"
	"main-service/src/handlers"
	"main-service/src/messages"
	"main-service/src/models/repositories"

	"github.com/gofiber/fiber/v2"
)

func ProductCreateOne(c *fiber.Ctx, dto *request.CreateProduct) error {
	var result error
	handlers.Error{
		Try: func() {
			repositories.InsertProduct(dto)
			result = c.JSON(fiber.Map{
				"message": messages.DataSaveSuccess,
			})
		},
		Catch: func(e handlers.Exception) {
			fmt.Printf("Caught error : %v\n", e)
			result = c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": e,
			})
		},
		Finally: func() {},
	}.Do()

	return result
}

func ProductReadAllData(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"data": repositories.ReadProducts(),
	})
}

func ProductById(c *fiber.Ctx, id int) error {
	var result error
	handlers.Error{
		Try: func() {
			result = c.JSON(repositories.ReadProductById(id))
		},
		Catch: func(e handlers.Exception) {
			fmt.Printf("Caught error : %v\n", e)
			result = c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": fmt.Sprintf("%v", e),
			})
		},
		Finally: func() {},
	}.Do()

	return result
}

func ProductByCode(c *fiber.Ctx, code string) error {
	var result error
	handlers.Error{
		Try: func() {
			result = c.JSON(repositories.ReadProductByCode(code))
		},
		Catch: func(e handlers.Exception) {
			fmt.Printf("Caught error : %v\n", e)
			result = c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": fmt.Sprintf("%v", e),
			})
		},
		Finally: func() {},
	}.Do()

	return result
}

func ProductUpdateById(c *fiber.Ctx, dto *request.UpdateProduct) error {
	var result error
	handlers.Error{
		Try: func() {
			repositories.UpdateProduct(dto)
			result = c.JSON(fiber.Map{
				"message": messages.DataUpdateSuccess,
			})
		},
		Catch: func(e handlers.Exception) {
			fmt.Printf("Caught error : %v\n", e)
			result = c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": e,
			})
		},
		Finally: func() {},
	}.Do()

	return result
}

func ProductDeleteById(c *fiber.Ctx, id int) error {
	var result error
	handlers.Error{
		Try: func() {
			repositories.DeleteProductById(id)
			result = c.JSON(fiber.Map{
				"message": messages.DataDeleteSuccess,
			})
		},
		Catch: func(e handlers.Exception) {
			fmt.Printf("Caught error : %v\n", e)
			result = c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": e,
			})
		},
		Finally: func() {},
	}.Do()

	return result
}
