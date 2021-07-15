package routers

import (
	"booking-car/server/http/handlers"
	"github.com/gofiber/fiber/v2"
)

type UserRouters struct{
	RouteGroup fiber.Router
	Handler handlers.HandlerContract
}

func NewUserRouters(routeGroup fiber.Router,handler handlers.HandlerContract) IRouters{
	return &UserRouters{
		RouteGroup: routeGroup,
		Handler:    handler,
	}
}

func (r UserRouters) RegisterRouter() {
	handler := handlers.NewUserHandler(r.Handler)

	userRouters := r.RouteGroup.Group("/user")
	userRouters.Get("",handler.Browse)
	userRouters.Get("/:id",handler.GetUserByID)
	userRouters.Put("/:id",handler.Edit)
	userRouters.Post("",handler.Add)
	userRouters.Delete("/:id",handler.Delete)
}

