package view_models

import "booking-car/domain/models"

type CarBrandListVm struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func NewCarBrandListVm(model *models.CarBrands) CarBrandListVm {
	return CarBrandListVm{
		ID:   model.Id(),
		Name: model.Name(),
	}
}
