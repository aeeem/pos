package model

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	CustomerTotalDebt        float64 `json:"customer_total_debt"`
	CustomerTotalTransaction float64 `json:"customer_total_transaction"`
	CustomerName             string  `json:"customer_name" `
	PhoneNumber              string  `json:"phone_number"`
}
