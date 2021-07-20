package commands

import (
	"booking-car/domain/models"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAddTransaction(t *testing.T) {
	db, mock := NewMock()
	defer func() {
		db.Close()
	}()

	now := time.Now().UTC()
	model := models.NewTransactionModel().SetUserId(gofakeit.UUID()).SetTransactionType("on_going").SetTransactionNumber(gofakeit.UUID()).
		SetTotalAmount(gofakeit.Price(1000000000, 120000000)).SetCreatedAt(now).SetUpdatedAt(now)
	id := gofakeit.UUID()

	cmd := TransactionCommandMock{
		db:    db,
		model: model,
	}
	rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
	statement := `INSERT INTO transactions (user_id,transaction_type,transaction_number,total_amount,created_at,updated_at) ` +
		`VALUES($1,$2,$3,$4,$5,$6) RETURNING id`
	mock.ExpectBegin()
	mock.ExpectQuery(statement).WithArgs(model.UserId(), model.TransactionType(), model.TransactionNumber(), model.TotalAmount(), model.CreatedAt(),
		model.UpdatedAt()).WillReturnRows(rows)
	mock.ExpectCommit()
	res, err := cmd.Add()

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestAddTransactionError(t *testing.T) {
	db, mock := NewMock()
	defer func() {
		db.Close()
	}()

	now := time.Now().UTC()
	model := models.NewTransactionModel().SetTransactionType("on_going").SetTransactionNumber(gofakeit.UUID()).SetCreatedAt(now).SetUpdatedAt(now)

	cmd := TransactionCommandMock{
		db:    db,
		model: model,
	}
	statement := `INSERT INTO transactions (user_id,transaction_type,transaction_number,total_amount,created_at,updated_at) ` +
		`VALUES($1,$2,$3,$4,$5,$6) RETURNING id`
	mock.ExpectBegin()
	mock.ExpectQuery(statement).WithArgs(model.UserId(), model.TransactionType(), model.TransactionNumber(), model.TotalAmount(), model.CreatedAt(),
		model.UpdatedAt()).WillReturnError(sql.ErrNoRows)
	mock.ExpectRollback()
	res, err := cmd.Add()

	assert.Error(t, err)
	assert.Empty(t, res)
}

func TestEditPaymentReceived(t *testing.T) {
	db, mock := NewMock()
	defer func() {
		db.Close()
	}()

	now := time.Now().UTC()
	id := gofakeit.UUID()
	paymentReceived := gofakeit.Price(100000000,130000000)
	model := models.NewTransactionModel().SetTransactionType("success").SetPaymentReceived(paymentReceived).
		SetUpdatedAt(now).SetPaidAt(now).SetId(id)

	cmd := TransactionCommandMock{
		db:    db,
		model: model,
	}
	statement := `UPDATE transactions set transaction_type=$1,payment_received=$2,updated_at=$3,paid_at=$4 WHERE id=$5`
	mock.ExpectBegin()
	mock.ExpectExec(statement).WithArgs(model.TransactionType(), model.PaymentReceived().Float64, model.UpdatedAt(), model.PaidAt().Time, model.Id()).
		WillReturnResult(sqlmock.NewResult(0,1))
	mock.ExpectCommit()
	err := cmd.EditPaymentReceived()

	assert.NoError(t, err)
}

func TestEditPaymentReceivedError(t *testing.T) {
	db, mock := NewMock()
	defer func() {
		db.Close()
	}()

	now := time.Now().UTC()
	paymentReceived := gofakeit.Price(100000000,130000000)
	model := models.NewTransactionModel().SetTransactionType("success").SetPaymentReceived(paymentReceived).
		SetUpdatedAt(now).SetPaidAt(now)

	cmd := TransactionCommandMock{
		db:    db,
		model: model,
	}
	statement := `UPDATE transactions set transaction_type=$1,payment_received=$2,updated_at=$3,paid_at=$4 WHERE id=$5`
	mock.ExpectBegin()
	mock.ExpectExec(statement).WithArgs(model.TransactionType(), model.PaymentReceived().Float64, model.UpdatedAt(), model.PaidAt().Time, model.Id()).
		WillReturnResult(sqlmock.NewResult(0,0)).WillReturnError(sql.ErrNoRows)
	mock.ExpectRollback()
	err := cmd.EditPaymentReceived()

	assert.Error(t, err)
}

func TestPaymentCancel(t *testing.T) {
	db, mock := NewMock()
	defer func() {
		db.Close()
	}()

	now := time.Now().UTC()
	id := gofakeit.UUID()
	model := models.NewTransactionModel().SetTransactionType("canceled").SetUpdatedAt(now).SetCanceledAt(now).SetId(id)

	cmd := TransactionCommandMock{
		db:    db,
		model: model,
	}
	rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
	statement := `UPDATE transactions set transaction_type=$1,updated_at=$2,canceled_at=$3 WHERE id=$4 RETURNING id`
	mock.ExpectBegin()
	mock.ExpectQuery(statement).WithArgs(model.TransactionType(),model.UpdatedAt(),model.CanceledAt().Time,model.Id()).WillReturnRows(rows)
	mock.ExpectCommit()
	res, err := cmd.EditCancelPayment()

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestPaymentCancelError(t *testing.T) {
	db, mock := NewMock()
	defer func() {
		db.Close()
	}()

	now := time.Now().UTC()
	model := models.NewTransactionModel().SetTransactionType("canceled").SetUpdatedAt(now).SetCanceledAt(now)

	cmd := TransactionCommandMock{
		db:    db,
		model: model,
	}
	statement := `UPDATE transactions set transaction_type=$1,updated_at=$2,canceled_at=$3 WHERE id=$4 RETURNING id`
	mock.ExpectBegin()
	mock.ExpectQuery(statement).WithArgs(model.TransactionType(),model.UpdatedAt(),model.CanceledAt().Time,model.Id()).WillReturnError(sql.ErrNoRows)
	mock.ExpectRollback()
	res, err := cmd.EditCancelPayment()

	assert.Error(t, err)
	assert.Empty(t, res)
}
