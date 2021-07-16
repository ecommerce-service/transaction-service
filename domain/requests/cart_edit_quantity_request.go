package requests

type CartEditQuantityRequest struct {
	Quantity int `json:"quantity" validate:"required"`
}
