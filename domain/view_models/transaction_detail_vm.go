package view_models

import (
	"booking-car/domain/models"
	"strings"
	"time"
)

type TransactionDetailVm struct {
	ID                string                    `json:"id"`
	TransactionType   string                    `json:"transaction_type"`
	TransactionNumber string                    `json:"transaction_number"`
	TotalAmount       float64                   `json:"total_amount"`
	PaymentReceived   float64                   `json:"payment_received"`
	User              UserTransactionVm         `json:"user"`
	Details           []TransactionDetailListVm `json:"details"`
	CreatedAt         string                    `json:"created_at"`
	UpdatedAt         string                    `json:"updated_at"`
	PaidAt            string                    `json:"paid_at"`
	CanceledAt        string                    `json:"canceled_at"`
}

func NewTransactionDetailVm(model *models.Transactions) TransactionDetailVm {
	var transactionDetailListVm []TransactionDetailListVm
	transactionDetails := strings.Split(model.TransactionDetail(), ",")
	for _, transactionDetail := range transactionDetails {
		details := strings.Split(transactionDetail, ":")
		transactionDetailListVm = append(transactionDetailListVm, NewTransactionDetailListVm(details))
	}

	return TransactionDetailVm{
		ID:                model.Id(),
		TransactionType:   model.TransactionType(),
		TransactionNumber: model.TransactionNumber(),
		TotalAmount:       model.TotalAmount(),
		PaymentReceived:   model.PaymentReceived().Float64,
		User:              NewUserTransactionVm(model.User.Id(), model.User.FirstName(), model.User.LastName(), model.User.Email(), model.User.PhoneNumber()),
		Details:           transactionDetailListVm,
		CreatedAt:         model.CreatedAt().Format(time.RFC3339),
		UpdatedAt:         model.UpdatedAt().Format(time.RFC3339),
		PaidAt:            model.PaidAt().Time.Format(time.RFC3339),
		CanceledAt:        model.CanceledAt().Time.Format(time.RFC3339),
	}
}
