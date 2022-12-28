package server

import (
	"ecommerce/errors"

	"github.com/gofiber/fiber/v2"
)

var errorHandler = func(ctx *fiber.Ctx, err error) error {

	err = errors.ErrorHandler(ctx, err)
	if err != nil {
		return err
	}
	return nil
}
