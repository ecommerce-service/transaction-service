package view_models

import (
	"booking-car/domain/models"
	"time"
)

type CarVm struct {
	ID             string         `json:"id"`
	CarBrand       CarBrandListVm `json:"car_brand"`
	CarType        CarTypeListVm  `json:"car_type"`
	CarColor       CarColorListVm `json:"car_color"`
	ProductionYear string         `json:"production_year"`
	Price          float64        `json:"price"`
	Stock          int            `json:"stock"`
	CreatedAt      string         `json:"created_at"`
	UpdatedAt      string         `json:"updated_at"`
}

func NewCarVm(model *models.Cars) CarVm {
	return CarVm{
		ID:             model.Id(),
		CarBrand:       NewCarBrandListVm(model.CarBrands),
		CarType:        NewCarTypeListVm(model.CarTypes),
		CarColor:       NewCarColorListVm(model.CarColors),
		ProductionYear: model.ProductionYear(),
		Price:          model.Price(),
		Stock:          model.Stock(),
		CreatedAt:      model.CreatedAt().Format(time.RFC3339),
		UpdatedAt:      model.UpdatedAt().Format(time.RFC3339),
	}
}
