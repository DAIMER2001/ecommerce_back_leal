package server

import (
	"gitlab.com/DevelopmentAveonline/ILA/lib/utils"

	"github.com/gofiber/fiber/v2"
)

var errorHandler = func(ctx *fiber.Ctx, err error) error {

	err = utils.ErrorHandler(ctx, err)
	if err != nil {
		return err
	}
	return nil
}
