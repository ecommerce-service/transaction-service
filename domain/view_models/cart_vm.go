package view_models

import (
	"booking-car/domain/models"
	"time"
)

type CartVm struct {
	ID             string  `json:"id"`
	CarID          string  `json:"car_id"`
	CarBrand       string  `json:"car_brand"`
	CarType        string  `json:"car_type"`
	CarColor       string  `json:"car_color"`
	ProductionYear string  `json:"production_year"`
	Price          float64 `json:"price"`
	Quantity       int     `json:"quantity"`
	SubTotal       float64 `json:"sub_total"`
	CreatedAt      string  `json:"created_at"`
	UpdatedAt      string  `json:"updated_at"`
}

func NewCartVm(model *models.Carts) CartVm {
	return CartVm{
		ID:             model.Id(),
		CarID:          model.CarId(),
		CarBrand:       model.CarBrand(),
		CarType:        model.CarType(),
		CarColor:       model.CarColor(),
		ProductionYear: model.ProductionYear(),
		Price:          model.Price(),
		Quantity:       model.Quantity(),
		SubTotal:       model.SubTotal(),
		CreatedAt:      model.CreatedAt().Format(time.RFC3339),
		UpdatedAt:      model.UpdatedAt().Format(time.RFC3339),
	}
}
