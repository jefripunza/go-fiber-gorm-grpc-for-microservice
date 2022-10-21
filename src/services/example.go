package services

import (
	"main-service/src/utils/gofiber"

	"github.com/gofiber/fiber/v2"
)

func Example(msg string) gofiber.Response {
	var response gofiber.Response

	response.Code = fiber.StatusOK
	response.Render = fiber.Map{
		"message": msg,
	}

	return response
}
