package requests

type TransactionDetailRequest struct {
	CarID          string  `json:"car_id"`
	CarBrand       string  `json:"car_brand"`
	CarType        string  `json:"car_type"`
	CarColor       string  `json:"car_color"`
	ProductionYear string  `json:"production_year"`
	Price          float64 `json:"price"`
	Quantity       int     `json:"quantity"`
	SubTotal       float64 `json:"sub_total"`
}
