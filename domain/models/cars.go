package models

import (
	"database/sql"
	"time"
)

type Cars struct {
	id             string
	carTypeId      string
	carColorId     string
	productionYear string
	price          float64
	stock          int
	createdAt      time.Time
	updatedAt      time.Time
	deletedAt      sql.NullTime

	CarBrands *CarBrands
	CarTypes  *CarTypes
	CarColors *CarColors
}

func NewCarModel() *Cars {
	return &Cars{}
}

func (model *Cars) Id() string {
	return model.id
}

func (model *Cars) SetId(id string) *Cars {
	model.id = id

	return model
}

func (model *Cars) CarTypeId() string {
	return model.carTypeId
}

func (model *Cars) SetCarTypeId(carTypeId string) *Cars {
	model.carTypeId = carTypeId

	return model
}

func (model *Cars) CarColorId() string {
	return model.carColorId
}

func (model *Cars) SetCarColorId(carColorId string) *Cars {
	model.carColorId = carColorId

	return model
}

func (model *Cars) ProductionYear() string {
	return model.productionYear
}

func (model *Cars) SetProductionYear(productionYear string) *Cars {
	model.productionYear = productionYear

	return model
}

func (model *Cars) Price() float64 {
	return model.price
}

func (model *Cars) SetPrice(price float64) *Cars {
	model.price = price

	return model
}

func (model *Cars) Stock() int {
	return model.stock
}

func (model *Cars) SetStock(stock int) *Cars {
	model.stock = stock

	return model
}

func (model *Cars) CreatedAt() time.Time {
	return model.createdAt
}

func (model *Cars) SetCreatedAt(createdAt time.Time) *Cars {
	model.createdAt = createdAt

	return model
}

func (model *Cars) UpdatedAt() time.Time {
	return model.updatedAt
}

func (model *Cars) SetUpdatedAt(updatedAt time.Time) *Cars {
	model.updatedAt = updatedAt

	return model
}

func (model *Cars) DeletedAt() sql.NullTime {
	return model.deletedAt
}

func (model *Cars) SetDeletedAt(deletedAt time.Time) *Cars {
	model.deletedAt.Time = deletedAt

	return model
}

const (
	CarSelectStatement = `SELECT c.id,c.production_year,c.price,c.stock,c.created_at,c.updated_at,` +
		`b.id,b.name,ct.id,ct.name,cc.id,cc.name,cc.hex_code FROM cars c`
	CarCountSelectStatement = `SELECT COUNT(c.id) FROM cars c`
	CarJoinStatement        = `INNER JOIN car_types ct ON ct.id = c.car_type_id AND ct.deleted_at IS NULL ` +
		`INNER JOIN car_brands b ON b.id = ct.brand_id AND b.deleted_at IS NULL ` +
		`INNER JOIN car_colors cc ON cc.id = c.car_color_id AND cc.deleted_at IS NULL`
	CarDefaultWhereStatement = `WHERE c.deleted_at IS NULL`
)

func (model *Cars) ScanRows(rows *sql.Rows) (interface{}, error) {
	model.CarBrands = NewCarBrandModel()
	model.CarTypes = NewCarTypeModel()
	model.CarColors = NewCarColorModel()

	err := rows.Scan(&model.id, &model.productionYear, &model.price, &model.stock, &model.createdAt, &model.updatedAt, &model.CarBrands.id, &model.CarBrands.name,
		&model.CarTypes.id, &model.CarTypes.name, &model.CarColors.id, &model.CarColors.name, &model.CarColors.hexCode)
	if err != nil {
		return model, err
	}

	return model, nil
}

func (model *Cars) ScanRow(row *sql.Row) (interface{}, error) {
	model.CarBrands = NewCarBrandModel()
	model.CarTypes = NewCarTypeModel()
	model.CarColors = NewCarColorModel()

	err := row.Scan(&model.id, &model.productionYear, &model.price, &model.stock, &model.createdAt, &model.updatedAt, &model.CarBrands.id, &model.CarBrands.name,
		&model.CarTypes.id, &model.CarTypes.name, &model.CarColors.id, &model.CarColors.name, &model.CarColors.hexCode)
	if err != nil {
		return model, err
	}

	return model, nil
}
