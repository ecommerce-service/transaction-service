package commands

import (
	"booking-car/domain/models"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
	"time"
)

func TestAddTransactionDetail(t *testing.T){
	db, mock := NewMock()
	defer func() {
		db.Close()
	}()

	now := time.Now().UTC()
	quantity := int(gofakeit.Int8())
	price := gofakeit.Price(100000000,130000000)
	subTotal := float64(quantity) * price
	model := models.NewTransactionDetailModel().SetTransactionId(gofakeit.UUID()).SetCarId(gofakeit.UUID()).SetCarBrand(gofakeit.Car().Brand).SetCarType(gofakeit.CarType()).
		SetCarColor(gofakeit.Color()).SetProductionYear(strconv.Itoa(gofakeit.Year())).SetPrice(price).SetQuantity(quantity).
		SetSubTotal(subTotal).SetCreatedAt(now).SetUpdatedAt(now)

	cmd := TransactionDetailCommandMock{
		db:    db,
		model: model,
	}
	statement := `INSERT INTO transaction_details(transaction_id,car_id,car_brand,car_type,car_color,production_year,price,quantity,sub_total,created_at,updated_at) `+
		`VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)`
	mock.ExpectBegin()
	mock.ExpectExec(statement).WithArgs(model.TransactionId(),model.CarId(),model.CarBrand(),model.CarType(),model.CarColor(),model.ProductionYear(),
		model.Price(),model.Quantity(),model.SubTotal(),model.CreatedAt(),model.UpdatedAt()).WillReturnResult(sqlmock.NewResult(0,1))
	mock.ExpectCommit()
	err := cmd.Add()

	assert.NoError(t, err)
}

func TestAddTransactionDetailError(t *testing.T){
	db, mock := NewMock()
	defer func() {
		db.Close()
	}()

	now := time.Now().UTC()
	quantity := int(gofakeit.Int8())
	price := gofakeit.Price(100000000,130000000)
	subTotal := float64(quantity) * price
	model := models.NewTransactionDetailModel().SetCarBrand(gofakeit.Car().Brand).SetCarType(gofakeit.CarType()).
		SetCarColor(gofakeit.Color()).SetProductionYear(strconv.Itoa(gofakeit.Year())).SetPrice(price).SetQuantity(quantity).
		SetSubTotal(subTotal).SetCreatedAt(now).SetUpdatedAt(now)

	cmd := TransactionDetailCommandMock{
		db:    db,
		model: model,
	}
	statement := `INSERT INTO transaction_details(transaction_id,car_id,car_brand,car_type,car_color,production_year,price,quantity,sub_total,created_at,updated_at) `+
		`VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)`
	mock.ExpectBegin()
	mock.ExpectExec(statement).WithArgs(model.TransactionId(),model.CarId(),model.CarBrand(),model.CarType(),model.CarColor(),model.ProductionYear(),
		model.Price(),model.Quantity(),model.SubTotal(),model.CreatedAt(),model.UpdatedAt()).WillReturnResult(sqlmock.NewResult(0,0)).
		WillReturnError(sql.ErrNoRows)
	mock.ExpectRollback()
	err := cmd.Add()

	assert.Error(t, err)
}
