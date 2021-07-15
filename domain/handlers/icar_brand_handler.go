package handlers

import "github.com/gofiber/fiber/v2"

type ICarBrandHandler interface {
	IBaseHandler

	GetAll(ctx *fiber.Ctx) error
}
