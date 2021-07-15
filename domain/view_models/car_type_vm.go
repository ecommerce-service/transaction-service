package view_models

import (
	"booking-car/domain/models"
	"time"
)

type CarTypeVm struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	BrandID   string `json:"brand_id"`
	CreatedAt string `json:"created_at"`
	DeletedAt string `json:"deleted_at"`
}

func NewCarTypeVm(model *models.CarTypes) CarTypeVm{
	return CarTypeVm{
		ID:        model.Id(),
		Name:      model.Name(),
		BrandID:   model.BrandId(),
		CreatedAt: model.CreatedAt().Format(time.RFC3339),
		DeletedAt: model.UpdatedAt().Format(time.RFC3339),
	}
}
