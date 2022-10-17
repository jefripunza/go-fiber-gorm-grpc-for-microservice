package services

import (
	"main-service/src/remotes"

	"github.com/gofiber/fiber/v2"
)

func BasicAdd(c *fiber.Ctx, a uint64, b uint64) error {
	result, err_msg := remotes.BasicAdd(a, b)
	if err_msg != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err_msg.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"result": result,
	})
}

func BasicMultiply(c *fiber.Ctx, a uint64, b uint64) error {
	result, err_msg := remotes.BasicMultiply(a, b)
	if err_msg != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err_msg.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"result": result,
	})
}
