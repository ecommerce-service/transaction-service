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

type CarColorHandler struct {
	HandlerContract
}

func NewCarColorHandler(handler HandlerContract) handlers.ICarColorHandler {
	return &CarColorHandler{HandlerContract: handler}
}

func (h CarColorHandler) GetListWithPagination(ctx *fiber.Ctx) (err error) {
	search := ctx.Query("search")
	orderBy := ctx.Query("order_by")
	sort := ctx.Query("sort")
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	page, _ := strconv.Atoi(ctx.Query("page"))

	uc := usecases.NewCarColorUseCase(h.UseCaseContract)
	res, pagination, err := uc.GetListWithPagination(search, orderBy, sort, page, limit)

	return response.NewResponse(response.NewResponseWithMeta(res, pagination, err)).Send(ctx)
}

func (h CarColorHandler) GetAll(ctx *fiber.Ctx) error {
	search := ctx.Query("search")

	uc := usecases.NewCarColorUseCase(h.UseCaseContract)
	res, err := uc.GetAll(search)

	return response.NewResponse(response.NewResponseWithOutMeta(res, err, http.StatusOK)).Send(ctx)
}

func (h CarColorHandler) GetByID(ctx *fiber.Ctx) (err error) {
	id := ctx.Params("id")

	uc := usecases.NewCarColorUseCase(h.UseCaseContract)
	res, err := uc.GetByID(id)

	return response.NewResponse(response.NewResponseWithOutMeta(res, err, http.StatusOK)).Send(ctx)
}

func (h CarColorHandler) Edit(ctx *fiber.Ctx) (err error) {
	req := new(requests.CarColorRequest)
	id := ctx.Params("id")

	if err := ctx.BodyParser(req); err != nil {
		return response.NewResponse(response.NewResponseBadRequest(err)).Send(ctx)
	}
	if err := h.UseCaseContract.Config.Validator.GetValidator().Struct(req); err != nil {
		return response.NewResponse(response.NewResponseErrorValidator(err.(validator.ValidationErrors), h.UseCaseContract.Config.Validator.GetTranslator())).Send(ctx)
	}

	uc := usecases.NewCarColorUseCase(h.UseCaseContract)
	res, err := uc.Edit(req, id)

	return response.NewResponse(response.NewResponseWithOutMeta(res, err, http.StatusOK)).Send(ctx)
}

func (h CarColorHandler) Add(ctx *fiber.Ctx) (err error) {
	req := new(requests.CarColorRequest)

	if err := ctx.BodyParser(req); err != nil {
		return response.NewResponse(response.NewResponseBadRequest(err)).Send(ctx)
	}
	if err := h.UseCaseContract.Config.Validator.GetValidator().Struct(req); err != nil {
		return response.NewResponse(response.NewResponseErrorValidator(err.(validator.ValidationErrors), h.UseCaseContract.Config.Validator.GetTranslator())).Send(ctx)
	}

	uc := usecases.NewCarColorUseCase(h.UseCaseContract)
	res, err := uc.Add(req)

	return response.NewResponse(response.NewResponseWithOutMeta(res, err, http.StatusOK)).Send(ctx)
}

func (h CarColorHandler) Delete(ctx *fiber.Ctx) (err error) {
	id := ctx.Params("id")

	uc := usecases.NewCarColorUseCase(h.UseCaseContract)
	err = uc.Delete(id)

	return response.NewResponse(response.NewResponseWithOutMeta(nil, err, http.StatusOK)).Send(ctx)
}
