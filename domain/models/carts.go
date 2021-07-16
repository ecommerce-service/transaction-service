package models

import (
	"database/sql"
	"time"
)

type Carts struct {
	id             string
	userId         string
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

func NewCartModel() *Carts {
	return &Carts{}
}

func (model *Carts) Id() string {
	return model.id
}

func (model *Carts) SetId(id string) *Carts {
	model.id = id

	return model
}

func (model *Carts) UserId() string {
	return model.userId
}

func (model *Carts) SetUserId(userId string) *Carts {
	model.userId = userId

	return model
}

func (model *Carts) CarId() string {
	return model.carId
}

func (model *Carts) SetCarId(carId string) *Carts {
	model.carId = carId

	return model
}

func (model *Carts) CarBrand() string {
	return model.carBrand
}

func (model *Carts) SetCarBrand(carBrand string) *Carts {
	model.carBrand = carBrand

	return model
}

func (model *Carts) CarType() string {
	return model.carType
}

func (model *Carts) SetCarType(carType string) *Carts {
	model.carType = carType

	return model
}

func (model *Carts) CarColor() string {
	return model.carColor
}

func (model *Carts) SetCarColor(carColor string) *Carts {
	model.carColor = carColor

	return model
}

func (model *Carts) ProductionYear() string {
	return model.productionYear
}

func (model *Carts) SetProductionYear(productionYear string) *Carts {
	model.productionYear = productionYear

	return model
}

func (model *Carts) Price() float64 {
	return model.price
}

func (model *Carts) SetPrice(price float64) *Carts {
	model.price = price

	return model
}

func (model *Carts) Quantity() int {
	return model.quantity
}

func (model *Carts) SetQuantity(quantity int) *Carts {
	model.quantity = quantity

	return model
}

func (model *Carts) SubTotal() float64 {
	return model.subTotal
}

func (model *Carts) SetSubTotal(subTotal float64) *Carts {
	model.subTotal = subTotal

	return model
}

func (model *Carts) CreatedAt() time.Time {
	return model.createdAt
}

func (model *Carts) SetCreatedAt(createdAt time.Time) *Carts {
	model.createdAt = createdAt

	return model
}

func (model *Carts) UpdatedAt() time.Time {
	return model.updatedAt
}

func (model *Carts) SetUpdatedAt(updatedAt time.Time) *Carts {
	model.updatedAt = updatedAt

	return model
}

func (model *Carts) DeletedAt() sql.NullTime {
	return model.deletedAt
}

func (model *Carts) SetDeletedAt(deletedAt time.Time) *Carts {
	model.deletedAt.Time = deletedAt

	return model
}

const (
	CartSelectStatement = `SELECT c.id,c.car_id,c.car_brand,c.car_type,c.car_color,c.production_year,c.price,c.quantity,c.sub_total,` +
		`c.created_at,c.updated_at FROM carts c`
	CartCountSelectStatement  = `SELECT count(c.id) FROM carts c`
	CartDefaultWhereStatement = `WHERE c.deleted_at IS NULL`
	CartJoinStatement         = `INNER JOIN users u ON u.id = c.user_id AND u.deleted_at IS NULL`
)

func (model *Carts) ScanRows(rows *sql.Rows) (interface{}, error) {
	err := rows.Scan(&model.id, &model.carId, &model.carBrand, &model.carType, &model.carColor, &model.productionYear, &model.price, &model.quantity, &model.subTotal,
		&model.createdAt, &model.updatedAt)
	if err != nil {
		return model, err
	}

	return model, nil
}

func (model *Carts) ScanRow(row *sql.Row) (interface{}, error) {
	err := row.Scan(&model.id, &model.carId, &model.carBrand, &model.carType, &model.carColor, &model.productionYear, &model.price, &model.quantity, &model.subTotal,
		&model.createdAt, &model.updatedAt)
	if err != nil {
		return model, err
	}

	return model, nil
}
