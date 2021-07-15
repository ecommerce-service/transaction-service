package view_models

import (
	"booking-car/domain/models"
	"time"
)

type CarBrandVm struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func NewCarBrandVm(model *models.CarBrands) CarBrandVm{
	return CarBrandVm{
		ID:        model.Id(),
		Name:      model.Name(),
		CreatedAt: model.CreatedAt().Format(time.RFC3339),
		UpdatedAt: model.UpdatedAt().Format(time.RFC3339),
	}
}
