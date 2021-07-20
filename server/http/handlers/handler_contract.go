package handlers

import (
	"booking-car/usecases"
	"github.com/gofiber/fiber/v2"
)

type HandlerContract struct {
	UseCaseContract *usecases.UseCaseContract
	App             *fiber.App
}
