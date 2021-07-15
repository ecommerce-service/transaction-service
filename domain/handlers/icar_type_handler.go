package handlers

import "github.com/gofiber/fiber/v2"

type ICarTypeHandler interface {
	IBaseHandler

	GetAll(ctx *fiber.Ctx) error
}
