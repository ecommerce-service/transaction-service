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
		UseCaseContract: usecases.NewUseCaseContract(xid.New().String(), boot.Config),
	}

	//init route group
	rootRouter := boot.App.Group("/api")
	//check health
	rootRouter.Get("", func(ctx *fiber.Ctx) error {
		return ctx.Status(200).JSON("it's working")
	})

	//authentication routers
	authenticationRouters := routers.NewAuthenticationRouters(rootRouter,handler)
	authenticationRouters.RegisterRouter()

	//role routers
	roleRouters := routers.NewRoleRouters(rootRouter, handler)
	roleRouters.RegisterRouter()

	//user routers
	userRouters := routers.NewUserRouters(rootRouter, handler)
	userRouters.RegisterRouter()

	//car brand routers
	carBrandRouters := routers.NewCarBrandRouters(rootRouter,handler)
	carBrandRouters.RegisterRouter()
}
