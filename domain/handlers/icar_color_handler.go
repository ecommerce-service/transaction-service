package handlers

import "github.com/gofiber/fiber/v2"

type ICarColorHandler interface {
	IBaseHandler

	GetAll(ctx *fiber.Ctx) error
}
