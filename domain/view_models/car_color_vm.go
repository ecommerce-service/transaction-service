package view_models

import (
	"booking-car/domain/models"
	"time"
)

type CarColorVm struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	HexCode   string `json:"hex_code"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func NewCarColorVm(model *models.CarColors) CarColorVm{
	return CarColorVm{
		ID:        model.Id(),
		Name:      model.Name(),
		HexCode:   model.HexCode(),
		CreatedAt: model.CreatedAt().Format(time.RFC3339),
		UpdatedAt: model.UpdatedAt().Format(time.RFC3339),
	}
}
