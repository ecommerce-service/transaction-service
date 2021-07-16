package requests

type CarColorRequest struct {
	Name    string `json:"name" validate:"required"`
	HexCode string `json:"hex_code"`
}
