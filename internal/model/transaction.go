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
	Draft     Status = "draft"
)

type Transaction struct {
	gorm.Model
	CustomerID            uint           `json:"customer_id" gorm:"index:idx_customer_id"`
	Customer              Customer       `json:"customer" gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	CustomerTransactionNo uint           `json:"customer_transaction_no" gorm:"default:0"`
	CustomerName          string         `json:"customer_name" gorm:"index:idx_customer_name"`
	Status                Status         `gorm:"type:status" json:"status" gorm:"default:draft"`
	TotalPrice            float64        `json:"total_price" gorm:"default:0"`
	Cart                  []Cart         `json:"cart"`
	CartJson              datatypes.JSON `json:"-"`
}

func (T *Transaction) BeforeUpdate(tx *gorm.DB) (err error) {
	tOld := Transaction{}
	err = tx.Model(Transaction{}).Where("id = ?", T.ID).First(&tOld).Error
	if tOld.Status != T.Status {
		if tOld.Status == "draft" && (T.Status == "pending" || T.Status == "completed") {
			//update_customer transaction no
			last_tx_no := 0
			err = tx.Model(Transaction{}).Where("status = pending OR status = completed").Select("customer_transaction_no").Order("customer_transaction_no desc").Limit(1).Scan(&last_tx_no).Error
			T.CustomerTransactionNo = uint(last_tx_no + 1)

		}
		if tOld.Status == "pending" && (T.Status == "completed") {
			T.CustomerTransactionNo = tOld.CustomerTransactionNo
			//update_customer transaction no
		}
	}
	err = tx.Model(Cart{}).Where("transaction_id = ?", T.ID).Select("Coalasce(0,sum(sub_price))").Scan(&T.TotalPrice).Error
	//get transaction oldstate
	if err != nil {
		return
	}
	return
}
