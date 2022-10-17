package gofiber

import "github.com/gofiber/fiber/v2"

func BodyValidation(c *fiber.Ctx, dto interface{}) error {
	if err := c.BodyParser(dto); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return nil
}
