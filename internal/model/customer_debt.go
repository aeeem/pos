package model

import "gorm.io/gorm"

type DebtStatus string

const (
	Paid          DebtStatus = "completed"
	DebtCancelled DebtStatus = "cancelled"
	Unpaid        DebtStatus = "unpaid"
)

type CustomerDebt struct {
	gorm.Model
	CustomerID       uint       `json:"customer_id" gorm:"index:idx_customer_id"`
	Customer         Customer   `json:"customer" gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	TrxID            uint       `json:"trx_id"`
	DebtStatus       DebtStatus `json:"debt_status"`
	TotalTransaction float64    `json:"total_transaction"`
}
