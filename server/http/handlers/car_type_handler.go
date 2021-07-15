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

type CarTypeHandler struct {
	HandlerContract
}

func NewCarTypeHandler(handler HandlerContract) handlers.ICarTypeHandler {
	return &CarTypeHandler{HandlerContract: handler}
}

func (h CarTypeHandler) GetListWithPagination(ctx *fiber.Ctx) (err error) {
	search := ctx.Query("search")
	orderBy := ctx.Query("order_by")
	sort := ctx.Query("sort")
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	page, _ := strconv.Atoi(ctx.Query("page"))

	uc := usecases.NewCarTypeUseCase(h.UseCaseContract)
	res, pagination, err := uc.GetListWithPagination(search, orderBy, sort, page, limit)

	return response.NewResponse(response.NewResponseWithMeta(res, pagination, err)).Send(ctx)
}

func (h CarTypeHandler) GetAll(ctx *fiber.Ctx) error {
	search := ctx.Query("search")
	brandId := ctx.Query("brand_id")

	uc := usecases.NewCarTypeUseCase(h.UseCaseContract)
	res, err := uc.GetAll(search, brandId)

	return response.NewResponse(response.NewResponseWithOutMeta(res, err, http.StatusOK)).Send(ctx)
}

func (h CarTypeHandler) GetByID(ctx *fiber.Ctx) (err error) {
	id := ctx.Params("id")

	uc := usecases.NewCarTypeUseCase(h.UseCaseContract)
	res, err := uc.GetByID(id)

	return response.NewResponse(response.NewResponseWithOutMeta(res, err, http.StatusOK)).Send(ctx)
}

func (h CarTypeHandler) Edit(ctx *fiber.Ctx) (err error) {
	req := new(requests.CarTypeRequest)
	id := ctx.Params("id")

	if err := ctx.BodyParser(req); err != nil {
		return response.NewResponse(response.NewResponseBadRequest(err)).Send(ctx)
	}
	if err := h.UseCaseContract.Config.Validator.GetValidator().Struct(req); err != nil {
		return response.NewResponse(response.NewResponseErrorValidator(err.(validator.ValidationErrors), h.UseCaseContract.Config.Validator.GetTranslator())).Send(ctx)
	}

	uc := usecases.NewCarTypeUseCase(h.UseCaseContract)
	res, err := uc.Edit(req, id)

	return response.NewResponse(response.NewResponseWithOutMeta(res, err, http.StatusOK)).Send(ctx)
}

func (h CarTypeHandler) Add(ctx *fiber.Ctx) (err error) {
	req := new(requests.CarTypeRequest)

	if err := ctx.BodyParser(req); err != nil {
		return response.NewResponse(response.NewResponseBadRequest(err)).Send(ctx)
	}
	if err := h.UseCaseContract.Config.Validator.GetValidator().Struct(req); err != nil {
		return response.NewResponse(response.NewResponseErrorValidator(err.(validator.ValidationErrors), h.UseCaseContract.Config.Validator.GetTranslator())).Send(ctx)
	}

	uc := usecases.NewCarTypeUseCase(h.UseCaseContract)
	res, err := uc.Add(req)

	return response.NewResponse(response.NewResponseWithOutMeta(res, err, http.StatusOK)).Send(ctx)
}

func (h CarTypeHandler) Delete(ctx *fiber.Ctx) (err error) {
	id := ctx.Params("id")

	uc := usecases.NewCarTypeUseCase(h.UseCaseContract)
	err = uc.Delete(id)

	return response.NewResponse(response.NewResponseWithOutMeta(nil, err, http.StatusOK)).Send(ctx)
}
