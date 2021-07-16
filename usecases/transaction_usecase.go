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
	"fmt"
	"time"
)

type TransactionUseCase struct {
	*UseCaseContract
}

func NewTransactionUseCase(useCaseContract *UseCaseContract) usecases.ITransactionUseCase {
	return &TransactionUseCase{UseCaseContract: useCaseContract}
}

func (uc TransactionUseCase) GetListForAdminWithPagination(search, orderBy, sort, transactionType string, page, limit int) (res []view_models.TransactionListVm, pagination view_models.PaginationVm, err error) {
	q := queries.NewTransactionQuery(uc.Config.DB)
	offset, limit, page, orderBy, sort := uc.SetPaginationParameter(page, limit, orderBy, sort)

	transactions, err := q.Browse(search, orderBy, sort, transactionType, limit, offset)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-transaction-browse")
		return res, pagination, err
	}
	for _, transaction := range transactions.([]*models.Transactions) {
		res = append(res, view_models.NewTransactionListVm(transaction))
	}

	//set pagination
	totalCount, err := uc.Count(search, "", transactionType)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-transaction-count")
		return res, pagination, err
	}
	pagination = uc.SetPaginationResponse(page, limit, totalCount)

	return res, pagination, nil
}

func (uc TransactionUseCase) GetListForNormalUserWithPagination(search, orderBy, sort, transactionType string, page, limit int) (res []view_models.TransactionListVm, pagination view_models.PaginationVm, err error) {
	q := queries.NewTransactionQuery(uc.Config.DB)
	offset, limit, page, orderBy, sort := uc.SetPaginationParameter(page, limit, orderBy, sort)

	transactions, err := q.BrowseByUserId(search, orderBy, sort, uc.UserID, transactionType, limit, offset)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-transaction-browse")
		return res, pagination, err
	}
	for _, transaction := range transactions.([]*models.Transactions) {
		res = append(res, view_models.NewTransactionListVm(transaction))
	}

	//set pagination
	totalCount, err := uc.Count(search, uc.UserID, transactionType)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-transaction-count")
		return res, pagination, err
	}
	pagination = uc.SetPaginationResponse(page, limit, totalCount)

	return res, pagination, nil
}

func (uc TransactionUseCase) GetByID(id string) (res view_models.TransactionDetailVm, err error) {
	q := queries.NewTransactionQuery(uc.Config.DB)

	transaction, err := q.ReadBy("t.id", "=", id)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-transaction-readByID")
		return res, err
	}
	res = view_models.NewTransactionDetailVm(transaction.(*models.Transactions))

	return res, nil
}

func (uc TransactionUseCase) CancelPayment(id string) (res string, err error) {
	now := time.Now().UTC()

	model := models.NewTransactionModel().SetId(id).SetTransactionType(CancelTransactionType).SetCanceledAt(now).SetUpdatedAt(now)
	cmd := commands.NewTransactionCommand(uc.Config.DB, model)
	res, err = cmd.EditCancelPayment()
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-transaction-editCancelPayment")
		return res, err
	}

	return res, nil
}

func (uc TransactionUseCase) ConfirmPayment(req *requests.ConfirmPaymentRequest, id string) (res string, err error) {
	now := time.Now().UTC()

	transaction, err := uc.GetByID(id)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-transaction-getById")
		return res, err
	}

	carUc := NewCarUseCase(uc.UseCaseContract)
	for _, transactionDetail := range transaction.Details {
		car, err := carUc.GetByID(transactionDetail.CarID)
		if err != nil {
			logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-car-getById")
			return res, err
		}
		if car.Stock < transactionDetail.Quantity {
			return res, errors.New(messages.NotEnoughStock)
		}
	}

	model := models.NewTransactionModel().SetId(id).SetPaymentReceived(float64(req.PaymentAmount)).SetTransactionType(SuccessTransactionType).SetPaidAt(now).SetUpdatedAt(now)
	cmd := commands.NewTransactionCommand(uc.Config.DB, model)
	err = cmd.EditPaymentReceived()
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-transaction-editPaymentReceived")
		return res, err
	}
	res = id

	for _, transactionDetail := range transaction.Details {
		err = carUc.EditStock(transactionDetail.CarID, transactionDetail.Quantity)
		if err != nil {
			logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-car-editStock")
			return res, err
		}
	}

	if float64(req.PaymentAmount) > transaction.TotalAmount {
		amount := float64(req.PaymentAmount) - transaction.TotalAmount
		userUc := NewUserUseCase(uc.UseCaseContract)
		err = userUc.AddDepositBalance(uc.UserID, amount)
	}

	return res, nil
}

func (uc TransactionUseCase) Add() (res string, err error) {
	now := time.Now().UTC()

	cartUc := NewCartUseCase(uc.UseCaseContract)
	carts, err := cartUc.GetAllByUserId(uc.UserID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-cart-getAllByUserId")
		return res, err
	}
	carUc := NewCarUseCase(uc.UseCaseContract)
	for _, cart := range carts {
		car, err := carUc.GetByID(cart.CarID)
		if err != nil {
			logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-car-getById")
			return res, err
		}
		if car.Stock < cart.Quantity {
			return res, errors.New(messages.NotEnoughStock)
		}
	}

	totalAmount, transactionDetailRequest := uc.GetTotalAmountAndBuildTransactionDetailRequest(carts)
	transactionNumber := uc.GetTransactionNumber()
	model := models.NewTransactionModel().SetTransactionType(DefaultTransactionType).SetTransactionNumber(transactionNumber).
		SetTotalAmount(totalAmount).SetCreatedAt(now).SetUpdatedAt(now).SetUserId(uc.UserID)
	cmd := commands.NewTransactionCommand(uc.Config.DB, model)
	res, err = cmd.Add()
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-transaction-add")
		return res, err
	}

	transactionDetailUc := NewTransactionDetailUseCase(uc.UseCaseContract)
	err = transactionDetailUc.Store(transactionDetailRequest, res)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-transactionDetail-store")
		return res, err
	}

	err = cartUc.DeleteAllByUserId()
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-cart-deleteAllByUserId")
		return res, err
	}

	return res, nil
}

func (uc TransactionUseCase) Count(search, userId, transactionType string) (res int, err error) {
	q := queries.NewTransactionQuery(uc.Config.DB)

	res, err = q.Count(search, userId, transactionType)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-transaction-count")
		return res, err
	}

	return res, nil
}

func (uc TransactionUseCase) CountAll() (res int, err error) {
	q := queries.NewTransactionQuery(uc.Config.DB)

	res, err = q.CountAll()
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-transaction-countAll")
		return res, err
	}

	return res, nil
}

func (uc TransactionUseCase) GetTransactionNumber() (res string) {
	count, err := uc.CountAll()
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-transaction-countAll")
		return res
	}
	res = time.Now().UTC().Format("2006-01-02")
	res += fmt.Sprintf("%04d", count+1)

	return res
}

func (uc TransactionUseCase) GetTotalAmountAndBuildTransactionDetailRequest(carts []view_models.CartVm) (totalAmount float64, req []requests.TransactionDetailRequest) {
	for _, cart := range carts {
		totalAmount += cart.SubTotal
		req = append(req, requests.TransactionDetailRequest{
			CarID:          cart.CarID,
			CarBrand:       cart.CarBrand,
			CarType:        cart.CarType,
			CarColor:       cart.CarColor,
			ProductionYear: cart.ProductionYear,
			Price:          cart.Price,
			Quantity:       cart.Quantity,
			SubTotal:       cart.SubTotal,
		})
	}

	return totalAmount, req
}
