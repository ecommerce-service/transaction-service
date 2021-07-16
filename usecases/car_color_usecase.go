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

type CarColorUseCase struct {
	*UseCaseContract
}

func NewCarColorUseCase(useCaseContract *UseCaseContract) usecases.ICarColorUseCase {
	return &CarColorUseCase{UseCaseContract: useCaseContract}
}

func (uc CarColorUseCase) GetListWithPagination(search, orderBy, sort string, page, limit int) (res []view_models.CarColorVm, pagination view_models.PaginationVm, err error) {
	q := queries.NewCarColorQuery(uc.Config.DB.GetDbInstance())
	offset, limit, page, orderBy, sort := uc.SetPaginationParameter(page, limit, orderBy, sort)

	carColors, err := q.Browse(search, orderBy, sort, limit, offset)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-carColor-browse")
		return res, pagination, err
	}
	for _, carColor := range carColors.([]*models.CarColors) {
		res = append(res, view_models.NewCarColorVm(carColor))
	}

	//set pagination
	totalCount, err := uc.Count(search)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-carColor-count")
		return res, pagination, err
	}
	pagination = uc.SetPaginationResponse(page, limit, totalCount)

	return res, pagination, nil
}

func (uc CarColorUseCase) GetAll(search string) (res []view_models.CarColorVm, err error) {
	q := queries.NewCarColorQuery(uc.Config.DB.GetDbInstance())

	carColors, err := q.BrowseAll(search)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-carColor-browseAll")
		return res, err
	}
	for _, carColor := range carColors.([]*models.CarColors) {
		res = append(res, view_models.NewCarColorVm(carColor))
	}

	return res, nil
}

func (uc CarColorUseCase) GetByID(id string) (res view_models.CarColorVm, err error) {
	q := queries.NewCarColorQuery(uc.Config.DB.GetDbInstance())

	carColor, err := q.ReadBy("id", "=", id)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-carColor-readByID")
		return res, err
	}
	res = view_models.NewCarColorVm(carColor.(*models.CarColors))

	return res, nil
}

func (uc CarColorUseCase) Edit(req *requests.CarColorRequest, id string) (res string, err error) {
	now := time.Now().UTC()

	count, err := uc.CountBy("hex_code", "=", id, req.HexCode)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-carColor-countByName")
		return res, err
	}
	if count > 0 {
		logruslogger.Log(logruslogger.WarnLevel, messages.DataAlreadyExist, functioncaller.PrintFuncName(), "uc-carColor-countByName")
		return res, errors.New(messages.DataAlreadyExist)
	}

	model := models.NewCarColorModel().SetName(req.Name).SetHexCode(req.HexCode).SetUpdatedAt(now).SetId(id)
	cmd := commands.NewCarColorCommand(uc.Config.DB.GetDbInstance(), model)
	res, err = cmd.Edit()
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "cmd-carColor-edit")
		return res, err
	}

	return res, nil
}

func (uc CarColorUseCase) Add(req *requests.CarColorRequest) (res string, err error) {
	now := time.Now().UTC()

	count, err := uc.CountBy("hex_code", "=", "", req.HexCode)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-carColor-countByName")
		return res, err
	}
	if count > 0 {
		logruslogger.Log(logruslogger.WarnLevel, messages.DataAlreadyExist, functioncaller.PrintFuncName(), "uc-carColor-countByName")
		return res, errors.New(messages.DataAlreadyExist)
	}

	model := models.NewCarColorModel().SetName(req.Name).SetHexCode(req.HexCode).SetCreatedAt(now).SetUpdatedAt(now)
	cmd := commands.NewCarColorCommand(uc.Config.DB.GetDbInstance(), model)
	res, err = cmd.Add()
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "cmd-carColor-edit")
		return res, err
	}

	return res, nil
}

func (uc CarColorUseCase) Delete(id string) (err error) {
	now := time.Now().UTC()

	count, err := uc.CountBy("id", "=", "", id)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-carColor-countById")
		return err
	}
	if count > 0 {
		model := models.NewCarColorModel().SetUpdatedAt(now).SetDeletedAt(now).SetId(id)
		cmd := commands.NewCarColorCommand(uc.Config.DB.GetDbInstance(), model)
		_, err = cmd.Delete()
		if err != nil {
			logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "cmd-carColor-delete")
			return err
		}
	}

	return nil
}

func (uc CarColorUseCase) Count(search string) (res int, err error) {
	q := queries.NewCarColorQuery(uc.Config.DB.GetDbInstance())

	res, err = q.Count(search)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-carColor-count")
		return res, err
	}

	return res, nil
}

func (uc CarColorUseCase) CountBy(column, operator, id string, value interface{}) (res int, err error) {
	q := queries.NewCarColorQuery(uc.Config.DB.GetDbInstance())

	res, err = q.CountBy(column, operator, id, value)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-carColor-countBy")
		return res, err
	}

	return res, nil
}
