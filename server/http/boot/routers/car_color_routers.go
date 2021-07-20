package routers

import (
	"booking-car/server/http/handlers"
	"booking-car/server/http/middlewares"
	"github.com/gofiber/fiber/v2"
)

type CarColorRouters struct {
	RouteGroup fiber.Router
	Handler    handlers.HandlerContract
}

func NewCarColorRouters(routeGroup fiber.Router, handler handlers.HandlerContract) IRouters {
	return &CarColorRouters{
		RouteGroup: routeGroup,
		Handler:    handler,
	}
}

func (r CarColorRouters) RegisterRouter() {
	handler := handlers.NewCarColorHandler(r.Handler)
	jwt := middlewares.NewJwtMiddleware(r.Handler.UseCaseContract)
	adminMiddleware := middlewares.NewRoleAdminMiddleware(r.Handler.UseCaseContract)

	carColorRouters := r.RouteGroup.Group("/car-color")
	carColorRouters.Use(jwt.Use)
	carColorRouters.Use(adminMiddleware.Use)
	carColorRouters.Get("", handler.GetListWithPagination)
	carColorRouters.Get("/all", handler.GetAll)
	carColorRouters.Get("/:id", handler.GetByID)
	carColorRouters.Put("/:id", handler.Edit)
	carColorRouters.Post("", handler.Add)
	carColorRouters.Delete("/:id", handler.Delete)
}
