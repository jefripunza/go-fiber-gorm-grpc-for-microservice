package services

import (
	"fmt"
	"main-service/src/dto/request"
	"main-service/src/messages"
	"main-service/src/models/repositories"
	"main-service/src/utils/gofiber"
	"main-service/src/utils/handlers"

	"github.com/gofiber/fiber/v2"
)

func ProductCreateOne(dto *request.CreateProduct) gofiber.Response {
	var response gofiber.Response

	handlers.Error{
		Try: func() {
			repositories.InsertProducts(dto)
			response.Code = fiber.StatusOK
			response.Render = fiber.Map{
				"message": messages.DataSaveSuccess,
			}
		},
		Catch: func(e handlers.Exception) {
			fmt.Printf("Caught error : %v\n", e)
			response.Code = fiber.StatusBadRequest
			response.Render = fiber.Map{
				"error": e,
			}
		},
		Finally: func() {},
	}.Do()

	return response
}

func ProductReadAllData() gofiber.Response {
	var response gofiber.Response

	response.Code = fiber.StatusOK
	response.Render = fiber.Map{
		"data": repositories.ReadProducts(),
	}

	return response
}

func ProductById(id int) gofiber.Response {
	var response gofiber.Response

	handlers.Error{
		Try: func() {
			response.Code = fiber.StatusOK
			response.Render = fiber.Map{
				"data": repositories.ReadProductsById(id),
			}
		},
		Catch: func(e handlers.Exception) {
			fmt.Printf("Caught error : %v\n", e)
			response.Code = fiber.StatusBadRequest
			response.Render = fiber.Map{
				"error": e,
			}
		},
		Finally: func() {},
	}.Do()

	return response
}

func ProductByCode(code string) gofiber.Response {
	var response gofiber.Response

	handlers.Error{
		Try: func() {
			response.Code = fiber.StatusOK
			response.Render = fiber.Map{
				"data": repositories.ReadProductsByCode(code),
			}
		},
		Catch: func(e handlers.Exception) {
			fmt.Printf("Caught error : %v\n", e)
			response.Code = fiber.StatusBadRequest
			response.Render = fiber.Map{
				"error": e,
			}
		},
		Finally: func() {},
	}.Do()

	return response
}

func ProductUpdateById(dto *request.UpdateProduct) gofiber.Response {
	var response gofiber.Response

	handlers.Error{
		Try: func() {
			repositories.UpdateProducts(dto)
			response.Code = fiber.StatusOK
			response.Render = fiber.Map{
				"message": messages.DataUpdateSuccess,
			}
		},
		Catch: func(e handlers.Exception) {
			fmt.Printf("Caught error : %v\n", e)
			response.Code = fiber.StatusBadRequest
			response.Render = fiber.Map{
				"error": e,
			}
		},
		Finally: func() {},
	}.Do()

	return response
}

func ProductDeleteById(id int) gofiber.Response {
	var response gofiber.Response

	handlers.Error{
		Try: func() {
			repositories.DeleteProductsById(id)
			response.Code = fiber.StatusOK
			response.Render = fiber.Map{
				"message": messages.DataDeleteSuccess,
			}
		},
		Catch: func(e handlers.Exception) {
			fmt.Printf("Caught error : %v\n", e)
			response.Code = fiber.StatusBadRequest
			response.Render = fiber.Map{
				"error": e,
			}
		},
		Finally: func() {},
	}.Do()

	return response
}
