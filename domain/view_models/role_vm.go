package view_models

import (
	"booking-car/domain/models"
)

type RoleVm struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func NewRoleVm(model *models.Roles) RoleVm {
	return RoleVm{
		ID:   model.Id(),
		Name: model.Name(),
	}
}
