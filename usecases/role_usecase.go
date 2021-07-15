package usecases

import (
	"booking-car/domain/usecases"
	"booking-car/domain/view_models"
	"booking-car/pkg/functioncaller"
	"booking-car/pkg/logruslogger"
	"booking-car/repository/queries"
)

type RoleUseCase struct {
	*UseCaseContract
}

func NewRoleUseCase(useCaseContract *UseCaseContract) usecases.IRoleUseCase {
	return &RoleUseCase{UseCaseContract: useCaseContract}
}

func (uc RoleUseCase) BrowseAll(search string) (res []view_models.RoleVm, err error) {
	q := queries.NewRoleQuery(uc.Config.DB.GetDbInstance())

	roles, err := q.BrowseAll(search)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-role-browse")
		return res, err
	}
	for _, role := range roles {
		res = append(res, view_models.NewRoleVm(role))
	}

	return res, nil
}
