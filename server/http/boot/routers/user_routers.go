package routers

import (
	"booking-car/server/http/handlers"
	"booking-car/server/http/middlewares"
	"github.com/gofiber/fiber/v2"
)

type UserRouters struct {
	RouteGroup fiber.Router
	Handler    handlers.HandlerContract
}

func NewUserRouters(routeGroup fiber.Router, handler handlers.HandlerContract) IRouters {
	return &UserRouters{
		RouteGroup: routeGroup,
		Handler:    handler,
	}
}

func (r UserRouters) RegisterRouter() {
	handler := handlers.NewUserHandler(r.Handler)
	jwt := middlewares.NewJwtMiddleware(r.Handler.UseCaseContract)

	userRouters := r.RouteGroup.Group("/user")
	userRouters.Use(jwt.Use)
	userRouters.Get("", handler.GetListWithPagination)
	userRouters.Get("/current",handler.GetCurrentUser)
	userRouters.Get("/:id", handler.GetUserByID)
	userRouters.Put("/:id", handler.Edit)
	userRouters.Post("", handler.Add)
	userRouters.Delete("/:id", handler.Delete)
}
