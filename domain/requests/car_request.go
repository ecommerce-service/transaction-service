package requests

type CarRequest struct {
	CarTypeID      string  `json:"car_type_id" validate:"required"`
	CarColorID     string  `json:"car_color_id" validate:"required"`
	ProductionYear string  `json:"production_year" validate:"required"`
	Price          float64 `json:"price" validate:"required"`
	Stock          int     `json:"stock" validate:"required"`
}
