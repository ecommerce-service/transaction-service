package requests

type CartRequest struct {
	CarID    string `json:"car_id" validate:"required"`
	Quantity int    `json:"quantity" validate:"required"`
}
