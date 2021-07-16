package usecases

import (
	"booking-car/domain/requests"
	"booking-car/domain/view_models"
)

type ITransactionUseCase interface {
	GetListForAdminWithPagination(search, orderBy, sort, transactionType string, page, limit int) (res []view_models.TransactionListVm, pagination view_models.PaginationVm, err error)

	GetListForNormalUserWithPagination(search, orderBy, sort, transactionType string, page, limit int) (res []view_models.TransactionListVm, pagination view_models.PaginationVm, err error)

	GetByID(id string) (res view_models.TransactionDetailVm, err error)

	CancelPayment(id string) (res string, err error)

	ConfirmPayment(req *requests.ConfirmPaymentRequest, id string) (res string, err error)

	Add() (res string, err error)

	Count(search, userId, transactionType string) (res int, err error)

	CountAll() (res int, err error)

	GetTransactionNumber() string

	GetTotalAmountAndBuildTransactionDetailRequest(carts []view_models.CartVm) (totalAmount float64,req []requests.TransactionDetailRequest)
}
