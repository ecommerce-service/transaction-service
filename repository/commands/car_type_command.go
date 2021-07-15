package commands

import (
	"booking-car/domain/commands"
	"booking-car/domain/models"
	"database/sql"
)

type CarTypeCommand struct{
	db *sql.DB
	model *models.CarTypes
}

func NewCarTypeCommand(db *sql.DB,model *models.CarTypes) commands.IBaseCommand{
	return &CarTypeCommand{
		db:    db,
		model: model,
	}
}

func (c CarTypeCommand) Add() (res string, err error) {
	statement := `INSERT INTO car_types(name,brand_id,created_at,updated_at) VALUES($1,$2,$3,$4) RETURNING id`

	err = c.db.QueryRow(statement,c.model.Name(),c.model.BrandId(),c.model.CreatedAt(),c.model.UpdatedAt()).Scan(&res)
	if err != nil {
		return res,err
	}

	return res,nil
}

func (c CarTypeCommand) Edit() (res string, err error) {
	statement := `UPDATE car_types SET name=$1,brand_id=$2,updated_at=$3 WHERE id=$4 RETURNING id`

	err = c.db.QueryRow(statement,c.model.Name(),c.model.BrandId(),c.model.UpdatedAt(),c.model.Id()).Scan(&res)
	if err != nil {
		return res,err
	}

	return res,nil
}

func (c CarTypeCommand) Delete() (res string, err error) {
	statement := `UPDATE car_types SET updated_at=$1,deleted_at=$2 WHERE id=$3 RETURNING id`

	err = c.db.QueryRow(statement,c.model.UpdatedAt(),c.model.DeletedAt().Time,c.model.Id()).Scan(&res)
	if err != nil {
		return res,err
	}

	return res,nil
}

