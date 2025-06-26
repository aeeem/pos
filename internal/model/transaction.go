package model

import (
	"errors"

	"github.com/rs/zerolog/log"
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

func (T *Transaction) BeforeCreate(tx *gorm.DB) (err error) {
	log.Debug().Msg("masuk before insert")
	if T.Status == "" {
		log.Debug().Msg("masuk equals to draft")
		T.Status = "draft"
		return
	}

	if T.Status != "draft" {
		if T.Status == Status("cancelled") {
			err = errors.New("new transaction status cannot be cancelled")
			return
		}

		last_tx_no := 0
		err = tx.Model(Transaction{}).Where("status = 'pending' OR status = 'completed'").Select("customer_transaction_no").Order("customer_transaction_no desc").Limit(1).Scan(&last_tx_no).Error
		if err != nil {
			log.Debug().Err(err).Msg("sql Error")
		}
		log.Debug().Err(err).Msg("cek")
		T.CustomerTransactionNo = uint(last_tx_no + 1)
		return
	}
	return

}
func (T *Transaction) BeforeUpdate(tx *gorm.DB) (err error) {
	tOld := Transaction{}
	err = tx.Model(Transaction{}).Where("id = ?", T.ID).First(&tOld).Error
	if T.CustomerTransactionNo != tOld.CustomerTransactionNo {
		//check if customer transaction no is ok to update
		cnt := int64(0)
		err = tx.Table("transactions").Where("id < ?", T.ID).Count(&cnt).Error
		if err != nil {
			return
		} // count data  where id is less than t id
		/*
			I.E
			1 2 3 4 5

			change data where customer id is 4 to 1 is not valid
			because
			4 index is 3
			counting data before 4 is 3

			1 - 3 = -2 <- it cant be happend so throw error
			other example if
			4 changed to 6
			data count before 4 value is 3
			then
			6-3 = 3 <- valid
			because customer transaction no cant be 0 then throw error

		*/
		if cnt > 0 {
			if (int64(T.CustomerTransactionNo) - cnt) <= 0 {
				return errors.New("customer transaction no cant be this number other number can be negative or zero")
			}
		}
		if T.Status == "cancelled" && tOld.Status == "cancelled" {
			return errors.New("status already cancelled cannot update customer")
			//update_customer transaction no
		}
		err = tx.Table("transactions").Where("id <> ?", T.ID).Update("customer_transaction_no", gorm.Expr("(customer_transaction_no - ?) + ?", tOld.CustomerTransactionNo, T.CustomerTransactionNo)).Error
		if err != nil {
			log.Debug().Err(err).Msg("Error when updating customer transaction no")
			return
		}
	}

	if tOld.Status != T.Status {
		if tOld.Status == "draft" && (T.Status == "pending" || T.Status == "completed") {
			//update_customer transaction no
			last_tx_no := 0
			err = tx.Model(Transaction{}).Where("status = 'pending' OR status = 'completed'").Select("customer_transaction_no").Order("customer_transaction_no desc").Limit(1).Scan(&last_tx_no).Error
			T.CustomerTransactionNo = uint(last_tx_no + 1)
			print("update_customer_transaction_no", T.CustomerTransactionNo)
		}
		if tOld.Status == "pending" && (T.Status == "completed") {
			T.CustomerTransactionNo = tOld.CustomerTransactionNo
			//update_customer transaction no
		}
		if T.Status == "cancelled" && tOld.Status != "cancelled" {
			err = tx.Table("transactions").Where("id > ?", T.ID).Update("customer_transaction_no", gorm.Expr("customer_transaction_no - ?", 1)).Error
			if err != nil {
				log.Debug().Err(err).Msg("Error when updating customer transaction no")
				return
			}
			T.CustomerTransactionNo = 0
			//update_customer transaction no
		}
	}
	return
}
