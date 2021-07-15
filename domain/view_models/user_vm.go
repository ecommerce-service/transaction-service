package view_models

import (
	"booking-car/domain/models"
	"time"
)

type UserVm struct {
	ID            string  `json:"id"`
	FirstName     string  `json:"first_name"`
	LastName      string  `json:"last_name"`
	Email         string  `json:"email"`
	UserName      string  `json:"user_name"`
	Address       string  `json:"address"`
	PhoneNumber   string  `json:"phone_number"`
	DepositAmount float64 `json:"deposit_amount"`
	Role          RoleVm  `json:"role"`
	CreatedAt     string  `json:"created_at"`
	UpdatedAt     string  `json:"updated_at"`
}

func NewUserVm(model *models.Users) UserVm {
	return UserVm{
		ID:            model.Id(),
		FirstName:     model.FirstName(),
		LastName:      model.LastName(),
		Email:         model.Email(),
		UserName:      model.UserName(),
		Address:       model.Address().String,
		PhoneNumber:   model.PhoneNumber(),
		DepositAmount: model.DepositAmount().Float64,
		Role:          NewRoleVm(model.Role),
		CreatedAt:     model.CreatedAt().Format(time.RFC3339),
		UpdatedAt:     model.UpdatedAt().Format(time.RFC3339),
	}
}
