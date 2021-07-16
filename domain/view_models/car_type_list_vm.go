package view_models

import "booking-car/domain/models"

type CarTypeListVm struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func NewCarTypeListVm(model *models.CarTypes) CarTypeListVm {
	return CarTypeListVm{
		ID:   model.Id(),
		Name: model.Name(),
	}
}
