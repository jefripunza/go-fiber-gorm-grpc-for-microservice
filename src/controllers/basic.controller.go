package controllers

import (
	"main-service/src/services"
	"main-service/src/utils/gofiber"

	"github.com/gofiber/fiber/v2"
)

func BasicAdd(c *fiber.Ctx) error {
	// Check, if received JSON data is valid.
	a, err_a := gofiber.GetParameterUnsignedInteger(c, "a")
	if err_a != nil {
		return err_a
	}
	b, err_b := gofiber.GetParameterUnsignedInteger(c, "b")
	if err_b != nil {
		return err_b
	}
	return gofiber.ResponseService(c, services.BasicAdd(a, b))
}

func BasicMultiply(c *fiber.Ctx) error {
	// Check, if received JSON data is valid.
	a, err_a := gofiber.GetParameterUnsignedInteger(c, "a")
	if err_a != nil {
		return err_a
	}
	b, err_b := gofiber.GetParameterUnsignedInteger(c, "b")
	if err_b != nil {
		return err_b
	}
	return gofiber.ResponseService(c, services.BasicMultiply(a, b))
}
