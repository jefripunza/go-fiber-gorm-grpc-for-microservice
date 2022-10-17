package controllers

import (
	"main-service/src/dto/request"
	"main-service/src/services"
	"main-service/src/utils/gofiber"

	"github.com/gofiber/fiber/v2"
)

func ProductCreateOne(c *fiber.Ctx) error {
	dto := &request.CreateProduct{}
	err := gofiber.BodyValidation(c, dto)
	if err != nil {
		return err
	}

	return services.ProductCreateOne(c, dto)
}

func ProductReadAllData(c *fiber.Ctx) error {
	return services.ProductReadAllData(c)
}

func ProductById(c *fiber.Ctx) error {
	id, err_id := gofiber.GetParameterInteger(c, "id")
	if err_id != nil {
		return err_id
	}

	return services.ProductById(c, id)
}

func ProductByCode(c *fiber.Ctx) error {
	code := gofiber.GetParameterString(c, "code")

	return services.ProductByCode(c, code)
}

func ProductUpdateById(c *fiber.Ctx) error {
	dto := &request.UpdateProduct{}
	err := gofiber.BodyValidation(c, dto)
	if err != nil {
		return err
	}

	return services.ProductUpdateById(c, dto)
}

func ProductDeleteById(c *fiber.Ctx) error {
	id, err_id := gofiber.GetParameterInteger(c, "id")
	if err_id != nil {
		return err_id
	}

	return services.ProductDeleteById(c, id)
}
