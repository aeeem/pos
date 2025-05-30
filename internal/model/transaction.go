package model

import (
	"database/sql/driver"

	"gorm.io/gorm"
)

type Status string

const (
	Active   Status = "completed"
	Inactive Status = "cancelled"
	Pending  Status = "pending"
)

type Transaction struct {
	gorm.Model
	CustomerName string  `json:"customer_name"`
	Status       Status  `gorm:"type:status" json:"status" `
	TotalPrice   float64 `json:"total_price"`
	Cart         []Cart  `json:"cart"`
}

func (p *Status) Scan(value interface{}) error {
	*p = Status(value.([]byte))
	return nil
}

func (p Status) Value() (driver.Value, error) {
	return string(p), nil
}
