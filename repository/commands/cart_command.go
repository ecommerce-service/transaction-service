package commands

import (
	"booking-car/domain/commands"
	"booking-car/domain/models"
	"booking-car/pkg/postgresql"
)

type CartCommand struct {
	db    postgresql.IConnection
	model *models.Carts
}

func NewCartCommand(db postgresql.IConnection, model *models.Carts) commands.ICartCommand {
	return &CartCommand{
		db:    db,
		model: model,
	}
}

func (c CartCommand) Add() (res string, err error) {
	statement := `INSERT INTO carts(user_id,car_id,car_brand,car_type,car_color,production_year,price,quantity,sub_total,created_at,updated_at) ` +
		`VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11) RETURNING id`

	err = c.db.GetDbInstance().QueryRow(statement, c.model.UserId(), c.model.CarId(), c.model.CarBrand(), c.model.CarType(), c.model.CarColor(), c.model.ProductionYear(), c.model.Price(),
		c.model.Quantity(), c.model.SubTotal(), c.model.CreatedAt(), c.model.UpdatedAt()).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (c CartCommand) Edit() (res string, err error) {
	statement := `UPDATE carts SET car_id=$1,car_brand=$2,car_type=$3,car_color=$4,production_year=$5,price=$6,quantity=$7,sub_total=$8,updated_at=$9 WHERE id=$10 RETURNING id`

	err = c.db.GetDbInstance().QueryRow(statement, c.model.CarId(), c.model.CarBrand(), c.model.CarType(), c.model.CarColor(), c.model.ProductionYear(), c.model.Price(),
		c.model.Quantity(), c.model.SubTotal(), c.model.UpdatedAt(), c.model.Id()).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (c CartCommand) Delete() (res string, err error) {
	statement := `UPDATE carts SET updated_at=$1,deleted_at=$2 WHERE id=$3 RETURNING id`

	err = c.db.GetDbInstance().QueryRow(statement, c.model.UpdatedAt(), c.model.DeletedAt().Time, c.model.Id()).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (c CartCommand) EditQuantity() (res string, err error) {
	statement := `UPDATE carts SET quantity=$1,sub_total=$2,updated_at=$3 WHERE id=$4 RETURNING id`

	err = c.db.GetDbInstance().QueryRow(statement, c.model.Quantity(), c.model.SubTotal(), c.model.UpdatedAt(), c.model.Id()).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (c CartCommand) DeleteAllByUserId() (err error) {
	statement := `UPDATE carts SET updated_at=$1,deleted_at=$2 WHERE user_id=$3 RETURNING id`

	_, err = c.db.GetTx().Exec(statement, c.model.UpdatedAt(), c.model.DeletedAt().Time, c.model.UserId())
	if err != nil {
		return err
	}

	return nil
}
