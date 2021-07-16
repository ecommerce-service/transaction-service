package models

import (
	"database/sql"
	"time"
)

type Transactions struct {
	id                string
	userId            string
	transactionType   string
	transactionNumber string
	totalAmount       float64
	paymentReceived   sql.NullFloat64
	createdAt         time.Time
	updatedAt         time.Time
	paidAt            sql.NullTime
	canceledAt        sql.NullTime
	deletedAt         sql.NullTime

	transactionDetail string

	User *Users
}

func NewTransactionModel() *Transactions {
	return &Transactions{}
}

func (model *Transactions) Id() string {
	return model.id
}

func (model *Transactions) SetId(id string) *Transactions {
	model.id = id

	return model
}

func (model *Transactions) UserId() string {
	return model.userId
}

func (model *Transactions) SetUserId(userId string) *Transactions {
	model.userId = userId

	return model
}

func (model *Transactions) TransactionType() string {
	return model.transactionType
}

func (model *Transactions) SetTransactionType(transactionType string) *Transactions {
	model.transactionType = transactionType

	return model
}

func (model *Transactions) TransactionNumber() string {
	return model.transactionNumber
}

func (model *Transactions) SetTransactionNumber(transactionNumber string) *Transactions {
	model.transactionNumber = transactionNumber

	return model
}

func (model *Transactions) TotalAmount() float64 {
	return model.totalAmount
}

func (model *Transactions) SetTotalAmount(totalAmount float64) *Transactions {
	model.totalAmount = totalAmount

	return model
}

func (model *Transactions) PaymentReceived() sql.NullFloat64 {
	return model.paymentReceived
}

func (model *Transactions) SetPaymentReceived(paymentReceived float64) *Transactions {
	model.paymentReceived.Float64 = paymentReceived

	return model
}

func (model *Transactions) CreatedAt() time.Time {
	return model.createdAt
}

func (model *Transactions) SetCreatedAt(createdAt time.Time) *Transactions {
	model.createdAt = createdAt

	return model
}

func (model *Transactions) UpdatedAt() time.Time {
	return model.updatedAt
}

func (model *Transactions) SetUpdatedAt(updatedAt time.Time) *Transactions {
	model.updatedAt = updatedAt

	return model
}

func (model *Transactions) PaidAt() sql.NullTime {
	return model.paidAt
}

func (model *Transactions) SetPaidAt(paidAt time.Time) *Transactions {
	model.paidAt.Time = paidAt

	return model
}

func (model *Transactions) CanceledAt() sql.NullTime {
	return model.canceledAt
}

func (model *Transactions) SetCanceledAt(canceledAt time.Time) *Transactions {
	model.canceledAt.Time = canceledAt

	return model
}

func (model *Transactions) DeletedAt() sql.NullTime {
	return model.deletedAt
}

func (model *Transactions) SetDeletedAt(deletedAt time.Time) *Transactions {
	model.deletedAt.Time = deletedAt

	return model
}

func (model *Transactions) TransactionDetail() string {
	return model.transactionDetail
}

const (
	TransactionSelectListStatement = `SELECT t.id,t.user_id,t.transaction_type,t.transaction_number,t.total_amount,t.payment_received,t.created_at,t.updated_at,t.paid_at,t.canceled_at,` +
		`u.first_name,u.last_name,u.email,u.phone_number`
	TransactionSelectDetailStatement = `,ARRAY_TO_STRING(ARRAY_AGG(td.id ||':'|| td.car_id ||':'|| td.car_brand ||':'|| td.car_type ||':'|| td.car_color ||':'|| td.production_year ||':'||` +
		` td.price ||':'|| td.quantity ||':'|| td.sub_total),',')`
	TransactionSelectCountStatement  = `SELECT COUNT(t.id) FROM transactions t`
	TransactionListJoinStatement     = `INNER JOIN users u ON u.id = t.user_id AND u.deleted_at IS NULL`
	TransactionDetailJoinStatement   = `INNER JOIN transaction_details td ON td.transaction_id = t.id AND td.deleted_at IS NULL`
	TransactionDefaultWhereStatement = `WHERE t.deleted_at IS NULL`
	TransactionGroupByStatement      = `GROUP BY t.id,u.id`
)

func (model *Transactions) ScanRows(rows *sql.Rows) (interface{}, error) {
	model.User = NewUserModel()
	err := rows.Scan(&model.id, &model.userId, &model.transactionType, &model.transactionNumber, &model.totalAmount, &model.paymentReceived, &model.createdAt, &model.updatedAt,
		&model.paidAt, &model.canceledAt, &model.User.firstName, &model.User.lastName, &model.User.email, &model.User.phoneNumber)
	if err != nil {
		return model, err
	}

	return model, nil
}

func (model *Transactions) ScanRow(row *sql.Row) (interface{}, error) {
	model.User = NewUserModel()
	err := row.Scan(&model.id, &model.userId, &model.transactionType, &model.transactionNumber, &model.totalAmount, &model.paymentReceived, &model.createdAt, &model.updatedAt,
		&model.paidAt, &model.canceledAt, &model.User.firstName, &model.User.lastName, &model.User.email, &model.User.phoneNumber, &model.transactionDetail)
	if err != nil {
		return model, err
	}

	return model, nil
}
