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
	CustomerTransactionNo uint           `json:"customer_transaction_no" gorm:"default:0"`
	CustomerName          string         `json:"customer_name" gorm:"index:idx_customer_name"`
	Status                Status         `gorm:"type:status" json:"status" gorm:"default:draft"`
	TotalPrice            float64        `json:"total_price" gorm:"default:0"`
	Cart                  []Cart         `json:"cart"`
	CartJson              datatypes.JSON `json:"-"`
}
