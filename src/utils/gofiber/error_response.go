package gofiber

import "github.com/gofiber/fiber/v2"

func ResponseBadRequest(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"message": message,
	})
}
