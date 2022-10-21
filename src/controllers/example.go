package controllers

import (
	"main-service/src/services"
	"main-service/src/utils/gofiber"

	"github.com/gofiber/fiber/v2"
)

func Example(c *fiber.Ctx) error {
	msg := gofiber.GetQueryString(c, "msg")
	return gofiber.ResponseService(c, services.Example(msg))
}
