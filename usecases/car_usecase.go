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

type CarUseCase struct {
	*UseCaseContract
}

func NewCarUseCase(useCaseContract *UseCaseContract) usecases.ICarUseCase {
	return &CarUseCase{UseCaseContract: useCaseContract}
}

func (uc CarUseCase) GetListWithPagination(search, orderBy, sort string, page, limit int) (res []view_models.CarVm, pagination view_models.PaginationVm, err error) {
	q := queries.NewCarQuery(uc.Config.DB.GetDbInstance())
	offset, limit, page, orderBy, sort := uc.SetPaginationParameter(page, limit, orderBy, sort)

	cars, err := q.Browse(search, orderBy, sort, limit, offset)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-car-browse")
		return res, pagination, err
	}
	for _, car := range cars.([]*models.Cars) {
		res = append(res, view_models.NewCarVm(car))
	}

	//set pagination
	totalCount, err := uc.Count(search)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-car-count")
		return res, pagination, err
	}
	pagination = uc.SetPaginationResponse(page, limit, totalCount)

	return res, pagination, nil
}

func (uc CarUseCase) GetByID(id string) (res view_models.CarVm, err error) {
	q := queries.NewCarQuery(uc.Config.DB.GetDbInstance())

	car, err := q.ReadBy("c.id", "=", id)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-car-readByID")
		return res, err
	}
	res = view_models.NewCarVm(car.(*models.Cars))

	return res, nil
}

func (uc CarUseCase) Edit(req *requests.CarRequest, id string) (res string, err error) {
	now := time.Now().UTC()

	isDuplicate, err := uc.ValidateDuplication(req.ProductionYear, req.CarTypeID, req.CarColorID, id)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-car-validateDuplication")
		return res, err
	}
	if isDuplicate {
		return res, err
	}

	model := models.NewCarModel().SetCarTypeId(req.CarTypeID).SetCarColorId(req.CarColorID).
		SetProductionYear(req.ProductionYear).SetPrice(req.Price).SetStock(req.Stock).SetId(id).
		SetUpdatedAt(now)
	cmd := commands.NewCarCommand(uc.Config.DB, model)
	res, err = cmd.Edit()
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "command-car-edit")
		return res, err
	}

	return res, nil
}

func (uc CarUseCase) EditStock(id string, reduceStock int) (err error) {
	now := time.Now().UTC()

	car, err := uc.GetByID(id)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-car-getByID")
		return err
	}

	model := models.NewCarModel().SetStock(car.Stock - reduceStock).SetUpdatedAt(now).SetId(id)
	cmd := commands.NewCarCommand(uc.Config.DB, model)
	err = cmd.EditStock()
	if err != nil {
		return err
	}

	return nil
}

func (uc CarUseCase) Add(req *requests.CarRequest) (res string, err error) {
	now := time.Now().UTC()

	isDuplicate, err := uc.ValidateDuplication(req.ProductionYear, req.CarTypeID, req.CarColorID, "")
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-car-validateDuplication")
		return res, err
	}
	if isDuplicate {
		return res, err
	}

	model := models.NewCarModel().SetCarTypeId(req.CarTypeID).SetCarColorId(req.CarColorID).
		SetProductionYear(req.ProductionYear).SetPrice(req.Price).SetStock(req.Stock).SetCreatedAt(now).
		SetUpdatedAt(now)
	cmd := commands.NewCarCommand(uc.Config.DB, model)
	res, err = cmd.Add()
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "command-car-add")
		return res, err
	}

	return res, nil
}

func (uc CarUseCase) Delete(id string) (err error) {
	now := time.Now().UTC()

	count, err := uc.CountBy("c.id", "=", "", id)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-car-countByID")
		return err
	}
	if count > 0 {
		model := models.NewCarModel().SetUpdatedAt(now).SetDeletedAt(now).SetId(id)
		cmd := commands.NewCarCommand(uc.Config.DB, model)
		_, err = cmd.Delete()
		if err != nil {
			logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "command-car-delete")
			return err
		}
	}

	return nil
}

func (uc CarUseCase) Count(search string) (res int, err error) {
	q := queries.NewCarQuery(uc.Config.DB.GetDbInstance())

	res, err = q.Count(search)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-car-count")
		return res, err
	}

	return res, nil
}

func (uc CarUseCase) CountBy(column, operator, id string, value interface{}) (res int, err error) {
	q := queries.NewCarQuery(uc.Config.DB.GetDbInstance())

	res, err = q.CountBy(column, operator, id, value)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-car-countBy")
		return res, err
	}

	return res, nil
}

func (uc CarUseCase) ValidateDuplication(productionYear, carTypeId, carColorId, id string) (bool, error) {
	countProductionYear, err := uc.CountBy("c.production_year", "=", id, productionYear)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-car-countByProductionYear")
		return true, err
	}

	countCarType, err := uc.CountBy("c.car_type_id", "=", id, carTypeId)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-car-countByCarTypeId")
		return true, err
	}

	countCarColor, err := uc.CountBy("c.car_color_id", "=", id, carColorId)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-car-countByCarColorId")
		return true, err
	}

	if countProductionYear > 0 && countCarType > 0 && countCarColor > 0 {
		return true, errors.New(messages.DataAlreadyExist)
	}

	return false, nil
}

func (uc CarUseCase) ReduceStock(ids []string, reducedBy int) (err error) {
	for _, id := range ids {
		err = uc.EditStock(id, reducedBy)
		if err != nil {
			logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-car-editStock")
			return err
		}
	}

	return nil
}
