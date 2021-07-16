package models

import (
	"database/sql"
	"time"
)

type TransactionDetails struct {
	id             string
	transactionId  string
	carId          string
	carBrand       string
	carType        string
	carColor       string
	productionYear string
	price          float64
	quantity       int
	subTotal       float64
	createdAt      time.Time
	updatedAt      time.Time
	deletedAt      sql.NullTime
}

func NewTransactionDetailModel() *TransactionDetails {
	return &TransactionDetails{}
}

func (model *TransactionDetails) Id() string {
	return model.id
}

func (model *TransactionDetails) SetId(id string) *TransactionDetails {
	model.id = id

	return model
}

func (model *TransactionDetails) TransactionId() string {
	return model.transactionId
}

func (model *TransactionDetails) SetTransactionId(transactionId string) *TransactionDetails {
	model.transactionId = transactionId

	return model
}

func (model *TransactionDetails) CarId() string {
	return model.carId
}

func (model *TransactionDetails) SetCarId(carId string) *TransactionDetails {
	model.carId = carId

	return model
}

func (model *TransactionDetails) CarBrand() string {
	return model.carBrand
}

func (model *TransactionDetails) SetCarBrand(carBrand string) *TransactionDetails {
	model.carBrand = carBrand

	return model
}

func (model *TransactionDetails) CarType() string {
	return model.carType
}

func (model *TransactionDetails) SetCarType(carType string) *TransactionDetails {
	model.carType = carType

	return model
}

func (model *TransactionDetails) CarColor() string {
	return model.carColor
}

func (model *TransactionDetails) SetCarColor(carColor string) *TransactionDetails {
	model.carColor = carColor

	return model
}

func (model *TransactionDetails) ProductionYear() string {
	return model.productionYear
}

func (model *TransactionDetails) SetProductionYear(productionYear string) *TransactionDetails {
	model.productionYear = productionYear

	return model
}

func (model *TransactionDetails) Price() float64 {
	return model.price
}

func (model *TransactionDetails) SetPrice(price float64) *TransactionDetails {
	model.price = price

	return model
}

func (model *TransactionDetails) Quantity() int {
	return model.quantity
}

func (model *TransactionDetails) SetQuantity(quantity int) *TransactionDetails {
	model.quantity = quantity

	return model
}

func (model *TransactionDetails) SubTotal() float64 {
	return model.subTotal
}

func (model *TransactionDetails) SetSubTotal(subTotal float64) *TransactionDetails {
	model.subTotal = subTotal

	return model
}

func (model *TransactionDetails) CreatedAt() time.Time {
	return model.createdAt
}

func (model *TransactionDetails) SetCreatedAt(createdAt time.Time) *TransactionDetails {
	model.createdAt = createdAt

	return model
}

func (model *TransactionDetails) UpdatedAt() time.Time {
	return model.updatedAt
}

func (model *TransactionDetails) SetUpdatedAt(updatedAt time.Time) *TransactionDetails {
	model.updatedAt = updatedAt

	return model
}

func (model *TransactionDetails) DeletedAt() sql.NullTime {
	return model.deletedAt
}

func (model *TransactionDetails) SetDeletedAt(deletedAt time.Time) *TransactionDetails {
	model.deletedAt.Time = deletedAt

	return model
}

const (
	TransactionDetailSelectStatement = `SELECT id,transaction_id,car_id,car_brand,car_type,car_color,production_year,price,quantity,sub_total,` +
		`updated_at,created_at FROM transaction_details`
	TransactionDetailDefaultWhereStatement = `WHERE deleted_at IS NULL`
)

func (model *TransactionDetails) ScanRows(rows *sql.Rows) (interface{}, error) {
	err := rows.Scan(&model.id, &model.transactionId, &model.carId, &model.carBrand, &model.carType, &model.carColor, &model.productionYear, &model.price, &model.quantity, &model.subTotal,
		&model.createdAt, &model.updatedAt)
	if err != nil {
		return model, err
	}

	return model, nil
}
