package model

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Status string

const (
	Completed Status = "completed"
	Cancelled Status = "cancelled"
	Pending   Status = "pending"
)

type Transaction struct {
	gorm.Model
	CustomerID            uint           `json:"customer_id" gorm:"index:idx_customer_id"`
	CustomerTransactionNo uint           `json:"customer_transaction_no"`
	CustomerName          string         `json:"customer_name" gorm:"index:idx_customer_name"`
	Status                Status         `gorm:"type:status" json:"status" `
	TotalPrice            float64        `json:"total_price"`
	Cart                  []Cart         `json:"cart"`
	CartJson              datatypes.JSON `json:"-"`
}
