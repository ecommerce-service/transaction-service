package usecases

import (
	"booking-car/domain/models"
	"booking-car/domain/requests"
	"booking-car/domain/usecases"
	"booking-car/domain/view_models"
	"booking-car/pkg/functioncaller"
	"booking-car/pkg/logruslogger"
	"booking-car/pkg/messages"
	"booking-car/repository/commands"
	"booking-car/repository/queries"
	"errors"
	"time"
)

type CarBrandUseCase struct {
	*UseCaseContract
}

func NewCarBrandUseCase(useCaseContract *UseCaseContract) usecases.ICarBrandUseCase {
	return &CarBrandUseCase{UseCaseContract: useCaseContract}
}

func (uc CarBrandUseCase) GetListWithPagination(search, orderBy, sort string, page, limit int) (res []view_models.CarBrandVm, pagination view_models.PaginationVm, err error) {
	q := queries.NewCarBrandQuery(uc.Config.DB.GetDbInstance())
	offset, limit, page, orderBy, sort := uc.SetPaginationParameter(page, limit, orderBy, sort)

	carBrands, err := q.Browse(search, orderBy, sort, limit, offset)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-carBrand-browse")
		return res, pagination, err
	}
	for _, carBrand := range carBrands.([]*models.CarBrands) {
		res = append(res, view_models.NewCarBrandVm(carBrand))
	}

	//set pagination
	totalCount, err := uc.Count(search)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-carBrand-count")
		return res, pagination, err
	}
	pagination = uc.SetPaginationResponse(page, limit, totalCount)

	return res, pagination, nil
}

func (uc CarBrandUseCase) GetAll(search string) (res []view_models.CarBrandVm, err error) {
	q := queries.NewCarBrandQuery(uc.Config.DB.GetDbInstance())

	carBrands, err := q.BrowseAll(search)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-carBrand-browseAll")
		return res, err
	}
	for _, carBrand := range carBrands.([]*models.CarBrands) {
		res = append(res, view_models.NewCarBrandVm(carBrand))
	}

	return res, nil
}

func (uc CarBrandUseCase) GetByID(id string) (res view_models.CarBrandVm, err error) {
	q := queries.NewCarBrandQuery(uc.Config.DB.GetDbInstance())

	carBrand, err := q.ReadBy("id", "=", id)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-carBrand-readByID")
		return res, err
	}
	res = view_models.NewCarBrandVm(carBrand.(*models.CarBrands))

	return res, nil
}

func (uc CarBrandUseCase) Edit(req *requests.CarBrandRequest, id string) (res string, err error) {
	now := time.Now().UTC()

	count, err := uc.CountBy("name", "=", id, req.Name)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-carBrand-countByName")
		return res, err
	}
	if count > 0 {
		logruslogger.Log(logruslogger.WarnLevel, messages.DataAlreadyExist, functioncaller.PrintFuncName(), "uc-carBrand-countByName")
		return res, errors.New(messages.DataAlreadyExist)
	}

	model := models.NewCarBrandModel().SetName(req.Name).SetUpdatedAt(now).SetId(id)
	cmd := commands.NewCarBrandCommand(uc.Config.DB.GetDbInstance(), model)
	res, err = cmd.Edit()
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "cmd-carBrand-edit")
		return res, err
	}

	return res, nil
}

func (uc CarBrandUseCase) Add(req *requests.CarBrandRequest) (res string, err error) {
	now := time.Now().UTC()

	count, err := uc.CountBy("name", "=", "", req.Name)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-carBrand-countByName")
		return res, err
	}
	if count > 0 {
		logruslogger.Log(logruslogger.WarnLevel, messages.DataAlreadyExist, functioncaller.PrintFuncName(), "uc-carBrand-countByName")
		return res, errors.New(messages.DataAlreadyExist)
	}

	model := models.NewCarBrandModel().SetName(req.Name).SetCreatedAt(now).SetUpdatedAt(now)
	cmd := commands.NewCarBrandCommand(uc.Config.DB.GetDbInstance(), model)
	res, err = cmd.Add()
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "cmd-carBrand-add")
		return res, err
	}

	return res, nil
}

func (uc CarBrandUseCase) Delete(id string) (err error) {
	now := time.Now().UTC()

	count, err := uc.CountBy("id", "=", "", id)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-carBrand-countById")
		return err
	}
	if count > 0 {
		model := models.NewCarBrandModel().SetUpdatedAt(now).SetDeletedAt(now).SetId(id)
		cmd := commands.NewCarBrandCommand(uc.Config.DB.GetDbInstance(), model)
		_, err = cmd.Delete()
		if err != nil {
			logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "cmd-carBrand-delete")
			return err
		}
	}

	return nil
}

func (uc CarBrandUseCase) Count(search string) (res int, err error) {
	q := queries.NewCarBrandQuery(uc.Config.DB.GetDbInstance())

	res, err = q.Count(search)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-carBrand-count")
		return res, err
	}

	return res, nil
}

func (uc CarBrandUseCase) CountBy(column, operator, id string, value interface{}) (res int, err error) {
	q := queries.NewCarBrandQuery(uc.Config.DB.GetDbInstance())

	res, err = q.CountBy(column, operator, id, value)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-carBrand-countBy")
		return res, err
	}

	return res, nil
}
