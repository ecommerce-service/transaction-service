package usecases

import (
	"booking-car/domain/models"
	"booking-car/domain/requests"
	"booking-car/domain/usecases"
	"booking-car/pkg/functioncaller"
	"booking-car/pkg/logruslogger"
	"booking-car/repository/commands"
	"time"
)

type TransactionDetailUseCase struct {
	*UseCaseContract
}

func NewTransactionDetailUseCase(useCaseContract *UseCaseContract) usecases.ITransactionDetailUseCase {
	return &TransactionDetailUseCase{UseCaseContract: useCaseContract}
}

func (uc TransactionDetailUseCase) Add(req requests.TransactionDetailRequest, transactionId string) (err error) {
	now := time.Now().UTC()

	model := models.NewTransactionDetailModel().SetTransactionId(transactionId).SetCarId(req.CarID).SetCarBrand(req.CarBrand).SetCarType(req.CarType).SetCarColor(req.CarColor).
		SetProductionYear(req.ProductionYear).SetPrice(req.Price).SetQuantity(req.Quantity).SetSubTotal(req.SubTotal).SetCreatedAt(now).SetUpdatedAt(now)
	cmd := commands.NewTransactionDetailCommand(uc.Config.DB, model)
	err = cmd.Add()
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "command-transactionDetail-add")
		return err
	}

	return nil
}

func (uc TransactionDetailUseCase) Store(reqs []requests.TransactionDetailRequest, transactionId string) (err error) {
	for _, req := range reqs {
		err = uc.Add(req, transactionId)
		if err != nil {
			logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-transactionDetail-add")
			return err
		}
	}

	return nil
}
