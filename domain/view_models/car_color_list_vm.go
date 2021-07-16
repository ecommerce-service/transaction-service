package view_models

import "booking-car/domain/models"

type CarColorListVm struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	HexCode string `json:"hex_code"`
}

func NewCarColorListVm(model *models.CarColors) CarColorListVm {
	return CarColorListVm{
		ID:      model.Id(),
		Name:    model.Name(),
		HexCode: model.HexCode(),
	}
}
