package view_models

import (
	"booking-car/domain/models"
	"time"
)

type TransactionListVm struct {
	ID                string            `json:"id"`
	TransactionType   string            `json:"transaction_type"`
	TransactionNumber string            `json:"transaction_number"`
	TotalAmount       float64           `json:"total_amount"`
	PaymentReceived   float64           `json:"payment_received"`
	User              UserTransactionVm `json:"user"`
	CreatedAt         string            `json:"created_at"`
	UpdatedAt         string            `json:"updated_at"`
	PaidAt            string            `json:"paid_at"`
	CanceledAt        string            `json:"canceled_at"`
}

func NewTransactionListVm(model *models.Transactions) TransactionListVm {
	return TransactionListVm{
		ID:                model.Id(),
		TransactionType:   model.TransactionType(),
		TransactionNumber: model.TransactionNumber(),
		TotalAmount:       model.TotalAmount(),
		PaymentReceived:   model.PaymentReceived().Float64,
		User:              NewUserTransactionVm(model.User.Id(), model.User.FirstName(), model.User.LastName(), model.User.Email(), model.User.PhoneNumber()),
		CreatedAt:         model.CreatedAt().Format(time.RFC3339),
		UpdatedAt:         model.UpdatedAt().Format(time.RFC3339),
		PaidAt:            model.PaidAt().Time.Format(time.RFC3339),
		CanceledAt:        model.CanceledAt().Time.Format(time.RFC3339),
	}
}
