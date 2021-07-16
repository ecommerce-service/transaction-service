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

type CartUseCase struct {
	*UseCaseContract
}

func NewCartUseCase(useCaseContract *UseCaseContract) usecases.ICartUseCase {
	return &CartUseCase{UseCaseContract: useCaseContract}
}

func (uc CartUseCase) GetListWithPagination(search, orderBy, sort string, page, limit int) (res []view_models.CartVm, pagination view_models.PaginationVm, err error) {
	q := queries.NewCartQuery(uc.Config.DB)
	offset, limit, page, orderBy, sort := uc.SetPaginationParameter(page, limit, orderBy, sort)

	cart, err := q.Browse(search, orderBy, sort, uc.UserID, limit, offset)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-cart-browse")
		return res, pagination, err
	}
	for _, cart := range cart.([]*models.Carts) {
		res = append(res, view_models.NewCartVm(cart))
	}

	//set pagination
	totalCount, err := uc.Count(search)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-cart-count")
		return res, pagination, err
	}
	pagination = uc.SetPaginationResponse(page, limit, totalCount)

	return res, pagination, nil
}

func (uc CartUseCase) GetAllByUserId(userId string) (res []view_models.CartVm, err error) {
	q := queries.NewCartQuery(uc.Config.DB)

	cart, err := q.BrowseAllByUser(uc.UserID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-cart-browseAllByUser")
		return res, err
	}
	for _, cart := range cart.([]*models.Carts) {
		res = append(res, view_models.NewCartVm(cart))
	}

	return res, nil
}

func (uc CartUseCase) GetByID(id string) (res view_models.CartVm, err error) {
	q := queries.NewCartQuery(uc.Config.DB)

	cart, err := q.ReadBy("c.id", "=", uc.UserID, id)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-cart-readByID")
		return res, err
	}
	res = view_models.NewCartVm(cart.(*models.Carts))

	return res, nil
}

func (uc CartUseCase) Edit(req *requests.CartRequest, id string) (res string, err error) {
	now := time.Now().UTC()
	model := models.NewCartModel()

	carUc := NewCarUseCase(uc.UseCaseContract)
	car, err := carUc.GetByID(req.CarID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-car-readByID")
		return res, err
	}
	if car.Stock < req.Quantity {
		logruslogger.Log(logruslogger.WarnLevel, messages.NotEnoughStock, functioncaller.PrintFuncName(), "uc-car-readByID")
		return res, errors.New(messages.NotEnoughStock)
	}

	cart, err := uc.GetByID(id)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-cart-getByID")
		return res, err
	}

	if cart.CarID != req.CarID {
		subTotal := float64(req.Quantity) * car.Price
		model.SetCarId(req.CarID).SetCarBrand(car.CarBrand.Name).SetCarType(car.CarType.Name).SetCarColor(car.CarColor.Name).SetProductionYear(car.ProductionYear).
			SetPrice(car.Price).SetQuantity(req.Quantity).SetSubTotal(subTotal).SetUpdatedAt(now).SetId(id)
	} else {
		subTotal := float64(req.Quantity) * cart.Price
		model.SetCarId(req.CarID).SetCarBrand(cart.CarBrand).SetCarType(cart.CarType).SetCarColor(cart.CarColor).SetProductionYear(cart.ProductionYear).
			SetPrice(cart.Price).SetQuantity(req.Quantity).SetSubTotal(subTotal).SetUpdatedAt(now).SetId(id)
	}

	cmd := commands.NewCartCommand(uc.Config.DB, model)
	res, err = cmd.Edit()
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "command-car-edit")
		return res, err
	}

	return res, nil
}

func (uc CartUseCase) EditQuantity(req *requests.CartEditQuantityRequest, id string) (res string, err error) {
	now := time.Now().UTC()

	cart, err := uc.GetByID(id)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-cart-readByID")
		return res, err
	}

	subTotal := float64(req.Quantity) * cart.Price
	model := models.NewCartModel().SetQuantity(req.Quantity).SetSubTotal(subTotal).SetUpdatedAt(now).SetId(id)
	cmd := commands.NewCartCommand(uc.Config.DB, model)
	res, err = cmd.EditQuantity()
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "command-cart-editQuantity")
		return res, err
	}

	return res, nil
}

