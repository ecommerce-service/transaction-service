package models

import (
	"database/sql"
	"time"
)

type CarBrands struct {
	id        string
	name      string
	createdAt time.Time
	updatedAt time.Time
	deletedAt sql.NullTime
}

func NewCarBrandModel() *CarBrands {
	return &CarBrands{}
}

func (model *CarBrands) Id() string {
	return model.id
}

func (model *CarBrands) SetId(id string) *CarBrands {
	model.id = id

	return model
}

func (model *CarBrands) Name() string {
	return model.name
}

func (model *CarBrands) SetName(name string) *CarBrands {
	model.name = name

	return model
}

func (model *CarBrands) CreatedAt() time.Time {
	return model.createdAt
}

func (model *CarBrands) SetCreatedAt(createdAt time.Time) *CarBrands {
	model.createdAt = createdAt

	return model
}

func (model *CarBrands) UpdatedAt() time.Time {
	return model.updatedAt
}

func (model *CarBrands) SetUpdatedAt(updatedAt time.Time) *CarBrands {
	model.updatedAt = updatedAt

	return model
}

func (model *CarBrands) DeletedAt() sql.NullTime {
	return model.deletedAt
}

func (model *CarBrands) SetDeletedAt(deletedAt time.Time) *CarBrands {
	model.deletedAt.Time = deletedAt

	return model
}

const (
	BrandSelectStatement       = `SELECT id,name,created_at,updated_at FROM car_brands`
	BrandDefaultWhereStatement = `WHERE deleted_at IS NULL`
	BrandSelectCountStatement  = `SELECT COUNT(id) FROM car_brands`
)

func (model *CarBrands) ScanRows(rows *sql.Rows) (interface{}, error) {
	err := rows.Scan(&model.id, &model.name, &model.createdAt, &model.updatedAt)
	if err != nil {
		return model, err
	}

	return model, nil
}

func (model *CarBrands) ScanRow(row *sql.Row) (interface{}, error) {
	err := row.Scan(&model.id, &model.name, &model.createdAt, &model.updatedAt)
	if err != nil {
		return model, err
	}

	return model, nil
}
