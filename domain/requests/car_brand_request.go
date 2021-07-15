package requests

type CarBrandRequest struct {
	Name string `json:"name" validate:"required"`
}
