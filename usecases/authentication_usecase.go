package usecases

import (
	"booking-car/domain/models"
	"booking-car/domain/requests"
	"booking-car/domain/usecases"
	"booking-car/domain/view_models"
	"booking-car/pkg/functioncaller"
	"booking-car/pkg/hashing"
	"booking-car/pkg/logruslogger"
	"booking-car/pkg/messages"
	"booking-car/repository/queries"
	"errors"
	"os"
)

type AuthenticationUseCase struct {
	*UseCaseContract
}

func NewAuthenticationUseCase(useCaseContract *UseCaseContract) usecases.IAuthenticationUseCase {
	return &AuthenticationUseCase{UseCaseContract: useCaseContract}
}

func (uc AuthenticationUseCase) Login(req *requests.LoginRequest) (res view_models.LoginVm, err error) {
	user, valid := uc.ValidateCredential(req.Username, req.Password)
	if !valid {
		logruslogger.Log(logruslogger.WarnLevel, messages.CredentialDoNotMatch, functioncaller.PrintFuncName(), "uc-authentication-validateCredential")
		return res, errors.New(messages.CredentialDoNotMatch)
	}

	//generate jwt payload and encrypted with jwe
	payload := map[string]interface{}{
		"id":     user.ID,
		"roleID": user.Role.ID,
		"role":   user.Role.Name,
	}
	jwePayload, err := uc.Config.Jwe.GenerateJwePayload(payload)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-generate-jwe-payload")
		return res, errors.New(messages.CredentialDoNotMatch)
	}

	//generate jwt token
	res, err = uc.GenerateJWT(req.Username, jwePayload)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-generate-jwt-token")
		return res, errors.New(messages.CredentialDoNotMatch)
	}

	userLoggedIn := map[string]interface{}{
		"email": user.Email,
		"role": map[string]interface{}{
			"id":   user.Role.ID,
			"name": user.Role.Name,
		},
	}
	err = uc.Config.Redis.StoreWithExpired(res.Token, userLoggedIn, os.Getenv("TOKEN_EXP_TIME")+`h`)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-generate-jwt-token")
		return res, err
	}

	return res, nil
}

func (uc AuthenticationUseCase) GenerateJWT(issuer, payload string) (res view_models.LoginVm, err error) {
	res.Token, res.TokenExpiredAt, err = uc.Config.Jwt.GenerateToken(issuer, payload)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "pkg-jwt-getToken")
		return res, err
	}

	res.RefreshToken, res.RefreshTokenExpiredAt, err = uc.Config.Jwt.GenerateRefreshToken(issuer, payload)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "pkg-jwt-getRefreshToken")
		return res, err
	}

	return res, nil
}

func (uc AuthenticationUseCase) ValidateCredential(username, password string) (view_models.UserVm, bool) {
	q := queries.NewUserQuery(uc.Config.DB.GetDbInstance())
	var res view_models.UserVm

	user, err := q.ReadBy("u.username", "=", username)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-user-readByID")
		return res, false
	}

	model := user.(*models.Users)
	isValid := hashing.CheckHashString(password, model.Password())
	res = view_models.NewUserVm(model)

	return res, isValid
}