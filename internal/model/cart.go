package model

import (
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	ItemID        uint         `json:"item_id"`
	Item          *Item        `json:"item" gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	TransactionID uint         `json:"transaction_id" `
	Transaction   *Transaction `json:"transaction" gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	ItemName      string       `json:"item_name"`
	Unit          string       `json:"unit"`
	Quantity      float64      `json:"quantity"`
	ItemPrice     float64      `json:"item_price"`
	PriceID       uint         `json:"price_id"`
	Price         Price        `json:"price" gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	SubPrice      float64      `json:"sub_price"`
}

func (C *Cart) AfterUpdate(tx *gorm.DB) (err error) {
	Total := float64(0)
	err = tx.Table("carts").Where("transaction_id = ?  AND deleted_at is null", C.TransactionID).Select("sum(sub_price)").Scan(&Total).Error
	if err != nil {
		return
	}

	err = tx.Table("transactions").Where("id = ?", C.TransactionID).Update("total_price", Total).Error
	if err != nil {
		return
	}
	T := Transaction{}
	err = tx.Model(Transaction{}).Where("id = ?", C.TransactionID).First(&T).Error
	if err != nil {
		return
	}
	if T.Status == "pending" {
		err = tx.Model(CustomerDebt{}).Where("trx_id=?", T.ID).Update("total_transaction", Total).Error
		if err != nil {
			return
		}
		TotalDebt := float64(0)
		err = tx.Table("customer_debts").Where("customer_id = ?", T.CustomerID).Where("debt_status = ? OR debt_status = ?", "unpaid", "half_paid").Select("sum(total_transaction)").Scan(&TotalDebt).Error
		if err != nil {
			return
		}
		err = tx.Table("customers").Where("id = ?", T.CustomerID).Update("customer_total_debt", TotalDebt).Error
		if err != nil {
			return
		}
	}
	CustomerTotalTransaction := float64(0)
	err = tx.Table("transactions").Where("customer_id = ?", T.CustomerID).Select("sum(total_price)").Scan(&CustomerTotalTransaction).Error
	if err != nil {
		return
	}
	err = tx.Table("customers").Where("id = ?", T.CustomerID).Update("customer_total_transaction", CustomerTotalTransaction).Error
	if err != nil {
		return
	}
	return
}

func (C *Cart) AfterCreate(tx *gorm.DB) (err error) {
	Total := float64(0)
	err = tx.Table("carts").Where("transaction_id = ?  AND deleted_at is null", C.TransactionID).Select("sum(sub_price)").Scan(&Total).Error
	if err != nil {
		return
	}
	err = tx.Table("transactions").Where("id = ?", C.TransactionID).Update("total_price", Total).Error
	//get transactions
	T := Transaction{}
	err = tx.Model(Transaction{}).Where("id = ?", C.TransactionID).First(&T).Error
	if err != nil {
		return
	}
	if T.Status == "pending" {
		err = tx.Model(CustomerDebt{}).Where("trx_id=?", T.ID).Update("total_transaction", Total).Error
		if err != nil {
			return
		}
		TotalDebt := float64(0)
		err = tx.Table("customer_debts").Where("customer_id = ?", T.CustomerID).Where("debt_status = ? OR debt_status = ?", "unpaid", "half_paid").Select("sum(total_transaction)").Scan(&TotalDebt).Error
		if err != nil {
			return
		}
		err = tx.Table("customers").Where("id = ?", T.CustomerID).Update("customer_total_debt", TotalDebt).Error
		if err != nil {
			return
		}
	}

	if T.Status == "completed" { //kalau complete update customer debt
		//jadikan half paid karena sudah ada sebagian yang dibayar sebelumnya
		err = tx.Model(CustomerDebt{}).Where("trx_id=?", T.ID).Update("debt_status", "half_paid").Update("total_transaction", Total).Error
		if err != nil {
			return
		}
		TotalDebt := float64(0)
		//kalkulasi ulang Total Debt
		err = tx.Table("customer_debts").Where("customer_id = ?", T.CustomerID).Where("debt_status = ? OR debt_status = ?", "unpaid", "half_paid").Select("sum(total_transaction)").Scan(&TotalDebt).Error
		if err != nil {
			return
		}
		//getcustomerdebt and paid for these customers
		//kalkulasi ulang total debt customer
		err = tx.Table("customers").Where("id = ?", T.CustomerID).Update("customer_total_debt", TotalDebt).Error
		if err != nil {
			return
		}
		T.Status = "pending"
		err = tx.Model(&Transaction{}).Updates(&T).Error
	}
	CustomerTotalTransaction := float64(0)
	err = tx.Table("transactions").Where("customer_id = ?", T.CustomerID).Select("sum(total_price)").Scan(&CustomerTotalTransaction).Error
	if err != nil {
		return
	}
	err = tx.Table("customers").Where("id = ?", T.CustomerID).Update("customer_total_transaction", CustomerTotalTransaction).Error
	if err != nil {
		return
	}
	return
}

func (C *Cart) AfterDelete(tx *gorm.DB) (err error) {
	//get detail cart
	DetailedCart := Cart{}
	err = tx.Unscoped().Where("id = ?", C.ID).Find(&DetailedCart).Error
	if err != nil {
		return
	}
	log.Info().Any("Cart", DetailedCart).Msg("carts")
	Total := float64(0)
	err = tx.Table("carts").
		Where("transaction_id = ? AND deleted_at is null", DetailedCart.TransactionID).Select("sum(sub_price) as total").
		Scan(&Total).Error
	if err != nil {
		return
	}
	err = tx.Table("transactions").Where("id = ?", DetailedCart.TransactionID).Update("total_price", Total).Error
	if err != nil {
		return
	}
	T := Transaction{}
	err = tx.Model(Transaction{}).Where("id = ?", DetailedCart.TransactionID).First(&T).Error
	if err != nil {
		log.Error().Err(err).Msg("error")
		return
	}
	if T.Status == "pending" {
		err = tx.Model(CustomerDebt{}).Where("trx_id=?", T.ID).Update("total_transaction", Total).Error
		if err != nil {
			return
		}
		TotalDebt := float64(0)
		err = tx.Table("customer_debts").Where("customer_id = ?", T.CustomerID).Where("debt_status = ? OR debt_status = ?", "unpaid", "half_paid").Select("sum(total_transaction)").Scan(&TotalDebt).Error
		if err != nil {
			return
		}
		err = tx.Table("customers").Where("id = ?", T.CustomerID).Update("customer_total_debt", TotalDebt).Error
		if err != nil {
			return
		}
	}
	CustomerTotalTransaction := float64(0)
	err = tx.Table("transactions").Where("customer_id = ?", T.CustomerID).Select("sum(total_price)").Scan(&CustomerTotalTransaction).Error
	if err != nil {
		return
	}
	err = tx.Table("customers").Where("id = ?", T.CustomerID).Update("customer_total_transaction", CustomerTotalTransaction).Error
	if err != nil {
		return
	}
	return
}
