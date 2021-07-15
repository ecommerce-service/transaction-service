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

type CarTypeUseCase struct {
	*UseCaseContract
}

func NewCarTypeUseCase(useCaseContract *UseCaseContract) usecases.ICarTypeUseCase {
	return &CarTypeUseCase{UseCaseContract: useCaseContract}
}

func (uc CarTypeUseCase) GetListWithPagination(search, orderBy, sort string, page, limit int) (res []view_models.CarTypeVm, pagination view_models.PaginationVm, err error) {
	q := queries.NewCarTypeQuery(uc.Config.DB.GetDbInstance())
	offset, limit, page, orderBy, sort := uc.SetPaginationParameter(page, limit, orderBy, sort)

	carTypes, err := q.Browse(search, orderBy, sort, limit, offset)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-carType-browse")
		return res, pagination, err
	}
	for _, carType := range carTypes.([]*models.CarTypes) {
		res = append(res, view_models.NewCarTypeVm(carType))
	}

	//set pagination
	totalCount, err := uc.Count(search)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-carType-count")
		return res, pagination, err
	}
	pagination = uc.SetPaginationResponse(page, limit, totalCount)

	return res, pagination, nil
}

func (uc CarTypeUseCase) GetAll(search, brandId string) (res []view_models.CarTypeVm, err error) {
	q := queries.NewCarTypeQuery(uc.Config.DB.GetDbInstance())

	carTypes, err := q.BrowseAll(search, brandId)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-carType-browseAll")
		return res, err
	}
	for _, carType := range carTypes.([]*models.CarTypes) {
		res = append(res, view_models.NewCarTypeVm(carType))
	}

	return res, nil
}

func (uc CarTypeUseCase) GetByID(id string) (res view_models.CarTypeVm, err error) {
	q := queries.NewCarTypeQuery(uc.Config.DB.GetDbInstance())

	carType, err := q.ReadBy("id", "=", id)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-carType-readByID")
		return res, err
	}
	res = view_models.NewCarTypeVm(carType.(*models.CarTypes))

	return res, nil
}

func (uc CarTypeUseCase) Edit(req *requests.CarTypeRequest, id string) (res string, err error) {
	now := time.Now().UTC()

	count, err := uc.CountBy("name", "=", id, req.Name)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-carType-countByName")
		return res, err
	}
	if count > 0 {
		logruslogger.Log(logruslogger.WarnLevel, messages.DataAlreadyExist, functioncaller.PrintFuncName(), "uc-carType-countByName")
		return res, errors.New(messages.DataAlreadyExist)
	}

	model := models.NewCarType().SetName(req.Name).SetBrandId(req.BrandID).SetUpdatedAt(now).SetId(id)
	cmd := commands.NewCarTypeCommand(uc.Config.DB.GetDbInstance(), model)
	res, err = cmd.Edit()
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "cmd-carType-edit")
		return res, err
	}

	return res, nil
}

func (uc CarTypeUseCase) Add(req *requests.CarTypeRequest) (res string, err error) {
	now := time.Now().UTC()

	count, err := uc.CountBy("name", "=", "", req.Name)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-carType-countByName")
		return res, err
	}
	if count > 0 {
		logruslogger.Log(logruslogger.WarnLevel, messages.DataAlreadyExist, functioncaller.PrintFuncName(), "uc-carType-countByName")
		return res, errors.New(messages.DataAlreadyExist)
	}

	model := models.NewCarType().SetName(req.Name).SetBrandId(req.BrandID).SetCreatedAt(now).SetUpdatedAt(now)
	cmd := commands.NewCarTypeCommand(uc.Config.DB.GetDbInstance(), model)
	res, err = cmd.Add()
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "cmd-carType-add")
		return res, err
	}

	return res, nil
}

func (uc CarTypeUseCase) Delete(id string) (err error) {
	now := time.Now().UTC()

	count, err := uc.CountBy("id", "=", "", id)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-carType-countById")
		return err
	}
	if count > 0 {
		model := models.NewCarType().SetUpdatedAt(now).SetDeletedAt(now).SetId(id)
		cmd := commands.NewCarTypeCommand(uc.Config.DB.GetDbInstance(), model)
		_, err = cmd.Delete()
		if err != nil {
			logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "cmd-carType-delete")
			return err
		}
	}

	return nil
}

func (uc CarTypeUseCase) Count(search string) (res int, err error) {
	q := queries.NewCarTypeQuery(uc.Config.DB.GetDbInstance())

	res, err = q.Count(search)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-carType-count")
		return res, err
	}

	return res, nil
}

func (uc CarTypeUseCase) CountBy(column, operator, id string, value interface{}) (res int, err error) {
	q := queries.NewCarTypeQuery(uc.Config.DB.GetDbInstance())

	res, err = q.CountBy(column, operator, id, value)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-carType-countBy")
		return res, err
	}

	return res, nil
}
