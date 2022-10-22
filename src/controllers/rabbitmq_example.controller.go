package controllers

import (
	"microservice/src/services"
	"microservice/src/utils/gofiber"

	"github.com/gofiber/fiber/v2"
)

func RabbitExample(c *fiber.Ctx) error {
	msg := gofiber.GetQueryString(c, "msg")
	return gofiber.ResponseService(c, services.RabbitExample(msg))
}
