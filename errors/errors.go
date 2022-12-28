package errors

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

var SQLSyntax = 1
var ParseJSONBody = 2
var ResourceNotFound = 3
var BadBodyArgument = 4
var InvalidCredentials = 5
var QueryParamIsRequired = 6
var BadQueryArgument = 7

func (e *Error) Error() string {
	return fmt.Sprintf("%d %s", e.Code, e.Message)
}

type Error struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Target  string   `json:"target,omitempty"`
	Details []*Error `json:"details,omitempty"`
}

func newError(code int, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}
func ErrBadQueryArgument(errs ...error) error {
	return newError(
		BadQueryArgument,
		"Bad Query Argument.",
	).AddDetails(errs...)
}

func ErrQueryParamIsRequired(target string) error {
	return newWithTarget(
		QueryParamIsRequired,
		"Query Param is Required.",
		fmt.Sprintf("'%s' Parametro es requerido.", target),
	)
}

func newWithTarget(code int, message, target string) error {
	return &Error{
		Code:    code,
		Message: message,
		Target:  target,
	}
}

func ErrSQLSyntax(target string) error {
	return newWithTarget(
		SQLSyntax,
		"SQL Syntax",
		target,
	)
}

func ErrParseJSONBody(target string) error {
	return newWithTarget(
		ParseJSONBody,
		"Parse JSON Body Request",
		fmt.Sprintf("Error: %s", target),
	)
}

func ErrResourceNotFound(target string) error {
	return newWithTarget(
		ResourceNotFound,
		"Resource Not Found",
		fmt.Sprintf("Recurso '%s' no encontrado.", target),
	)
}
func (e *Error) AddDetails(errs ...error) error {
	for _, err := range errs {
		e.Details = append(e.Details, err.(*Error))
	}
	return e
}

func ErrBadBodyArgument(errs ...error) error {
	return newError(
		BadBodyArgument,
		"Bad Body Argument.",
	).AddDetails(errs...)
}

func ErrorHandler(ctx *fiber.Ctx, err interface{}) error {

	e, ok := err.(*Error)
	if !ok {
		er := err.(error)
		return ctx.Status(fiber.StatusInternalServerError).SendString(er.Error())
	}
	switch e.Code {
	case ResourceNotFound:
		return ctx.Status(fiber.StatusNotFound).JSON(e)
	case BadBodyArgument,
		ParseJSONBody:

		return ctx.Status(fiber.StatusBadRequest).JSON(e)

	case InvalidCredentials:

		return ctx.Status(fiber.StatusUnauthorized).JSON(e)
	}

	return ctx.Status(fiber.StatusInternalServerError).JSON(e)
}
