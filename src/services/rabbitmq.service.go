package services

import (
	"fmt"
	"microservice/src/configs"
	"microservice/src/utils/gofiber"
	"microservice/src/utils/rabbitmq"

	"github.com/gofiber/fiber/v2"
)

func RabbitExample(msg string) gofiber.Response {
	var response gofiber.Response
	example_exchange := configs.ExampleExchange()
	err := rabbitmq.Send(example_exchange, msg)

	if err != nil {
		response.Code = fiber.StatusBadRequest
		response.Render = fiber.Map{
			"message": err.Error(),
		}
	} else {
		response.Code = fiber.StatusOK
		response.Render = fiber.Map{
			"message": fmt.Sprintf("success insert : \"%v\"", msg),
		}
	}

	return response
}
