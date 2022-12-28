package server

import (
	"ecommerce/errors"
	"ecommerce/models"

	"github.com/gofiber/fiber/v2"
)

func (s Server) CreateProduct(ctx *fiber.Ctx) error {
	var productWithValues models.Product
	err := ctx.BodyParser(&productWithValues)
	if err != nil {
		return errors.ErrorHandler(ctx, errors.ErrParseJSONBody(err.Error()))
	}

	err = s.productService.CreateProduct(productWithValues)
	if err != nil {
		return errors.ErrorHandler(ctx, err)
	}

	return ctx.Status(fiber.StatusCreated).JSON(productWithValues)
}

func (s Server) UpdateProduct(ctx *fiber.Ctx) error {
	var productWithValues models.Product
	err := ctx.BodyParser(&productWithValues)

	err = s.productService.UpdateProduct(productWithValues)
	if err != nil {
		return errors.ErrorHandler(ctx, err)
	}

	return ctx.JSON(fiber.StatusOK)
}

func (s Server) FindAllProducts(ctx *fiber.Ctx) error {

	response, err := s.productService.FindProducts()

	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (s Server) DeleteProduct(ctx *fiber.Ctx) error {
	id := ctx.Params("id_product", "")
	err := s.productService.DeleteProduct(id)

	if err != nil {
		return err
	}
	return ctx.JSON(fiber.StatusOK)
}
