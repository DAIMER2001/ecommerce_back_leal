package server

import (
	"ecommerce/errors"
	"ecommerce/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (s Server) CreateClient(ctx *fiber.Ctx) error {
	var clientWithValues models.ClientCreate
	err := ctx.BodyParser(&clientWithValues)
	if err != nil || clientWithValues.Password == "" || clientWithValues.Name == "" || clientWithValues.Role == "" {
		return errors.ErrorHandler(ctx, errors.ErrParseJSONBody(err.Error()))
	}

	err = s.clientService.CreateClient(clientWithValues)
	if err != nil {
		return errors.ErrorHandler(ctx, err)
	}
	return ctx.Status(fiber.StatusCreated).JSON(clientWithValues)
}

func (s Server) UpdateClient(ctx *fiber.Ctx) error {
	price := ctx.Query("price", "")
	idClient := ctx.Query("id_client", "")

	accumulation, err := strconv.ParseFloat(price, 32)
	if err != nil {
		return errors.ErrorHandler(ctx, errors.ErrParseJSONBody(err.Error()))
	}

	err = s.clientService.UpdateClient(int(accumulation*0.10), idClient)
	if err != nil {
		return errors.ErrorHandler(ctx, err)
	}

	return ctx.JSON(fiber.StatusOK)
}

func (s Server) FindAllClientById(ctx *fiber.Ctx) error {

	id := ctx.Params("id_client", "")
	response, err := s.clientService.FindByIdClient(id)

	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (s Server) FindAllClients(ctx *fiber.Ctx) error {

	response, err := s.clientService.FindAllClient()

	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (s Server) AuthClient(ctx *fiber.Ctx) error {

	var clientWithValues models.ClientAuth
	err := ctx.BodyParser(&clientWithValues)
	if err != nil {
		return err
	}
	var errs []error

	if clientWithValues.Name == "" {
		errs = append(errs, errors.ErrQueryParamIsRequired("name"))
	}
	if clientWithValues.Password == "" {
		errs = append(errs, errors.ErrQueryParamIsRequired("password"))
	}

	if len(errs) > 0 {
		return errors.ErrBadQueryArgument(errs...)
	}

	response, err := s.clientService.AuthClient(clientWithValues.Name, clientWithValues.Password)

	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (s Server) DeleteClient(ctx *fiber.Ctx) error {
	id := ctx.Query("id_client", "")
	err := s.clientService.DeleteClient(id)

	if err != nil {
		return err
	}
	return ctx.JSON(fiber.StatusOK)
}
