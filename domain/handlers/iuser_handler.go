package handlers

import "github.com/gofiber/fiber/v2"

type IUserHandler interface {
	IBaseHandler

	EditDepositAmount(ctx *fiber.Ctx) (err error)
}