func (uc CartUseCase) Add(req *requests.CartRequest) (res string, err error) {
	now := time.Now().UTC()

	carUc := NewCarUseCase(uc.UseCaseContract)
	car, err := carUc.GetByID(req.CarID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-car-readByID")
		return res, err
	}
	if car.Stock < req.Quantity {
		logruslogger.Log(logruslogger.WarnLevel, messages.NotEnoughStock, functioncaller.PrintFuncName(), "uc-car-readByID")
		return res, errors.New(messages.NotEnoughStock)
	}

	count, err := uc.CountBy("c.car_id", "=", req.CarID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-cart-countByCarId")
		return res, err
	}

	if count > 0 {
		cart, err := uc.GetBy("c.car_id", "=", req.CarID)
		if err != nil {
			logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-cart-getByCarId")
			return res, err
		}

		editQuantityRequest := requests.CartEditQuantityRequest{Quantity: req.Quantity + cart.Quantity}
		res, err = uc.EditQuantity(&editQuantityRequest, cart.ID)
		if err != nil {
			logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-cart-editQuantity")
			return res, err
		}
	} else {
		subTotal := float64(req.Quantity) * car.Price
		model := models.NewCartModel().SetUserId(uc.UserID).SetCarId(req.CarID).SetCarBrand(car.CarBrand.Name).SetCarType(car.CarType.Name).SetCarColor(car.CarColor.Name).
			SetProductionYear(car.ProductionYear).SetPrice(car.Price).SetQuantity(req.Quantity).SetSubTotal(subTotal).SetCreatedAt(now).SetUpdatedAt(now)
		cmd := commands.NewCartCommand(uc.Config.DB, model)
		res, err = cmd.Add()
		if err != nil {
			logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "command-cart-add")
			return res, err
		}
	}

	return res, nil
}

func (uc CartUseCase) Delete(id string) (err error) {
	now := time.Now().UTC()

	count, err := uc.CountBy("c.id", "=", id)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-cart-countBy")
		return err
	}
	if count > 0 {
		model := models.NewCartModel().SetUpdatedAt(now).SetDeletedAt(now).SetId(id)
		cmd := commands.NewCartCommand(uc.Config.DB, model)
		_, err = cmd.Delete()
		if err != nil {
			logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "command-cart-delete")
			return err
		}
	}

	return nil
}

func (uc CartUseCase) DeleteAllByUserId() (err error) {
	now := time.Now().UTC()

	count,err := uc.CountBy("c.user_id","=",uc.UserID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-cart-countBy")
		return err
	}

	if count > 0 {
		model := models.NewCartModel().SetUpdatedAt(now).SetDeletedAt(now).SetUserId(uc.UserID)
		cmd := commands.NewCartCommand(uc.Config.DB, model)
		err = cmd.DeleteAllByUserId()
		if err != nil {
			logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "command-cart-deleteAllByUserId")
			return err
		}
	}

	return nil
}

func (uc CartUseCase) Count(search string) (res int, err error) {
	q := queries.NewCartQuery(uc.Config.DB)

	res, err = q.Count(search, uc.UserID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-cart-count")
		return res, err
	}

	return res, nil
}

func (uc CartUseCase) CountBy(column, operator string, value interface{}) (res int, err error) {
	q := queries.NewCartQuery(uc.Config.DB)

	res, err = q.CountBy(column, operator, uc.UserID, value)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-cart-countBy")
		return res, err
	}

	return res, nil
}

func (uc CartUseCase) GetBy(column, operator string, value interface{}) (res view_models.CartVm, err error) {
	q := queries.NewCartQuery(uc.Config.DB)

	cart, err := q.ReadBy(column, operator, uc.UserID, value)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-cart-getBy")
		return res, err
	}
	res = view_models.NewCartVm(cart.(*models.Carts))

	return res, nil
}
