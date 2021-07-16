package commands

import (
	"booking-car/domain/commands"
	"booking-car/domain/models"
	"booking-car/pkg/postgresql"
	"fmt"
)

type CarCommand struct {
	db    postgresql.IConnection
	model *models.Cars
}

func NewCarCommand(db postgresql.IConnection, model *models.Cars) commands.ICarCommand {
	return &CarCommand{
		db:    db,
		model: model,
	}
}

func (c CarCommand) Add() (res string, err error) {
	statement := `INSERT INTO cars (car_type_id,car_color_id,production_year,price,stock,created_at,updated_at) VALUES($1,$2,$3,$4,$5,$6,$7) RETURNING id`
	fmt.Println(c.model.ProductionYear())

	err = c.db.GetDbInstance().QueryRow(statement, c.model.CarTypeId(), c.model.CarColorId(), c.model.ProductionYear(), c.model.Price(), c.model.Stock(),
		c.model.CreatedAt(), c.model.UpdatedAt()).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (c CarCommand) Edit() (res string, err error) {
	statement := `UPDATE cars SET car_type_id=$1,car_color_id=$2,production_year=$3,price=$4,stock=$5,updated_at=$6 WHERE id=$7 RETURNING id`

	err = c.db.GetDbInstance().QueryRow(statement, c.model.CarTypeId(), c.model.CarColorId(), c.model.ProductionYear(), c.model.Price(), c.model.Stock(),
		c.model.UpdatedAt(), c.model.Id()).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (c CarCommand) Delete() (res string, err error) {
	statement := `UPDATE cars SET updated_at=$1, deleted_at=$2 WHERE id=$3 RETURNING id`

	err = c.db.GetDbInstance().QueryRow(statement, c.model.UpdatedAt(), c.model.DeletedAt().Time, c.model.Id()).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (c CarCommand) EditStock() (err error) {
	statement := `UPDATE cars SET stock=$1,updated_at=$2 WHERE id=$3 RETURNING id`

	_, err = c.db.GetTx().Exec(statement, c.model.Stock(), c.model.UpdatedAt(), c.model.Id())
	if err != nil {
		return err
	}

	return nil
}
