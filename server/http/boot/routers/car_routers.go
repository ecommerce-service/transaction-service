package routers

import (
	"booking-car/server/http/handlers"
	"booking-car/server/http/middlewares"
	"github.com/gofiber/fiber/v2"
)

type CarRouters struct {
	RouteGroup fiber.Router
	Handler    handlers.HandlerContract
}

func NewCarRouters(routeGroup fiber.Router, handler handlers.HandlerContract) IRouters {
	return &CarRouters{
		RouteGroup: routeGroup,
		Handler:    handler,
	}
}

func (r CarRouters) RegisterRouter() {
	handler := handlers.NewCarHandler(r.Handler)
	jwt := middlewares.NewJwtMiddleware(r.Handler.UseCaseContract)
	adminMiddleware := middlewares.NewRoleAdminMiddleware(r.Handler.UseCaseContract)

	carRouters := r.RouteGroup.Group("/car")
	carRouters.Use(jwt.Use)
	carRouters.Get("", handler.GetListWithPagination)
	carRouters.Get("/:id", handler.GetByID)
	carRouters.Use(adminMiddleware.Use).Put("/:id", handler.Edit)
	carRouters.Use(adminMiddleware.Use).Post("", handler.Add)
	carRouters.Use(adminMiddleware.Use).Delete("/:id", handler.Delete, adminMiddleware.Use)
}
