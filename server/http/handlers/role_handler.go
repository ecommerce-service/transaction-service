package handlers

import (
	"booking-car/domain/handlers"
	"booking-car/pkg/response"
	"booking-car/usecases"
	"github.com/gofiber/fiber/v2"
)

type RoleHandler struct{
	Handler HandlerContract
}

func NewRoleHandler(handler HandlerContract) handlers.IRoleHandler{
	return &RoleHandler{Handler: handler}
}

func (h RoleHandler) BrowseAll(ctx *fiber.Ctx) (err error) {
	search := ctx.Query("search")

	uc := usecases.NewRoleUseCase(h.Handler.UcContract)
	res,err := uc.BrowseAll(search)

	return response.NewResponse(response.NewResponseWithOutMeta(res,err)).Send(ctx)
}

