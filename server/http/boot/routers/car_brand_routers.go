package routers

import (
	"booking-car/server/http/handlers"
	"booking-car/server/http/middlewares"
	"github.com/gofiber/fiber/v2"
)

type CarBrandRouters struct {
	RouteGroup fiber.Router
	Handler    handlers.HandlerContract
}

func NewCarBrandRouters(routeGroup fiber.Router, handler handlers.HandlerContract) IRouters {
	return &CarBrandRouters{
		RouteGroup: routeGroup,
		Handler:    handler,
	}
}

func (r CarBrandRouters) RegisterRouter() {
	handler := handlers.NewCarBrandHandler(r.Handler)
	jwt := middlewares.NewJwtMiddleware(r.Handler.UseCaseContract)

	carBrandRouters := r.RouteGroup.Group("/car-brand")
	carBrandRouters.Use(jwt.Use)
	carBrandRouters.Get("", handler.GetListWithPagination)
	carBrandRouters.Get("/all", handler.GetAll)
	carBrandRouters.Get("/:id", handler.GetByID)
	carBrandRouters.Put("/:id", handler.Edit)
	carBrandRouters.Post("", handler.Add)
	carBrandRouters.Delete("/:id", handler.Delete)
}
