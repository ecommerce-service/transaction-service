package commands

import (
	"booking-car/domain/models"
	"database/sql"
)

type TransactionDetailCommandMock struct {
	db    *sql.DB
	model *models.TransactionDetails
}

func (c TransactionDetailCommandMock) Add() (err error) {
	statement := `INSERT INTO transaction_details(transaction_id,car_id,car_brand,car_type,car_color,production_year,price,quantity,sub_total,created_at,updated_at) `+
		`VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)`
	tx,_ := c.db.Begin()

	_,err = tx.Exec(statement,c.model.TransactionId(),c.model.CarId(),c.model.CarBrand(),c.model.CarType(),c.model.CarColor(),c.model.ProductionYear(),c.model.Price(),
		c.model.Quantity(),c.model.SubTotal(),c.model.CreatedAt(),c.model.UpdatedAt())
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}
