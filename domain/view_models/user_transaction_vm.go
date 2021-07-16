package view_models

type UserTransactionVm struct {
	ID          string `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

func NewUserTransactionVm(id, firstName, lastName, email, phoneNumber string) UserTransactionVm {
	return UserTransactionVm{
		ID:          id,
		FirstName:   firstName,
		LastName:    lastName,
		Email:       email,
		PhoneNumber: phoneNumber,
	}
}
