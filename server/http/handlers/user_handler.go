package handlers

import (
	"booking-car/domain/handlers"
	"booking-car/domain/requests"
	"booking-car/pkg/response"
	"booking-car/usecases"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type UserHandler struct {
	HandlerContract
}

func NewUserHandler(handler HandlerContract) handlers.IUserHandler {
	return &UserHandler{HandlerContract: handler}
}

func (h UserHandler) Browse(ctx *fiber.Ctx) (err error) {
	search := ctx.Query("search")
	orderBy := ctx.Query("order_by")
	sort := ctx.Query("sort")
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	page, _ := strconv.Atoi(ctx.Query("page"))

	uc := usecases.NewUserUseCase(h.UcContract)
	res, pagination, err := uc.GetListWithPagination(search, orderBy, sort, page, limit)

	return response.NewResponse(response.NewResponseWithMeta(res, pagination, err)).Send(ctx)
}

func (h UserHandler) GetUserByID(ctx *fiber.Ctx) (err error) {
	id := ctx.Params("id")

	uc := usecases.NewUserUseCase(h.UcContract)
	res, err := uc.GetByID(id)

	return response.NewResponse(response.NewResponseWithOutMeta(res, err)).Send(ctx)
}

func (h UserHandler) Edit(ctx *fiber.Ctx) (err error) {
	req := new(requests.UserEditRequest)
	id := ctx.Params("id")

	if err := ctx.BodyParser(req); err != nil {
		return response.NewResponse(response.NewResponseBadRequest(err)).Send(ctx)
	}
	if err := h.UcContract.Config.Validator.GetValidator().Struct(req); err != nil {
		return response.NewResponse(response.NewResponseErrorValidator(err.(validator.ValidationErrors), h.UcContract.Config.Validator.GetTranslator())).Send(ctx)
	}

	uc := usecases.NewUserUseCase(h.UcContract)
	res, err := uc.Edit(req, id)

	return response.NewResponse(response.NewResponseWithOutMeta(res, err)).Send(ctx)
}

func (h UserHandler) Add(ctx *fiber.Ctx) (err error) {
	req := new(requests.UserAddRequest)

	if err := ctx.BodyParser(req); err != nil {
		return response.NewResponse(response.NewResponseBadRequest(err)).Send(ctx)
	}
	if err := h.UcContract.Config.Validator.GetValidator().Struct(req); err != nil {
		return response.NewResponse(response.NewResponseErrorValidator(err.(validator.ValidationErrors), h.UcContract.Config.Validator.GetTranslator())).Send(ctx)
	}

	uc := usecases.NewUserUseCase(h.UcContract)
	res, err := uc.Add(req)

	return response.NewResponse(response.NewResponseWithOutMeta(res, err)).Send(ctx)
}

func (h UserHandler) Delete(ctx *fiber.Ctx) (err error) {
	id := ctx.Params("id")

	uc := usecases.NewUserUseCase(h.UcContract)
	err = uc.Delete(id)

	return response.NewResponse(response.NewResponseWithOutMeta(nil, err)).Send(ctx)
}

func (UserHandler) EditDepositAmount(ctx *fiber.Ctx) (err error) {
	panic("implement me")
}
