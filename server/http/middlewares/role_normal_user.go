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

type RoleNormalUser struct {
	*usecases.UseCaseContract
}

func NewRoleNormalUser(useCaseContract *usecases.UseCaseContract) RoleNormalUser {
	return RoleNormalUser{UseCaseContract: useCaseContract}
}

func (m RoleNormalUser) Use(ctx *fiber.Ctx) error {
	if m.UseCaseContract.RoleID != 2 {
		logruslogger.Log(logruslogger.WarnLevel, messages.Unauthorized, functioncaller.PrintFuncName(), "uc-roleId")
		return response.NewResponse(response.NewResponseUnauthorized(errors.New(messages.Unauthorized))).Send(ctx)
	}

	return ctx.Next()
}
