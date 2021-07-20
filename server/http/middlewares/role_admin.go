package middlewares

import (
	"booking-car/pkg/functioncaller"
	"booking-car/pkg/logruslogger"
	"booking-car/pkg/messages"
	"booking-car/pkg/response"
	"booking-car/usecases"
	"errors"
	"github.com/gofiber/fiber/v2"
)

type RoleAdmin struct {
	*usecases.UseCaseContract
}

func NewRoleAdminMiddleware(useCaseContract *usecases.UseCaseContract) RoleAdmin {
	return RoleAdmin{UseCaseContract: useCaseContract}
}

func (m RoleAdmin) Use(ctx *fiber.Ctx) error {
	if m.UseCaseContract.RoleID != 1 {
		logruslogger.Log(logruslogger.WarnLevel, messages.Unauthorized, functioncaller.PrintFuncName(), "uc-roleId")
		return response.NewResponse(response.NewResponseUnauthorized(errors.New(messages.Unauthorized))).Send(ctx)
	}

	return ctx.Next()
}
