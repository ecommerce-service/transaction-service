package handlers

import "github.com/gofiber/fiber/v2"

type ICartHandler interface {
	IBaseHandler

	EditQuantity(ctx *fiber.Ctx) (err error)
}
