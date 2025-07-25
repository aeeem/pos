package domain

import (
	"errors"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type DebtStatus string

const (
	Paid     DebtStatus = "paid"
	HalfPaid DebtStatus = "half_paid"
	Unpaid   DebtStatus = "unpaid"
)

type CustomerDebt struct {
	gorm.Model
	CustomerID       uint       `json:"customer_id" gorm:"index:idx_customer_id"`
	Customer         *Customer  `json:"customer" gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	TrxID            uint       `json:"trx_id"`
	DebtStatus       DebtStatus `json:"debt_status" gorm:"type:debt_status;default:unpaid;not null"`
	TotalTransaction float64    `json:"total_transaction"`
	PaidAmount       float64    `json:"paid_amount" gorm:"default:0;not null"`
	UnpaidAmount     float64    `json:"unpaid_amount" gorm:"default:0;not null"`
	Mutation         []Mutation `json:"mutation" gorm:"many2many:customer_debt_mutations;"`
}

func (c *CustomerDebt) AfterUpdate(tx *gorm.DB) (err error) {
	//get customer debt

	cd := CustomerDebt{}
	err = tx.Model(&CustomerDebt{}).Where("trx_id = ?", c.TrxID).First(&cd).Error
	if err != nil {
		return
	}
	//get the customer debt
	err = tx.Table("customer_debts").Where("trx_id = ?", c.TrxID).Find(&c).Error
	if err != nil {
		return
	}
	log.Info().Any("cd", cd).Msg("Msg")

	if cd.DebtStatus == Paid {
		trx := Transaction{}
		trx.ID = cd.TrxID
		trx.CustomerID = cd.CustomerID
		err = tx.Model(&trx).Where("id = ?", cd.TrxID).Update("status", "completed").Error
		if err != nil {
			err = errors.New(err.Error() + "Transaction Not found")
			return
		}
	} else if cd.DebtStatus == HalfPaid {
		customer := Customer{}
		customer.ID = cd.CustomerID
		err = tx.Model(&customer).Where("id = ?", cd.CustomerID).Update("customer_total_debt", cd.UnpaidAmount).Error
		if err != nil {
			err = errors.New(err.Error() + "Customer Not found")
			return
		}

	}

	//updating customer total debt even if they paid in half

	return
}
