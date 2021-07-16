package commands

import (
	"booking-car/domain/commands"
	"booking-car/domain/models"
	"database/sql"
)

type CarColorCommand struct {
	db    *sql.DB
	model *models.CarColors
}

func NewCarColorCommand(db *sql.DB, model *models.CarColors) commands.IBaseCommand {
	return &CarColorCommand{
		db:    db,
		model: model,
	}
}

func (c CarColorCommand) Add() (res string, err error) {
	statement := `INSERT INTO car_colors (name,hex_code,created_at,updated_at) VALUES($1,$2,$3,$4) RETURNING id`

	err = c.db.QueryRow(statement, c.model.Name(), c.model.HexCode(), c.model.CreatedAt(), c.model.UpdatedAt()).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (c CarColorCommand) Edit() (res string, err error) {
	statement := `UPDATE car_colors SET name=$1,hex_code=$2,updated_at=$3 WHERE id=$4 RETURNING id`

	err = c.db.QueryRow(statement, c.model.Name(), c.model.HexCode(), c.model.UpdatedAt(), c.model.Id()).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (c CarColorCommand) Delete() (res string, err error) {
	statement := `UPDATE car_colors SET updated_at=$1,deleted_at=$2 WHERE id=$3 RETURNING id`

	err = c.db.QueryRow(statement, c.model.UpdatedAt(), c.model.DeletedAt().Time, c.model.Id()).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}
