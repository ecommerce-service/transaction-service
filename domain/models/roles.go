package models

import (
	"database/sql"
	"time"
)

type Roles struct {
	id        int64
	name      string
	createdAt time.Time
	updatedAt time.Time
	deletedAt sql.NullTime
}

func NewRoleModel() *Roles {
	return &Roles{}
}

func (model *Roles) Id() int64 {
	return model.id
}

func (model *Roles) SetId(id int64) *Roles {
	model.id = id

	return model
}

func (model *Roles) Name() string {
	return model.name
}

func (model *Roles) SetName(name string) *Roles {
	model.name = name

	return model
}

func (model *Roles) CreatedAt() time.Time {
	return model.createdAt
}

func (model *Roles) SetCreatedAt(createdAt time.Time) *Roles {
	model.createdAt = createdAt

	return model
}

func (model *Roles) UpdatedAt() time.Time {
	return model.updatedAt
}

func (model *Roles) SetUpdatedAt(updatedAt time.Time) *Roles {
	model.updatedAt = updatedAt

	return model
}

func (model *Roles) DeletedAt() sql.NullTime {
	return model.deletedAt
}

func (model *Roles) SetDeletedAt(deletedAt sql.NullTime) *Roles {
	model.deletedAt = deletedAt

	return model
}

const (
	RoleSelectStatement = `SELECT id,name FROM roles`
	RoleWhereStatement  = `WHERE LOWER(name) LIKE $1`
	RoleOrderStatement  = `ORDER BY id ASC`
)

func (model *Roles) ScanRows(rows *sql.Rows) (interface{}, error) {
	err := rows.Scan(&model.id, &model.name)
	if err != nil {
		return model, err
	}

	return model, nil
}

func (model *Roles) ScanRow(row *sql.Row) (interface{}, error) {
	err := row.Scan(&model.id, &model.name)
	if err != nil {
		return model, err
	}

	return model, nil
}
