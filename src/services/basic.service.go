package services

import (
	"main-service/src/remotes"
	"main-service/src/utils/gofiber"

	"github.com/gofiber/fiber/v2"
)

func BasicAdd(a uint64, b uint64) gofiber.Response {
	var response gofiber.Response
	result, err_msg := remotes.BasicAdd(a, b)

	if err_msg != nil {
		response.Code = fiber.StatusBadRequest
		response.Render = fiber.Map{
			"message": err_msg.Error(),
		}
	} else {
		response.Code = fiber.StatusOK
		response.Render = fiber.Map{
			"result": result,
		}
	}

	return response
}

func BasicMultiply(a uint64, b uint64) gofiber.Response {
	var response gofiber.Response
	result, err_msg := remotes.BasicMultiply(a, b)

	if err_msg != nil {
		response.Code = fiber.StatusBadRequest
		response.Render = fiber.Map{
			"message": err_msg.Error(),
		}
	} else {
		response.Code = fiber.StatusOK
		response.Render = fiber.Map{
			"result": result,
		}
	}

	return response
}
