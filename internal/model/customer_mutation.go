package model

import "gorm.io/gorm"

type CustomerDebtMutations struct {
	gorm.Model
	MutationID     uint    `json:"mutation_id" gorm:"primaryKey"`
	CustomerDebtID uint    `json:"customer_debt_id" gorm:"primaryKey"`
	CurrentDebt    float64 `json:"current_debt"`
	CurrentPayment float64 `json:"current_payment"`
	TotalPayment   float64 `json:"total_payment"`
}
