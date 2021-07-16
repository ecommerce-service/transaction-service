package handlers

import (
	"booking-car/domain/handlers"
	"booking-car/domain/requests"
	"booking-car/pkg/response"
	"booking-car/usecases"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
)

type CarHandler struct {
	HandlerContract
}

func NewCarHandler(handler HandlerContract) handlers.IBaseHandler {
	return &CarHandler{HandlerContract: handler}
}

func (h CarHandler) GetListWithPagination(ctx *fiber.Ctx) (err error) {
	search := ctx.Query("search")
	orderBy := ctx.Query("order_by")
	sort := ctx.Query("sort")
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	page, _ := strconv.Atoi(ctx.Query("page"))

	uc := usecases.NewCarUseCase(h.UseCaseContract)
	res, pagination, err := uc.GetListWithPagination(search, orderBy, sort, page, limit)

	return response.NewResponse(response.NewResponseWithMeta(res, pagination, err)).Send(ctx)
}

func (h CarHandler) GetByID(ctx *fiber.Ctx) (err error) {
	id := ctx.Params("id")

	uc := usecases.NewCarUseCase(h.UseCaseContract)
	res, err := uc.GetByID(id)

	return response.NewResponse(response.NewResponseWithOutMeta(res, err, http.StatusOK)).Send(ctx)
}

func (h CarHandler) Edit(ctx *fiber.Ctx) (err error) {
	req := new(requests.CarRequest)
	id := ctx.Params("id")

	if err := ctx.BodyParser(req); err != nil {
		return response.NewResponse(response.NewResponseBadRequest(err)).Send(ctx)
	}
	if err := h.UseCaseContract.Config.Validator.GetValidator().Struct(req); err != nil {
		return response.NewResponse(response.NewResponseErrorValidator(err.(validator.ValidationErrors), h.UseCaseContract.Config.Validator.GetTranslator())).Send(ctx)
	}

	uc := usecases.NewCarUseCase(h.UseCaseContract)
	res, err := uc.Edit(req, id)

	return response.NewResponse(response.NewResponseWithOutMeta(res, err, http.StatusOK)).Send(ctx)
}

func (h CarHandler) Add(ctx *fiber.Ctx) (err error) {
	req := new(requests.CarRequest)

	if err := ctx.BodyParser(req); err != nil {
		return response.NewResponse(response.NewResponseBadRequest(err)).Send(ctx)
	}
	if err := h.UseCaseContract.Config.Validator.GetValidator().Struct(req); err != nil {
		return response.NewResponse(response.NewResponseErrorValidator(err.(validator.ValidationErrors), h.UseCaseContract.Config.Validator.GetTranslator())).Send(ctx)
	}

	uc := usecases.NewCarUseCase(h.UseCaseContract)
	res, err := uc.Add(req)

	return response.NewResponse(response.NewResponseWithOutMeta(res, err, http.StatusOK)).Send(ctx)
}

func (h CarHandler) Delete(ctx *fiber.Ctx) (err error) {
	id := ctx.Params("id")

	uc := usecases.NewCarUseCase(h.UseCaseContract)
	err = uc.Delete(id)

	return response.NewResponse(response.NewResponseWithOutMeta(nil, err, http.StatusOK)).Send(ctx)
}
