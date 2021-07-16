package usecases

import "booking-car/domain/requests"

type ITransactionDetailUseCase interface {

	Add(req requests.TransactionDetailRequest,transactionId string) (err error)

	Store(reqs []requests.TransactionDetailRequest,transactionId string) (err error)
}
