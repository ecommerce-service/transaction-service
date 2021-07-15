package routers

import (
	"booking-car/server/http/handlers"
	"github.com/gofiber/fiber/v2"
)

type RoleRouters struct{
	RouteGroup fiber.Router
	Handler handlers.HandlerContract
}

func NewRoleRouters(routeGroup fiber.Router,handler handlers.HandlerContract) IRouters{
	return &RoleRouters{
		RouteGroup: routeGroup,
		Handler:    handler,
	}
}

func (r RoleRouters) RegisterRouter() {
	handler := handlers.NewRoleHandler(r.Handler)

	roleRouters := r.RouteGroup.Group("/role")
	roleRouters.Get("",handler.BrowseAll)
}

