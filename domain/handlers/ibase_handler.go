package handlers

import "github.com/gofiber/fiber/v2"

type IBaseHandler interface {
	Browse(ctx *fiber.Ctx) (err error)

	GetUserByID(ctx *fiber.Ctx) (err error)

	Edit(ctx *fiber.Ctx) (err error)

	Add(ctx *fiber.Ctx) (err error)

	Delete(ctx *fiber.Ctx) (err error)
}