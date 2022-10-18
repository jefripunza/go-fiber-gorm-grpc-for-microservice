package gofiber

import "github.com/gofiber/fiber/v2"

type Response struct {
	Code int
	Data map[string]interface{}
}

func ResponseService(c *fiber.Ctx, response Response) error {
	return c.Status(response.Code).JSON(response.Data)
}
