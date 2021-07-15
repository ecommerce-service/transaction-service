package boot

import (
	"booking-car/domain/configs"
	"github.com/gofiber/fiber/v2"
)

type Boot struct {
	App    *fiber.App
	Config *configs.Config
}

func NewBoot(app *fiber.App, config *configs.Config) *Boot {
	return &Boot{
		App:    app,
		Config: config,
	}
}
