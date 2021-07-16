package models

import (
	"database/sql"
	"time"
)

type CarColors struct {
	id        string
	name      string
	hexCode   string
	createdAt time.Time
	updatedAt time.Time
	deletedAt sql.NullTime
}

func NewCarColorModel() *CarColors{
	return &CarColors{}
}

func (model *CarColors) Id() string {
	return model.id
}

func (model *CarColors) SetId(id string)*CarColors {
	model.id = id

	return model
}

func (model *CarColors) Name() string {
	return model.name
}

func (model *CarColors) SetName(name string) *CarColors{
	model.name = name

	return model
}

func (model *CarColors) HexCode() string {
	return model.hexCode
}

func (model *CarColors) SetHexCode(hexCode string) *CarColors {
	model.hexCode = hexCode

	return model
}

func (model *CarColors) CreatedAt() time.Time {
	return model.createdAt
}

func (model *CarColors) SetCreatedAt(createdAt time.Time) *CarColors{
	model.createdAt = createdAt

	return model
}

func (model *CarColors) UpdatedAt() time.Time {
	return model.updatedAt
}

func (model *CarColors) SetUpdatedAt(updatedAt time.Time) *CarColors {
	model.updatedAt = updatedAt

	return model
}

func (model *CarColors) DeletedAt() sql.NullTime {
	return model.deletedAt
}

func (model *CarColors) SetDeletedAt(deletedAt time.Time) *CarColors {
	model.deletedAt.Time = deletedAt

	return model
}

const (
	CarColorSelectStatement = `SELECT id,name,hex_code,created_at,updated_at FROM car_colors`
	CarColorCountSelectStatement = `SELECT COUNT(id) FROM car_colors`
	CarColorDefaultWhereStatement = `WHERE deleted_at IS NULL`
)

func (model *CarColors) ScanRows(rows *sql.Rows) (interface{}, error) {
	err := rows.Scan(&model.id,&model.name,&model.hexCode,&model.createdAt,&model.updatedAt)
	if err != nil {
		return model,err
	}

	return model,nil
}

func (model *CarColors) ScanRow(row *sql.Row) (interface{}, error) {
	err := row.Scan(&model.id,&model.name,&model.hexCode,&model.createdAt,&model.updatedAt)
	if err != nil {
		return model,err
	}

	return model,nil
}
