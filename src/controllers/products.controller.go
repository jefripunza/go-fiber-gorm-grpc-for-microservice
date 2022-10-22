package controllers

import (
	"microservice/src/dto/request"
	"microservice/src/helpers"
	"microservice/src/services"
	"microservice/src/utils/gofiber"

	"github.com/gofiber/fiber/v2"
)

// @Title        Membuat Product Baru
// @Description  menambahkan product baru kedalam database
// @Tags         managements
// @Router       /api/main/v1/add [post]
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        product  body      request.CreateProduct   true  "Content New Product"
// @Success      200      {object}  response.ProductAdd200  "success add product"
// @Failure      400      {object}  response.ProductAdd400  "(code harus di isi !!) (harga harus lebih dari 0 !!)"
func ProductCreateOne(c *fiber.Ctx) error {
	dto := &request.CreateProduct{}
	// validation body is json
	err := gofiber.BodyValidation(c, dto)
	if err != nil {
		return err
	}

	// validation body per key
	if helpers.LengthOfString(dto.Code) == 0 {
		return gofiber.ResponseBadRequest(c, "code harus di isi !!")
	}
	if dto.Price == 0 {
		return gofiber.ResponseBadRequest(c, "harga harus lebih dari 0 !!")
	}

	return gofiber.ResponseService(c, services.ProductCreateOne(dto))
}

func ProductReadAllData(c *fiber.Ctx) error {
	return gofiber.ResponseService(c, services.ProductReadAllData())
}

func ProductById(c *fiber.Ctx) error {
	id, err_id := gofiber.GetParameterInteger(c, "id")
	if err_id != nil {
		return err_id
	}

	return gofiber.ResponseService(c, services.ProductById(id))
}

func ProductByCode(c *fiber.Ctx) error {
	code := gofiber.GetParameterString(c, "code")

	return gofiber.ResponseService(c, services.ProductByCode(code))
}

func ProductUpdateById(c *fiber.Ctx) error {
	dto := &request.UpdateProduct{}
	err := gofiber.BodyValidation(c, dto)
	if err != nil {
		return err
	}

	return gofiber.ResponseService(c, services.ProductUpdateById(dto))
}

func ProductDeleteById(c *fiber.Ctx) error {
	id, err_id := gofiber.GetParameterInteger(c, "id")
	if err_id != nil {
		return err_id
	}

	return gofiber.ResponseService(c, services.ProductDeleteById(id))
}
