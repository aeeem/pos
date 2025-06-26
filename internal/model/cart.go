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
	err = tx.Model(Cart{}).Where("transaction_id = ?", C.TransactionID).Select("sum(sub_price)").Scan(&Total).Error
	if err != nil {
		return
	}
	err = tx.Table("transactions").Where("id = ?", C.TransactionID).Update("total_price", Total).Error
	return
}

func (C *Cart) AfterCreate(tx *gorm.DB) (err error) {
	Total := float64(0)
	err = tx.Model(Cart{}).Where("transaction_id = ?", C.TransactionID).Select("sum(sub_price)").Scan(&Total).Error
	if err != nil {
		return
	}
	err = tx.Table("transactions").Where("id = ?", C.TransactionID).Update("total_price", Total).Error
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
	err = tx.Model(&Cart{}).
		Where("transaction_id = ?", DetailedCart.TransactionID).Select("sum(sub_price) as total").
		Scan(&Total).Error
	if err != nil {
		return
	}
	log.Info().Any("Total", Total).Msg("Total")
	log.Info().Any("TransactionID", DetailedCart.TransactionID).Msg("TransactionID")
	err = tx.Table("transactions").Where("id = ?", DetailedCart.TransactionID).Update("total_price", Total).Error
	if err != nil {
		return
	}
	return
}
