package usecases

import (
	"booking-car/domain/requests"
	"booking-car/domain/view_models"
)

type IAuthenticationUseCase interface {
	Login(req *requests.LoginRequest) (res view_models.LoginVm, err error)

	GenerateJWT(issuer, payload string) (res view_models.LoginVm, err error)

	ValidateCredential(username, password string) (view_models.UserVm, bool)
}
