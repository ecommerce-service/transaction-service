package models

import (
	"database/sql"
	"time"
)

type CarTypes struct {
	id        string
	brandId   string
	name      string
	createdAt time.Time
	updatedAt time.Time
	deletedAt sql.NullTime
}

func NewCarTypeModel() *CarTypes{
	return &CarTypes{}
}

func (model *CarTypes) Id() string {
	return model.id
}

func (model *CarTypes) SetId(id string) *CarTypes {
	model.id = id

	return model
}

func (model *CarTypes) BrandId() string {
	return model.brandId
}

func (model *CarTypes) SetBrandId(brandId string) *CarTypes {
	model.brandId = brandId

	return model
}

func (model *CarTypes) Name() string {
	return model.name
}

func (model *CarTypes) SetName(name string) *CarTypes {
	model.name = name

	return model
}

func (model *CarTypes) CreatedAt() time.Time {
	return model.createdAt
}

func (model *CarTypes) SetCreatedAt(createdAt time.Time) *CarTypes {
	model.createdAt = createdAt

	return model
}

func (model *CarTypes) UpdatedAt() time.Time {
	return model.updatedAt
}

func (model *CarTypes) SetUpdatedAt(updatedAt time.Time) *CarTypes {
	model.updatedAt = updatedAt

	return model
}

func (model *CarTypes) DeletedAt() sql.NullTime {
	return model.deletedAt
}

func (model *CarTypes) SetDeletedAt(deletedAt time.Time) *CarTypes {
	model.deletedAt.Time = deletedAt

	return model
}

const(
	CarTypeSelectStatement = `SELECT id,brand_id,name,created_at,updated_at FROM car_types`
	CarTypeSelectCountStatement = `SELECT COUNT(id) FROM car_types`
	CarTypeDefaultWhereStatement = `WHERE deleted_at IS NULL`
)

func (model *CarTypes) ScanRows(rows *sql.Rows) (interface{}, error) {
	err := rows.Scan(&model.id,&model.brandId,&model.name,&model.createdAt,&model.updatedAt)
	if err != nil {
		return model,err
	}

	return model,nil
}

func (model *CarTypes) ScanRow(row *sql.Row) (interface{}, error) {
	err := row.Scan(&model.id,&model.brandId,&model.name,&model.createdAt,&model.updatedAt)
	if err != nil {
		return model,err
	}

	return model,nil
}
