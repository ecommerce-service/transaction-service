package commands

import (
	"booking-car/domain/commands"
	"booking-car/domain/models"
	"booking-car/pkg/postgresql"
)

type TransactionCommand struct {
	db    postgresql.IConnection
	model *models.Transactions
}

func NewTransactionCommand(db postgresql.IConnection, model *models.Transactions) commands.ITransactionCommand {
	return &TransactionCommand{
		db:    db,
		model: model,
	}
}

func (c TransactionCommand) Add() (res string, err error) {
	statement := `INSERT INTO transactions (user_id,transaction_type,transaction_number,total_amount,created_at,updated_at) ` +
		`VALUES($1,$2,$3,$4,$5,$6) RETURNING id`

	err = c.db.GetTx().QueryRow(statement, c.model.UserId(), c.model.TransactionType(), c.model.TransactionNumber(), c.model.TotalAmount(),
		c.model.CreatedAt(), c.model.UpdatedAt()).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (c TransactionCommand) EditPaymentReceived() (err error) {
	statement := `UPDATE transactions set transaction_type=$1,payment_received=$2,updated_at=$3,paid_at=$4 WHERE id=$5`

	_, err = c.db.GetTx().Exec(statement, c.model.TransactionType(), c.model.PaymentReceived().Float64, c.model.UpdatedAt(), c.model.PaidAt().Time, c.model.Id())
	if err != nil {
		return err
	}

	return nil
}

func (c TransactionCommand) EditCancelPayment() (res string, err error) {
	statement := `UPDATE transactions set transaction_type=$1,updated_at=$2,canceled_at=$3 WHERE id=$4 RETURNING id`

	err = c.db.GetDbInstance().QueryRow(statement, c.model.TransactionType(), c.model.UpdatedAt(), c.model.CanceledAt().Time, c.model.Id()).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}
