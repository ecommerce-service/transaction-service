package commands

import (
	"booking-car/domain/models"
	"database/sql"
)

type TransactionCommandMock struct {
	db    *sql.DB
	model *models.Transactions

}

func (c TransactionCommandMock) Add() (res string, err error) {
	statement := `INSERT INTO transactions (user_id,transaction_type,transaction_number,total_amount,created_at,updated_at) ` +
		`VALUES($1,$2,$3,$4,$5,$6) RETURNING id`
	tx, _ := c.db.Begin()

	err = tx.QueryRow(statement, c.model.UserId(), c.model.TransactionType(), c.model.TransactionNumber(), c.model.TotalAmount(),
		c.model.CreatedAt(), c.model.UpdatedAt()).Scan(&res)
	if err != nil {
		tx.Rollback()
		return res, err
	}
	tx.Commit()

	return res, nil
}

func (c TransactionCommandMock) EditPaymentReceived() (err error) {
	statement := `UPDATE transactions set transaction_type=$1,payment_received=$2,updated_at=$3,paid_at=$4 WHERE id=$5`
	tx, _ := c.db.Begin()
	_, err = tx.Exec(statement, c.model.TransactionType(), c.model.PaymentReceived().Float64, c.model.UpdatedAt(), c.model.PaidAt().Time, c.model.Id())
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}

func (c TransactionCommandMock) EditCancelPayment() (res string, err error) {
	statement := `UPDATE transactions set transaction_type=$1,updated_at=$2,canceled_at=$3 WHERE id=$4 RETURNING id`
	tx, _ := c.db.Begin()

	err = tx.QueryRow(statement, c.model.TransactionType(), c.model.UpdatedAt(), c.model.CanceledAt().Time, c.model.Id()).Scan(&res)
	if err != nil {
		tx.Rollback()
		return res, err
	}
	tx.Commit()

	return res, nil
}
