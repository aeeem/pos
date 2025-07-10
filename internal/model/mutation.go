package model

import (
	"errors"

	"gorm.io/gorm"
)

type Mutation struct {
	gorm.Model
	MutationUUID    string         `json:"mutation_uuid" gorm:"unique"`
	CustomerID      uint           `json:"customer_id"`
	Customer        *Customer      `json:"customer"`
	MutationType    string         `json:"mutation_type"`
	Amount          float64        `json:"amount"`
	UsedAmount      float64        `json:"used_amount"`
	CustomerBalance float64        `json:"customer_balance"`
	Description     string         `json:"description"`
	MutationDate    string         `json:"mutation_date"`
	CustomerDebt    []CustomerDebt `json:"customer_debt" gorm:"many2many:customer_debt_mutations"`
}

func (m *Mutation) AfterCreate(tx *gorm.DB) (err error) {

	err = tx.Model(&Customer{}).Where("id = ?", m.CustomerID).Update("customer_balance", m.CustomerBalance).Error
	if err != nil {
		err = errors.New("error update customer balance " + err.Error())
		return
	}

	return
}
