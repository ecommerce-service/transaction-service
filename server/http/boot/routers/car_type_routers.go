package routers

import (
	"booking-car/server/http/handlers"
	"booking-car/server/http/middlewares"
	"github.com/gofiber/fiber/v2"
)

type CarTypeRouters struct {
	RouteGroup fiber.Router
	Handler    handlers.HandlerContract
}

func NewCarTypeRouters(routeGroup fiber.Router, handler handlers.HandlerContract) IRouters {
	return &CarTypeRouters{
		RouteGroup: routeGroup,
		Handler:    handler,
	}
}

func (r CarTypeRouters) RegisterRouter() {
	handler := handlers.NewCarTypeHandler(r.Handler)
	jwt := middlewares.NewJwtMiddleware(r.Handler.UseCaseContract)
	adminMiddleware := middlewares.NewRoleAdminMiddleware(r.Handler.UseCaseContract)

	carTypeRouters := r.RouteGroup.Group("/car-type")
	carTypeRouters.Use(jwt.Use)
	carTypeRouters.Use(adminMiddleware.Use)
	carTypeRouters.Get("", handler.GetListWithPagination)
	carTypeRouters.Get("/all", handler.GetAll)
	carTypeRouters.Get("/:id", handler.GetByID)
	carTypeRouters.Put("/:id", handler.Edit)
	carTypeRouters.Post("", handler.Add)
	carTypeRouters.Delete("/:id", handler.Delete)
}
