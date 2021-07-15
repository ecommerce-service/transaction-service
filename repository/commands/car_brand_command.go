package commands

import (
	"booking-car/domain/commands"
	"booking-car/domain/models"
	"database/sql"
)

type CarBrandCommand struct {
	db    *sql.DB
	model *models.CarBrands
}

func NewCarBrandCommand(db *sql.DB, model *models.CarBrands) commands.IBaseCommand {
	return &CarBrandCommand{
		db:    db,
		model: model,
	}
}

func (c CarBrandCommand) Add() (res string, err error) {
	statement := `INSERT INTO car_brands (name,created_at,updated_at) VALUES($1,$2,$3) RETURNING id`

	err = c.db.QueryRow(statement, c.model.Name(), c.model.CreatedAt(), c.model.UpdatedAt()).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (c CarBrandCommand) Edit() (res string, err error) {
	statement := `UPDATE car_brands SET name=$1,updated_at=$2 WHERE id=$3 RETURNING id`

	err = c.db.QueryRow(statement, c.model.Name(), c.model.UpdatedAt(), c.model.Id()).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (c CarBrandCommand) Delete() (res string, err error) {
	statement := `UPDATE car_brands SET updated_at=$1,deleted_at=$2 WHERE id=$3 RETURNING id`

	err = c.db.QueryRow(statement, c.model.UpdatedAt(), c.model.DeletedAt().Time, c.model.Id()).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}
