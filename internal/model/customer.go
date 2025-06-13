package model

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	CustomerName string `json:"customer_name" `
	PhoneNumber  string `json:"phone_number"`
}
