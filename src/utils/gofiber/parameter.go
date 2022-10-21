package gofiber

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func GetParameterUnsignedInteger(c *fiber.Ctx, name string) (uint64, error) {
	result, err := strconv.ParseUint(c.Params(name), 10, 64)
	if err != nil {
		err_msg := err.Error()
		if strings.HasPrefix(err_msg, "strconv.ParseUint") {
			err_msg = fmt.Sprintf("%v harus angka!", name)
		}
		return 0, c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err_msg,
		})
	}
	return result, nil
}

func GetParameterInteger(c *fiber.Ctx, name string) (int, error) {
	result, err := strconv.Atoi(c.Params(name))
	if err != nil {
		err_msg := err.Error()
		if strings.HasPrefix(err_msg, "strconv.Atoi") {
			err_msg = fmt.Sprintf("%v harus angka!", name)
		}
		return 0, c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err_msg,
		})
	}
	return result, nil
}

func GetParameterString(c *fiber.Ctx, name string) string {
	return c.Params(name)
}

func GetQueryString(c *fiber.Ctx, name string) string {
	return c.Query(name)
}
