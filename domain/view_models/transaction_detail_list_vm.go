package view_models

import "booking-car/pkg/str"

type TransactionDetailListVm struct {
	ID             string  `json:"id"`
	CarID          string  `json:"car_id"`
	CarBrand       string  `json:"car_brand"`
	CarType        string  `json:"car_type"`
	CarColor       string  `json:"car_color"`
	ProductionYear string  `json:"production_year"`
	Price          float64 `json:"price"`
	Quantity       int     `json:"quantity"`
	SubTotal       float64 `json:"sub_total"`
}

func NewTransactionDetailListVm(transactionDetails []string) TransactionDetailListVm {
	return TransactionDetailListVm{
		ID:             transactionDetails[0],
		CarID:          transactionDetails[1],
		CarBrand:       transactionDetails[2],
		CarType:        transactionDetails[3],
		CarColor:       transactionDetails[4],
		ProductionYear: transactionDetails[5],
		Price:          float64(str.StringToInt(transactionDetails[6])),
		Quantity:       str.StringToInt(transactionDetails[7]),
		SubTotal:       float64(str.StringToInt(transactionDetails[8])),
	}
}
