package boot

import (
	"booking-car/server/http/boot/routers"
	"booking-car/server/http/handlers"
	"booking-car/usecases"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/xid"
)

func (boot *Boot) RegisterAllRouters() {
	handler := handlers.HandlerContract{
		UcContract: usecases.NewUseCaseContract(xid.New().String(), boot.Config),
	}

	rootRouter := boot.App.Group("/api")
	rootRouter.Get("", func(ctx *fiber.Ctx) error {
		return ctx.Status(200).JSON("it's working")
	})

	roleRouters := routers.NewRoleRouters(rootRouter, handler)
	roleRouters.RegisterRouter()

	userRouters := routers.NewUserRouters(rootRouter, handler)
	userRouters.RegisterRouter()
}
